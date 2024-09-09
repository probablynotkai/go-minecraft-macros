package helpers

import "fmt"

var (
	SmallLocations map[int][2]int = map[int][2]int{}
)

func RegisterSmallInventoryLocations() {
	minX, maxX, minY := 2650, 3100, 340

	fmt.Println("Registering small inventory locations.")

	slotSize := ((maxX - minX) / 9) + 4

	for i := 0; i < 9; i++ {
		x := minX + (i * slotSize)

		for j := 0; j < 3; j++ {
			y := minY + (j * slotSize)
			SmallLocations[i+(9*j)] = [2]int{x, y}
		}
	}

	fmt.Println("Registered 3 location rows for small inventory.")
}
