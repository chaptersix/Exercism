package tree

import (
	"errors"
	"sort"
)

//Record is the record sorted in the db
type Record struct {
	ID, Parent int
}

//Node is a tree element containing the  id and its children
type Node struct {
	ID       int
	Children []*Node
}

//Build builds the tree form an unsorted list of records
func Build(records []Record) (root *Node, err error) {
	rootNode := &Node{ID: -1}
	nodeList := make([]*Node, 0)
	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	for index, record := range records {
		if index != record.ID {
			return nil, errors.New("not continuous")
		}
		if record.ID > 0 && record.ID <= record.Parent {
			return nil, errors.New("cycle")
		}
		currentNode := &Node{ID: record.ID}
		nodeList = append(nodeList, currentNode)
		if record.ID == 0 {
			rootNode = currentNode
		}
	}
	if rootNode.ID == -1 {
		return nil, nil
	}
	for _, record := range records {
		if record.ID == 0 {
			if record.Parent > 0 {
				return nil, errors.New("zero has parent")
			}
			continue
		}
		for _, node := range nodeList {
			if node.ID == record.Parent {
				var child *Node
				for _, n := range nodeList {
					if n.ID == record.ID {
						child = n
						break
					}
				}
				node.Children = append(node.Children, child)
			}
		}
	}
	return rootNode, nil
}
