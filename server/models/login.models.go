package models

import (
	"github.com/golang-jwt/jwt"
)

type ClientCredentials struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RestaurantCredentials struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
}

type ClientClaims struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type RestaurantClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
