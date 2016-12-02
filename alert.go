package alert

type alert struct {
	cap int
	i   int
	buf []bool
}

func NewSlidingWindowAlert(cap int) *alert {
	a := alert{}
	a.i = 0
	a.cap = cap
	a.buf = make([]bool, cap)
	return &a
}

func (a *alert) Append(b bool) {
	a.i = (a.i + 1) % a.cap
	a.buf[a.i] = b
}

func (a *alert) AllTrue() bool {
	for _, v := range a.buf {
		if v == false {
			return false
		}
	}
	return true
}
