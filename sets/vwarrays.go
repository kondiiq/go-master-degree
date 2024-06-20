package sets

import "knapsackProblem/handler"

//No. Items is 15

func Value15() []int {
	return []int{135, 139, 149, 150, 156, 163, 173, 184, 192, 201, 210, 214, 221, 229, 240}
}

func Weight15() []int {
	return []int{70, 73, 77, 80, 82, 87, 90, 94, 98, 106, 110, 113, 115, 118, 120}
}

func MaxCapacity15() int {
	return 750
}

func Penalty15() int {
	return 1500
}

func Temp15() float32 {
	return 1000.0
}


//No. Items is 50
func Value50(sPath string) {
	handler.ReadFile(sPath)
}

func Weight50(sPath string) {
	handler.ReadFile(sPath)
}

func MaxCapacity50() int {
	return 750
}

func Penalty50() int {
	return 1500
}

func Temp50() float32 {
	return 1000.0
}


//No. Items is 100


