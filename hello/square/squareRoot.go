package square

import "fmt"

func SquareRoot(x float64) (z float64) {
	z = 1
	for i := 1; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}
