package two

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var left int = (m + n + 1) / 2
	var right int = (m + n + 2) / 2
	if left == right {
		return float64(findK(nums1, 0, nums2, 0, left))
	}

	return (float64(findK(nums1, 0, nums2, 0, left)) + float64(findK(nums1, 0, nums2, 0, right))) / 2

}

func findK(nums1 []int, i int, nums2 []int, j int, k int) int {
	fmt.Println(nums1, i, nums2, j, k)
	m := len(nums1)
	n := len(nums2)
	// 如果nums1已经到了最后一位
	if i >= m {
		return nums2[j+k-1]
	}
	// 如果nums2已经到了最后一位
	if j >= n {
		return nums1[i+k-1]
	}
	// 如果只余一位
	if k == 1 {
		if nums1[i] <= nums2[j] {
			return nums1[i]
		}

		return nums2[j]
	}

	key := (k / 2) - 1
	del1 := (k / 2)
	del2 := (k / 2)
	var mid1 int
	var mid2 int
	if i+key >= len(nums1) {
		mid1 = nums1[m-1]
		del1 = m - i
	} else {
		mid1 = nums1[i+key]
	}

	if j+key >= len(nums2) {
		mid2 = nums2[n-1]
		del2 = n - j
	} else {
		mid2 = nums2[j+key]
	}

	if mid1 < mid2 {
		return findK(nums1, i+(k/2), nums2, j, k-del1)
	}
	return findK(nums1, i, nums2, j+(k/2), k-del2)

}
