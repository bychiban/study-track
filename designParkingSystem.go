package main

import "fmt"

func main() {

	obj := Constructor(1, 2, 3)
	param_1 := obj.AddCar(1)
	fmt.Println(param_1)

}

type ParkingSystem struct {
	big, medium, small             int
	currBig, currMedium, currSmall int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small, 0, 0, 0}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.currBig < this.big {
			this.currBig++
			return true
		}
		return false
	case 2:
		if this.currMedium < this.medium {
			this.currMedium++
			return true
		}
		return false
	case 3:
		if this.currSmall < this.small {
			this.currSmall++
			return true
		}
		return false
	default:
		return false
	}
}
