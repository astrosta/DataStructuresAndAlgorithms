package binarysearch

/**
 *@title    	BinarySearchRepeatFrist
 *@description	查找第一个等于给定值的元素
 *@auth   		astrosta
 *@param
	[]int	有序数组
	int 	待查找的值
 *@return
*/
func BinarySearchRepeatFrist(ia []int, value int)int{
	var low int = 0
	var high int = len(ia)-1
	var mid int = low + ((high-low) >> 1)

	for low <= high{
		if ia[mid] >= value{
			high = mid -1

		}else{
			low = mid + 1
		}

		mid = low + ((high-low) >> 1)
	}

	if low <  len(ia) && ia[low] == value{
		return low
	}else{
		return -1
	}
}

/**
 *@title    	BinarySearchRepeatLast
 *@description	查找最后一个等于给定值的元素
 *@auth   		astrosta
 *@param
	[]int	有序数组
	int 	待查找的值
 *@return
*/
func BinarySearchRepeatLast(ia []int, value int)int{
	var low int = 0
	var high int = len(ia)-1
	var mid int = low + ((high-low) >> 1)

	for low <= high{
		if ia[mid] > value{
			high = mid - 1
		}else{
			low = mid +1
		}

		mid = low + ((high-low) >> 1)
	}

	if high >=0 && ia[high] == value{
		return high
	}else{
		return -1
	}
}

/**
 *@title    	BinarySearchFirstBigEqual
 *@description	查找第一个大于等于给定值的元素
 *@auth   		astrosta
 *@param
	[]int	有序数组
	int 	待查找的值
 *@return
*/
func BinarySearchFirstBigEqual(ia []int, value int)int{
	var low int = 0
	var high int = len(ia)-1
	var mid int = low + ((high-low) >> 1)

	for low <= high{
		if ia[mid] < value{
			low = mid +1
		}else{
			high = mid - 1
		}

		mid = low + ((high-low) >> 1)
	}

	if low < len(ia) && ia[low] >= value{
		return low
	}else{
		return -1
	}
}

/**
 *@title    	BinarySearchLastSmallEqual
 *@description	查找最后一个小于等于给定值的元素
 *@auth   		astrosta
 *@param
	[]int	有序数组
	int 	待查找的值
 *@return
*/
func BinarySearchLastSmallEqual(ia []int, value int)int{
	var low int = 0
	var high int = len(ia)-1
	var mid int = low + ((high-low) >> 1)

	for low <= high{
		if ia[mid] > value{
			high = mid - 1
		}else{
			low = mid +1
		}

		mid = low + ((high-low) >> 1)
	}

	if high >=0 && ia[high] <= value{
		return high
	}else{
		return -1
	}
}

