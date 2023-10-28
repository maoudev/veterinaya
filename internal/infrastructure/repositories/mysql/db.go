package mysql

import (
	"sync"

	"github.com/maoudev/veterinaya/internal/config"
	"github.com/maoudev/veterinaya/internal/pkg/domain"
	"golang.org/x/exp/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	onceDBload sync.Once
	tables     = []interface{}{
		&domain.User{},
		&domain.Pet{},
		&domain.Veterinarian{},
		&domain.Appointment{},
	}
)

func connect() *gorm.DB {
	dsn := config.DB_URL
	onceDBload.Do(func() {
		var err error

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect: " + dsn)
		}

		slog.Info("Connected to db!")
		migrate()
	})

	return db
}

func migrate() {
	for _, t := range tables {
		db.AutoMigrate(t)
	}
}
