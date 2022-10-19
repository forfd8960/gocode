package skiplist

import (
	"fmt"
	"testing"
)

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
				Tail:     tt.fields.Tail,
				MaxLevel: tt.fields.MaxLevel,
				P:        tt.fields.P,
			}

			if got := sl.randomLevel(); got < tt.want {
				t.Errorf("SkipList.randomLevel() = %v, want %v", got, tt.want)
			}
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
