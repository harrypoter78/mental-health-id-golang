package models

import "time"

type Riwayat struct {
    ID          int        `gorm:"primaryKey;column:id_riwayat;autoIncrement"`
    UserID      uint       `gorm:"column:user_id"`
    NamaPenyakit string     `gorm:"column:nama_penyakit;size:50;not null"`
    Tanggal     time.Time  `gorm:"column:tanggal;type:date"`
    User        *User      `gorm:"foreignKey:UserID;references:ID"`
}

func (Riwayat) TableName() string {
    return "riwayat"
}
