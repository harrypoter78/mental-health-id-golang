package handlers

import (
    "testing"

    "github.com/gorilla/sessions"
    "gorm.io/gorm"
)

func TestTemplatesParse(t *testing.T) {
    app, err := NewApp(&gorm.DB{}, sessions.NewCookieStore([]byte("testsecret")))
    if err != nil {
        t.Fatalf("failed to parse templates: %v", err)
    }
    if app.TemplateDir == "" {
        t.Fatal("template directory was not initialized")
    }
    if app.FuncMap == nil {
        t.Fatal("template function map was not initialized")
    }
}
