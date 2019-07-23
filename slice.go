package cosmos

//Delete with origin Order
func DeleteIndex(a interface{}, index int) interface{} {
	switch a.(type) {
	case []int:
		return DeleteIndexInt(a.([]int), index)
	default:
		return a
	}
}

//Delete Without presevering Order
func DeleteIndexWithoutOrder(a interface{}, index int) interface{} {
	switch a.(type) {
	case []int:
		return DeleteIndexWithoutOrderInt(a.([]int), index)
	default:
		return a
	}
}

func DeleteIndexInt(a []int, index int) []int {
	a = append(a[:index], a[index+1:]...)
	return a
}

func DeleteIndexWithoutOrderInt(a []int, index int) []int {
	a[index] = a[len(a)-1]
	a = a[:len(a)-1]
	return a
}

func DeleteIndexInt32(a []int32, index int) []int32 {
	a = append(a[:index], a[index+1:]...)
	return a
}

func DeleteIndexWithoutOrderInt32(a []int32, index int) []int32 {
	a[index] = a[len(a)-1]
	a = a[:len(a)-1]
	return a
}
