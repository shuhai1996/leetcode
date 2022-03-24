package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1}
	nums2 := []int{2,3,4,5,6}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var k int
	var result float64
	if (len(nums1) + len(nums2))%2 == 0 { //为偶数时
		k = (len(nums1) + len(nums2))/2
		result = float64(getKNum(nums1, nums2, k) + getKNum(nums1, nums2, k+1)) /2
	} else { //为奇数时
		k = (len(nums1) + len(nums2))/2
		result = float64(getKNum(nums1, nums2, k+1))
	}
	return result
}

/** 二分查找 求第k小的数*/

func getKNum(nums1 []int, nums2 []int, k int) int {
	var index int
	var del int
	if k <=1 {
		index = 0
	} else {
		index = k/2 -1
	}
	fmt.Println(k)
	if k>=1 && len(nums1) ==0 && len(nums2) >0{
		return nums2[k-1]
	}
	if k>=1 && len(nums2) ==0 && len(nums1) >0{
		return nums1[k-1]
	}
	if len(nums1)>0 && len(nums2)>0 {
		if k>1 {
			var cn,cm int
			newN,newM := index,index
			if (index+1) < len(nums1) {
				cn = nums1[index]
			} else {
				cn = nums1[len(nums1)-1]
				newN = len(nums1)-1
			}
			if (index+1) < len(nums2) {
				cm = nums2[index]
			} else {
				cm = nums2[len(nums2)-1]
				newM = len(nums2)-1
			}
			if cn <= cm{
				nums1 = nums1[newN+1:] //删除前index+1个元素
				del = newN+1
			}else {
				nums2 = nums2[newM+1:] //删除前index+1个元素
				del = newM+1
			}
		} else {
			if nums1[0] < nums2[0]{
				return nums1[0]
			}
			return nums2[0]
		}
	}
	k = k-del
	return getKNum(nums1, nums2, k)
}


/** 归并排序 */
func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	for {
		if len(nums1)>0 && len(nums2)>0 {
			if nums1[0] <= nums2[0] {
				nums3 = append(nums3, nums1[0])
				nums1 = nums1[1:] //删除第一个元素
			} else {
				nums3 = append(nums3, nums2[0])
				nums2 = nums2[1:] //删除第一个元素
			}
		}

		if len(nums1) ==0 {
			nums3 = append(nums3, nums2...)
			break
		}
		if len(nums2) ==0 {
			nums3 = append(nums3, nums1...)
			break
		}
	}
	index := len(nums3)/2
	var result float64
	if len(nums3)>0 {
		if index>=1 && len(nums3)%2 == 0 {
			result = float64(nums3[index-1] + nums3[index])/2
		} else {
			result = float64(nums3[index])
		}
	}

	return result
}