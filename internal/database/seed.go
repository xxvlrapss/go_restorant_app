package database

import (
	"fmt"

	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
	"github.com/xxvlrapss/go_restorant_app.git/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(
		&model.MenuItem{},
		&model.Order{},
		&model.ProductOrder{},
		&model.User{},
	)

	foodMenu := []model.MenuItem{
		{
			Name:      "Indomie",
			OrderCode: "indomie",
			Price:     30000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam lada hitam",
			OrderCode: "ayam_lada_hitam",
			Price:     40000,
			Type:      constant.MenuTypeFood,
		},
	}
	drinksMenu := []model.MenuItem{
		{
			Name:      "Teh Manis",
			OrderCode: "teh_manis",
			Price:     12000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Jeruk",
			OrderCode: "jus_jeruk",
			Price:     15000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("Seeding db data...")
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}

}