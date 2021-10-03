package main

import (
	".main.go/assemblyspot"
	".main.go/factory"
)

const carsAmount = 100

func main() {
	factory := factory.New()
	vehiclesAssembled := make(chan *assemblyspot.AssemblySpot, carsAmount)
	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	factory.StartAssemblingProcess(carsAmount, vehiclesAssembled)
}
