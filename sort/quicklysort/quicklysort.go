package quicklysort

/**
 *@title    	QuicklySort
 *@description	快速排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func QuicklySort(ia []int) {
	if len(ia) < 2 {
		return
	}
	quicklySortC(ia, 0, len(ia)-1)
	return
}

func quicklySortC(ia []int, beg, end int) {
	if beg >= end {
		return
	}

	//获取分区点
	pivotIndex := partion(ia, beg, end)

	quicklySortC(ia, beg, pivotIndex-1)
	quicklySortC(ia, pivotIndex+1, end)

	return
}

func partion(ia []int, beg, end int) int {
	pivot := ia[end]
	i := beg

	for j := beg; j < end; j++ {
		if ia[j] < pivot {
			ia[i], ia[j] = ia[j], ia[i]
			i++
		}
	}

	ia[i], ia[end] = ia[end], ia[i]

	return i
}
