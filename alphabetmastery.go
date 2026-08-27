package sprint
func AlphabetMastery(n int) string {

	a := ""

	for i := 0; i < n; i++ {
		a += string(rune('a' + i))
	}
	return a
}
  