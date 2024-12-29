package sliceDel


import (
	"testing"
	"reflect"
)

func TestConstructor(t *testing.T) {
	m := Constructor("hello", "you", "world")
	for _, item := range m.Slice {
		t.Logf("item: %+v\n", item)
	}
	if !reflect.DeepEqual(m.Slice, []string{"hello", "you", "world"}) {
		t.Fatal("err")
	}
}


func TestDeleteAtIndex(t *testing.T) {
	m := Constructor("hello", "you", "world")
	m.DeleteAtIndex(1)
	if !reflect.DeepEqual(m.Slice, []string{"hello", "world"}) {
		t.Fatal("err")
	}
}

