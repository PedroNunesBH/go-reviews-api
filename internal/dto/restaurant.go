package dto 


type RestaurantDTO struct {
	Name 	string	`json:"name"`
	Cnpj 	string	`json:"cnpj"`
	Address string	`json:"address"`
}