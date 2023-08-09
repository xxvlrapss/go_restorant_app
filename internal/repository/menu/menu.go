package menu

import (
	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenuList(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem

	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil

}

func (m *menuRepo) 	GetMenu(orderCode string) (model.MenuItem, error) {
	var menuData model.MenuItem

	if err := m.db.Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}
	return menuData, nil
}