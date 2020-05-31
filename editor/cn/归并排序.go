// 归并排序：分治 2路归并
// 1) 把长度为n的输入序列分为两个长度为n/2的子序列
// 2) 对这两个子序列进行归并排序
// 3) 将两个排好序的子序列合并成一个最终的排序序列

package cn

func MergeSort(arr *[]int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) >> 1
	MergeSort(arr, left, mid)
	MergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)
}

func merge(arr *[]int, left, mid, right int) {
	temp := make([]int, right - left + 1)
	i, j, k := left, mid, 0 // 左子序列的起始位置，右子序列的起始位置，temp的起始位置

	// 合并两个数组
	for i <= mid && j <= right {
		if (*arr)[i] > (*arr)[j] {
			temp[k] = (*arr)[j]
			j++
		} else {
			temp[k] = (*arr)[i]
			i++
		}
		k++
	}

	// 合并剩余数组
	for i <= mid {
		temp[k] = (*arr)[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = (*arr)[j]
		k++
		j++
	}

	for p := 0; p < len(temp); p++ {
		(*arr)[left + p] = temp[p]
	}
}