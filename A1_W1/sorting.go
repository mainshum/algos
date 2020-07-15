package A1_W1

func halveAr(xs []int) ([]int, []int) {
	sliceInd := len(xs) / 2
	return xs[0:sliceInd], xs[sliceInd:len(xs)]
}

func mergeArs(xs []int, ys []int) []int {
	var merged []int
	k, l := 0, 0
	for k < len(xs) {
		for l < len(ys) {
			if k >= len(xs) {
				merged = append(merged, ys[l])
				l++
				continue
			}
			if xs[k] <= ys[l] {
				merged = append(merged, xs[k])
				k++
				continue
			} else {
				merged = append(merged, ys[l])
				l++
				continue
			}
		}
		if k < len(xs) {
			merged = append(merged, xs[k])
			k++
		}
	}
	return merged
}

func MergeSort(xs []int) []int {
	var v1, v2 []int
	if len(xs) > 1 {
		// split phase
		ar1, ar2 := halveAr(xs)
		v1 = MergeSort(ar1)
		v2 = MergeSort(ar2)
	} else {
		return xs
	}

	return mergeArs(v1, v2)
}

func InsertionSort(xs []int) []int {
	for k := 1; k < len(xs); k++ {
		elemToMove := xs[k]
		j := k - 1
		for j >= 0 && elemToMove <= xs[j] {
			xs[j+1] = xs[j]
			j--
		}
		xs[j+1] = elemToMove
	}
	return xs
}

func BubbleSort(xs []int) []int {
	swapsMade := true
	for swapsMade == true {
		swapsMade = false
		for k := 0; k < len(xs)-1; k++ {
			if xs[k] > xs[k+1] {
				temp := xs[k]
				xs[k] = xs[k+1]
				xs[k+1] = temp
				swapsMade = true
			}
		}
	}
	return xs
}
