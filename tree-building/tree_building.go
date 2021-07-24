package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID,
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(recs []Record) (*Node, error) {
	nodes := make(map[int]*Node, len(recs))

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].ID < recs[j].ID
	})

	for i, r := range recs {
		if r.ID != i || r.Parent > i || i > 0 && r.Parent == i {
			return nil, errors.New("invalid tree")
		}

		nodes[i] = &Node{ID: i}

		if i == 0 {
			continue
		}

		nodes[r.Parent].Children = append(
			nodes[r.Parent].Children,
			nodes[i],
		)
	}

	return nodes[0], nil
}
