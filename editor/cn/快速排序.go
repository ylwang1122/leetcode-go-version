package cn

// 快速排序思想： 利用分治思想，选择一个基数，所有数据都和基数进行比较，将大于基数和小于基数
// 的数据进行分开，再对分开的两组数据进行处理

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	quickSort(&arr, 0, len(arr) - 1)
	return arr
}

func quickSort(arr *[]int, low, high int) {
	if low == high {
		return
	}

	// 获取基数的索引值
	index := getIndex(arr, low, high)
	quickSort(arr, low, index - 1)
	quickSort(arr, index + 1, high)
}

func getIndex(arr *[]int, low, high int) int {
	tmp := (*arr)[low]
	for low < high {
		for low < high && (*arr)[high] >= tmp {
			high--
		}
		(*arr)[low] = (*arr)[high]
		for low < high && (*arr)[low] <= tmp {
			low++
		}
		(*arr)[high] = (*arr)[low]
	}
	(*arr)[low] = tmp
	return low
}