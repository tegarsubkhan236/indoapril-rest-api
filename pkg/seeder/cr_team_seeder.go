package seeder

import (
	"example/pkg/model"
	"gorm.io/gorm"
)

func CrTeamSeeder(db *gorm.DB) {
	teams := []model.CrTeam{
		{
			Name: "Central Team",
		},
	}

	result := db.Create(&teams)
	if result.Error != nil {
		panic("failed to seed CoreRole data")
	}
}
