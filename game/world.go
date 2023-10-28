package game

import (
	"github.com/calvinlarimore/factory/belt"
	"github.com/calvinlarimore/factory/inserter"
	"github.com/calvinlarimore/factory/machine"
)

var machines = []machine.Machine{}
var inserters = []*inserter.Inserter{}
var belts = []*belt.Belt{}

func InitWorld() {
	// belts = append(belts, belt.NewBelt(30, 2, 0))
	// belts = append(belts, belt.NewBelt(35, 2, 1))
	// belts = append(belts, belt.NewBelt(40, 2, 2))
	// belts = append(belts, belt.NewBelt(45, 2, 3))

	b := PlaceBelt(30, 5, 1)
	PlaceBelt(31, 5, 1)

	PlaceBelt(32, 5, 2)
	PlaceBelt(32, 6, 2)

	PlaceBelt(32, 7, 3)
	PlaceBelt(31, 7, 3)

	PlaceBelt(30, 7, 0)
	PlaceBelt(30, 6, 0)

	b.TempSetItemDebugDeleteThisFunction(69)

	PlaceMiner(40, 10)
}

func Tick() {

	for _, i := range inserters {
		i.Tick()
	}
	for _, m := range machines {
		m.Tick()
	}

	for _, b := range belts {
		b.Tick()
	}

	for _, b := range belts {
		b.Validate()
	}

	for _, b := range belts {
		b.Flush()
	}
}

func Belts() []*belt.Belt {
	return belts
}
