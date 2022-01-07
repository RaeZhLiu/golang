package qsort

//QuickSort 快速排序算法
func quickSort(values []int, left, right int) {
	//挖坑-填数-分治
	//随机选择 基准值, 初始化 坑 的索引值
	baseValue := values[left]
	p := left
	//标记左右两个游标
	i, j := left, right
	// 左边游标的索引小于右边游标的索引时，做比较
	for i < j {
		//先判断右值，若小于基准值，则挪到左侧坑位
		//如果 右值 >= 基准值，则右游标一直左移
		for j >= p && values[j] >= baseValue {
			j--
		}
		//右游标所对应的值小于基准值，将此值赋值到 左侧坑位， 并把此值对应的索引设置为坑位
		if j >= p {
			values[p] = values[j]
			p = j
		}
		//再判断左值，若大于基准值，则挪到右侧坑位
		//如果 左值 <= 基准值，则左游标一直右移
		for i <= p && values[i] <= baseValue {
			i++
		}
		//左游标所对应的值大于基准值，将此值赋值到 右侧坑位， 并把此值对应的索引设置为坑位
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = baseValue
	//递归排序 left至p-1, p+1至right
	if p > left+1 {
		quickSort(values, left, p-1)
	}
	if p < right-1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
