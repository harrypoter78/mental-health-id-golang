# 🤝 Contributing to Mental Health ID

Terima kasih telah tertarik untuk berkontribusi pada Mental Health ID! Panduan ini menjelaskan cara berkontribusi pada project ini.

## 📋 Aturan Kontribusi

### Sebelum Mulai

1. Pastikan Anda memiliki akun GitHub
2. Fork repository ini
3. Clone fork Anda: `git clone https://github.com/yourusername/mental-health-id-golang.git`
4. Buat branch untuk fitur Anda: `git checkout -b feature/amazing-feature`

### Proses Kontribusi

1. **Buat Branch Feature**
   ```bash
   git checkout -b feature/your-feature-name
   # Gunakan prefix: feature/ untuk fitur baru, bugfix/ untuk bug fixes, docs/ untuk dokumentasi
   ```

2. **Buat Perubahan**
   - Ikuti konvensi code yang ada
   - Pastikan code terformat dengan `go fmt ./...`
   - Tambahkan comments untuk logika kompleks
   - Update dokumentasi jika perlu

3. **Testing**
   ```bash
   go test ./...
   go vet ./...
   ```

4. **Commit Perubahan**
   ```bash
   git commit -m "feat: add amazing feature"
   # Gunakan conventional commits: feat:, fix:, docs:, style:, refactor:, test:, chore:
   ```

5. **Push ke Fork**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Buat Pull Request**
   - Deskripsi perubahan yang jelas
   - Link ke related issues jika ada
   - Pastikan semua checks pass

## 🐛 Melaporkan Bugs

Jika Anda menemukan bug, silakan buat issue dengan informasi:

- **Deskripsi**: Jelaskan bug yang Anda temukan
- **Reproduksi**: Langkah-langkah untuk mereproduksi
- **Expected**: Apa yang seharusnya terjadi
- **Actual**: Apa yang benar-benar terjadi
- **Environment**: Go version, OS, MySQL version, dll

## 💡 Saran Fitur

Untuk saran fitur baru:

1. Buka issue dengan label `enhancement`
2. Jelaskan use case dan manfaat
3. Diskusikan implementasi yang mungkin

## 📝 Commit Message Convention

Gunakan Conventional Commits format:

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: Fitur baru
- `fix`: Bug fix
- `docs`: Perubahan dokumentasi
- `style`: Formatting, missing semicolons, etc
- `refactor`: Refactoring code
- `test`: Adding tests
- `chore`: Build process, dependencies

**Contoh:**
```
feat(auth): add password reset functionality

Implement email-based password reset flow with token validation.
Adds new endpoint POST /api/reset-password.

Closes #123
```

## 🎨 Code Style

- Ikuti Go idioms dan conventions
- Run `go fmt ./...` sebelum commit
- Gunakan `go vet ./...` untuk deteksi issues
- Tambahkan error handling yang proper
- Tulis comments untuk exported functions

## 🧪 Testing Requirements

- Unit tests untuk fitur baru
- Maintain atau tingkatkan code coverage
- Test edge cases dan error scenarios

## 📚 Documentation

- Update README.md jika ada perubahan API
- Dokumentasikan environment variables baru
- Tambahkan comments untuk logika kompleks
- Update CHANGELOG.md jika ada

## 🔄 PR Review Process

1. Maintainers akan review PR Anda
2. Mungkin ada feedback atau requested changes
3. Diskusikan dan implementasikan feedback
4. Setelah approval, PR akan di-merge

## ❓ Questions?

- Buka GitHub Discussions untuk pertanyaan umum
- Cek existing issues/PRs sebelum membuat yang baru
- Tanyakan di issue sebelum mulai kerja besar

## 📄 Lisensi

Dengan berkontribusi, Anda setuju bahwa kontribusi Anda akan dilisensikan di bawah MIT License yang sama dengan project ini.

---

**Happy Contributing! 🎉**
