package piscine

func Swap(a *int, b *int) {
	A := *b
	B := *a

	*a = A
	*b = B
}
