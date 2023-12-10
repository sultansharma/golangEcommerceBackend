package models

import (
	"time"

	"gorm.io/gorm"
)

func Photo() [0]string {
	//photo := [...]int{}
	return [...]string{}
}

type Product struct {
	gorm.Model
	ProductName      string    `json:"product_name"`
	ThumbnailImage   string    `json:"thumbnail_image"`
	Photos           string    `json:"photos" gorm:"default:[]"`
	Description      string    `json:"description"`
	FullDescription  string    `json:"full_description"`
	ShopId           int       `json:"shop_id"`
	CategoryId       int       `json:"category_id"`
	SubCategoryId    int       `json:"sub_category_id"`
	MrpPrice         int       `json:"mrp_price"`
	DiscountedPrice  int       `json:"discounted_price"`
	Stock            int       `json:"stock"`
	Status           string    `json:"status" gorm:"default:pending"`
	FlipkartPrice    int       `json:"flipkart_price"`
	AmazonPrice      int       `json:"amazon_price"`
	FlipkartLink     int       `json:"flipkart_link"`
	AmazonLink       int       `json:"amazon_link"`
	ComparePriceDate time.Time `json:"compare_price_date"`
}

type ProductVariation struct {
	gorm.Model
	VariationType    string    `json:"variation_type"`
	Name             string    `json:"name"`
	ProductId        int       `json:"product_id"`
	Image            int       `json:"image"`
	MrpPrice         int       `json:"mrp_price"`
	DiscountedPrice  int       `json:"discounted_price"`
	Stock            int       `json:"stock"`
	Status           string    `json:"status" gorm:"default:pending"`
	FlipkartPrice    int       `json:"flipkart_price"`
	AmazonPrice      int       `json:"amazon_price"`
	FlipkartLink     int       `json:"flipkart_link"`
	AmazonLink       int       `json:"amazon_link"`
	ComparePriceDate time.Time `json:"compare_price_date"`
}
