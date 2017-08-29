package common

import (
	"context"

	behavior "github.com/szyhf/go-behavior"
)

type Node struct {
	ID       int
	Parent   behavior.Noder
	Children []behavior.Noder
	//BeforeUnionType
	BeforeAttachments []behavior.Attachment
	AfterAttachments  []behavior.Attachment
}

func NewNoder(id int) behavior.Noder {
	return &Node{
		ID:       id,
		Children: make([]behavior.Noder, 0),
	}
}

func (this *Node) Run(ctx context.Context) behavior.Status {
	panic("not imp")
}

func (this *Node) SetID(id int) {
	this.ID = id
}

func (this *Node) GetID() int {
	return this.ID
}

func (this *Node) SetParent(n behavior.Noder) {
	this.Parent = n
}

func (this *Node) GetParent() behavior.Noder {
	return this.Parent
}

func (this *Node) SetChildren(children ...behavior.Noder) {
	this.Children = children
	for _, child := range this.Children {
		if child != nil {
			child.SetParent(this)
		}
	}
}

func (this *Node) GetChildren() []behavior.Noder {
	return this.Children
}

func (this *Node) SetChildAtIndex(i int, child behavior.Noder) {
	child.SetParent(this)
	this.Children[i] = child
}
func (this *Node) GetChildrenCount() int {
	return len(this.Children)
}
func (this *Node) GetChildByIndex(index int) behavior.Noder {
	return this.Children[index]
}

func (this *Node) RunBeforeAttachments(ctx context.Context) behavior.Status {
	for _, bef := range this.BeforeAttachments {
		if status := bef.Run(ctx); status != behavior.StatusSuccess {
			return status
		}
	}
	return behavior.StatusSuccess
}

func (this *Node) RunAfterAttachments(ctx context.Context) behavior.Status {
	for _, bef := range this.AfterAttachments {
		if status := bef.Run(ctx); status != behavior.StatusSuccess {
			return status
		}
	}
	return behavior.StatusSuccess
}
