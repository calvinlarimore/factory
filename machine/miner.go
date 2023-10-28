package machine

type Miner struct {
	buffer buffer

	x, y int
}

func NewMiner(x, y int) *Miner {
	return &Miner{
		buffer: buffer{
			item: 0, count: 0,
		},
		x: x, y: y,
	}
}

func (m *Miner) Tick() {
	m.buffer.count++
}

func (m *Miner) Extract() int {
	if m.buffer.count <= 0 {
		return 0
	}

	m.buffer.count--
	return m.buffer.item
}

func (m *Miner) Insert(item int) bool {
	return false
}

func (m *Miner) Pos() (int, int) {
	return m.x, m.y
}

func (m *Miner) Sprite() string {
	s := "===" + "\n"
	s += "==="

	return s
}
