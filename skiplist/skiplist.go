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
	currentMaxLevel := nodeLevel(sl.Head)

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
func nodeLevel(n *Node) int {
	currentLevel := 1
	nilKey := math.MaxInt

	if n.key == nilKey {
		return currentLevel
	}

	for _, forward := range n.forward {
		if forward != nil && forward.key != nilKey {
			currentLevel++
		} else {
			break
		}
	}

	return currentLevel
}
