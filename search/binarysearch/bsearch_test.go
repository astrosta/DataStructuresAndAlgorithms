package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []struct{
		ia []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5,6,7,8},5,4},
		{[]int{1,2,3,4,5,6,7,8},10,-1},
		{[]int{1,2,3,4,5,6,7,8},-2,-1},
	}

	for _, d := range data{
		if index := BinarySearch(d.ia, d.value); index != d.index{
			t.Errorf("got:%d, want:%d", index, d.index)
		}
	}

	return
}


func TestBinarySearchRepeatFrist(t *testing.T) {
	data := []struct{
		ia []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5,5,5,6,7,8},5,4},
		{[]int{5,5,5,5,5,5,5},5,0},
		{[]int{6,6,6,6,6,6,6},5,-1},
	}

	for _, d := range data{
		if index := BinarySearchRepeatFrist(d.ia, d.value); index != d.index{
			t.Errorf("got:%d, want:%d", index, d.index)
		}
	}

	return
}

func TestBinarySearchRepeatLast(t *testing.T) {
	data := []struct{
		ia []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5,5,5,6,7,8},5,6},
		{[]int{5,5,5,5,5,5,5},5,6},
		{[]int{5,6,6,6,6,6,6},5,0},
		{[]int{6,6,6,6,6,6,6},5,-1},
		{[]int{4,4,4,4,4,4,4},5,-1},
	}

	for _, d := range data{
		if index := BinarySearchRepeatLast(d.ia, d.value); index != d.index{
			t.Errorf("got:%d, want:%d", index, d.index)
		}
	}

	return
}

func TestBinarySearchRepeatFristBigEqual(t *testing.T) {
	data := []struct{
		ia []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5,5,5,6,7,8},5,4},
		{[]int{1,2,3,4,5,5,5,6,7,8},6,7},
		{[]int{5,5,5,5,5,5,5},5,0},
		{[]int{6,6,6,6,6,6,6},5,0},
		{[]int{4,4,4,4,4,4,4},5,-1},
	}

	for _, d := range data{
		if index := BinarySearchFirstBigEqual(d.ia, d.value); index != d.index{
			t.Errorf("got:%d, want:%d", index, d.index)
		}
	}

	return
}

func TestBinarySearchLastSmallEqual(t *testing.T) {
	data := []struct{
		ia []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5,5,5,6,7,8},4,3},
		{[]int{1,2,3,4,5,5,5,6,7,8},6,7},
		{[]int{5,5,5,5,5,5,5},5,6},
		{[]int{6,6,6,6,6,6,6},5,-1},
		{[]int{4,4,4,4,4,4,4},5,6},
	}

	for _, d := range data{
		if index := BinarySearchLastSmallEqual(d.ia, d.value); index != d.index{
			t.Errorf("got:%d, want:%d", index, d.index)
		}
	}

	return
}

