package bubblesort

/**
 *@title    	BubbleSort
 *@description	冒泡排序
 *@auth   		astrosta
 *@param
	[]int	待排序数组
 *@return
*/
func BubbleSort(ia []int) {
	if len(ia) <= 1 {
		return
	}

	for i := 0; i < len(ia)-1; i++ {
		//提前退出标志
		flag := false
		for j := 0; j < len(ia)-i-1; j++ {
			if ia[j] > ia[j+1] {
				ia[j], ia[j+1] = ia[j+1], ia[j]
				flag = true
			}
		}

		if !flag {
			//若无数据交换，说明数组已经是有序数组，可以直接推出
			break
		}
	}

	return
}
