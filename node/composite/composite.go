package composite

import (
	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

type Composite struct {
	common.Node
}

func (c *Composite) Push(node ...behavior.Noder) {
	c.Node.Children = append(c.Node.Children, node...)
}
