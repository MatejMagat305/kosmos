package help

type Position struct {
	X, Y float64
}


func SwapIfNeed(num1, num2 int) (int, int) {
	if num2 > num1 {
		return num2, num1
	}
	return num1, num2
}