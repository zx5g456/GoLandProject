package geeServer

import (
	"strings"
)

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由的一部分
	children []*node //子节点
	isWild   bool    //是否精准匹配
}

// 返回第一个匹配的节点
func (n *node) matchChild(part string) *node {
	//简单的遍历查询
	for _, child := range n.children {
		if child.pattern == part || child.isWild {
			return child
		}
	}
	return nil
}

// 返回所有匹配的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// trie树的插入
// pattern 将要插入的节点
// parts pattern 切割后的的字符串数组
// height 当前层数
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
