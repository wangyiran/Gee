package gee

//the struction of the route node
type Node struct {
	pattern  string
	part     string
	children []*Node
	isWild   bool
}

func (n *Node) searchFirstChild(part string) *Node {
	tag := 0
	res := &Node{}
	for _, child := range n.children {
		if child.part == part {
			return child
		} else if child.part[0] == ':' {
			res = child
			tag = 1
		}
	}
	if tag == 1 {
		return res
	} else {
		return nil
	}
}

func (n *Node) searchAllChild(part string) []*Node {
	res := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			res = append(res, child)
		}
	}
	return res
}

func (n *Node) insertRoute(pattern string, parts []string, level int) {
	if level == len(parts) {
		n.pattern = pattern
		return
	}
	firstChild := n.searchFirstChild(parts[level])
	if firstChild == nil {
		//没有直接插入
		firstChild = &Node{
			part:   parts[level],
			isWild: parts[level][0] == ':' || parts[level][0] == '*',
		}
		n.children = append(n.children, firstChild)

	}

	firstChild.insertRoute(pattern, parts, level+1)
}

func (n *Node) searchRoute(parts []string, level int) *Node {
	if len(parts) == level || n.part[0] == '*' {
		if n.pattern == "" {
			return nil
		}
		return n

	}

	part := parts[level]
	children := n.searchAllChild(part)
	for _, child := range children {
		res := child.searchRoute(parts, level+1)
		if res != nil {
			return res
		}
	}
	return nil
}
