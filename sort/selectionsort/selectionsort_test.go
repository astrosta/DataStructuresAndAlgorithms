package selectionsort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSelectionSortCorrectness(t *testing.T) {
	ia := []int{3,1,6,7,3,9,2,5}
	SelectionSort(ia)
	ia1 := []int{1,2,3,3,5,6,7,9}

	fmt.Println(reflect.DeepEqual(ia, ia1))
}

func TestSelectionSortPerformance(t *testing.T) {
	ia := make([]int, 100000)
	for i:= 0; i<100000;i++{
		ia[i] = rand.Int()
	}

	beg:=time.Now()
	SelectionSort(ia)
	fmt.Println("selectionsort takes ", time.Since(beg))
}

