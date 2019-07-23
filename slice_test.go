package cosmos

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	a := []int{1, 2, 3}
	a = DeleteIndexWithoutOrder(a, 1).([]int)
	if !reflect.DeepEqual(a, []int{1, 3}) {
		t.Errorf("TestDeleteIndexWithouOrder failed, %v", a)
	}

	a = []int{1, 2, 3}
	a = DeleteIndex(a, 2).([]int)
	if !reflect.DeepEqual(a, []int{1, 2}) {
		t.Errorf("TestDeleteIndexWithouOrder failed, %v", a)
	}
}
