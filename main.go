package main

import (
	"fmt"
	"go-orderFood/core"
	"go-orderFood/services"
)

var logger = core.NewFileLogger()

func main() {
	logger.Info().Msg("Start of main")

	x, y := services.GetOrder()

	finalPrice := x + y
	fmt.Printf("Final Price: %d\n", finalPrice)

	if x == 0 && y == 0 {
		fmt.Println("the order was not completed")
		logger.Warn().Str("Restaurant", "main").Msg("the order was not completed")
	} else {
		fmt.Println("The order was successfully placed")
	}

	logger.Info().Msg("End of main")

}