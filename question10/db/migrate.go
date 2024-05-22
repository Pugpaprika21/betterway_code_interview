package db

import (
	"github.com/Pugpaprika21/question10/models"
	"gorm.io/gorm"
)

type migration struct {
	db *gorm.DB
}

func NewMigration(db *gorm.DB) *migration {
	return &migration{
		db: db,
	}
}

func (m *migration) Run(dst ...interface{}) {
	if len(dst) == 0 {
		dst = m.runbatch()
	}

	for _, model := range dst {
		if err := m.db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
}

// register models here!!
func (m *migration) runbatch() []interface{} {
	return []interface{}{
		&models.Employee{},
		&models.Department{},
	}
}