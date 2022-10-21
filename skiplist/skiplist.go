package skiplist

import (
	"math"
	"math/rand"
	"time"
)

var (
	headKey = math.MinInt
	tailKey = math.MaxInt
)

type Node struct {
	key     int
	value   string
	forward []*Node
}

func makeNode(key int, value string, level int) *Node {
	n := &Node{
		key:     key,
		value:   value,
		forward: nil,
	}

	n.forward = make([]*Node, level)
	for i := 0; i < level; i++ {
		n.forward[i] = nil
	}
	return n
}

type SkipList struct {
	Head     *Node
	Tail     *Node
	MaxLevel int
	P        float32
}

func NewSkipList(maxLevel int, p float32) *SkipList {
	sl := &SkipList{
		MaxLevel: maxLevel,
		P:        p,
	}

	sl.Head = makeNode(headKey, "HEAD", maxLevel)
	sl.Tail = makeNode(tailKey, "NIL", maxLevel)

	for i := 0; i < maxLevel; i++ {
		sl.Head.forward[i] = sl.Tail
	}

	return sl
}

// Find find the skip node according key
func (sl *SkipList) Find(key int) *Node {
	currentNode := sl.Head
	currentMaxLevel := len(sl.Head.forward)

	for level := currentMaxLevel; level > 0; level-- {
		for currentNode.forward[level] != nil && currentNode.forward[level].key < key {
			currentNode = currentNode.forward[level]
		}
	}

	currentNode = currentNode.forward[0]
	if currentNode.key == key {
		return currentNode
	}

	return nil
}

// Insert insert key, value
func (sl *SkipList) Insert(key int, value string) {
	n := sl.Find(key)
	if n != nil {
		n.value = value
		return
	}

	updateForward := make([]*Node, len(sl.Head.forward))
	copy(updateForward, sl.Head.forward)

	currentNode := sl.Head
	currentMaxLevel := len(sl.Head.forward) - 1
	for level := currentMaxLevel; level > 0; level-- {
		for currentNode.forward[level] != nil && currentNode.forward[level].key < key {
			currentNode = currentNode.forward[level]
		}
		updateForward[level] = currentNode
	}

	currentNode = currentNode.forward[0]
	newNodeLevel := 1

	if currentNode.key != key {
		newNodeLevel = sl.randomLevel()
		currentLevel := nodeLevel(updateForward)

		if newNodeLevel > currentLevel {
			for i := currentLevel + 1; i < newNodeLevel; i++ {
				updateForward[i] = sl.Head
			}
		}

		currentNode = makeNode(key, value, newNodeLevel)
	}

	for i := 0; i < newNodeLevel; i++ {
		currentNode.forward[i] = updateForward[i].forward[i]
		updateForward[i].forward[i] = currentNode
	}
}

func (sl *SkipList) Delete(key int) {
	update := make([]*Node, len(sl.Head.forward))
	currentLevel := len(sl.Head.forward) - 1
	currentNode := sl.Head

	for level := currentLevel; level > 0; level-- {
		for currentNode.forward[level] != nil && currentNode.forward[level].key < key {
			currentNode = currentNode.forward[level]
		}

		update[level] = currentNode
	}

	currentNode = currentNode.forward[0]
	if currentNode == nil || currentNode.key != key {
		return
	}

	for i := 0; i < len(currentNode.forward); i++ {
		update[i].forward[i] = currentNode.forward[i]
	}
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for randFloat() < sl.P && level < sl.MaxLevel {
		level++
	}

	return level
}

func randFloat() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

// nodeLevel returns the number of non-null pointers
// corresponding to the level of the current node.
func nodeLevel(forward []*Node) int {
	currentLevel := 1
	if forward[0].key == tailKey {
		return currentLevel
	}

	for _, forward := range forward {
		if forward != nil && forward.key != tailKey {
			currentLevel++
		} else {
			break
		}
	}

	return currentLevel
}
