func wordPattern(pattern string, s string) bool {
    d := make(map[byte]string)
	c := make(map[string]byte)
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}
	for i := range words {
		key := pattern[i]
		v, vok := d[key]
		cc, cok := c[words[i]]
		if vok && cok && cc == key && v == words[i] || !vok && !cok {
			d[key] = words[i]
			c[words[i]] = key
		} else {
			return false
		}
	}
	return true
}