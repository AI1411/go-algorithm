package search

func Binary(arr []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(arr) == 0 {
		return -1, ErrNotFound
	}
	mid := lowIndex + (highIndex-lowIndex)/2
	if arr[mid] > target {
		return Binary(arr, target, lowIndex, mid-1)
	} else if arr[mid] < target {
		return Binary(arr, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

func BinaryIterative(arr []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(arr) - 1
	var mid int
	for startIndex <= endIndex {
		mid = startIndex + (endIndex-startIndex)/2
		if arr[mid] > target {
			endIndex = mid - 1
		} else if arr[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}

func LowerBound(arr []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(arr) - 1
	var mid int
	for startIndex <= endIndex {
		mid = startIndex + (endIndex-startIndex)/2
		if arr[mid] < target {
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}

	if startIndex >= len(arr) {
		return -1, ErrNotFound
	}
	return startIndex, nil
}

func UpperBound(arr []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(arr) - 1
	var mid int
	for startIndex <= endIndex {
		mid = startIndex + (endIndex-startIndex)/2
		if arr[mid] > target {
			endIndex = mid - 1
		} else {
			startIndex = mid + 1
		}
	}

	if startIndex >= len(arr) {
		return -1, ErrNotFound
	}
	return startIndex, nil
}
