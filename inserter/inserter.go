package inserter

type Insertable interface {
	Extract() int
	Insert(item int) bool
}

type Inserter struct {
	source      Insertable
	destination Insertable

	item int

	x, y     int
	rotation int
}

func NewInserter(src Insertable, dest Insertable, rot int) Inserter {
	return Inserter{
		source:      src,
		destination: dest,
		item:        0,
		rotation:    rot,
	}
}

func (i *Inserter) Tick() {
	if i.item == 0 {
		i.item = i.source.Extract()
	} else {
		if i.destination.Insert(i.item) {
			i.item = 0
		}
	}
}

func (i *Inserter) Item() int {
	return i.item
}

func (i *Inserter) Pos() (int, int) {
	return i.x, i.y
}

func (i *Inserter) Rotation() int {
	return i.rotation
}

func (i *Inserter) SetSource(src Insertable) {
	i.source = src
}

func (i *Inserter) SetDestination(dest Insertable) {
	i.destination = dest
}
