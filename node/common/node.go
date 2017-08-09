package common

import (
	"context"

	behavior "github.com/szyhf/go-behavior"
)

type Node struct {
	Children []behavior.Noder
}

func NewNoder(name string) behavior.Noder {
	panic("Not imp")
}

func (this *Node) Run(ctx context.Context) behavior.Status {
	panic("Not imp")
}

func (this *Node) GetChildren() []behavior.Noder {
	return this.Children
}
