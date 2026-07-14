package models

type Rule struct {
    ID          int      `gorm:"primaryKey;column:id_rule;autoIncrement"`
    KodeRule    int      `gorm:"column:kode_rule;not null"`
    KodePenyakit string   `gorm:"column:kode_penyakit;size:3;not null"`
    KodeGejala  string   `gorm:"column:kode_gejala;size:3;not null"`
    Penyakit    *Penyakit `gorm:"foreignKey:KodePenyakit;references:KodePenyakit"`
    Gejala      *Gejala   `gorm:"foreignKey:KodeGejala;references:KodeGejala"`
}

func (Rule) TableName() string {
    return "rule"
}
