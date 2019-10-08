package pascal

func generate(row int, previous []int) (ret []int) {
	// firstly
	ret = append(ret, 1)
	// logic
	len := len(previous)
	for i := 0; i < len-1; i++ {
		ret = append(ret, previous[i]+previous[i+1])
	}
	// finaly
	ret = append(ret, 1)
	return
}

// Triangle a
func Triangle(size int) (ret [][]int) {
	// row 1 is fixed: [1]
	ret = append(ret, []int{1})
	if size > 1 {
		for i := 1; i < size; i++ {
			if i == 1 {
				// row 2 is fixed: [1, 1]
				ret = append(ret, []int{1, 1})
				continue
			}
			ret = append(ret, generate(i, ret[i-1]))
		}
	}
	return
}
