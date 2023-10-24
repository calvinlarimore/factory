package machine

type miner struct {
	buffer buffer
}

func (m *miner) Tick() {
	m.buffer.count++
}

func (m *miner) Extract() int {
	if m.buffer.count <= 0 {
		return 0
	}

	m.buffer.count--
	return m.buffer.item
}

func (m *miner) Insert(item int) bool {
	return false
}
