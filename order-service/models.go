package main

import (
    "gorm.io/gorm"
)

// Модель заказа
type Order struct {
    gorm.Model
    DrinkID   uint    `json:"drink_id"`
    DrinkName string  `json:"drink_name"` // Копируем из Menu Service
    Price     float64 `json:"price"`      // Копируем из Menu Service
    Quantity  int     `json:"quantity"`
    Total     float64 `json:"total"`
    Status    string  `json:"status"` // pending, completed, cancelled
}

// DTO для создания заказа
type CreateOrderRequest struct {
    DrinkID  uint `json:"drink_id"`
    Quantity int  `json:"quantity"`
}

// Структура для ответа от Menu Service
type DrinkFromMenu struct {
    ID      uint    `json:"ID"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    InStock bool    `json:"in_stock"`
}