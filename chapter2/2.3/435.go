/**
 * 题目描述
 * 给定多个区间，计算让这些区间互不重叠所需要移除区间的最少个数，起止相连不算重叠。
 *
 * 输入输出样例
 * 输入是一个数组，数组由多个长度固定为 2 的数组组成，表示区间的开始和结尾。输出一个整数，表示需要移除的区间数量。
 * Input: [[1, 2], [2, 4], [1, 3]]
 * Output: 1
 * 在这个样例中，我们可以移除区间 [1, 3]，使得剩余的区间 [[1, 2], [2, 4]] 互不重叠。
 */
package main

import (
	"fmt"
	"sort"
)

func _435(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	// 以右端点为准，进行从小到大的排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 将第一个元素的右端点作为基准，循环剩余元素
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	fmt.Println(n - ans)
	return n - ans
}
