package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//冒泡排序
func bubbleSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	for i := 0; i < arrayLen-1; i++ {
		for j := 0; j < arrayLen-1-i; j++ {
			sum++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//冒泡排序优化
func bubbleSortOptimization(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	flag := 0
	arrayLen := len(array)
	k := arrayLen - 1
	pos := 0

	for i := 0; i < arrayLen-1; i++ {
		flag = 1
		for j := 0; j < k; j++ {
			sum++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = 0
				pos = j
			}
		}
		if flag == 1 {
			break
		}
		k = pos
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//冒泡排序优化：双向冒泡排序
func bilateralBubbleSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	flag := 0
	arrayLen := len(array)
	cyclesMax := arrayLen - 1
	cyclesMin := 0
	posMax := arrayLen - 1
	posMin := 0

	for i := 0; i < arrayLen-1; i++ {
		flag = 1
		for j := posMin; j < cyclesMax; j++ {
			sum++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = 0
				posMax = j
			}
		}
		if flag == 1 {
			break
		}
		cyclesMax = posMax

		flag = 1
		for k := posMax; k > cyclesMin; k-- {
			sum++
			if array[k] < array[k-1] {
				array[k], array[k-1] = array[k-1], array[k]
				flag = 0
				posMin = k
			}
		}
		if flag == 1 {
			break
		}
		cyclesMin = posMin
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//选择排序
func selectSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	for i := 0; i < arrayLen-1; i++ {
		for j := i + 1; j < arrayLen; j++ {
			sum++
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//选择排序优化，双向选择排序
func selectSortOptimization(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	rightPos := arrayLen - 1
	leftPos := 0
	for leftPos < rightPos {
		for j := leftPos; j <= rightPos; j++ {
			sum++
			if array[leftPos] > array[j] {
				array[leftPos], array[j] = array[j], array[leftPos]
			}
		}
		leftPos++
		for k := rightPos; k >= leftPos; k-- {
			sum++
			if array[rightPos] < array[k] {
				array[rightPos], array[k] = array[k], array[rightPos]
			}
		}
		rightPos--
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//直接插入排序
func insertSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	var tmp, j, flag int
	for i := 1; i < arrayLen; i++ {
		tmp = array[i]
		flag = 0
		for j = i; j > 0 && tmp < array[j-1]; j-- {
			sum++
			array[j] = array[j-1]
			flag = 1
		}
		if flag == 1 {
			array[j] = tmp
		}
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//直接交换变量，更加耗时
// //直接插入排序
// func insertSort(array []int) (int, time.Duration) {
// 	startTime := time.Now()
// 	sum := 0
// 	arrayLen := len(array)
// 	for i := 1; i < arrayLen; i++ {
// 		for j := i; j > 0 && array[j] < array[j-1]; j-- {
// 			sum++
// 			array[j], array[j-1] = array[j-1], array[j]
// 		}
// 	}
// 	timeCost := time.Since(startTime)
// 	return sum, timeCost
// }

// //希尔排序
// func shellSort(array []int) (int, time.Duration) {
// 	startTime := time.Now()
// 	sum := 0
// 	arrayLen := len(array)
// 	for i := arrayLen / 2; i > 0; i /= 2 {
// 		for j := i; j < arrayLen; j += i {
// 			for k := j; k > 0 && array[k] < array[k-i]; k -= i {
// 				sum++
// 				array[k], array[k-i] = array[k-i], array[k]
// 			}
// 		}
// 	}
// 	timeCost := time.Since(startTime)
// 	return sum, timeCost
// }

//希尔排序
func shellSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	var tmp, k, flag int
	for i := arrayLen / 2; i > 0; i /= 2 {
		for j := i; j < arrayLen; j += i {
			tmp = array[j]
			flag = 0
			for k = j; k > 0 && tmp < array[k-i]; k -= i {
				sum++
				array[k] = array[k-i]
				flag = 1
			}
			if flag == 1 {
				array[k] = tmp
			}
		}
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//堆排序
func heapSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)

	for i := arrayLen - 1; i > 0; i-- {
		for j := (i - 1) / 2; j >= 0; j-- {
			sum++
			if array[j] < array[j*2+1] {
				array[j], array[j*2+1] = array[j*2+1], array[j]
			}
			if j*2+2 <= i && array[j] < array[j*2+2] {
				array[j], array[j*2+2] = array[j*2+2], array[j]
			}
		}
		array[i], array[0] = array[0], array[i]
	}

	timeCost := time.Since(startTime)
	return sum, timeCost
}

//快速排序
func quickSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	tmp := array[0]
	i, j := 0, arrayLen-1
	for i != j {
		sum++
		if array[j] > tmp {
			j--
			continue
		}

		if array[i] <= tmp {
			i++
			continue
		}
		array[i], array[j] = array[j], array[i]
	}
	array[0], array[i] = array[i], array[0]
	if j > 0 {
		quickSort(array[:j])
	}
	if arrayLen-1-j > 0 {
		quickSort(array[j+1:])
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//归并排序
func mergerSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	arrayLen := len(array)
	tmp := make([]int, arrayLen)
	if arrayLen > 2 {
		mergerSort(array[:arrayLen/2])
		mergerSort(array[arrayLen/2:])
	}
	for i, k, j := 0, 0, arrayLen/2; i < arrayLen; i++ {
		sum++
		switch {
		case k >= arrayLen/2:
			tmp[i] = array[j]
			j++
		case j >= arrayLen:
			tmp[i] = array[k]
			k++
		case array[k] <= array[j]:
			tmp[i] = array[k]
			k++
		case array[k] > array[j]:
			tmp[i] = array[j]
			j++
		}
	}
	for l := 0; l < arrayLen; l++ {
		sum++
		array[l] = tmp[l]
	}
	timeCost := time.Since(startTime)
	return sum, timeCost
}

//基数排序
func baseSort(array []int) (int, time.Duration) {
	startTime := time.Now()
	sum := 0

	timeCost := time.Since(startTime)
	return sum, timeCost
}

//有序程度检查
func orderedCheck(array []int) string {
	t, f := 0, 0
	var result string
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			t = 1
		}
		if array[i] < array[i+1] {
			f = 1
		}
	}

	switch {
	case t == 0 && f == 1:
		result = "正序"
	case t == 1 && f == 0:
		result = "反序"
	default:
		result = "无序"
	}
	return result
}

func main() {
	var srcArray [10000]int

	for i := 0; i < len(srcArray); i++ {
		rand.Seed(time.Now().UnixNano() * int64(i))
		srcArray[i] = rand.Intn(10000)
	}

	fmt.Println("Start>>>>>>>")
	fmt.Println("srcArray: [", srcArray[0], srcArray[1], srcArray[2], "....", srcArray[9998], srcArray[9999], "]", "\t有序验证: ", orderedCheck(srcArray[:]))

	opArray := srcArray
	startTime := time.Now()
	sort.Ints(opArray[:])
	ctimes := time.Since(startTime)
	sum := 0
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles:  nil", "\t\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: systeamSort")

	opArray = srcArray
	sum, ctimes = bubbleSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: BubbleSort")

	opArray = srcArray
	sum, ctimes = bubbleSortOptimization(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: BubbleSortOptimization")

	opArray = srcArray
	sum, ctimes = bilateralBubbleSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: bilateralBubbleSort")

	opArray = srcArray
	sum, ctimes = selectSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: selectSort")

	opArray = srcArray
	sum, ctimes = selectSortOptimization(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: selectSortOptimization")

	opArray = srcArray
	sum, ctimes = insertSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: insertSort")

	opArray = srcArray
	sum, ctimes = shellSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: shellSort")

	opArray = srcArray
	sum, ctimes = heapSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: heapSort")

	opArray = srcArray
	sum, ctimes = quickSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\t\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: quickSort")

	opArray = srcArray
	sum, ctimes = mergerSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\t\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: mergerSort")

	opArray = srcArray
	sum, ctimes = baseSort(opArray[:])
	fmt.Print("recArray: [", opArray[0], opArray[1], opArray[2], "....", opArray[9998], opArray[9999], "]")
	fmt.Println("\tCycles: ", sum, "\t\tTime consuming: ", ctimes, "\t有序验证: ", orderedCheck(opArray[:]), "\tAlgorithm: baseSort")
}
