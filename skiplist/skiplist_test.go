package skiplist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList_NewSkipList(t *testing.T) {
	cases := []struct {
		name     string
		maxLevel int
		p        float32
		expectSL *SkipList
	}{
		{
			name:     "get one level sl",
			maxLevel: 1,
			p:        0.5,
			expectSL: &SkipList{
				Head: &Node{
					key:   headKey,
					value: "HEAD",
					forward: []*Node{
						makeNode(tailKey, "NIL", 1),
					},
				},
				MaxLevel: 1,
				P:        0.5,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sl := NewSkipList(tt.maxLevel, tt.p)
			assert.Equal(t, tt.expectSL, sl)
		})
	}
}

func TestSkipList_randomLevel(t *testing.T) {
	type fields struct {
		Head     *Node
		Tail     *Node
		MaxLevel int
		P        float32
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "happy path",
			fields: fields{
				MaxLevel: 6,
				P:        0.5,
			},
			want: 1,
		},
		{
			name: "happy path - 1",
			fields: fields{
				MaxLevel: 6,
				P:        0.5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &SkipList{
				Head:     tt.fields.Head,
				MaxLevel: tt.fields.MaxLevel,
				P:        tt.fields.P,
			}

			if got := sl.randomLevel(); got < tt.want {
				t.Errorf("SkipList.randomLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SkipList_Find(t *testing.T) {
	cases := []struct {
		name        string
		buildList   func() *SkipList
		key         int
		expectValue *Node
	}{
		{
			name: "find from a empty list",
			buildList: func() *SkipList {
				return NewSkipList(10, 0.5)
			},
			key:         1,
			expectValue: nil,
		},
		{
			name: "success get value - 1",
			buildList: func() *SkipList {
				sl := NewSkipList(10, 0.5)
				node0 := &Node{
					key:     0,
					value:   "test",
					forward: make([]*Node, 1),
				}
				node1 := &Node{
					key:     1,
					value:   "A",
					forward: make([]*Node, 2),
				}
				node0.forward[0] = sl.Head.forward[0]
				sl.Head.forward[0] = node0

				node1.forward[0] = node0.forward[0]
				node1.forward[1] = sl.Head.forward[1]
				node0.forward[0] = node1
				sl.Head.forward[1] = node1
				return sl
			},
			key: 1,
			expectValue: &Node{
				key:   1,
				value: "A",
				forward: []*Node{
					makeNode(tailKey, "NIL", 10),
					makeNode(tailKey, "NIL", 10),
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sl := tt.buildList()
			node := sl.Find(tt.key)
			assert.Equal(t, tt.expectValue, node)
		})
	}
}

func Test_randFloat(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := randFloat()
		fmt.Println(i, ", random: ", got)
		fmt.Printf("%f less than 0.5: %t\n", got, got < 0.5)
	}
}

func TestSkipList_Insert(t *testing.T) {
	type fields struct {
		buildList func() *SkipList
	}
	type args struct {
		key   int
		value string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectValue *Node
	}{
		{
			name: "happy path - insert to empty list",
			fields: fields{
				buildList: func() *SkipList {
					return NewSkipList(10, 0.5)
				},
			},
			args: args{
				key:   1,
				value: "A",
			},
			expectValue: &Node{
				key:   1,
				value: "A",
			},
		},
		{
			name: "happy path - insert to non-empty list",
			fields: fields{
				buildList: func() *SkipList {
					sl := NewSkipList(10, 0.5)
					node0 := &Node{
						key:     0,
						value:   "test",
						forward: make([]*Node, 1),
					}
					node1 := &Node{
						key:     6,
						value:   "A",
						forward: make([]*Node, 2),
					}
					node0.forward[0] = sl.Head.forward[0]
					sl.Head.forward[0] = node0

					node1.forward[0] = node0.forward[0]
					node1.forward[1] = sl.Head.forward[1]
					node0.forward[0] = node1
					sl.Head.forward[1] = node1
					return sl
				},
			},
			args: args{
				key:   2,
				value: "B",
			},
			expectValue: &Node{
				key:   2,
				value: "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := tt.fields.buildList()
			sl.Insert(tt.args.key, tt.args.value)

			fmt.Println(sl.String())

			node := sl.Find(tt.args.key)
			assert.NotNil(t, node)
			assert.Equal(t, tt.expectValue.key, node.key)
			assert.Equal(t, tt.expectValue.value, node.value)
		})
	}
}

func TestSkipList_Delete(t *testing.T) {
	type fields struct {
		Head     *Node
		MaxLevel int
		P        float32
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &SkipList{
				Head:     tt.fields.Head,
				MaxLevel: tt.fields.MaxLevel,
				P:        tt.fields.P,
			}
			sl.Delete(tt.args.key)
		})
	}
}
