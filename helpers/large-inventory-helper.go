package helpers

import "fmt"

var (
	LargeLocations map[int][2]int = map[int][2]int{}
)

func RegisterLargeInventoryLocations() {
	minX, maxX, minY := 2650, 3100, 290

	fmt.Println("Registering large inventory locations.")

	slotSize := ((maxX - minX) / 9) + 4

	for i := 0; i < 9; i++ {
		x := minX + (i * slotSize)

		for j := 0; j < 5; j++ {
			y := minY + (j * slotSize)

			//fmt.Printf("%d x %d @ %d\n", x, y, i+(9*j))
			LargeLocations[i+(9*j)] = [2]int{x, y}
		}
	}

	LargeLocations[69] = [2]int{2650, 290}

	fmt.Println("Registered 5 location rows for large inventory.")
}
