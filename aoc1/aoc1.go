package aoc1

func FuelByMass(mass int) int {
	return mass / 3 -2
}

func recurseFuelWeigth(mass int) int {
	fuelMassToAdd := FuelByMass(mass)
	if fuelMassToAdd > 0 {
		return fuelMassToAdd + recurseFuelWeigth(fuelMassToAdd)
	}
	return 0
}

func FuelRequiredByMass(mass int) int{
	return recurseFuelWeigth(mass)
}

func AccumulateFuel() func(int) int {
	accumulator := 0
	return func(add int) int {
		accumulator += add
		return accumulator
	}
}
