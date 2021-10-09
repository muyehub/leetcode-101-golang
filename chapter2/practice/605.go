package main

import "fmt"

func _605(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	// 按照种花的规则，最大数量计算方式为 cnt/2 + 1
	cnt := len(flowerbed)
	max := cnt/2 + 1
	if n > max {
		return false
	}

	for i := 0; i < cnt; i++ {
		// 如果当前值、上一个值、下一个值皆为 0，则当前值改为 1
		// 当特殊情况存在，如果是第一个值和最后一个值，则第一个值的下一个和当前值为 0，或者最后一个值和上一个值为 0，当前值改为 1
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == cnt-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		if n <= 0 {
			fmt.Println(n <= 0)
			return true
		}
	}
	fmt.Println(n <= 0)
	return n <= 0
}
