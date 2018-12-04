package graph

import (
	"testing"
)

func TestGraph_PrintGraph(t *testing.T) {
	g := NewGraph()

	tests := []struct {
		name string
		mock func()
		want string
	}{
		{
			name: "t1",
			mock: func() {
				user1 := &Node{Value: "Ikhsan"}
				user2 := &Node{Value: "Ahmad"}
				user3 := &Node{Value: "Faruqi"}
				g.AddNode(user1)
				g.AddNode(user2)
				g.AddNode(user3)

			},
			want: "Ikhsan -> Ahmad, Faruqi,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g.PrintGraph()
		})
	}
}

func TestGraph_GetNeighbour(t *testing.T) {
	g := NewGraph()

	user1 := Node{Value: "Ikhsan"}
	user2 := Node{Value: "Ahmad"}
	user3 := Node{Value: "Faruqi"}

	g.AddNode(&user1)
	g.AddNode(&user2)
	g.AddNode(&user3)

	type args struct {
		source Node
	}
	tests := []struct {
		name string
		args args
		mock func()
		want string
	}{
		{
			name: "Test 1",
			args: args{
				source: user1,
			},
			mock: func() {
				g.AddEdge(user1, &user2)
				g.AddEdge(user1, &user3)
			},
			want: "Ikhsan -> Ahmad, Faruqi, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := g.GetNeighbour(tt.args.source); got != tt.want {
				t.Errorf("Graph.GetNeighbour() = %v, want %v", got, tt.want)
			}
		})
	}
}
