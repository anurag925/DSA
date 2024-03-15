package dsa

func productExceptSelf(nums []int) []int {
	product := 1
	hasZero := false
	allZero := 0
	for _, i := range nums {
		// fmt.Println(i)
		if i == 0 {
			hasZero = true
			allZero++
		} else {
			product *= i
		}
		// fmt.Println("p ", product)
	}

	// fmt.Println("hola ", hasZero, allZero, product, nums)
	if allZero >= 2 {
		product = 0
	}

	res := make([]int, len(nums))
	for i, e := range nums {
		// fmt.Println("hola ", hasZero, allZero, product, nums, i, e)
		if hasZero {
			if e != 0 {
				res[i] = 0
			} else {
				// fmt.Println("hola ",2, i, product)
				res[i] = product
				// fmt.Println("hola ",2, i, product, res[i])
			}
		} else {
			res[i] = product / e
		}
	}

	return res
}
