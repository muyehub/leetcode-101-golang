/**
 * 题目描述：
 * 有一群孩子和一堆饼干，每个孩子有一个饥饿度，每个饼干有一个大小。
 * 每个孩子只能吃一个饼干，且只有饼干的大小不小于孩子的饥饿度时，这个孩子才能吃饱。
 * 求解最多有多少孩子可以吃饱。
 *
 * 输入输出样例：
 * 输入两个数组，分别代表孩子的饥饿度和饼干的大小。输出最多有多少孩子可以吃饱的数量。
 * Input：[1, 2], [1, 2, 3]
 * Output：2
 * 在这个样例中，我们可以给两个孩子喂[1, 2], [1, 3], [2, 3] 这三种组合的任意一种。
 */
package main

import (
	"sort"
)

func _455(children, cookies []int) int {
	// 首先对孩子和饼干进行由小到大的排序
	sort.Ints(children)
	sort.Ints(cookies)
	var (
		child  int
		cookie int
	)
	for {
		if child >= len(children) || cookie >= len(cookies) {
			break
		}
		if children[child] <= cookies[cookie] {
			child++
		}
		cookie++
	}
	println(child)
	return child
}
