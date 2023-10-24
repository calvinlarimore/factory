package inserter

type Insertable interface {
	Extract() int
	Insert(item int) bool
}

type inserter struct {
	source      Insertable
	destination Insertable

	item int
}

func (i *inserter) Tick() {
	if i.item == 0 {
		i.item = i.source.Extract()
	} else {
		if i.destination.Insert(i.item) {
			i.item = 0
		}
	}
}
