
# 📦 Proyek REST API Golang – GO-TEST

RESTful API sederhana menggunakan **Golang**, **Echo**, **GORM**, **MySQL**, dan **JWT**. Proyek ini mendukung autentikasi, manajemen user, produk, dan kategori lengkap dengan fitur pagination dan pencarian.

---

## 📁 Struktur Folder Proyek

```
GO-TEST/
├── config/         # Konfigurasi database
├── controllers/    # Handler untuk setiap endpoint API
├── logs/           # File log
├── middleware/     # Middleware untuk JWT & logging
├── models/         # Model untuk GORM
├── routes/         # Routing semua endpoint
├── utils/          # Fungsi utilitas (contoh: hash password)
├── .env            # Variabel environment
├── main.go         # Entry point aplikasi
├── go.mod          # File modul Go
└── README.md       # Dokumentasi proyek
```

---

## 🚀 Cara Menjalankan Aplikasi

### 1. Clone repository
```bash
git clone https://github.com/username/go-test.git
cd go-test
```

### 2. Install dependency
```bash
go mod tidy
```

### 3. Konfigurasi file `.env`
```env
DB_USER=root
DB_PASS=passwordmu
DB_NAME=test_db
DB_HOST=127.0.0.1
DB_PORT=3306
JWT_SECRET=rahasia_super
```

### 4. Jalankan aplikasi
```bash
go run main.go
```

---

## ✅ Fitur

- 🔐 Autentikasi JWT (Login & Register)
- 👤 CRUD untuk User
- 📦 CRUD untuk Produk (dengan kategori)
- 🗂️ CRUD untuk Kategori
- 🔍 Pencarian produk berdasarkan nama/kategori
- 📄 Pagination
- 📋 GORM Hooks
- 📝 Logging ke file
- 🧼 Struktur folder clean code

---

## 🔐 Autentikasi

### Register
```http
POST /register
```
**Body:**
```json
{
  "name": "Nama Kamu",
  "email": "email@example.com",
  "password": "123456"
}
```

### Login
```http
POST /login
```
**Body:**
```json
{
  "email": "email@example.com",
  "password": "123456"
}
```

**Response:**
```json
{
  "token": "token_jwt_kamu"
}
```

> Gunakan token di header Authorization:
```
Authorization: Bearer token_jwt_kamu
```

---

## 👥 Endpoint User (`/api/users`)
| Method | Endpoint         | Keterangan         |
|--------|------------------|--------------------|
| POST   | `/api/users`     | Tambah user        |
| GET    | `/api/users/:id` | Ambil user by ID   |
| PUT    | `/api/users/:id` | Update user        |
| DELETE | `/api/users/:id` | Hapus user         |

---

## 📦 Endpoint Produk (`/api/products`)
| Method | Endpoint             | Keterangan                     |
|--------|----------------------|--------------------------------|
| POST   | `/api/products`      | Tambah produk                  |
| GET    | `/api/products/:id`  | Ambil produk by ID             |
| PUT    | `/api/products/:id`  | Update produk                  |
| DELETE | `/api/products/:id`  | Hapus produk                   |
| GET    | `/api/products`      | Ambil semua produk (dengan filter dan pagination) |

**Query Parameters (opsional):**
- `page`: halaman keberapa
- `limit`: jumlah data per halaman
- `name`: filter berdasarkan nama produk
- `category`: filter berdasarkan nama kategori

---

## 🗂️ Endpoint Kategori (`/api/categories`)
| Method | Endpoint               | Keterangan             |
|--------|------------------------|------------------------|
| POST   | `/api/categories`      | Tambah kategori        |
| GET    | `/api/categories`      | Ambil semua kategori   |
| GET    | `/api/categories/:id`  | Ambil kategori by ID   |
| PUT    | `/api/categories/:id`  | Update kategori        |
| DELETE | `/api/categories/:id`  | Hapus kategori         |

---

## 🧪 Testing API

Gunakan tools seperti:
- [Postman](https://www.postman.com/)
- [curl](https://curl.se/)
- Swagger (dalam pengembangan)

---

## 📝 Log File

Semua log disimpan otomatis ke dalam:
```
logs/server.log
```

---

## 👨‍💻 Developer

**@dim.devs** – Influencer Coding (TikTok & Instagram)  
Website: [https://dimdevs.dev](https://dimdevs.dev)
