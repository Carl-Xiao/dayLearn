package main

import (
	"fmt"
)

func partition(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}

//寻找第K大元素
func findKthLargest(nums []int, k int) int {
	len := len(nums)
	left := 0
	right := len - 1
	// 转换一下，第 k 大元素的索引是 len - k
	target := len - k
	for {
		index := Partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}
func Partition(nums []int, left, right int) int {
	//随机选择一个元素
	pivot := nums[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			// 小于 pivot 的元素都被交换到前面
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[j], nums[left] = nums[left], nums[j]
	return j
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
