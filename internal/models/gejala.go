package models

type Gejala struct {
    ID         int    `gorm:"primaryKey;column:id_gejala;autoIncrement"`
    KodeGejala string `gorm:"column:kode_gejala;size:3;not null"`
    NamaGejala string `gorm:"column:nama_gejala;size:100;not null"`
}

func (Gejala) TableName() string {
    return "gejala"
}
