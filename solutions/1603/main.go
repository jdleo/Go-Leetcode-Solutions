package main

type ParkingSystem struct {
	slots []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	// initialize slots be capacities passed in
	return ParkingSystem{[]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	// decrease capacity of car type
	this.slots[carType-1]--
	// check if there is still some capacity
	return this.slots[carType-1] >= 0
}
