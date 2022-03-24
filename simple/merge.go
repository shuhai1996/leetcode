package main

// NO.88
func main() {
	merge([]int{1,2,3,4,5,6,7,0,0},7, []int{4,8},2)
}

//归并排序，双指针
func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums3:= make([]int, 0, m+n)
	l,r := 0,0
	for{
		if l==m {
			nums3 = append(nums3, nums2[r:]...)
			break
		}
		if r==n {
			nums3 = append(nums3, nums1[l:]...)
			break
		}

		if nums1[l]<nums2[r] {
			nums3 = append(nums3,nums1[l])
			l++
		} else {
			nums3 = append(nums3,nums2[r])
			r++
		}
	}
	copy(nums1, nums3)
}
