package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList_Put(t *testing.T) {
	list := NewSkipList()

	val := []byte("test_val")

	list.Put([]byte("a"), val)
	list.Put([]byte("b"), val)
	list.Put([]byte("d"), val)
	list.Put([]byte("c"), val)
}

func TestSkipList_Get(t *testing.T) {
	list := NewSkipList()

	list.Put([]byte("a"), []byte("test_val_a"))
	list.Put([]byte("b"), []byte("test_val_b"))
	list.Put([]byte("d"), []byte("test_val_d"))
	list.Put([]byte("c"), []byte("test_val_c"))

	fmt.Printf("%s\n", list.Get([]byte("a")).value)
	fmt.Printf("%s\n", list.Get([]byte("b")).value)
	fmt.Printf("%s\n", list.Get([]byte("d")).value)
	fmt.Printf("%s\n", list.Get([]byte("c")).value)

}

func TestSkipList_Remove(t *testing.T) {
	list := NewSkipList()

	list.Put([]byte("a"), []byte("test_val_a"))
	list.Put([]byte("b"), []byte("test_val_b"))
	list.Put([]byte("d"), []byte("test_val_d"))
	list.Put([]byte("c"), []byte("test_val_c"))

	list.Remove([]byte("c"))

}

func TestSkipList_Foreach(t *testing.T) {
	list := NewSkipList()

	list.Put([]byte("a"), []byte("test_val_a"))
	list.Put([]byte("b"), []byte("test_val_b"))
	list.Put([]byte("d"), []byte("test_val_d"))
	list.Put([]byte("c"), []byte("test_val_c"))

	list.Foreach()

}
