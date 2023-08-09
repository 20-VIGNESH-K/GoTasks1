package main

import "fmt"

func main() {
	var num, temp, dig, n, sum, count, mul int64
	fmt.Scan(&num)
	temp = num
	sum = 0
	count = 0
	for temp != 0 {
		temp /= 10
		count += 1
	}
	temp = num
	if count >= 13 && count <= 16 {
		for temp != 0 {
			dig = temp % 100
			n = dig % 10
			mul = (dig / 10) * 2
			if mul > 9 {
				mul = mul - 9
			}
			sum = sum + mul + n
			temp /= 100
		}
		if sum%10 == 0 {
			fmt.Println("Valid card")
		} else {
			fmt.Println("Not valid card")
		}

	} else {
		fmt.Println("---Not valid card---")
	}
}
