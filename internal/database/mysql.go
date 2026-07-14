package database

import (
    "fmt"

    "github.com/example/mental-health-id/internal/config"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
    statements := []string{
        `CREATE TABLE IF NOT EXISTS gejala (
            id_gejala INT(3) NOT NULL AUTO_INCREMENT,
            kode_gejala VARCHAR(3) NOT NULL,
            nama_gejala VARCHAR(100) NOT NULL,
            PRIMARY KEY (id_gejala),
            KEY kode_gejala (kode_gejala)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;`,
        `CREATE TABLE IF NOT EXISTS penyakit (
            id_penyakit INT(3) NOT NULL AUTO_INCREMENT,
            kode_penyakit VARCHAR(3) NOT NULL,
            nama_penyakit VARCHAR(50) NOT NULL,
            deskripsi TEXT DEFAULT NULL,
            solusi_obat TEXT DEFAULT NULL,
            solusi_lain TEXT DEFAULT NULL,
            PRIMARY KEY (id_penyakit),
            KEY kode_penyakit (kode_penyakit)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;`,
        `CREATE TABLE IF NOT EXISTS rule (
            id_rule INT(3) NOT NULL AUTO_INCREMENT,
            kode_rule INT(3) NOT NULL,
            kode_penyakit VARCHAR(3) NOT NULL,
            kode_gejala VARCHAR(3) NOT NULL,
            PRIMARY KEY (id_rule),
            KEY FK_rule_penyakit (kode_penyakit),
            KEY FK_rule_gejala (kode_gejala)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;`,
        `CREATE TABLE IF NOT EXISTS users (
            id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            role ENUM('admin','user') NOT NULL DEFAULT 'user',
            email_verified_at TIMESTAMP NULL DEFAULT NULL,
            password VARCHAR(255) NOT NULL,
            remember_token VARCHAR(100) DEFAULT NULL,
            created_at TIMESTAMP NULL DEFAULT NULL,
            updated_at TIMESTAMP NULL DEFAULT NULL,
            PRIMARY KEY (id),
            UNIQUE KEY users_email_unique (email)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS riwayat (
            id_riwayat INT(3) NOT NULL AUTO_INCREMENT,
            user_id BIGINT(20) UNSIGNED DEFAULT NULL,
            nama_penyakit VARCHAR(50) NOT NULL,
            tanggal DATE DEFAULT NULL,
            PRIMARY KEY (id_riwayat),
            KEY riwayat_user_id_foreign (user_id)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;`,
    }

    for _, statement := range statements {
        if err := db.Exec(statement).Error; err != nil {
            return err
        }
    }

    return nil
}
