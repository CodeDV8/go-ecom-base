package EComBase

import (
	"fmt"
	"strings"
	"sync"
)

// TODO - Add mutex?

type LinkedTree struct {
	root  *LinkedTreeNode
	count int64
	mux   sync.Mutex
}

func (tree *LinkedTree) Add(key string, data interface{}) (bool, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root == nil {
		tree.root = &LinkedTreeNode{
			Key:  key,
			Data: data,
		}
		tree.count = 1
		return true, nil
	}
	ok, err := tree.root.Add(key, data)
	if ok == true {
		tree.count += 1
	}
	return ok, err
}

func (tree *LinkedTree) Delete(key string) (bool, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	node, err := tree.FindNode(key)
	if err != nil {
		return false, err
	}
	node.Deleted = true
	tree.count--
	return true, nil
}

func (tree *LinkedTree) Balance() (*LinkedTreeNode, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root != nil {
		node, err, nodeCount := tree.root.Balance(tree.count)
		tree.count = nodeCount
		if err != nil {
			return nil, err
		}
		tree.root = node
		return node, nil
	}
	return nil, fmt.Errorf("Can't balance a tree without a root node")
}

func (tree *LinkedTree) GetRootNode() (*LinkedTreeNode, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root != nil {
		return tree.root, nil
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) GetFirstNode() (*LinkedTreeNode, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root != nil {
		return tree.root.GetFirstNode()
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) GetLastNode() (*LinkedTreeNode, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root != nil {
		return tree.root.GetLastNode()
	}
	return nil, fmt.Errorf("Can't get a node from a tree without a root node")
}

func (tree *LinkedTree) FindNode(key string) (*LinkedTreeNode, error) {
	tree.mux.Lock()
	defer tree.mux.Unlock()
	if tree.root != nil {
		return tree.root.FindNode(key)
	}
	return nil, fmt.Errorf("Can't search a tree without a root node")
}

type LinkedTreeNode struct {
	Parent   *LinkedTreeNode
	Next     *LinkedTreeNode
	Right    *LinkedTreeNode
	Left     *LinkedTreeNode
	Prev     *LinkedTreeNode
	Key      string
	Data     interface{}
	Wildcard bool
	Deleted  bool
}

func (lt *LinkedTreeNode) Add(key string, data interface{}) (bool, error) {
	if lt.Parent != nil {
		return lt.Parent.Add(key, data)
	}
	return lt.add(key, data)
}

func (lt *LinkedTreeNode) add(key string, data interface{}) (bool, error) {
	var wildcard = false
	var addKey = ""
	if key[len(key)-1:] == "*" {
		addKey = key[:len(key)-1]
		wildcard = true
	} else {
		addKey = key
	}
	cmp := strings.Compare(lt.Key, key)
	if cmp == 0 {
		if lt.Deleted == true {
			lt.Deleted = false
			lt.Data = data
			return true, nil
		}
		return false, nil
	}
	if cmp > 0 {
		if lt.Left == nil {
			lt.Left = &LinkedTreeNode{
				Parent:   lt,
				Prev:     lt.Prev,
				Next:     lt,
				Key:      addKey,
				Data:     data,
				Wildcard: wildcard,
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
			Parent:   lt,
			Next:     lt.Next,
			Prev:     lt,
			Key:      addKey,
			Data:     data,
			Wildcard: wildcard,
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
	var cmp int
	if lt.Wildcard == true {
		if len(key) >= len(lt.Key) {
			cmp = strings.Compare(lt.Key, key[:len(lt.Key)])
		} else {
			cmp = strings.Compare(lt.Key, key)
		}
	} else {
		cmp = strings.Compare(lt.Key, key)
	}
	if cmp == 0 {
		if lt.Deleted == true {
			return nil, fmt.Errorf("Node is deleted")
		}
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

func (lt *LinkedTreeNode) Balance(size int64) (*LinkedTreeNode, error, int64) {
	// list := make([]*LinkedTreeNode, size)
	list := []*LinkedTreeNode{}
	node, err := lt.GetFirstNode()
	if err != nil {
		return nil, err, 0
	}
	var nodeCount int64 = 0
	for node != nil {
		node.Parent = nil
		node.Left = nil
		node.Right = nil
		if lt.Deleted == false {
			list = append(list, node)
			nodeCount++
		}
		node = node.Next
	}
	node, err = lt.balance(list)
	return node, err, nodeCount
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
