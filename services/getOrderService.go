package services

import (
	"errors"
	"fmt"
	"go-orderFood/calculatorPriceService"
	"go-orderFood/core"
)

var logger = core.NewFileLogger()

func GetOrder() (x, y int) {
	logger.Info().Str("getOrderService", "GetOrder").Msg("Start of GetOrder")

	x = GetFood()
	y = GetDrink()

	logger.Info().Str("getOrderService", "GetOrder").Msg("End of GetOrder")

	return x, y
}

func GetFood() (f int) {
	logger.Info().Str("getOrderService", "GetFood").Msg("Start of GetFood")

	fmt.Printf("List foods: %v\n", foods)

	food := ""
	fmt.Print("Select Food: ")
	fmt.Scanln(&food)
	for _, fod := range foods {
		if fod.Name == food && fod.Status {
			f = calculatorPriceService.CalculatePriceFood(fod.Price)
			break
		} else if fod.Name == food && !fod.Status {
			fmt.Println("this food is not available")
		    logger.Warn().Str("getOrderService", "GetFood").Msg("Warn is at GetFood")
			break
		} else if food == "" {
			err := errors.New("select food is empty")
		    logger.Error().Err(err).Str("getOrderService", "GetFood").Msg("Error is at GetFood")
			break
		}
	}

	logger.Info().Str("getOrderService", "GetFood").Msg("End of GetFood")

	return
}

func GetDrink() (d int) {
	logger.Info().Str("getOrderService", "GetDrink").Msg("Start of GetDrink")

	fmt.Printf("List drinks: %v\n", drinks)

	drink := ""
	fmt.Print("Select Drink: ")
	fmt.Scanln(&drink)
	for _, dri := range drinks {
		if dri.Name == drink && dri.Status {
			d = calculatorPriceService.CalculatePriceDrink(dri.Price)
			break
		} else if dri.Name == drink && !dri.Status {
			fmt.Println("this drink is not available")
			logger.Warn().Str("getOrderService", "GetDrink").Msg("Warn is at GetDrink")
			break
		} else if drink == "" {
			err := errors.New("select drink is empty")
			logger.Error().Err(err).Str("getOrderService", "GetDrink").Msg("Error is at GetDrink")
			break
		}
	}

	logger.Info().Str("getOrderService", "GetDrink").Msg("End of GetDrink")

	return
}
