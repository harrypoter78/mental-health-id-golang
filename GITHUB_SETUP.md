# 🚀 Setup Guide untuk Push ke GitHub

Berikut panduan lengkap untuk push project ke GitHub:

## 1. Create Repository di GitHub

1. Buka https://github.com/new
2. **Repository name**: `mental-health-id-golang`
3. **Description**: "Sistem pakar diagnosa gangguan mental berbasis rule-based menggunakan Go"
4. **Visibility**: Public (agar bisa dilihat orang lain)
5. **Initialize this repository with**: Kosongkan (jangan centang)
6. Klik **Create repository**

## 2. Push ke GitHub

Setelah repository dibuat, jalankan command berikut:

```bash
# Ganti YOUR_USERNAME dengan username GitHub Anda
git remote add origin https://github.com/YOUR_USERNAME/mental-health-id-golang.git

# Set main branch
git branch -M main

# Push ke GitHub
git push -u origin main
```

## 3. Verify Repository

Setelah push berhasil:
1. Buka https://github.com/YOUR_USERNAME/mental-health-id-golang
2. Pastikan semua file terlihat:
   - ✅ README.md (dengan dokumentasi lengkap)
   - ✅ LICENSE (MIT License)
   - ✅ CONTRIBUTING.md (panduan kontribusi)
   - ✅ CHANGELOG.md (versi history)
   - ✅ .gitignore (untuk exclude file tidak penting)
   - ✅ go.mod, go.sum (dependencies)
   - ✅ .env.example (template config)
   - ✅ cmd/, internal/, web/ (source code)
3. Pastikan `.env` dan `mental-health-id.exe` TIDAK ada (sudah di-gitignore)

## ✅ Verifikasi File yang Terpublikasi

### File PENTING yang dipublikasikan:
- ✅ Source code (cmd/, internal/)
- ✅ Templates (web/templates/)
- ✅ go.mod, go.sum (untuk dependencies)
- ✅ .env.example (template config)
- ✅ README.md (dokumentasi)
- ✅ LICENSE (lisensi project)
- ✅ CONTRIBUTING.md (panduan kontribusi)
- ✅ CHANGELOG.md (version history)

### File yang TIDAK dipublikasikan (.gitignore):
- ❌ .env (file konfigurasi dengan sensitive data)
- ❌ *.exe (executable files)
- ❌ *.log (log files)
- ❌ .vscode/ (IDE settings)
- ❌ vendor/ (jika ada)
- ❌ build/ (compiled files)

## 🔒 Security Checklist

Sebelum push ke GitHub, pastikan:

- [ ] `.env` sudah di-gitignore ✅
- [ ] `.env.example` berisi placeholder, bukan real credentials ✅
- [ ] Tidak ada password atau API keys di-commit
- [ ] Tidak ada file besar (> 100MB)
- [ ] README.md lengkap dan jelas ✅
- [ ] LICENSE file ada ✅
- [ ] Source code sudah di-review

## 📝 Repository Description & Topics

Anda juga bisa tambahkan di GitHub Settings > About:

**Description:**
```
Sistem pakar diagnosa gangguan mental berbasis rule-based dengan algoritma scoring, 
dibangun dengan Go dan MySQL.
```

**Topics:**
- golang
- go
- diagnosis
- mental-health
- rule-based-system
- webapp
- mysql
- bootstrap

## 🤝 Kolaborasi

Untuk invite collaborators:
1. Settings > Collaborators
2. Add people dengan username GitHub mereka
3. Pilih permission level (write untuk developer, maintain untuk admin)

## 📊 Repository Stats

Setelah push, GitHub akan menampilkan:
- Commit history
- Contributors
- Language stats (Go %)
- Network graph

---

**Project Anda siap untuk dipublikasikan! 🎉**
