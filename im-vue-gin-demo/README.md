# Simple IM Project (Gin + Vue 3)

A simple Instant Messaging application built with Go (Gin) and Vue 3.

## Tech Stack
- **Backend:** Go, Gin, WebSocket, GORM (MySQL), go-redis (Redis)
- **Frontend:** Vue 3, Vite
- **Database:** MySQL
- **Cache:** Redis

## Setup

### Backend
```bash
cd backend
go mod tidy
go run main.go
```
Listens on `:8080`.

### Frontend
```bash
cd frontend
npm install
npm run dev
```
Listens on `:5173` (or similar).

## Configuration
Update `backend/main.go` with your MySQL and Redis credentials if needed.
Current defaults:
- MySQL: `root:wch123456.@tcp(8.129.109.155:3306)/im_db`
- Redis: `8.129.109.155:6379`
