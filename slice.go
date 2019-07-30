package cosmos

import "math/rand"

//Delete with origin Order
func DeleteIndex(a interface{}, index int) interface{} {
	switch a.(type) {
	case []int:
		return DeleteIndexInt(a.([]int), index)
	default:
		panic("not support type")
	}
}

//Delete Without presevering Order
func DeleteIndexWithoutOrder(a interface{}, index int) interface{} {
	switch a.(type) {
	case []int:
		return DeleteIndexWithoutOrderInt(a.([]int), index)
	default:
		panic("not support type")
	}
}

//Reverse
func Reverse(a interface{}) interface{} {
	switch a.(type) {
	case []int:
		return ReverseInt(a.([]int))
	default:
		panic("not support type")
	}
}

// Shuffle slice
func Shuffle(a interface{}) {
	switch a.(type) {
	case []int:
		ShuffleInt(a.([]int))
	default:
		panic("not support type")
	}
}

// Del First value of slice
func SliceDelFirstVal(a interface{}, val interface{}) interface{} {
	switch a.(type) {
	case []int:
		return SliceDelFirstValInt(a.([]int), val.(int))
	default:
		panic("not support type")
	}
}

func ShuffleInt(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
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

func ReverseInt(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

func SliceDelFirstValInt(a []int, val int) []int {
	for i, v := range a {
		if v == val {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	return a
}
