package services

import (
	"go-orderFood/entities"
)

var foods []entities.MenuRestaurant = Foods()
var drinks []entities.MenuRestaurant = Drinks()

func Foods() []entities.MenuRestaurant {

	food := []entities.MenuRestaurant{}
	food = append(food, entities.MenuRestaurant{Name: "Pizza", Price: 130, Status: true})
	food = append(food, entities.MenuRestaurant{Name: "Potato", Price: 56, Status: true})
	food = append(food, entities.MenuRestaurant{Name: "PizzaBomb", Price: 180, Status: false})
	food = append(food, entities.MenuRestaurant{Name: "ChickenHam", Price: 75, Status: true})
	food = append(food, entities.MenuRestaurant{Name: "Hotdog", Price: 98, Status: false})
	food = append(food, entities.MenuRestaurant{Name: "Pasta", Price: 200, Status: true})
	food = append(food, entities.MenuRestaurant{Name: "Falafel", Price: 50, Status: true})
	food = append(food, entities.MenuRestaurant{Name: "Ham", Price: 70, Status: false})
	food = append(food, entities.MenuRestaurant{Name: "Burger", Price: 100, Status: true})

	
	return food

}

func Drinks() []entities.MenuRestaurant {

	drink := []entities.MenuRestaurant{}
	drink = append(drink, entities.MenuRestaurant{Name: "CocaCola", Price: 20, Status: true})
	drink = append(drink, entities.MenuRestaurant{Name: "Water", Price: 5, Status: true})
	drink = append(drink, entities.MenuRestaurant{Name: "ButterMilk", Price: 5, Status: false})

	return drink

}
