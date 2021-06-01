package bubblesort

/**
 *@title    	BubbleSort
 *@description	冒泡排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func BubbleSort(ia []int){
	for i := 0; i < len(ia)-1; i++{
		for j := 0; j<len(ia)-i-1; j++{
			if ia[j] > ia[j+1]{
				ia[j], ia[j+1] = ia[j+1], ia[j]
			}
		}
	}

	return
}
