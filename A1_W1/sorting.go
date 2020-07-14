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
