package main

import (
    "log"
    "net/http"

    "github.com/example/mental-health-id/internal/config"
    "github.com/example/mental-health-id/internal/database"
    "github.com/example/mental-health-id/internal/handlers"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    db, err := database.Connect(cfg)
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    if err := database.Migrate(db); err != nil {
        log.Fatalf("database migration failed: %v", err)
    }

    store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
    store.Options = &sessions.Options{
        Path:     "/",
        HttpOnly: true,
        MaxAge:   86400 * 7,
    }

    app, err := handlers.NewApp(db, store)
    if err != nil {
        log.Fatalf("failed to initialize app: %v", err)
    }

    r := mux.NewRouter()
    r.Use(app.MethodOverrideMiddleware)
    r.Use(app.LoadUserMiddleware)

    r.HandleFunc("/", app.IndexHandler).Methods("GET")
    r.HandleFunc("/diagnosis/kuis", app.DiagnosisKuisHandler).Methods("GET")
    r.HandleFunc("/diagnosis/proses", app.DiagnosisProsesHandler).Methods("GET", "POST")
    r.HandleFunc("/diagnosis/riwayat", app.DiagnosisRiwayatHandler).Methods("GET")

    r.HandleFunc("/login", app.LoginPageHandler).Methods("GET")
    r.HandleFunc("/login", app.LoginHandler).Methods("POST")
    r.HandleFunc("/register", app.RegisterPageHandler).Methods("GET")
    r.HandleFunc("/register", app.RegisterHandler).Methods("POST")
    r.HandleFunc("/logout", app.LogoutHandler).Methods("POST")

    r.HandleFunc("/profile", app.RequireAuth(app.ProfilePageHandler)).Methods("GET")
    r.HandleFunc("/profile/update", app.RequireAuth(app.ProfileUpdateHandler)).Methods("POST")
    r.HandleFunc("/profile/change-password", app.RequireAuth(app.ChangePasswordHandler)).Methods("POST")

    admin := r.PathPrefix("/admin").Subrouter()
    admin.Use(app.RequireAdminMiddleware)
    admin.HandleFunc("/dashboard", app.AdminDashboardHandler).Methods("GET")

    admin.HandleFunc("/gejala", app.AdminGejalaIndexHandler).Methods("GET")
    admin.HandleFunc("/gejala/create", app.AdminGejalaCreateHandler).Methods("GET")
    admin.HandleFunc("/gejala", app.AdminGejalaStoreHandler).Methods("POST")
    admin.HandleFunc("/gejala/{id:[0-9]+}/edit", app.AdminGejalaEditHandler).Methods("GET")
    admin.HandleFunc("/gejala/{id:[0-9]+}", app.AdminGejalaUpdateHandler).Methods("PUT")
    admin.HandleFunc("/gejala/{id:[0-9]+}/update", app.AdminGejalaUpdateHandler).Methods("POST")
    admin.HandleFunc("/gejala/{id:[0-9]+}/delete", app.AdminGejalaDeleteHandler).Methods("POST")
    admin.HandleFunc("/gejala/{id:[0-9]+}", app.AdminGejalaDeleteHandler).Methods("DELETE")

    admin.HandleFunc("/penyakit", app.AdminPenyakitIndexHandler).Methods("GET")
    admin.HandleFunc("/penyakit/create", app.AdminPenyakitCreateHandler).Methods("GET")
    admin.HandleFunc("/penyakit", app.AdminPenyakitStoreHandler).Methods("POST")
    admin.HandleFunc("/penyakit/{id:[0-9]+}/edit", app.AdminPenyakitEditHandler).Methods("GET")
    admin.HandleFunc("/penyakit/{id:[0-9]+}", app.AdminPenyakitUpdateHandler).Methods("PUT")
    admin.HandleFunc("/penyakit/{id:[0-9]+}/update", app.AdminPenyakitUpdateHandler).Methods("POST")
    admin.HandleFunc("/penyakit/{id:[0-9]+}/delete", app.AdminPenyakitDeleteHandler).Methods("POST")
    admin.HandleFunc("/penyakit/{id:[0-9]+}", app.AdminPenyakitDeleteHandler).Methods("DELETE")

    admin.HandleFunc("/rule", app.AdminRuleIndexHandler).Methods("GET")
    admin.HandleFunc("/rule/create", app.AdminRuleCreateHandler).Methods("GET")
    admin.HandleFunc("/rule", app.AdminRuleStoreHandler).Methods("POST")
    admin.HandleFunc("/rule/{id:[0-9]+}/edit", app.AdminRuleEditHandler).Methods("GET")
    admin.HandleFunc("/rule/{id:[0-9]+}", app.AdminRuleUpdateHandler).Methods("PUT")
    admin.HandleFunc("/rule/{id:[0-9]+}/update", app.AdminRuleUpdateHandler).Methods("POST")
    admin.HandleFunc("/rule/{id:[0-9]+}/delete", app.AdminRuleDeleteHandler).Methods("POST")
    admin.HandleFunc("/rule/{id:[0-9]+}", app.AdminRuleDeleteHandler).Methods("DELETE")

    admin.HandleFunc("/riwayat", app.AdminRiwayatIndexHandler).Methods("GET")
    admin.HandleFunc("/riwayat/{id:[0-9]+}/delete", app.AdminRiwayatDeleteHandler).Methods("POST")
    admin.HandleFunc("/riwayat/{id:[0-9]+}", app.AdminRiwayatDeleteHandler).Methods("DELETE")

    admin.HandleFunc("/users", app.AdminUsersIndexHandler).Methods("GET")
    admin.HandleFunc("/users/create", app.AdminUsersCreateHandler).Methods("GET")
    admin.HandleFunc("/users", app.AdminUsersStoreHandler).Methods("POST")
    admin.HandleFunc("/users/{id:[0-9]+}/edit", app.AdminUsersEditHandler).Methods("GET")
    admin.HandleFunc("/users/{id:[0-9]+}", app.AdminUsersUpdateHandler).Methods("PUT")
    admin.HandleFunc("/users/{id:[0-9]+}/update", app.AdminUsersUpdateHandler).Methods("POST")
    admin.HandleFunc("/users/{id:[0-9]+}/delete", app.AdminUsersDeleteHandler).Methods("POST")
    admin.HandleFunc("/users/{id:[0-9]+}", app.AdminUsersDeleteHandler).Methods("DELETE")

    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
