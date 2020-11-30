package constant1

import (
	"builder/po"
	"container/list"
	"fmt"
	"regexp"
	"sync"
)

var WaitGroup = sync.WaitGroup{}
var Queue = NewQueue()
var AlgoliasMap = map[string]po.Algolia{}
var CacheAlgoliasMap = map[string]po.Algolia{}
var Md5Map = po.NewConcurrentMap(make(map[string]interface{}))
var NeedArticleList = []*po.Article{}
var NeedAlgoliasList = []*po.Algolia{}
var ArticleMap = po.NewConcurrentMap(make(map[string]interface{}))
var StopArray = []string{}
var HtmlReg, _ = regexp.Compile("<.{0,200}?>")
var PointReg, _ = regexp.Compile("\n|\t|\r")
var NumberReg, _ = regexp.Compile("[0-9]+|[0-9]+\\.+[0-9]+")

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

func (q *QueueNode) Size() int {
	return q.data.Len()
}
