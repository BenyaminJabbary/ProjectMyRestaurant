package calculatorPriceService

import (
	"fmt"
	"go-orderFood/core"
)

var logger = core.NewFileLogger()

func CalculatePriceFood(priceFood int) (price1 int) {
	logger.Info().Str("calculatorPrice", "CalculatePriceFood").Msg("Start of CalculatePriceFood")

	var numberFood int
	fmt.Print("Number Food: ")
	fmt.Scanln(&numberFood)
	price1 = priceFood * numberFood
	fmt.Printf("Price: %d\n", price1)

	logger.Info().Str("calculatorPrice", "CalculatePriceFood").Msg("End of CalculatePriceFood")

	return
}

func CalculatePriceDrink(priceDrink int) (price2 int) {
	logger.Info().Str("calculatorPrice", "CalculatePriceDrink").Msg("Start of CalculatePriceDrink")

	var numberDrink int
	fmt.Print("Number Drink: ")
	fmt.Scanln(&numberDrink)
	price2 = priceDrink * numberDrink
	fmt.Printf("Price: %d\n", price2)

	logger.Info().Str("calculatorPrice", "CalculatePriceDrink").Msg("End of CalculatePriceDrink")

	return
}


