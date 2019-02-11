package stringutil

func Reverse(s string) string {
	p := []rune(s)
	k := []rune(s)
	j := 0
	for i := len(p)-1; i >= 0; i = i-1 {
		k[j] = p[i]
		j = j+1
	}
	return string(k)
}