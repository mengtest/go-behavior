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
		ID:                id,
		Children:          make([]behavior.Noder, 0),
		BeforeAttachments: make([]behavior.Attachment, 0),
		AfterAttachments:  make([]behavior.Attachment, 0),
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

func (this *Node) GetBeforeAttachments() []behavior.Attachment {
	return this.BeforeAttachments
}

func (this *Node) GetAfterAttachments() []behavior.Attachment {
	return this.AfterAttachments
}

func (this *Node) SetBeforeAttachmentAtIndex(i int, atcm behavior.Attachment) {
	this.BeforeAttachments[i] = atcm
}

func (this *Node) SetBeforeAttachments(attachments ...behavior.Attachment) {
	this.BeforeAttachments = attachments
}

func (this *Node) SetAfterAttachmentAtIndex(i int, atcm behavior.Attachment) {
	this.AfterAttachments[i] = atcm
}

func (this *Node) SetAfterAttachments(attachments ...behavior.Attachment) {
	this.AfterAttachments = attachments
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
