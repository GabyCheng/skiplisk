package skiplist

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	maxLevel    int     = 3
	probability float64 = 1 / math.E
)

type (
	Node struct {
		next []*Element
	}

	Element struct {
		Node
		key   []byte
		value interface{}
	}

	SkipList struct {
		Node
		maxLevel       int
		Len            int
		randSource     rand.Source
		probability    float64
		probTable      []float64
		prevNodesCache []*Node
	}
)

// NewSkipList create a new skip list.
func NewSkipList() *SkipList {
	return &SkipList{
		Node:           Node{next: make([]*Element, maxLevel)},
		prevNodesCache: make([]*Node, maxLevel),
		maxLevel:       maxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    probability,
		probTable:      probabilityTable(probability, maxLevel),
	}
}

func (t *SkipList) Put(key []byte, value interface{}) *Element {
	var element *Element
	prev := t.backNodes(key)

	if element = prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		element.value = value
		return element
	}

	element = &Element{
		Node: Node{
			//随机一个 level，如果不随机，时间复杂度会退化成链表
			next: make([]*Element, t.randomLevel()),
		},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = prev[i].next[i]
		// 比如插入的key比之前的都要大，那么这里的perv[0].next[0] 则是nil
		// 比如插入的key是比之前小的，比如 c < d 那么这里 prev[i].next[i] 是d节点，然后相互替换位置
		prev[i].next[i] = element
	}

	t.Len++

	return element
}

func (t *SkipList) Get(key []byte) *Element {
	var prev = &t.Node
	var next *Element

	for i := t.maxLevel - 1; i >= 0; i-- {

		next = prev.next[i]
		// 当前层级找，找不到就下一层开始找，符合了下面的if 就从当前节点开始再继续遍历。
		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
	}

	if next != nil && bytes.Compare(key, next.key) <= 0 {
		return next
	}

	return nil
}

func (t *SkipList) Remove(key []byte) *Element {
	prev := t.backNodes(key)

	if element := prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		for k, v := range element.next {
			prev[k].next[k] = v
		}

		t.Len--
		return element
	}

	return nil
}

func (t *SkipList) Foreach() {
	for p := t.next[0]; p != nil; p = p.next[0] {
		fmt.Printf("%s\n", p.value)
	}
}

// 找到这个 key 前面的节点
func (t *SkipList) backNodes(key []byte) []*Node {
	var prev = &t.Node
	var next *Element

	prevs := t.prevNodesCache

	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		//这里很重要，如果插入的key比之前的大需要改变指针位置 &next.Node 然后插入到合适的位置
		for next != nil && bytes.Compare(key, next.key) > 0 {

			prev = &next.Node
			next = next.next[i]
		}

		prevs[i] = prev
	}

	return prevs
}
// 随机的 level，避免退化成链表
func (t *SkipList) randomLevel() (level int) {
	r := float64(t.randSource.Int63()) / (1 << 63)
	level = 1
	for level < t.maxLevel && r < t.probTable[level] {
		level++
	}
	return
}

func probabilityTable(probability float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}
