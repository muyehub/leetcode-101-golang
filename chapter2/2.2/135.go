/**
 * 题目描述：
 * 一群孩子站成一排，每一个孩子有自己的评分。现在需要给这些孩子发糖果，规则是如果一个孩子的评分比自己身旁的一个孩子要高，那么这个孩子就必须
 * 得到比身旁孩子更多的糖果；所有孩子至少要有一个糖果。求解最少需要多少个糖果。
 *
 * 输入输出样例：
 * 输入是一个数组，表示孩子的评分。输出是最少糖果的数量。
 * Input: [1, 0, 2]
 * Output: 5
 * 在这个样例中，最少的糖果分法是[2, 1, 2]。
 */
package main

func _135(ratings []int) int {
	// 计算参与评分的小朋友的个数，初始化糖果变量
	candy := make([]int, len(ratings))
	// 初始化总糖果数量
	sum := 0
	// 初始化糖果数量，每一个都为 1
	for i := 0; i < len(ratings); i++ {
		candy[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			candy[i+1] = candy[i] + 1
		}
	}

	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			if candy[i-1] <= candy[i]+1 {
				candy[i-1] = candy[i] + 1
			}
		}
	}

	for i := 0; i < len(candy); i++ {
		sum += candy[i]
	}

	println(sum)
	return sum
}
