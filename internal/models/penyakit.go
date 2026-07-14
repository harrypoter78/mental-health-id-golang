package models

type Penyakit struct {
    ID          int    `gorm:"primaryKey;column:id_penyakit;autoIncrement"`
    KodePenyakit string `gorm:"column:kode_penyakit;size:3;not null"`
    NamaPenyakit string `gorm:"column:nama_penyakit;size:50;not null"`
    Deskripsi    string `gorm:"column:deskripsi;type:text"`
    SolusiObat   string `gorm:"column:solusi_obat;type:text"`
    SolusiLain   string `gorm:"column:solusi_lain;type:text"`
}

func (Penyakit) TableName() string {
    return "penyakit"
}
