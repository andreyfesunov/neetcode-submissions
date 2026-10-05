var r = regexp.MustCompile(`[\d]{10}[MFO]([\d]{2})[\d]{2}`)

func countSeniors(details []string) int {
    c := 0

	for _, v := range details {
		match := r.FindStringSubmatch(v)
		if match != nil {
			d, _ := strconv.Atoi(match[1])
			if d > 60 {
				c++
			}
		}
	}

	return c
}