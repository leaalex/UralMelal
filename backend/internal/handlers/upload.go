package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"

	"uralmelal/backend/internal/config"
)

var allowedExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
}

var allowedMime = map[string]bool{
	"image/jpeg": true, "image/png": true, "image/gif": true, "image/webp": true,
}

type UploadHandler struct {
	cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) *UploadHandler {
	return &UploadHandler{cfg: cfg}
}

func (h *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, `{"error":"no file"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExt[ext] {
		http.Error(w, `{"error":"invalid file type"}`, http.StatusBadRequest)
		return
	}

	maxSize := int64(h.cfg.MaxUploadMB) * 1024 * 1024
	if header.Size > maxSize {
		http.Error(w, fmt.Sprintf(`{"error":"file too large, max %dMB"}`, h.cfg.MaxUploadMB), http.StatusBadRequest)
		return
	}

	// Validate content
	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	buf = buf[:n]
	mime := http.DetectContentType(buf)
	if !allowedMime[mime] {
		http.Error(w, `{"error":"invalid file content"}`, http.StatusBadRequest)
		return
	}

	if err := os.MkdirAll(h.cfg.UploadPath, 0755); err != nil {
		http.Error(w, `{"error":"storage error"}`, http.StatusInternalServerError)
		return
	}

	tmpPath := filepath.Join(h.cfg.UploadPath, fmt.Sprintf(".tmp%d", time.Now().UnixNano()))
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		http.Error(w, `{"error":"storage error"}`, http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(tmpFile, io.MultiReader(bytes.NewReader(buf), file)); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		http.Error(w, `{"error":"storage error"}`, http.StatusInternalServerError)
		return
	}
	tmpFile.Close()

	img, err := imaging.Open(tmpPath)
	os.Remove(tmpPath)
	if err != nil {
		http.Error(w, `{"error":"invalid image"}`, http.StatusBadRequest)
		return
	}
	// Resize if too large (max 1920px width)
	if img.Bounds().Dx() > 1920 {
		img = imaging.Resize(img, 1920, 0, imaging.Lanczos)
	}
	name := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
	path := filepath.Join(h.cfg.UploadPath, name)
	if err := imaging.Save(img, path, imaging.JPEGQuality(85)); err != nil {
		http.Error(w, `{"error":"storage error"}`, http.StatusInternalServerError)
		return
	}
	url := "/storage/" + name
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"url":"%s"}`, url)))
}
