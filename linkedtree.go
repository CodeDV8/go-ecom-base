package EComBase

import (
	"fmt"
	"strings"
)

// TODO - Add mutex?

type LinkedTree struct {
	root *LinkedTreeNode
}

func (tree *LinkedTree) Add(key string, data interface{}) (bool, error) {
	if tree.root == nil {
		tree.root = &LinkedTreeNode{
			Key:  key,
			Data: data,
		}
		return true, nil
	}
	return tree.root.Add(key, data)
}

func (tree *LinkedTree) Balance() (*LinkedTreeNode, error) {
	if tree.root != nil {
		node, err := tree.root.Balance()
		if err != nil {
			return nil, err
		}
		tree.root = node
		return node, nil
	}
	return nil, fmt.Errorf("Can't balance a tree without a root node")
}

func (tree *LinkedTree) GetRootNode() (*LinkedTreeNode, error) {
	if tree.root != nil {
		return tree.root, nil
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) GetFirstNode() (*LinkedTreeNode, error) {
	if tree.root != nil {
		return tree.root.GetFirstNode()
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) GetLastNode() (*LinkedTreeNode, error) {
	if tree.root != nil {
		return tree.root.GetLastNode()
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) FindNode(key string) (*LinkedTreeNode, error) {
	if tree.root != nil {
		return tree.root.FindNode(key)
	}
	return nil, fmt.Errorf("Can't search a tree without a root node")
}

type LinkedTreeNode struct {
	Parent *LinkedTreeNode
	Next   *LinkedTreeNode
	Right  *LinkedTreeNode
	Left   *LinkedTreeNode
	Prev   *LinkedTreeNode
	Key    string
	Data   interface{}
}

func (lt *LinkedTreeNode) Add(key string, data interface{}) (bool, error) {
	if lt.Parent != nil {
		return lt.Parent.Add(key, data)
	}
	return lt.add(key, data)
}

func (lt *LinkedTreeNode) add(key string, data interface{}) (bool, error) {
	cmp := strings.Compare(lt.Key, key)
	if cmp == 0 {
		return false, nil
	}
	if cmp > 0 {
		if lt.Left == nil {
			lt.Left = &LinkedTreeNode{
				Parent: lt,
				Prev:   lt.Prev,
				Next:   lt,
				Key:    key,
				Data:   data,
			}
			if lt.Prev != nil {
				lt.Prev.Next = lt.Left
			}
			lt.Prev = lt.Left
			return true, nil
		}
		return lt.Left.add(key, data)
	}
	if lt.Right == nil {
		lt.Right = &LinkedTreeNode{
			Parent: lt,
			Next:   lt.Next,
			Prev:   lt,
			Key:    key,
			Data:   data,
		}
		if lt.Next != nil {
			lt.Next.Prev = lt.Right
		}
		lt.Next = lt.Right
		return true, nil
	}
	return lt.Right.add(key, data)
}

func (lt *LinkedTreeNode) GetFirstNode() (*LinkedTreeNode, error) {
	root, err := lt.GetRootNode()
	if err != nil {
		return nil, err
	}
	return root.getFirstNode()
}

func (lt *LinkedTreeNode) getFirstNode() (*LinkedTreeNode, error) {
	if lt.Left == nil {
		return lt, nil
	}
	return lt.Left.getFirstNode()
}

func (lt *LinkedTreeNode) GetRootNode() (*LinkedTreeNode, error) {
	return lt.getRootNode()
}

func (lt *LinkedTreeNode) getRootNode() (*LinkedTreeNode, error) {
	if lt.Parent == nil {
		return lt, nil
	}
	return lt.Parent.getRootNode()
}

func (lt *LinkedTreeNode) GetLastNode() (*LinkedTreeNode, error) {
	root, err := lt.GetRootNode()
	if err != nil {
		return nil, err
	}
	return root.getLastNode()
}

func (lt *LinkedTreeNode) getLastNode() (*LinkedTreeNode, error) {
	if lt.Right == nil {
		return lt, nil
	}
	return lt.Right.getLastNode()
}

func (lt *LinkedTreeNode) FindNode(key string) (*LinkedTreeNode, error) {
	root, err := lt.GetRootNode()
	if err != nil {
		return nil, err
	}
	return root.findNode(key)
}

func (lt *LinkedTreeNode) findNode(key string) (*LinkedTreeNode, error) {
	cmp := strings.Compare(lt.Key, key)
	if cmp == 0 {
		return lt, nil
	}
	if cmp > 0 {
		if lt.Left != nil {
			return lt.Left.findNode(key)
		}
		return nil, nil
	}
	if lt.Right != nil {
		return lt.Right.findNode(key)
	}
	return nil, nil
}

func (lt *LinkedTreeNode) Balance() (*LinkedTreeNode, error) {
	var list []*LinkedTreeNode

	node, err := lt.GetFirstNode()
	if err != nil {
		return nil, err
	}
	for node != nil {
		node.Parent = nil
		node.Left = nil
		node.Right = nil
		list = append(list, node)
		node = node.Next
	}
	return lt.balance(list)
}

func (lt *LinkedTreeNode) balance(list []*LinkedTreeNode) (*LinkedTreeNode, error) {
	listLength := len(list)
	if listLength < 1 {
		return nil, fmt.Errorf("List is empty")
	}
	if listLength == 1 {
		return list[0], nil
	}
	var newList []*LinkedTreeNode

	base := 0
	for base < listLength {
		if base+3 < listLength {
			list[base+1].Left = list[base]
			list[base].Parent = list[base+1]
			list[base+1].Right = list[base+2]
			list[base+2].Parent = list[base+1]
			newList = append(newList, list[base+1])
			newList = append(newList, list[base+3])
			base += 4
		} else if base+2 < listLength {
			list[base+1].Left = list[base]
			list[base].Parent = list[base+1]
			list[base+1].Right = list[base+2]
			list[base+2].Parent = list[base+1]
			newList = append(newList, list[base+1])
			base += 3
		} else if base+1 < listLength {
			list[base+1].Left = list[base]
			list[base].Parent = list[base+1]
			newList = append(newList, list[base+1])
			base += 2
		} else if base < listLength {
			newList = append(newList, list[base])
			base += 1
		}
	}
	return lt.balance(newList)
}
