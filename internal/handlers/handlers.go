package handlers

import (
    "bytes"
    "context"
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"

    "github.com/example/mental-health-id/internal/models"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type contextKey string

const sessionName = "mhid_session"
const userContextKey = contextKey("currentUser")

type App struct {
    DB           *gorm.DB
    TemplateDir  string
    FuncMap      template.FuncMap
    SessionStore *sessions.CookieStore
}

type TemplateData struct {
    Title        string
    CurrentUser  *models.User
    Flash        string
    FlashError   string
    FormMode     string
    Form         map[string]string
    Errors       map[string]string
    Data         map[string]interface{}
    CurrentPath  string
    Styles       template.HTML
    Scripts      template.HTML
}

type Pagination struct {
    Page       int
    PageSize   int
    TotalRows  int64
    TotalPages int
    Pages      []int
    StartIndex int
}

type DiagnosisResult struct {
    Penyakit   *models.Penyakit
    Persentase float64
    Cocok      int
    Total      int
}

func NewApp(db *gorm.DB, store *sessions.CookieStore) (*App, error) {
    templateDir, err := findTemplateDir()
    if err != nil {
        return nil, err
    }

    funcMap := template.FuncMap{
        "eq": func(a, b interface{}) bool { return a == b },
        "ne": func(a, b interface{}) bool { return a != b },
        "add": func(a, b int) int { return a + b },
        "sub": func(a, b int) int { return a - b },
        "hasPrefix": func(s, prefix string) bool { return strings.HasPrefix(s, prefix) },
        "contains": func(slice interface{}, value string) bool {
            if slice == nil {
                return false
            }
            switch v := slice.(type) {
            case []string:
                for _, item := range v {
                    if item == value {
                        return true
                    }
                }
            case []interface{}:
                for _, item := range v {
                    if str, ok := item.(string); ok && str == value {
                        return true
                    }
                }
            }
            return false
        },
    }

    return &App{DB: db, TemplateDir: templateDir, FuncMap: funcMap, SessionStore: store}, nil
}

func findTemplateDir() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "web/templates", err
    }

    for {
        candidateDir := filepath.Join(dir, "web", "templates")
        if _, err := os.Stat(candidateDir); err == nil {
            return candidateDir, nil
        }

        parent := filepath.Dir(dir)
        if parent == dir {
            break
        }
        dir = parent
    }

    return "web/templates", nil
}

func (app *App) getSession(r *http.Request) (*sessions.Session, error) {
    return app.SessionStore.Get(r, sessionName)
}

func (app *App) setFlash(w http.ResponseWriter, r *http.Request, message string, errorMessage string) {
    session, _ := app.getSession(r)
    if message != "" {
        session.Values["flash_success"] = message
    }
    if errorMessage != "" {
        session.Values["flash_error"] = errorMessage
    }
    _ = session.Save(r, w)
}

func (app *App) popFlash(w http.ResponseWriter, r *http.Request) (string, string) {
    session, _ := app.getSession(r)
    success, _ := session.Values["flash_success"].(string)
    errorMessage, _ := session.Values["flash_error"].(string)
    delete(session.Values, "flash_success")
    delete(session.Values, "flash_error")
    _ = session.Save(r, w)
    return success, errorMessage
}

func (app *App) currentUserFromSession(r *http.Request) (*models.User, error) {
    session, err := app.getSession(r)
    if err != nil {
        return nil, err
    }
    rawID, ok := session.Values["user_id"]
    if !ok || rawID == nil {
        return nil, nil
    }

    var userID uint
    switch v := rawID.(type) {
    case uint:
        userID = v
    case int:
        if v >= 0 {
            userID = uint(v)
        }
    case int64:
        if v >= 0 {
            userID = uint(v)
        }
    case float64:
        if v >= 0 {
            userID = uint(v)
        }
    case string:
        if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
            userID = uint(parsed)
        }
    default:
        return nil, nil
    }

    if userID == 0 {
        return nil, nil
    }
    var user models.User
    if err := app.DB.First(&user, userID).Error; err != nil {
        return nil, nil
    }
    return &user, nil
}

func (app *App) LoadUserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, _ := app.currentUserFromSession(r)
        ctx := context.WithValue(r.Context(), userContextKey, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func (app *App) GetCurrentUser(r *http.Request) *models.User {
    if user, ok := r.Context().Value(userContextKey).(*models.User); ok {
        return user
    }
    return nil
}

func (app *App) renderTemplate(w http.ResponseWriter, r *http.Request, layout string, pageTemplate string, data *TemplateData) {
    if data == nil {
        data = &TemplateData{}
    }

    data.CurrentUser = app.GetCurrentUser(r)
    data.CurrentPath = r.URL.Path
    data.Flash, data.FlashError = app.popFlash(w, r)
    data.Form = ensureMap(data.Form)
    data.Errors = ensureMap(data.Errors)
    data.Data = ensureMapInterface(data.Data)

    if pageTemplate == "" {
        http.Error(w, "template not specified", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.New("app").Funcs(app.FuncMap).ParseFiles(
        filepath.Join(app.TemplateDir, layout+".html"),
        filepath.Join(app.TemplateDir, pageTemplate+".html"),
    )
    if err != nil {
        http.Error(w, fmt.Sprintf("template parse error: %v", err), http.StatusInternalServerError)
        return
    }

    var buf bytes.Buffer
    if err := tmpl.ExecuteTemplate(&buf, layout, data); err != nil {
        http.Error(w, fmt.Sprintf("template error: %v", err), http.StatusInternalServerError)
        return
    }

    if _, err := w.Write(buf.Bytes()); err != nil {
        http.Error(w, fmt.Sprintf("template write error: %v", err), http.StatusInternalServerError)
    }
}

func ensureMap(m map[string]string) map[string]string {
    if m == nil {
        return make(map[string]string)
    }
    return m
}

func ensureMapInterface(m map[string]interface{}) map[string]interface{} {
    if m == nil {
        return make(map[string]interface{})
    }
    return m
}

func (app *App) MethodOverrideMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            if err := r.ParseForm(); err == nil {
                if method := r.FormValue("_method"); method != "" {
                    r.Method = strings.ToUpper(method)
                }
            }
        }
        next.ServeHTTP(w, r)
    })
}

func (app *App) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if app.GetCurrentUser(r) == nil {
            app.setFlash(w, r, "", "Silakan login untuk mengakses halaman ini.")
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next(w, r)
    }
}

func (app *App) RequireAdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := app.GetCurrentUser(r)
        if user == nil || user.Role != "admin" {
            app.setFlash(w, r, "", "Halaman admin hanya dapat diakses oleh administrator.")
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
    data := &TemplateData{Title: "Sistem Diagnosis Kesehatan Mental"}
    app.renderTemplate(w, r, "layout", "diagnosis_index", data)
}

func (app *App) DiagnosisKuisHandler(w http.ResponseWriter, r *http.Request) {
    var gejala []models.Gejala
    app.DB.Order("id_gejala").Find(&gejala)
    data := &TemplateData{
        Title: "Kuis Diagnosis Gejala",
        Data: map[string]interface{}{"Gejala": gejala},
    }
    app.renderTemplate(w, r, "layout", "diagnosis_kuis", data)
}

func (app *App) DiagnosisProsesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        http.Redirect(w, r, "/diagnosis/kuis", http.StatusSeeOther)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid submission", http.StatusBadRequest)
        return
    }

    selected := r.Form["gejala"]
    var gejalaList []models.Gejala
    app.DB.Order("id_gejala").Find(&gejalaList)

    if len(selected) == 0 {
        data := &TemplateData{
            Title: "Kuis Diagnosis Gejala",
            Data: map[string]interface{}{"Gejala": gejalaList, "Selected": selected},
            Errors: map[string]string{"gejala": "Silakan pilih minimal satu gejala."},
        }
        app.renderTemplate(w, r, "layout", "diagnosis_kuis", data)
        return
    }

    var penyakit []models.Penyakit
    app.DB.Find(&penyakit)
    var results []DiagnosisResult

    for _, p := range penyakit {
        var rules []models.Rule
        app.DB.Where("kode_penyakit = ?", p.KodePenyakit).Find(&rules)
        if len(rules) == 0 {
            continue
        }
        total := len(rules)
        cocok := 0
        ruleMap := make(map[string]struct{}, len(rules))
        for _, rule := range rules {
            ruleMap[rule.KodeGejala] = struct{}{}
        }
        for _, kode := range selected {
            if _, ok := ruleMap[kode]; ok {
                cocok++
            }
        }
        persentase := 0.0
        if total > 0 {
            persentase = roundFloat(float64(cocok)/float64(total)*100, 2)
        }
        if persentase > 0 {
            results = append(results, DiagnosisResult{Penyakit: &p, Persentase: persentase, Cocok: cocok, Total: total})
        }
    }

    if len(results) > 1 {
        sortResults(results)
    }

    var best *DiagnosisResult
    if len(results) > 0 {
        best = &results[0]
    }

    if best != nil && app.GetCurrentUser(r) != nil {
        riwayat := models.Riwayat{
            UserID:      app.GetCurrentUser(r).ID,
            NamaPenyakit: best.Penyakit.NamaPenyakit,
            Tanggal:     time.Now(),
        }
        app.DB.Create(&riwayat)
    }

    nama := "Tamu"
    if user := app.GetCurrentUser(r); user != nil {
        nama = user.Name
    }

    data := &TemplateData{
        Title: "Hasil Diagnosis",
        Data: map[string]interface{}{
            "NamaPasien":       nama,
            "DiagnosisTertinggi": best,
            "HasilDiagnosis":   results,
        },
    }
    app.renderTemplate(w, r, "layout", "diagnosis_hasil", data)
}

func sortResults(results []DiagnosisResult) {
    for i := 0; i < len(results)-1; i++ {
        for j := i + 1; j < len(results); j++ {
            if results[j].Persentase > results[i].Persentase {
                results[i], results[j] = results[j], results[i]
            }
        }
    }
}

func roundFloat(value float64, precision int) float64 {
    format := fmt.Sprintf("%%.%df", precision)
    s := fmt.Sprintf(format, value)
    f, _ := strconv.ParseFloat(s, 64)
    return f
}

func (app *App) DiagnosisRiwayatHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var riwayat []models.Riwayat
    var totalRows int64
    query := app.DB.Order("tanggal desc").Preload("User")
    user := app.GetCurrentUser(r)
    if user != nil {
        query = query.Where("user_id = ?", user.ID)
    }
    query.Model(&models.Riwayat{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&riwayat)

    data := &TemplateData{
        Title: "Riwayat Diagnosis",
        Data: map[string]interface{}{"Riwayat": riwayat, "Pagination": pagination},
    }
    app.renderTemplate(w, r, "layout", "diagnosis_riwayat", data)
}

func parsePage(r *http.Request) (int, int) {
    page := 1
    pageSize := 10
    if p := r.URL.Query().Get("page"); p != "" {
        if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
            page = parsed
        }
    }
    return page, pageSize
}

func makePagination(page, pageSize int, totalRows int64) Pagination {
    totalPages := int((totalRows + int64(pageSize) - 1) / int64(pageSize))
    if totalPages == 0 {
        totalPages = 1
    }
    pages := make([]int, 0, totalPages)
    for i := 1; i <= totalPages; i++ {
        pages = append(pages, i)
    }
    startIndex := (page-1)*pageSize + 1
    if startIndex < 1 {
        startIndex = 1
    }
    return Pagination{Page: page, PageSize: pageSize, TotalRows: totalRows, TotalPages: totalPages, Pages: pages, StartIndex: startIndex}
}

func (app *App) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
    if app.GetCurrentUser(r) != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    data := &TemplateData{Title: "Login", FormMode: "login", Data: map[string]interface{}{}}
    app.renderTemplate(w, r, "layout", "auth_login", data)
}

func (app *App) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
    if app.GetCurrentUser(r) != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    data := &TemplateData{Title: "Daftar", FormMode: "register", Data: map[string]interface{}{}}
    app.renderTemplate(w, r, "layout", "auth_login", data)
}

func (app *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form submission", http.StatusBadRequest)
        return
    }

    email := strings.TrimSpace(r.FormValue("email"))
    password := r.FormValue("password")
    errors := make(map[string]string)
    if email == "" {
        errors["email"] = "Email wajib diisi"
    }
    if password == "" {
        errors["password"] = "Password wajib diisi"
    }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Login", FormMode: "login", Form: map[string]string{"email": email}, Errors: errors}
        app.renderTemplate(w, r, "layout", "auth_login", data)
        return
    }

    var user models.User
    err := app.DB.Where("email = ?", email).First(&user).Error
    if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        data := &TemplateData{Title: "Login", FormMode: "login", Form: map[string]string{"email": email}, Errors: map[string]string{"email": "Email atau password tidak sesuai."}}
        app.renderTemplate(w, r, "layout", "auth_login", data)
        return
    }

    session, _ := app.getSession(r)
    session.Values["user_id"] = user.ID
    _ = session.Save(r, w)

    if user.Role == "admin" {
        http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
        return
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form submission", http.StatusBadRequest)
        return
    }
    name := strings.TrimSpace(r.FormValue("name"))
    email := strings.TrimSpace(r.FormValue("email"))
    password := r.FormValue("password")
    passwordConfirmation := r.FormValue("password_confirmation")

    errors := make(map[string]string)
    if name == "" {
        errors["name"] = "Nama wajib diisi"
    }
    if email == "" {
        errors["email"] = "Email wajib diisi"
    }
    if password == "" {
        errors["password"] = "Password wajib diisi"
    }
    if password != passwordConfirmation {
        errors["password_confirmation"] = "Konfirmasi password tidak sesuai"
    }
    if len(password) > 0 && len(password) < 6 {
        errors["password"] = "Password minimal 6 karakter"
    }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Daftar", FormMode: "register", Form: map[string]string{"name": name, "email": email}, Errors: errors}
        app.renderTemplate(w, r, "layout", "auth_login", data)
        return
    }

    var existing models.User
    if err := app.DB.Where("email = ?", email).First(&existing).Error; err == nil {
        errors["email"] = "Email sudah terdaftar"
        data := &TemplateData{Title: "Daftar", FormMode: "register", Form: map[string]string{"name": name, "email": email}, Errors: errors}
        app.renderTemplate(w, r, "layout", "auth_login", data)
        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Tidak dapat membuat pengguna baru", http.StatusInternalServerError)
        return
    }
    user := models.User{Name: name, Email: email, Role: "user", Password: string(hashed), CreatedAt: ptrTime(time.Now()), UpdatedAt: ptrTime(time.Now())}
    if err := app.DB.Create(&user).Error; err != nil {
        http.Error(w, "Tidak dapat membuat pengguna baru", http.StatusInternalServerError)
        return
    }

    session, _ := app.getSession(r)
    session.Values["user_id"] = user.ID
    _ = session.Save(r, w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ptrTime(t time.Time) *time.Time {
    return &t
}

func (app *App) LogoutHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := app.getSession(r)
    session.Options.MaxAge = -1
    _ = session.Save(r, w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *App) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
    user := app.GetCurrentUser(r)
    if user == nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    data := &TemplateData{Title: "Edit Profil", Form: map[string]string{"name": user.Name, "email": user.Email}}
    app.renderTemplate(w, r, "layout", "user_profile", data)
}

func (app *App) ProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {
    user := app.GetCurrentUser(r)
    if user == nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    name := strings.TrimSpace(r.FormValue("name"))
    email := strings.TrimSpace(r.FormValue("email"))
    errors := make(map[string]string)
    if name == "" {
        errors["name"] = "Nama wajib diisi"
    }
    if email == "" {
        errors["email"] = "Email wajib diisi"
    }
    if email != "" {
        var existing models.User
        if err := app.DB.Where("email = ? AND id <> ?", email, user.ID).First(&existing).Error; err == nil {
            errors["email"] = "Email sudah digunakan oleh akun lain"
        }
    }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Edit Profil", Form: map[string]string{"name": name, "email": email}, Errors: errors}
        app.renderTemplate(w, r, "layout", "user_profile", data)
        return
    }

    user.Name = name
    user.Email = email
    user.UpdatedAt = ptrTime(time.Now())
    app.DB.Save(user)
    app.setFlash(w, r, "Profil berhasil diperbarui", "")
    http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func (app *App) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
    user := app.GetCurrentUser(r)
    if user == nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    current := r.FormValue("current_password")
    newPassword := r.FormValue("new_password")
    confirmation := r.FormValue("new_password_confirmation")
    errors := make(map[string]string)
    if current == "" {
        errors["current_password"] = "Password saat ini wajib diisi"
    }
    if newPassword == "" {
        errors["new_password"] = "Password baru wajib diisi"
    }
    if newPassword != confirmation {
        errors["new_password_confirmation"] = "Konfirmasi password tidak sesuai"
    }
    if len(newPassword) > 0 && len(newPassword) < 6 {
        errors["new_password"] = "Password minimal 6 karakter"
    }
    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(current)) != nil {
        errors["current_password"] = "Password saat ini tidak sesuai"
    }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Edit Profil", Form: map[string]string{"name": user.Name, "email": user.Email}, Errors: errors}
        app.renderTemplate(w, r, "layout", "user_profile", data)
        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Tidak dapat memperbarui password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashed)
    user.UpdatedAt = ptrTime(time.Now())
    app.DB.Save(user)
    app.setFlash(w, r, "Password berhasil diubah", "")
    http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func (app *App) AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
    var totalGejala, totalPenyakit, totalRule, totalRiwayat, totalUser int64
    app.DB.Model(&models.Gejala{}).Count(&totalGejala)
    app.DB.Model(&models.Penyakit{}).Count(&totalPenyakit)
    app.DB.Model(&models.Rule{}).Count(&totalRule)
    app.DB.Model(&models.Riwayat{}).Count(&totalRiwayat)
    app.DB.Model(&models.User{}).Count(&totalUser)
    data := &TemplateData{Title: "Admin Dashboard", Data: map[string]interface{}{ "TotalGejala": totalGejala, "TotalPenyakit": totalPenyakit, "TotalRule": totalRule, "TotalRiwayat": totalRiwayat, "TotalUser": totalUser }}
    app.renderTemplate(w, r, "admin_layout", "admin_dashboard", data)
}

func (app *App) AdminGejalaIndexHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var totalRows int64
    var items []models.Gejala
    app.DB.Model(&models.Gejala{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    app.DB.Order("id_gejala").Offset((page-1)*pageSize).Limit(pageSize).Find(&items)
    data := &TemplateData{Title: "Kelola Gejala", Data: map[string]interface{}{"Gejala": items, "Pagination": pagination}}
    app.renderTemplate(w, r, "admin_layout", "admin_gejala_index", data)
}

func (app *App) AdminGejalaCreateHandler(w http.ResponseWriter, r *http.Request) {
    data := &TemplateData{Title: "Tambah Gejala", Data: map[string]interface{}{"FormTitle": "Tambah Gejala", "Action": "/admin/gejala", "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
}

func (app *App) AdminGejalaStoreHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    kode := strings.TrimSpace(r.FormValue("kode_gejala"))
    nama := strings.TrimSpace(r.FormValue("nama_gejala"))
    errors := make(map[string]string)
    if kode == "" {
        errors["kode_gejala"] = "Kode gejala wajib diisi"
    }
    if nama == "" {
        errors["nama_gejala"] = "Nama gejala wajib diisi"
    }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Tambah Gejala", Form: map[string]string{"kode_gejala": kode, "nama_gejala": nama}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Gejala", "Action": "/admin/gejala", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
        return
    }
    var existing models.Gejala
    if err := app.DB.Where("kode_gejala = ?", kode).First(&existing).Error; err == nil {
        errors["kode_gejala"] = "Kode gejala sudah digunakan"
        data := &TemplateData{Title: "Tambah Gejala", Form: map[string]string{"kode_gejala": kode, "nama_gejala": nama}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Gejala", "Action": "/admin/gejala", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
        return
    }
    item := models.Gejala{KodeGejala: kode, NamaGejala: nama}
    app.DB.Create(&item)
    app.setFlash(w, r, "Gejala berhasil ditambahkan", "")
    http.Redirect(w, r, "/admin/gejala", http.StatusSeeOther)
}

func (app *App) AdminGejalaEditHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Gejala
    if err := app.DB.First(&item, id).Error; err != nil {
        http.NotFound(w, r)
        return
    }
    data := &TemplateData{Title: "Edit Gejala", Form: map[string]string{"kode_gejala": item.KodeGejala, "nama_gejala": item.NamaGejala}, Data: map[string]interface{}{"FormTitle": "Edit Gejala", "Action": fmt.Sprintf("/admin/gejala/%d/update", item.ID), "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
}

func (app *App) AdminGejalaUpdateHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Gejala
    if err := app.DB.First(&item, id).Error; err != nil {
        http.NotFound(w, r)
        return
    }
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    kode := strings.TrimSpace(r.FormValue("kode_gejala"))
    nama := strings.TrimSpace(r.FormValue("nama_gejala"))
    errors := make(map[string]string)
    if kode == "" { errors["kode_gejala"] = "Kode gejala wajib diisi" }
    if nama == "" { errors["nama_gejala"] = "Nama gejala wajib diisi" }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Edit Gejala", Form: map[string]string{"kode_gejala": kode, "nama_gejala": nama}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Gejala", "Action": fmt.Sprintf("/admin/gejala/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
        return
    }
    var existing models.Gejala
    if err := app.DB.Where("kode_gejala = ? AND id_gejala <> ?", kode, item.ID).First(&existing).Error; err == nil {
        errors["kode_gejala"] = "Kode gejala sudah digunakan"
                data := &TemplateData{Title: "Edit Gejala", Form: map[string]string{"kode_gejala": kode, "nama_gejala": nama}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Gejala", "Action": fmt.Sprintf("/admin/gejala/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_gejala_form", data)
        return
    }
    item.KodeGejala = kode
    item.NamaGejala = nama
    app.DB.Save(&item)
    app.setFlash(w, r, "Gejala berhasil diperbarui", "")
    http.Redirect(w, r, "/admin/gejala", http.StatusSeeOther)
}

func (app *App) AdminGejalaDeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    app.DB.Delete(&models.Gejala{}, id)
    app.setFlash(w, r, "Gejala berhasil dihapus", "")
    http.Redirect(w, r, "/admin/gejala", http.StatusSeeOther)
}

func (app *App) AdminPenyakitIndexHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var totalRows int64
    var items []models.Penyakit
    app.DB.Model(&models.Penyakit{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    app.DB.Order("id_penyakit").Offset((page-1)*pageSize).Limit(pageSize).Find(&items)
    data := &TemplateData{Title: "Kelola Penyakit", Data: map[string]interface{}{"Penyakit": items, "Pagination": pagination}}
    app.renderTemplate(w, r, "admin_layout", "admin_penyakit_index", data)
}

func (app *App) AdminPenyakitCreateHandler(w http.ResponseWriter, r *http.Request) {
    data := &TemplateData{Title: "Tambah Penyakit", Data: map[string]interface{}{"FormTitle": "Tambah Penyakit", "Action": "/admin/penyakit", "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data)
}

func (app *App) AdminPenyakitStoreHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    kode := strings.TrimSpace(r.FormValue("kode_penyakit"))
    nama := strings.TrimSpace(r.FormValue("nama_penyakit"))
    deskripsi := strings.TrimSpace(r.FormValue("deskripsi"))
    solusiObat := strings.TrimSpace(r.FormValue("solusi_obat"))
    solusiLain := strings.TrimSpace(r.FormValue("solusi_lain"))
    errors := make(map[string]string)
    if kode == "" { errors["kode_penyakit"] = "Kode penyakit wajib diisi" }
    if nama == "" { errors["nama_penyakit"] = "Nama penyakit wajib diisi" }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Tambah Penyakit", Form: map[string]string{"kode_penyakit": kode, "nama_penyakit": nama, "deskripsi": deskripsi, "solusi_obat": solusiObat, "solusi_lain": solusiLain}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Penyakit", "Action": "/admin/penyakit", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data); return
    }
    var existing models.Penyakit
    if err := app.DB.Where("kode_penyakit = ?", kode).First(&existing).Error; err == nil {
        errors["kode_penyakit"] = "Kode penyakit sudah digunakan"
        data := &TemplateData{Title: "Tambah Penyakit", Form: map[string]string{"kode_penyakit": kode, "nama_penyakit": nama, "deskripsi": deskripsi, "solusi_obat": solusiObat, "solusi_lain": solusiLain}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Penyakit", "Action": "/admin/penyakit", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data); return
    }
    item := models.Penyakit{KodePenyakit: kode, NamaPenyakit: nama, Deskripsi: deskripsi, SolusiObat: solusiObat, SolusiLain: solusiLain}
    app.DB.Create(&item)
    app.setFlash(w, r, "Penyakit berhasil ditambahkan", "")
    http.Redirect(w, r, "/admin/penyakit", http.StatusSeeOther)
}

func (app *App) AdminPenyakitEditHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Penyakit
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    data := &TemplateData{Title: "Edit Penyakit", Form: map[string]string{"kode_penyakit": item.KodePenyakit, "nama_penyakit": item.NamaPenyakit, "deskripsi": item.Deskripsi, "solusi_obat": item.SolusiObat, "solusi_lain": item.SolusiLain}, Data: map[string]interface{}{"FormTitle": "Edit Penyakit", "Action": fmt.Sprintf("/admin/penyakit/%d/update", item.ID), "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data)
}

func (app *App) AdminPenyakitUpdateHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Penyakit
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    kode := strings.TrimSpace(r.FormValue("kode_penyakit"))
    nama := strings.TrimSpace(r.FormValue("nama_penyakit"))
    deskripsi := strings.TrimSpace(r.FormValue("deskripsi"))
    solusiObat := strings.TrimSpace(r.FormValue("solusi_obat"))
    solusiLain := strings.TrimSpace(r.FormValue("solusi_lain"))
    errors := make(map[string]string)
    if kode == "" { errors["kode_penyakit"] = "Kode penyakit wajib diisi" }
    if nama == "" { errors["nama_penyakit"] = "Nama penyakit wajib diisi" }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Edit Penyakit", Form: map[string]string{"kode_penyakit": kode, "nama_penyakit": nama, "deskripsi": deskripsi, "solusi_obat": solusiObat, "solusi_lain": solusiLain}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Penyakit", "Action": fmt.Sprintf("/admin/penyakit/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data); return
    }
    var existing models.Penyakit
    if err := app.DB.Where("kode_penyakit = ? AND id_penyakit <> ?", kode, item.ID).First(&existing).Error; err == nil {
        errors["kode_penyakit"] = "Kode penyakit sudah digunakan"
        data := &TemplateData{Title: "Edit Penyakit", Form: map[string]string{"kode_penyakit": kode, "nama_penyakit": nama, "deskripsi": deskripsi, "solusi_obat": solusiObat, "solusi_lain": solusiLain}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Penyakit", "Action": fmt.Sprintf("/admin/penyakit/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_penyakit_form", data); return
    }
    item.KodePenyakit = kode
    item.NamaPenyakit = nama
    item.Deskripsi = deskripsi
    item.SolusiObat = solusiObat
    item.SolusiLain = solusiLain
    app.DB.Save(&item)
    app.setFlash(w, r, "Penyakit berhasil diperbarui", "")
    http.Redirect(w, r, "/admin/penyakit", http.StatusSeeOther)
}

func (app *App) AdminPenyakitDeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    app.DB.Delete(&models.Penyakit{}, id)
    app.setFlash(w, r, "Penyakit berhasil dihapus", "")
    http.Redirect(w, r, "/admin/penyakit", http.StatusSeeOther)
}

func (app *App) AdminRuleIndexHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var totalRows int64
    var items []models.Rule
    app.DB.Model(&models.Rule{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    app.DB.Preload("Penyakit").Preload("Gejala").Order("id_rule").Offset((page-1)*pageSize).Limit(pageSize).Find(&items)
    data := &TemplateData{Title: "Kelola Rules", Data: map[string]interface{}{"Rules": items, "Pagination": pagination}}
    app.renderTemplate(w, r, "admin_layout", "admin_rule_index", data)
}

func (app *App) AdminRuleCreateHandler(w http.ResponseWriter, r *http.Request) {
    var penyakit []models.Penyakit
    var gejala []models.Gejala
    app.DB.Order("id_penyakit").Find(&penyakit)
    app.DB.Order("id_gejala").Find(&gejala)
    data := &TemplateData{Title: "Tambah Rule", Data: map[string]interface{}{"FormTitle": "Tambah Rule", "Action": "/admin/rule", "Method": "POST", "Penyakit": penyakit, "Gejala": gejala}}
    app.renderTemplate(w, r, "admin_layout", "admin_rule_form", data)
}

func (app *App) AdminRuleStoreHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    kodeRule := strings.TrimSpace(r.FormValue("kode_rule"))
    kodePenyakit := strings.TrimSpace(r.FormValue("kode_penyakit"))
    kodeGejala := strings.TrimSpace(r.FormValue("kode_gejala"))
    errors := make(map[string]string)
    if kodeRule == "" { errors["kode_rule"] = "Kode rule wajib diisi" }
    if kodePenyakit == "" { errors["kode_penyakit"] = "Penyakit wajib dipilih" }
    if kodeGejala == "" { errors["kode_gejala"] = "Gejala wajib dipilih" }
    ruleInt := 0
    if kodeRule != "" {
        if parsed, err := strconv.Atoi(kodeRule); err != nil { errors["kode_rule"] = "Kode rule harus berupa angka" } else { ruleInt = parsed }
    }
    if len(errors) > 0 {
        var penyakit []models.Penyakit
        var gejala []models.Gejala
        app.DB.Order("id_penyakit").Find(&penyakit)
        app.DB.Order("id_gejala").Find(&gejala)
        data := &TemplateData{Title: "Tambah Rule", Form: map[string]string{"kode_rule": kodeRule, "kode_penyakit": kodePenyakit, "kode_gejala": kodeGejala}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Rule", "Action": "/admin/rule", "Method": "POST", "Penyakit": penyakit, "Gejala": gejala}}
        app.renderTemplate(w, r, "admin_layout", "admin_rule_form", data); return
    }
    item := models.Rule{KodeRule: ruleInt, KodePenyakit: kodePenyakit, KodeGejala: kodeGejala}
    app.DB.Create(&item)
    app.setFlash(w, r, "Rule berhasil ditambahkan", "")
    http.Redirect(w, r, "/admin/rule", http.StatusSeeOther)
}

func (app *App) AdminRuleEditHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Rule
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    var penyakit []models.Penyakit
    var gejala []models.Gejala
    app.DB.Order("id_penyakit").Find(&penyakit)
    app.DB.Order("id_gejala").Find(&gejala)
    data := &TemplateData{Title: "Edit Rule", Form: map[string]string{"kode_rule": strconv.Itoa(item.KodeRule), "kode_penyakit": item.KodePenyakit, "kode_gejala": item.KodeGejala}, Data: map[string]interface{}{"FormTitle": "Edit Rule", "Action": fmt.Sprintf("/admin/rule/%d/update", item.ID), "Method": "POST", "Penyakit": penyakit, "Gejala": gejala}}
    app.renderTemplate(w, r, "admin_layout", "admin_rule_form", data)
}

func (app *App) AdminRuleUpdateHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.Rule
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    kodeRule := strings.TrimSpace(r.FormValue("kode_rule"))
    kodePenyakit := strings.TrimSpace(r.FormValue("kode_penyakit"))
    kodeGejala := strings.TrimSpace(r.FormValue("kode_gejala"))
    errors := make(map[string]string)
    if kodeRule == "" { errors["kode_rule"] = "Kode rule wajib diisi" }
    if kodePenyakit == "" { errors["kode_penyakit"] = "Penyakit wajib dipilih" }
    if kodeGejala == "" { errors["kode_gejala"] = "Gejala wajib dipilih" }
    ruleInt := 0
    if kodeRule != "" {
        if parsed, err := strconv.Atoi(kodeRule); err != nil { errors["kode_rule"] = "Kode rule harus berupa angka" } else { ruleInt = parsed }
    }
    if len(errors) > 0 {
        var penyakit []models.Penyakit
        var gejala []models.Gejala
        app.DB.Order("id_penyakit").Find(&penyakit)
        app.DB.Order("id_gejala").Find(&gejala)
        data := &TemplateData{Title: "Edit Rule", Form: map[string]string{"kode_rule": kodeRule, "kode_penyakit": kodePenyakit, "kode_gejala": kodeGejala}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Rule", "Action": fmt.Sprintf("/admin/rule/%d/update", item.ID), "Method": "POST", "Penyakit": penyakit, "Gejala": gejala}}
        app.renderTemplate(w, r, "admin_layout", "admin_rule_form", data); return
    }
    item.KodeRule = ruleInt
    item.KodePenyakit = kodePenyakit
    item.KodeGejala = kodeGejala
    app.DB.Save(&item)
    app.setFlash(w, r, "Rule berhasil diperbarui", "")
    http.Redirect(w, r, "/admin/rule", http.StatusSeeOther)
}

func (app *App) AdminRuleDeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    app.DB.Delete(&models.Rule{}, id)
    app.setFlash(w, r, "Rule berhasil dihapus", "")
    http.Redirect(w, r, "/admin/rule", http.StatusSeeOther)
}

func (app *App) AdminRiwayatIndexHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var totalRows int64
    var items []models.Riwayat
    app.DB.Model(&models.Riwayat{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    app.DB.Preload("User").Order("tanggal desc").Offset((page-1)*pageSize).Limit(pageSize).Find(&items)
    data := &TemplateData{Title: "Kelola Riwayat", Data: map[string]interface{}{"Riwayat": items, "Pagination": pagination}}
    app.renderTemplate(w, r, "admin_layout", "admin_riwayat_index", data)
}

func (app *App) AdminRiwayatDeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    app.DB.Delete(&models.Riwayat{}, id)
    app.setFlash(w, r, "Riwayat berhasil dihapus", "")
    http.Redirect(w, r, "/admin/riwayat", http.StatusSeeOther)
}

func (app *App) AdminUsersIndexHandler(w http.ResponseWriter, r *http.Request) {
    page, pageSize := parsePage(r)
    var totalRows int64
    var items []models.User
    app.DB.Model(&models.User{}).Count(&totalRows)
    pagination := makePagination(page, pageSize, totalRows)
    app.DB.Order("id").Offset((page-1)*pageSize).Limit(pageSize).Find(&items)
    data := &TemplateData{Title: "Kelola Pengguna", Data: map[string]interface{}{"Users": items, "Pagination": pagination}}
    app.renderTemplate(w, r, "admin_layout", "admin_users_index", data)
}

func (app *App) AdminUsersCreateHandler(w http.ResponseWriter, r *http.Request) {
    data := &TemplateData{Title: "Tambah Pengguna", Data: map[string]interface{}{"FormTitle": "Tambah Pengguna", "Action": "/admin/users", "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_users_form", data)
}

func (app *App) AdminUsersStoreHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    name := strings.TrimSpace(r.FormValue("name"))
    email := strings.TrimSpace(r.FormValue("email"))
    password := r.FormValue("password")
    confirm := r.FormValue("password_confirmation")
    role := strings.TrimSpace(r.FormValue("role"))
    errors := make(map[string]string)
    if name == "" { errors["name"] = "Nama wajib diisi" }
    if email == "" { errors["email"] = "Email wajib diisi" }
    if password == "" { errors["password"] = "Password wajib diisi" }
    if password != confirm { errors["password_confirmation"] = "Konfirmasi password tidak sesuai" }
    if role != "user" && role != "admin" { errors["role"] = "Role tidak valid" }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Tambah Pengguna", Form: map[string]string{"name": name, "email": email, "role": role}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Pengguna", "Action": "/admin/users", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_users_form", data); return
    }
    var existing models.User
    if err := app.DB.Where("email = ?", email).First(&existing).Error; err == nil {
        errors["email"] = "Email sudah digunakan"
        data := &TemplateData{Title: "Tambah Pengguna", Form: map[string]string{"name": name, "email": email, "role": role}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Tambah Pengguna", "Action": "/admin/users", "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_users_form", data); return
    }
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil { http.Error(w, "Tidak dapat membuat pengguna baru", http.StatusInternalServerError); return }
    item := models.User{Name: name, Email: email, Role: role, Password: string(hashed), CreatedAt: ptrTime(time.Now()), UpdatedAt: ptrTime(time.Now())}
    app.DB.Create(&item)
    app.setFlash(w, r, "User berhasil ditambahkan", "")
    http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func (app *App) AdminUsersEditHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var item models.User
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    data := &TemplateData{Title: "Edit Pengguna", Form: map[string]string{"name": item.Name, "email": item.Email, "role": item.Role}, Data: map[string]interface{}{"FormTitle": "Edit Pengguna", "Action": fmt.Sprintf("/admin/users/%d/update", item.ID), "Method": "POST"}}
    app.renderTemplate(w, r, "admin_layout", "admin_users_form", data)
}

func (app *App) AdminUsersUpdateHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    currentUser := app.GetCurrentUser(r)
    var item models.User
    if err := app.DB.First(&item, id).Error; err != nil { http.NotFound(w,r); return }
    if err := r.ParseForm(); err != nil { http.Error(w, "Invalid request", http.StatusBadRequest); return }
    name := strings.TrimSpace(r.FormValue("name"))
    email := strings.TrimSpace(r.FormValue("email"))
    password := r.FormValue("password")
    confirm := r.FormValue("password_confirmation")
    role := strings.TrimSpace(r.FormValue("role"))
    errors := make(map[string]string)
    if name == "" { errors["name"] = "Nama wajib diisi" }
    if email == "" { errors["email"] = "Email wajib diisi" }
    if password != "" && password != confirm { errors["password_confirmation"] = "Konfirmasi password tidak sesuai" }
    if role != "user" && role != "admin" { errors["role"] = "Role tidak valid" }
    if len(errors) > 0 {
        data := &TemplateData{Title: "Edit Pengguna", Form: map[string]string{"name": name, "email": email, "role": role}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Pengguna", "Action": fmt.Sprintf("/admin/users/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_users_form", data); return
    }
    var existing models.User
    if err := app.DB.Where("email = ? AND id <> ?", email, item.ID).First(&existing).Error; err == nil {
        errors["email"] = "Email sudah digunakan"
        data := &TemplateData{Title: "Edit Pengguna", Form: map[string]string{"name": name, "email": email, "role": role}, Errors: errors, Data: map[string]interface{}{"FormTitle": "Edit Pengguna", "Action": fmt.Sprintf("/admin/users/%d/update", item.ID), "Method": "POST"}}
        app.renderTemplate(w, r, "admin_layout", "admin_users_form", data); return
    }
    item.Name = name
    item.Email = email
    item.Role = role
    if password != "" {
        hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil { http.Error(w, "Tidak dapat memperbarui password", http.StatusInternalServerError); return }
        item.Password = string(hashed)
    }
    if currentUser != nil && currentUser.ID == item.ID && item.Role != "admin" {
        item.Role = "admin"
    }
    item.UpdatedAt = ptrTime(time.Now())
    app.DB.Save(&item)
    app.setFlash(w, r, "User berhasil diperbarui", "")
    http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func (app *App) AdminUsersDeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    currentUser := app.GetCurrentUser(r)
    if currentUser != nil && uint(id) == currentUser.ID {
        app.setFlash(w, r, "", "Anda tidak dapat menghapus akun sendiri.")
        http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
        return
    }
    app.DB.Delete(&models.User{}, id)
    app.setFlash(w, r, "User berhasil dihapus", "")
    http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
