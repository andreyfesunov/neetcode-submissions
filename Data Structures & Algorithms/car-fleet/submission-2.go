type car struct {
	speed, pos int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	s := make([]float64, 0)
	c := make([]car, 0, n)

	for i := 0; i < n; i++ {
		c = append(c, car{
			speed: speed[i],
			pos: position[i],
		})
	}

	sort.Slice(c, func(i,j int) bool {
		return c[i].pos < c[j].pos
	})

	for i := 0; i < n; i++ {
		dist := target - c[i].pos
		time := float64(dist) / float64(c[i].speed)

		for len(s) != 0 && (math.Abs(time - s[len(s) - 1]) <= 1e-9 || time > s[len(s) - 1])  {
			s = s[:len(s) - 1]
		}

		s = append(s, time)
	}

	return len(s)
}
