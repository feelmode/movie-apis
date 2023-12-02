package movie

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID          uint8  `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" valid:"required"`
	Description string `json:"description" valid:"required"`
	Rating      uint8  `json:"rating" valid:"range(1|10)"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
}

type Response struct {
	ID          uint8  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      uint8  `json:"rating"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
