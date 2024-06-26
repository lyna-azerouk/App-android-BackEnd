package models

type Tags struct {
	Amenity string `json:"amenity"`
	Name    string `json:"name"`
	Tourism string `json:"tourism"`
}

type OverPassRestaurant struct {
	Type string  `json:"type"`
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Tags Tags    `json:"tags"`
}

type OverPassResponse struct {
	Version   float64              `json:"version"`
	Generator string               `json:"generator"`
	OSM3S     map[string]string    `json:"osm3s"`
	Elements  []OverPassRestaurant `json:"elements"`
}

type BouffluenceRestaurant struct {
	RestaurantDetails    OverPassRestaurant `json:"restaurantDetails"`
	Menu                 []Menu             `json:"menus"`
	OrderAverageDuration int                `json:"order_average_duration"`
}

type Menu struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	RestaurantID int     `json:"restaurent_id"`
	Description  string  `json:"description"`
	Image        string  `json:"url"`
}

type RestaurantDetails struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Affluence            string `json:"affluence"`
	OrderAverageDuration int    `json:"order_average_duration"`
}
