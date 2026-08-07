type stackEntry struct {
	v, m int
}

type MinStack struct {
	s []stackEntry
	m int
}

func Constructor() MinStack {
	return MinStack {
		s: make([]stackEntry, 0),
		m: 1 << 31 - 1,
	}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, stackEntry{
		v: val,
		m: this.m,
	})
	this.m = min(this.m, val)
}

func (this *MinStack) Pop() {
	l := len(this.s) - 1
	r := this.s[l]
	this.s = this.s[:l]
	this.m = r.m
}

func (this *MinStack) Top() int {
	return this.s[len(this.s) - 1].v
}

func (this *MinStack) GetMin() int {
	return this.m
}
