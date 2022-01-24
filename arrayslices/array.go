package arrayslices

func Sum(numbers []int) (res int) {
	for _, val := range numbers {
		res += val
	}
	return
}

func SumAll(slices ...[]int) (res []int) {
	for _, v := range slices {
		res = append(res, Sum(v))
	}
	return
}

func SumAllTails(slices ...[]int) (res []int) {
	for _, v := range slices {
		if len(v) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(v[1:]))
		}
	}
	return
}
