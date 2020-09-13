package mergeSort

func merge(sortedLeftArr []int32, sortedRightArr []int32) []int32 {
	resultArr := make([]int32, 0, len(sortedLeftArr)+len(sortedRightArr))
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(sortedLeftArr) && rightIndex < len(sortedRightArr) {
		if sortedLeftArr[leftIndex] <= sortedRightArr[rightIndex] {
			resultArr = append(resultArr, sortedLeftArr[leftIndex])
			leftIndex++
		} else {
			resultArr = append(resultArr, sortedRightArr[rightIndex])
			rightIndex++
		}
	}
	for leftIndex < len(sortedLeftArr) {
		resultArr = append(resultArr, sortedLeftArr[leftIndex])
		leftIndex++
	}
	for rightIndex < len(sortedRightArr) {
		resultArr = append(resultArr, sortedRightArr[rightIndex])
		rightIndex++
	}

	return resultArr
}

func MergeSort(arr []int32) []int32 {
	if len(arr) > 1 {
		middle := len(arr) / 2

		leftArr := arr[0:middle]
		rightArr := arr[middle:len(arr)]

		sortedLeft := MergeSort(leftArr)
		sortedRight := MergeSort(rightArr)

		return merge(sortedLeft, sortedRight)
	}

	return arr
}
