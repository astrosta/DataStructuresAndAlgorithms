package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

type skipListNode struct {
	//跳表保存的值
	value interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层节点指针
	forwards []*skipListNode
}

//新建跳表节点
func newSkipListNode(value interface{}, score, level int) *skipListNode {
	return &skipListNode{
		value:    value,
		score:    score,
		level:    level,
		forwards: make([]*skipListNode, level),
	}
}

type SkipList struct {
	head   *skipListNode //跳表头节点
	level  int           //跳表当前层数
	length int           //跳表长度
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   newSkipListNode(nil, math.MinInt32, MAX_LEVEL),
		level:  1,
		length: 0,
	}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) Insert(value interface{}, score int) error {
	if value == nil {
		return fmt.Errorf("param error")
	}

	//首先获取插入数据所在层数,本跳表理论上每隔3个节点取一个索引
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%3 == 1 {
			level++
		}
	}

	update := make([]*skipListNode, level)
	cur := sl.head
	//记录每层插入的前一个节点
	for i := level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].value == value {
				return fmt.Errorf("repeat value")
			}

			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}

			cur = cur.forwards[i]
		}

		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(value, score, level)

	//将跳表节点插入跳表中
	for i := 0; i < level; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	if level > sl.level {
		sl.level = level
	}

	sl.length++
	return nil
}

func (sl *SkipList) Find(value interface{}, score int) *skipListNode {
	cur := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].value == value {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

func (sl *SkipList) Delete(value interface{}, score int) error {
	if value == nil {
		return fmt.Errorf("value == nil")
	}

	cur := sl.head

	//记录前驱节点
	update := make([]*skipListNode, MAX_LEVEL)

	for i := sl.level - 1; i >= 0; i-- {

		update[i] = sl.head

		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].value == value {
				update[i] = cur
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]

	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		update[i].forwards[i] = update[i].forwards[i].forwards[i]
	}

	sl.length--
	return nil
}
