# УралМетал — Сайт по продаже металла

Монорепозиторий: клиентский сайт, админ-панель и Go backend для каталога металлопродукции с формой обратной связи.

## Стек

- **Фронтенд**: Vue 3 + Vite + Tailwind CSS (client + admin)
- **Бэкенд**: Go + Chi + GORM
- **БД**: SQLite (dev) / PostgreSQL (prod)
- **Инфраструктура**: Docker Compose, Nginx

## Быстрый старт (разработка)

1. Скопируйте `.env.example` в `.env`.
2. Запустите через Docker:
   ```bash
   docker compose up -d
   ```
3. Откройте:
   - **Сайт**: http://localhost:3000
   - **Админка**: http://localhost:3000/admin
   - **API**: http://localhost:3000/api

   Backend, client и admin работают в контейнерах. Nginx проксирует запросы.

   Без Docker (если нужна прямая разработка фронтенда):
   ```bash
   cd backend && go run ./cmd/server &
   cd client && npm install && npm run dev
   cd admin && npm install && npm run dev
   ```
   Клиент: http://localhost:5173, Админка: http://localhost:5174/admin

## Структура

```
├── client/     # Vue 3 — публичный сайт
├── admin/      # Vue 3 — админ-панель
├── backend/    # Go — API
├── nginx/      # Конфигурация Nginx
├── storage/    # Загруженные изображения
```

## Продакшн

1. Соберите фронтенды:
   ```bash
   cd client && npm run build
   cd admin && npm run build
   ```
2. Скопируйте `.env.prod.example` в `.env.prod` и задайте пароли.
3. Запустите:
   ```bash
   docker compose -f docker-compose.prod.yml up -d
   ```
   Сайт: http://localhost  
   Админка: http://localhost/admin
