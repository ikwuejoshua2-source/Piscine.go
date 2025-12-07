package piscine 

func NRune (s string n int) rune {
	if n > len (s) || n < = o {
		return o
	}
	str := [] rune (s)
	return str [n-1]
}