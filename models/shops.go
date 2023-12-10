package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	CategoryId  int32   `json:"category_id"`
	UserId      int     `json:"user_id"`
	ManagerId   int     `json:"manager_id"`
	MboyId      int     `json:"m_boy_id"`
	CityId      int     `json:"city_id"`
	ShopName    string  `json:"shop_name"`
	ShopAddress string  `json:"shop_address"`
	ShopLat     float64 `json:"shop_lat" gorm:"type:decimal(10,8)"`
	ShopLng     float64 `json:"shop_lng" gorm:"type:decimal(11,8)"`
	ShopStatus  string  `json:"shop_status" gorm:"default:active"`
	PackageId   int     `json:"package_id" gorm:"default:1"`
}
