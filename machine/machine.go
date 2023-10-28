package machine

type Machine interface {
	Tick()
	Pos() (int, int)
	Sprite() string
}

type buffer struct {
	item  int
	count int
}
