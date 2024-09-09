package helpers

import "fmt"

type Inventory struct {
	MappedLocations map[int][2]int
}

func GetLocationsForSizedInventory(rows int) Inventory {
	minX, maxX, minY := 2650, 3100, 340

	fmt.Printf("Registering locations for %d row inventory.\n", rows)

	locations := map[int][2]int{}
	slotSize := ((maxX - minX) / 9) + 4

	for i := 0; i < 9; i++ {
		x := minX + (i * slotSize)

		for j := 0; j < rows; j++ {
			y := minY + (j * slotSize)
			locations[i+(9*j)] = [2]int{x, y}
		}
	}

	fmt.Printf("Registered %d rows for inventory.\n", rows)

	return Inventory{locations}
}
