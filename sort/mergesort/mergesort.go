package mergesort

/**
 *@title    	MergeSort
 *@description	归并排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func MergeSort(ia []int) {
	if len(ia) < 2 {
		return
	}
	mergeSortC(ia, 0, len(ia)-1)
	return
}

func mergeSortC(ia []int, beg int, end int) {
	if beg >= end {
		return
	}

	middle := (beg + end) / 2
	mergeSortC(ia, beg, middle)
	mergeSortC(ia, middle+1, end)

	merge(ia, beg, middle, end)

	return

}

func merge(ia []int, beg, middle, end int) {
	tmp := make([]int, end-beg+1)
	i := beg
	j := middle + 1
	k := 0
	for i <= middle && j <= end {
		if ia[i] <= ia[j] {
			tmp[k] = ia[i]
			i++
		} else {
			tmp[k] = ia[j]
			j++
		}
		k++
	}

	if i == middle+1 {
		copy(tmp[k:], ia[j:end+1])
	} else {
		copy(tmp[k:], ia[i:middle+1])
	}

	copy(ia[beg:end+1], tmp)

	return
}

//func merge(ia1 []int, ia2 []int) []int{
//
//	//fmt.Println("ia1:",ia1)
//	//fmt.Println("ia2:",ia2)
//	i := 0
//	j := 0
//	k := 0
//	tmpArrayLen := len(ia1)+len(ia2)
//	tmpArray := make([]int, tmpArrayLen)
//
//	ia1 = append(ia1, int(^uint(0) >> 1))	//int最大值
//	ia2 = append(ia2, int(^uint(0) >> 1))	//int最大值
//	for ;k < tmpArrayLen; k++{
//		if ia1[i] <= ia2[j]{
//
//			tmpArray[k] = ia1[i]
//			i++
//		}else{
//			tmpArray[k] = ia2[j]
//			j++
//		}
//
//		//fmt.Println("k:",k)
//		//fmt.Println("i:",i)
//		//fmt.Println("j:",j)
//	}
//
//	//fmt.Println("tmpArray:",tmpArray)
//
//	return tmpArray
//}
