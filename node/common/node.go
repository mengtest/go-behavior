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
}

func NewNoder(id int) behavior.Noder {
	return &Node{
		ID:       id,
		Children: make([]behavior.Noder, 0),
	}
}

func (this *Node) Run(ctx context.Context) behavior.Status {
	panic("Not imp")
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

func (this *Node) GetChildren() []behavior.Noder {
	return this.Children
}
