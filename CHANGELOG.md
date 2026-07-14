# Changelog

Semua perubahan penting pada project ini akan didokumentasikan di file ini.

Format mengikuti [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
dan project ini mengikuti [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024

### Added
- Initial release dari Mental Health ID Go Edition
- Sistem autentikasi user (login, register, profil, ubah password)
- Admin dashboard dengan CRUD untuk:
  - Gejala (symptoms)
  - Penyakit (diseases)
  - Rule diagnosis
  - Riwayat diagnosa
  - Manajemen user
- Form diagnosis rule-based dengan algoritma scoring
- Session management dengan Gorilla Sessions
- Bootstrap 5 responsive UI
- MySQL database dengan GORM ORM
- Password hashing dengan bcrypt
- Environment-based configuration

### Features
- User authentication with secure password hashing
- Role-based access control (admin/user)
- Interactive diagnosis form
- Rule-based scoring algorithm
- User history tracking
- Admin management panel
- Responsive design
- Session management

### Technical
- Built with Go 1.26.4
- MySQL 5.7+ support
- RESTful API design
- Server-side rendering dengan html/template
- Gorilla framework (mux, sessions)
- GORM ORM integration

---

## Versioning

Format versi: `MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

---

## Upcoming

- [ ] API endpoint documentation
- [ ] Unit tests
- [ ] CI/CD pipeline
- [ ] Docker support
- [ ] Mobile app integration
- [ ] Enhanced diagnosis algorithm
- [ ] Multi-language support
