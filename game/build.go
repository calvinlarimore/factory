package game

import (
	"fmt"

	"github.com/calvinlarimore/factory/belt"
	"github.com/calvinlarimore/factory/machine"
)

func PlaceBelt(x, y, rot int) *belt.Belt {
	b := belt.NewBelt(x, y, rot)

	fx, fy := x, y

	switch b.Rotation() {
	case 0:
		fy--
	case 1:
		fx++
	case 2:
		fy++
	case 3:
		fx--
	}

	fmt.Printf("\nPlacing conveyor @ (%d, %d)\n", x, y)

	for _, other := range belts {
		if b.Next != nil && b.Last != nil {
			break
		}

		u, v := other.Pos()

		fmt.Printf("Checking conveyor @ (%d, %d)\n", u, v)

		if u == fx && v == fy {
			b.Next = other
			other.Last = b
			fmt.Println("\tFound next!")
			continue
		}

		if (u == x && v == y-1) || (u == x-1 && v == y) || (u == x && v == y+1) || (u == x+1 && v == y) {
			b.Last = other
			other.Next = b
			fmt.Println("\tFound last!")
		}
	}

	belts = append(belts, b)

	return b
}

func PlaceMiner(x, y int) *machine.Miner {
	m := machine.NewMiner(x, y)

	machines = append(machines, m)

	return m
}
