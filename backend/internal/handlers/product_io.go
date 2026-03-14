package handlers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"

	"gorm.io/gorm"
)

// MC CSV columns: Категория, Подкатегория, Наименование, Размер, Марка, Длина, Город, Остаток, Цена 1т, Цена 5т, Цена 10т, Артикул, Описание, Характеристики, Изображения, URL, Дубликаты
const (
	mcColCategory = 0
	mcColSubcat   = 1
	mcColName     = 2
	mcColSize     = 3
	mcColMark     = 4
	mcColLength   = 5
	mcColCity     = 6
	mcColStock    = 7
	mcColPrice1   = 8
	mcColPrice5   = 9
	mcColPrice10  = 10
	mcColSKU      = 11
	mcColDesc     = 12
	mcColChars    = 13
	mcColImages   = 14
	mcColURL      = 15
)

type ProductIOHandler struct {
	repo     *repository.ProductRepository
	catRepo  *repository.CategoryRepository
	db       *gorm.DB
}

func NewProductIOHandler(repo *repository.ProductRepository, catRepo *repository.CategoryRepository, db *gorm.DB) *ProductIOHandler {
	return &ProductIOHandler{repo: repo, catRepo: catRepo, db: db}
}

func (h *ProductIOHandler) Export(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json"
	}

	params := repository.ProductListParams{PerPage: 10000}
	result, err := h.repo.List(params)
	if err != nil {
		http.Error(w, `{"error":"export failed"}`, http.StatusInternalServerError)
		return
	}

	switch format {
	case "csv":
		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=products.csv")
		writer := csv.NewWriter(w)
		writer.Write([]string{"name", "sku", "category_id", "description", "image_path", "attributes"})
		for _, p := range result.Products {
			catID := ""
			if p.CategoryID != nil {
				catID = strconv.FormatUint(uint64(*p.CategoryID), 10)
			}
			attrsJSON := ""
			if len(p.Attributes) > 0 {
				b, _ := json.Marshal(p.Attributes)
				attrsJSON = string(b)
			}
			writer.Write([]string{p.Name, p.SKU, catID, p.Description, p.ImagePath, attrsJSON})
		}
		writer.Flush()
	case "mc":
		if h.catRepo == nil {
			http.Error(w, `{"error":"mc export requires category repo"}`, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=products_mc.csv")
		writer := csv.NewWriter(w)
		writer.Comma = ';'
		writer.Write([]string{"Категория", "Подкатегория", "Наименование", "Размер", "Марка", "Длина", "Город", "Остаток", "Цена 1т", "Цена 5т", "Цена 10т", "Артикул", "Описание", "Характеристики", "Изображения", "URL", "Дубликаты"})
		for _, p := range result.Products {
			catName, subcatName := "", ""
			if p.CategoryID != nil {
				catName, subcatName = h.catRepo.GetPathByID(*p.CategoryID)
			}
			chars := ""
			for _, a := range p.Attributes {
				if a.Key == "characteristics" {
					chars = a.Value
					break
				}
			}
			price1 := strings.Replace(strconv.FormatFloat(p.Price1t, 'f', -1, 64), ".", ",", 1)
			price5 := strings.Replace(strconv.FormatFloat(p.Price5t, 'f', -1, 64), ".", ",", 1)
			price10 := strings.Replace(strconv.FormatFloat(p.Price10t, 'f', -1, 64), ".", ",", 1)
			writer.Write([]string{
				catName, subcatName, p.Name, p.Size, p.Mark, p.Length, p.City,
				strconv.Itoa(p.Stock), price1, price5, price10,
				p.SKU, p.Description, chars, p.ImagePath, p.SourceURL, "",
			})
		}
		writer.Flush()
	case "json":
		fallthrough
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", "attachment; filename=products.json")
		json.NewEncoder(w).Encode(result.Products)
	}
}

type ImportProduct struct {
	Name        string            `json:"name"`
	SKU         string            `json:"sku"`
	CategoryID  *uint             `json:"category_id"`
	Description string            `json:"description"`
	ImagePath   string            `json:"image_path"`
	Attributes  []ProductAttrIn   `json:"attributes"`
}

type ProductAttrIn struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h *ProductIOHandler) Import(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json"
	}

	var bodyReader io.Reader = r.Body
	if file, _, err := r.FormFile("file"); err == nil {
		defer file.Close()
		bodyReader = file
	}

	var products []ImportProduct
	switch format {
	case "json":
		if err := json.NewDecoder(bodyReader).Decode(&products); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
	case "csv":
		reader := csv.NewReader(bodyReader)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, `{"error":"invalid csv"}`, http.StatusBadRequest)
			return
		}
		if len(records) < 2 {
			http.Error(w, `{"error":"empty csv"}`, http.StatusBadRequest)
			return
		}
		for i := 1; i < len(records); i++ {
			row := records[i]
			if len(row) < 2 {
				continue
			}
			p := ImportProduct{Name: row[0]}
			if len(row) > 1 {
				p.SKU = row[1]
			}
			if len(row) > 2 && row[2] != "" {
				if id, err := strconv.ParseUint(row[2], 10, 32); err == nil {
					uid := uint(id)
					p.CategoryID = &uid
				}
			}
			if len(row) > 3 {
				p.Description = row[3]
			}
			if len(row) > 4 {
				p.ImagePath = row[4]
			}
			if len(row) > 5 && row[5] != "" {
				var attrs []ProductAttrIn
				if json.Unmarshal([]byte(row[5]), &attrs) == nil {
					p.Attributes = attrs
				}
			}
			products = append(products, p)
		}
	case "mc":
		if h.catRepo == nil {
			http.Error(w, `{"error":"mc import requires category repo"}`, http.StatusInternalServerError)
			return
		}
		reader := csv.NewReader(bodyReader)
		reader.Comma = ';'
		reader.LazyQuotes = true
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, `{"error":"invalid mc csv: `+err.Error()+`"}`, http.StatusBadRequest)
			return
		}
		if len(records) < 2 {
			http.Error(w, `{"error":"empty csv"}`, http.StatusBadRequest)
			return
		}
		created, skipped := h.importMC(records)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"created": created, "skipped": skipped})
		return
	default:
		http.Error(w, `{"error":"unsupported format"}`, http.StatusBadRequest)
		return
	}

	created := 0
	for _, ip := range products {
		if strings.TrimSpace(ip.Name) == "" {
			continue
		}
		p := models.Product{
			Name:        ip.Name,
			SKU:         ip.SKU,
			CategoryID:  ip.CategoryID,
			Description: ip.Description,
			ImagePath:   ip.ImagePath,
		}
		for _, a := range ip.Attributes {
			p.Attributes = append(p.Attributes, models.ProductAttribute{Key: a.Key, Value: a.Value})
		}
		if err := h.repo.Create(&p); err == nil {
			created++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"created": created})
}

func (h *ProductIOHandler) importMC(records [][]string) (created, skipped int) {
	// Skip header row
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) <= mcColName {
			continue
		}
		name := strings.TrimSpace(at(row, mcColName))
		if name == "" {
			continue
		}
		sku := strings.TrimSpace(at(row, mcColSKU))
		if sku == "" {
			sku = strconv.Itoa(i)
		}
		// Skip if SKU already exists
		var existing models.Product
		if err := h.db.Where("sku = ?", sku).First(&existing).Error; err == nil {
			skipped++
			continue
		}
		cat, err := h.catRepo.FindOrCreateByPath(at(row, mcColCategory), at(row, mcColSubcat))
		if err != nil {
			continue
		}
		catID := cat.ID
		stock, _ := strconv.Atoi(strings.TrimSpace(at(row, mcColStock)))
		price1, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, mcColPrice1)), ",", ".", 1), 64)
		price5, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, mcColPrice5)), ",", ".", 1), 64)
		price10, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, mcColPrice10)), ",", ".", 1), 64)
		imgRaw := strings.TrimSpace(at(row, mcColImages))
		imgPath := ""
		if idx := strings.Index(imgRaw, ","); idx > 0 {
			imgPath = strings.TrimSpace(imgRaw[:idx])
		} else if imgRaw != "" {
			imgPath = imgRaw
		}
		chars := strings.TrimSpace(at(row, mcColChars))
		var attrs []models.ProductAttribute
		if chars != "" {
			attrs = []models.ProductAttribute{{Key: "characteristics", Value: chars}}
		}
		p := models.Product{
			Name:       name,
			SKU:        sku,
			CategoryID: &catID,
			Description: strings.TrimSpace(at(row, mcColDesc)),
			ImagePath:  imgPath,
			Size:       strings.TrimSpace(at(row, mcColSize)),
			Mark:       strings.TrimSpace(at(row, mcColMark)),
			Length:     strings.TrimSpace(at(row, mcColLength)),
			City:       strings.TrimSpace(at(row, mcColCity)),
			Stock:      stock,
			Price1t:    price1,
			Price5t:    price5,
			Price10t:   price10,
			SourceURL:  strings.TrimSpace(at(row, mcColURL)),
			Attributes: attrs,
		}
		if err := h.repo.Create(&p); err == nil {
			created++
		}
	}
	return created, skipped
}

func at(row []string, i int) string {
	if i < len(row) {
		return row[i]
	}
	return ""
}
