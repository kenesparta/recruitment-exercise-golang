package main

import (
	".main.go/factory"
	".main.go/vehicle"
	"fmt"
)

const carsAmount = 100

func main() {
	factory := factory.New()
	vehiclesAssembled := make(chan vehicle.Car, carsAmount)
	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	factory.StartAssemblingProcess(carsAmount, vehiclesAssembled)

	// Reading the channel of the vehiclesAssembled
	for car := 0; car < carsAmount; car++ {
		v := <-vehiclesAssembled
		fmt.Printf("TestingLogs: %s \n", v.TestingLog)
		fmt.Printf("ID: %d => AssembledLogs: %s \n", v.Id, v.AssembleLog)
		fmt.Println("-----------")
	}
	defer close(vehiclesAssembled)
}
