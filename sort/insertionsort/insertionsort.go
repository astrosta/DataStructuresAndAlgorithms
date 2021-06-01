package insertionsort

/**
 *@title    	InsertionSort
 *@description	插入排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func InsertionSort(ia []int){

	for i := 1; i< len(ia); i++{
		value := ia[i]
		j := i -1
		for ; j >= 0; j--{
			if value  < ia[j]{
				ia[j+1] = ia[j]
			}else{
				break
			}
		}
		ia[j + 1] = value
	}
	return
}
