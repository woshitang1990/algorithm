package two

import "fmt"

/**
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 你可以假设 nums1 和 nums2 不会同时为空。
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * 速度太慢了，还需要优化一下
 */

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var left = (m + n + 1) / 2
	var right = (m + n + 2) / 2
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
	del1 := k / 2
	del2 := k / 2
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
