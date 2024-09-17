package main

import (
	"backend/config"
	"backend/models"
	"backend/routes"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    cfg := config.LoadConfig()

    dsn := cfg.PostgresConn
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
    db.Exec(`
    DO $$
    BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
        CREATE TYPE organization_type AS ENUM (
            'IE', 
            'LLC', 
            'JSC');
    END IF;
    END $$;
    `)

    // db.AutoMigrate(&models.Employee{}, &models.Organization{}, &models.OrganizationResponsible{}, &models.Tender{}, &models.Bid{}, &models.BidReview{})
    db.AutoMigrate(&models.Tender{}, &models.Bid{}, &models.BidReview{})

    router := routes.SetupRouter(db)
    log.Printf("Запуск сервера на %s", cfg.ServerAddress)
    if err := router.Run(cfg.ServerAddress); err != nil {
        log.Fatalf("Ошибка запуска сервера: %v", err)
    }

    router.Run(cfg.ServerAddress)
}
