package main

import "gorm.io/gorm"

type Drink struct {
    gorm.Model
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    InStock bool    `json:"in_stock"`
}

// DTO для создания напитка
type CreateDrinkRequest struct {
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    InStock bool    `json:"in_stock"`
}