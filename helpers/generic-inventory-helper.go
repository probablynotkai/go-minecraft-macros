package helpers

import (
	"fmt"
)

// figure out how to calculate this dynamically
const (
	threeRowStartY = 340
	fiveRowStartY  = 290
)

type Inventory struct {
	MappedLocations map[int][2]int
}

func GetLocationsForSizedInventory(rows int) *Inventory {
	minX, maxX := 2650, 3100

	minY := -1
	switch rows {
	case 3:
		minY = threeRowStartY
	case 5:
		minY = fiveRowStartY
	}

	if minY == -1 {
		return nil
	}

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

	return &Inventory{locations}
}
