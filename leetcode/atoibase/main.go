package main

import "fmt"

func main() {
	s := "123345"
	base := "2"
	var nb, pow int
		rst := 1
		if pow < 0 {
			 fmt.Print(0)
			 return
		}
		if pow == 0 {
			fmt.Print(1)
			return
		}
		for i := 0; i < pow; i++ {
			rst *= nb
		}
		fmt.Print(rst)
	}

	if len(base) < 2 {
		fmt.Println(0)
		return
	}

	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				return 0
			}
		}
	}

	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		j := 0
		for ; j < len(base); j++ {
			if s[i] == base[j] {
				break
			}
		}
		result += j * Power(len(base), len(s)-1-i)
	}
	fmt.Println(result)
}

func Power(nb, pow int) int {
	rst := 1
	if pow < 0 {
		return 0
	}
	if pow == 0 {
		return 1
	}
	for i := 0; i < pow; i++ {
		rst *= nb
	}
	return rst
}
