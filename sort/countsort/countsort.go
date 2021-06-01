package countsort


/**
 *@title    	CountSort
 *@description	计数排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func CountSort(ia []int){
	if len(ia) <=1 {
		return
	}

	max := ia[0]

	for i := 1; i< len(ia); i++{
		if ia[i] >max{
			max = ia[i]
		}
	}

	countArray := make([]int, max+1)

	for i := 0; i < len(ia); i++{
		countArray[ia[i]]++
	}

	for i:= 1; i<max+1; i++{
		countArray[i] = countArray[i] + countArray[i-1]
	}

	tmpArray := make([]int, len(ia))
	for i:= len(ia)-1; i>=0;i--{
		index := countArray[ia[i]]
		countArray[ia[i]]--
		tmpArray[index-1] = ia[i]
	}

	copy(ia,tmpArray)

	return
}
