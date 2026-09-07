package looping

import "fmt"

func Window(n int) error {
	if n <= 0 {
		return fmt.Errorf(" ")
	}

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == 1 || x == n || y == 1 || y == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return nil
}
