import (
	"reflect"
)

type sortableBytes []byte

func (s sortableBytes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortableBytes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortableBytes) Len() int {
	return len(s)
}

func isAnagram(s string, t string) bool {
	sb, tb := sortableBytes(s), sortableBytes(t)

	sort.Sort(sb)
	sort.Sort(tb)

	return reflect.DeepEqual(sb, tb)
}
