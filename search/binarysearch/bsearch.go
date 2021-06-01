package binarysearch

/**
 *@title    	BinarySearch
 *@description	二分查找
 *@auth   		astrosta
 *@param
	[]int	有序无重复数值数组
	int 	待查找的值
 *@return
*/
func BinarySearch(ia []int, value int) int{
	index := binarySearchInternally(ia, 0, len(ia)-1, value)
	return index
}

func binarySearchInternally(ia[]int, beg, end, value int) int{
	if beg > end{
		return -1
	}

	mid := beg + ((end -beg) >> 1)
	if ia[mid] == value{
		return mid
	}else if ia[mid] < value{
		return binarySearchInternally(ia, mid + 1, end, value)
	}else if ia[mid] > value{
		return binarySearchInternally(ia, beg, mid -1, value)
	}

	return 0
}
