package models

import "time"

type User struct {
    ID              uint       `gorm:"primaryKey;column:id;autoIncrement"`
    Name            string     `gorm:"column:name;size:255;not null"`
    Email           string     `gorm:"column:email;size:255;not null;uniqueIndex"`
    Role            string     `gorm:"column:role;type:enum('admin','user');default:'user'"`
    EmailVerifiedAt *time.Time `gorm:"column:email_verified_at;type:timestamp"`
    Password        string     `gorm:"column:password;size:255;not null"`
    RememberToken   string     `gorm:"column:remember_token;size:100"`
    CreatedAt       *time.Time `gorm:"column:created_at;type:timestamp"`
    UpdatedAt       *time.Time `gorm:"column:updated_at;type:timestamp"`
    Riwayat         []Riwayat  `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
    return "users"
}
