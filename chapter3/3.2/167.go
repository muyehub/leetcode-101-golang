package main

// _167 最直观解法
func _167(numbers []int, target int) []int {
	len := len(numbers)

	for i := 0; i < len - 1; i++ {
		for j := i + 1; j < len; j++ {
			if numbers[i] + numbers[j] == target {
				println(i+1, j+1)
				return []int{i+1, j+1}
			}
		}
	}

	return []int{}
}

// _167_2 双指针法
func _167_2(numbers []int, target int) []int {
	cnt := len(numbers)
	l := 0
	r := cnt - 1
	for {
		if l == r {
			break
		}
		if numbers[l] + numbers[r] == target {
			println(l+1, r+1)
			return []int{l+1, r+1}
		}
		if numbers[l] + numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	return []int{}
}