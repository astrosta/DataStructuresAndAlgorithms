package selectionsort

/**
 *@title    	SelectionSort
 *@description	选择排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func SelectionSort(ia []int) {

	for i := 0; i < len(ia); i++ {
		smallest := i
		for j := i + 1; j < len(ia); j++ {
			if ia[j] < ia[smallest] {
				smallest = j
			}
		}

		ia[i], ia[smallest] = ia[smallest], ia[i]
	}

	return
}