package po

import (
	"sync"
		"fmt"
	"container/list"
)

var ParticiplesMap sync.Map
var WaitGroup = sync.WaitGroup{}
var Queue = NewQueue()
var AlgoliasMap=map[string]Algolia{}

const N int = 10

type QueueNode struct {
	figure  int
	digits1 [N]int
	digits2 [N]int
	sflag   bool
	data    *list.List
}

var lock sync.Mutex

func NewQueue() *QueueNode {
	q := new(QueueNode)
	q.data = list.New()
	return q
}
func (q *QueueNode) Push(v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	q.data.PushFront(v)
}
func (q *QueueNode) Dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}
func (q *QueueNode) Pop() interface{} {
	defer lock.Unlock()
	lock.Lock()
	iter := q.data.Back()
	v := iter.Value
	q.data.Remove(iter)
	return v
}