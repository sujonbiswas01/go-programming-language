// main.go
package main

import (
	"fmt"
	"restaurant-app/delivery"
	"restaurant-app/food"

	"github.com/fatih/color"
)

func Init() {
	fmt.Println("initializing db..........")
}
func main() {
	sujon := "sujon"
	color.Red("Prints text in cyan. %v\n", sujon)

	// kitchen.MakeFood("Burger")

	// payment.Pay(250)
	// delivery.Deliver("Sylhet")
	food.MakeFood("pizza")
	delivery.Deliver("sylhet Bangladesh")
	// Init()

}
