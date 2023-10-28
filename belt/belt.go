package belt

type Belt struct {
	input  int
	output int

	item int

	Next *Belt
	Last *Belt

	x, y     int
	rotation int
}

func NewBelt(x, y, rot int) *Belt {
	return &Belt{
		input: 0, output: 0,
		item: 0,

		Next: nil, Last: nil,

		x: x, y: y,
		rotation: rot,
	}
}

func (b *Belt) Tick() {
	if b.Next == nil {
		return
	}

	if b.item != 0 && b.Next.input == 0 {
		b.output = b.item
		b.Next.input = b.item
		b.item = 0
	}
}

func (b *Belt) Validate() {
	if !(b.item != 0 && b.input != 0) {
		return
	}

	b.input = 0
	b.Last.item = b.Last.output
	b.Last.output = 0

	b.Last.Validate()
}

func (b *Belt) Flush() {
	if b.input != 0 {
		b.item = b.input
		b.input = 0
	}

	b.output = 0
}

func (b *Belt) Item() int {
	return b.item
}

func (b *Belt) Pos() (int, int) {
	return b.x, b.y
}

func (b *Belt) Rotation() int {
	return b.rotation
}

func (b *Belt) Extract() int {
	item := b.item

	if b.item != 0 {
		b.item = 0
	}

	return item
}

func (b *Belt) Insert(item int) bool {
	if b.item != 0 {
		return false
	}

	b.item = item
	return true
}

func (b *Belt) TempSetItemDebugDeleteThisFunction(item int) {
	b.item = item
}
