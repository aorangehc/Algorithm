package main

import "fmt"

// 快速排序
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := arr[right]
	i := left
	// 将所有小于pivot的元素放到左边
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 将基准元素放到对应位置
	arr[i], arr[right] = arr[right], arr[i]

	quickSort(arr, left, i-1)  // 遍历左侧分区
	quickSort(arr, i+1, right) // 遍历右侧分区
}

func quickSortArray(nums []int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)

	quickSort(arr, 0, len(arr)-1)

	return arr
}

// 堆排序
// 堆化函数：维护最大堆的性质
// arr - 待排序数组
// n - 当前堆的大小
// i - 要堆化的字数的根节点索引
func heapSort(arr []int, n, i int) {
	// 初始化最大值为当前节点
	// 左子节点索引
	// 右子节点索引
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 比较左子节点和当前最大值
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 比较右子节点和当前最大值
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，需要交换并递归堆化
	if largest != i {
		// 交换当前节点与较大子节点
		arr[i], arr[largest] = arr[largest], arr[i]
		// 递归化受影响的子树
		heapSort(arr, n, largest)
	}
}

// 堆排序入口
func heapSortArray(nums []int) []int {
	// 创建副本，避免修改原数组
	arr := make([]int, len(nums))
	copy(arr, nums)
	n := len(arr)

	// 构建堆阶段，构建最大堆
	//从最后一个非叶子节点开始，自底向上堆化
	// 最后一个非叶子节点的索引是：n / 2 - 1
	for i := n/2 - 1; i >= 0; i-- {
		heapSort(arr, n, i)
	}

	// 堆排序阶段，逐个提取最大值
	// 将堆顶元素（最大值）与当前堆位交换
	// 缩小堆的范围，重新堆化
	for i := n - 1; i > 0; i-- {
		// 将当前最大值移到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		// 对剩余元素重新堆化，堆大小缩小为i
		heapSort(arr, i, 0)
	}

	return arr
}

// 归并排序
func mergeSort(left, right []int) []int {
	mergedArray := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			mergedArray = append(mergedArray, left[i])
			i += 1
		} else {
			mergedArray = append(mergedArray, right[j])
			j += 1
		}
	}

	mergedArray = append(mergedArray, left[i:]...)
	mergedArray = append(mergedArray, right[j:]...)

	return mergedArray
}

func mergeSortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSortArray(nums[:mid])
	right := mergeSortArray(nums[mid:])
	return mergeSort(left, right)
}

func main() {
	testCases := [][]int{
		[]int{5, 2, 3, 1},
		[]int{5, 1, 1, 2, 0, 0},
	}

	for _, testCase := range testCases {
		quickArray := quickSortArray(testCase)
		fmt.Println(quickArray)

		heapArray := heapSortArray(testCase)
		fmt.Println(heapArray)

		mergeArray := mergeSortArray(testCase)
		fmt.Println(mergeArray)
	}
}
