package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*mod = a % b
	*div = a / b
}
