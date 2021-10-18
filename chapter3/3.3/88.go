package main

func _88(nums1 []int, m int, nums2 []int, n int)  {
	res := make([]int, 0, m+n)
	mPoint, nPoint := 0, 0

	for {
		if (mPoint == m) {
			res = append(res, nums2[nPoint:]...)
			break
		}
		if (nPoint == n) {
			res = append(res, nums1[mPoint:]...)
			break
		}
		if nums1[mPoint] < nums2[nPoint] {
			res = append(res, nums1[mPoint])
			mPoint++
		} else {
			res = append(res, nums2[nPoint])
			nPoint++
		}
	}

	copy(nums1, res)
}

// _88_2 从后向前不开辟额外空间的方法
func _88_2(nums1 []int, m int, nums2 []int, n int)  {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
