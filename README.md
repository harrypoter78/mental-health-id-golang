# Mental Health ID - Go Edition

[![Go Version](https://img.shields.io/badge/go-1.26.4-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Sistem pakar diagnosa gangguan mental berbasis rule-based yang diimplementasikan dengan Go. Merupakan replikasi fungsionalitas dari `mental-health-id` berbasis Laravel dengan performa yang lebih baik.

## 🎯 Fitur Utama

- **Autentikasi User**: Sistem login, register, profil, dan ubah password yang aman
- **Admin Dashboard**: CRUD untuk gejala, penyakit, rule, riwayat diagnosa, dan manajemen user
- **Diagnosis Rule-Based**: Form diagnosis interaktif dengan algoritma scoring kecocokan
- **Responsive UI**: Template Bootstrap 5 yang modern dan mobile-friendly
- **Session Management**: Manajemen sesi user yang aman dengan Gorilla Sessions
- **Database MySQL**: Struktur database yang terukur dan efisien

## 📋 Prasyarat

- **Go**: 1.26.4 atau lebih baru ([Download Go](https://golang.org/dl/))
- **MySQL**: 5.7 atau lebih baru
- **Git**: Untuk cloning repository

## 🚀 Quick Start

### 1. Clone Repository
```bash
git clone https://github.com/yourusername/mental-health-id-golang.git
cd mental-health-id-golang
```

### 2. Setup Database
```bash
# Buat database MySQL
mysql -u root -p -e "CREATE DATABASE gangguanmental;"

# Import schema (jika tersedia)
mysql -u root -p gangguanmental < schema.sql
```

### 3. Konfigurasi Environment
```bash
# Salin file contoh konfigurasi
cp .env.example .env

# Edit .env dengan konfigurasi Anda
# Minimal perlu update:
# - DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME
# - SESSION_SECRET (gunakan string random yang kuat)
```

### 4. Install Dependencies
```bash
go mod download
go mod tidy
```

### 5. Jalankan Aplikasi
```bash
# Menggunakan Go run
go run ./cmd/mental-health-id

# Atau build executable terlebih dahulu
go build -o mental-health-id ./cmd/mental-health-id
./mental-health-id
```

### 6. Akses Aplikasi
Buka browser dan navigasi ke: `http://localhost:8080`

## 📁 Struktur Project

```
.
├── cmd/
│   └── mental-health-id/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/                  # Konfigurasi environment & database
│   ├── database/                # Koneksi dan inisialisasi database
│   ├── models/                  # Model GORM untuk entitas data
│   ├── handlers/                # HTTP handlers dan business logic
│   └── middleware/              # Middleware autentikasi & logging
├── web/
│   ├── templates/               # Template HTML (html/template)
│   └── static/                  # Assets (CSS, JS, images)
├── .env.example                 # Template file konfigurasi
├── go.mod & go.sum              # Go module dependencies
├── README.md                    # Dokumentasi ini
└── LICENSE                      # Lisensi project
```

## 🔧 Konfigurasi Environment

File `.env.example` berisi template konfigurasi:

```env
# Database MySQL
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=gangguanmental

# Session Secret (gunakan string random yang kuat!)
SESSION_SECRET=replace-with-a-secure-random-string
```

**Penting**: Jangan commit file `.env` ke repository. Gunakan `.env.example` sebagai template.

## 📚 Dependencies

Proyek ini menggunakan beberapa package Go penting:

- **gorilla/mux**: HTTP router yang powerful
- **gorilla/sessions**: Session management
- **gorm**: ORM untuk database
- **gorm/driver/mysql**: Driver MySQL untuk GORM
- **golang.org/x/crypto**: Package cryptography untuk password hashing
- **joho/godotenv**: Loading environment variables dari .env

Lihat `go.mod` untuk daftar lengkap dependencies.

## 📝 Development

### Running Tests (jika tersedia)
```bash
go test ./...
```

### Formatting Code
```bash
go fmt ./...
```

### Linting (menggunakan golangci-lint)
```bash
golangci-lint run ./...
```

## 🤝 Kontribusi

Kontribusi sangat diterima! Silakan:

1. Fork repository ini
2. Buat branch fitur (`git checkout -b feature/AmazingFeature`)
3. Commit perubahan (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buka Pull Request

## 📄 Lisensi

Project ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk detail.

## 📧 Kontak & Support

Jika ada pertanyaan atau issues, silakan buka [Issues](https://github.com/yourusername/mental-health-id-golang/issues) di GitHub.

## 📖 Referensi

- [Go Official Documentation](https://golang.org/doc/)
- [Gorilla Web Toolkit](https://www.gorillatoolkit.org/)
- [GORM Documentation](https://gorm.io/)
- [Mental Health ID Original (Laravel)](https://github.com/yourrepo/mental-health-id)

---

**Last Updated**: 2026
**Version**: 1.0.0
