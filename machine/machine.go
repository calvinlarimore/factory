package machine

type Machine interface {
	Tick()
}

type buffer struct {
	item  int
	count int
}
