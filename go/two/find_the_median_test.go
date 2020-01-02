package two

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 3}
	r := FindMedianSortedArrays(nums1, nums2)
	if r != float64(1.5) {
		t.Errorf("Got %f, expected 1.5.", r)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	r = FindMedianSortedArrays(nums1, nums2)
	if r != float64(2.5) {
		t.Errorf("Got %f, expected 2.5.", r)
	}

	nums1 = []int{1}
	nums2 = []int{2, 3, 4, 5, 6}
	r = FindMedianSortedArrays(nums1, nums2)
	if r != float64(3.5) {
		t.Errorf("Got %f, expected 3.5.", r)
	}

	nums1 = []int{1, 4}
	nums2 = []int{2, 3, 5, 6}
	r = FindMedianSortedArrays(nums1, nums2)
	if r != float64(3.5) {
		t.Errorf("Got %f, expected 3.5.", r)
	}
}
