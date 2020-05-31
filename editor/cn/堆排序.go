package cn

// 堆排序
// 1） 构造堆 从最后一个非叶子节点开始构造堆
// 2) 交换数据，调整堆

func HeapSort(arr []int) []int {
	// 构造堆 从最后一个非叶子节点开始构造堆
	for i := len(arr) >> 1 - 1; i >= 0; i-- {
		adjustHeap(&arr, i, len(arr))
	}

	// 交换数据 调整堆
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(&arr, 0, j)
	}
	return arr
}

func adjustHeap(arr *[]int, i, len int) {
	tmp := (*arr)[i]
	for k := 2*i+1; k < len; k = 2*k+1 { // 从i节点的左子结点开始
		if k + 1 < len && (*arr)[k] < (*arr)[k+1] {
			k++
		}

		if (*arr)[k] > tmp {
			(*arr)[i] = (*arr)[k]
			i = k
		} else {
			break
		}
		(*arr)[i] = tmp
	}
 }