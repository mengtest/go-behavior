package decorator

import (
	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

// 装饰节点作为控制分支节点，必须且只接受一个子节点。装饰节点的执行首先执行子节点，并根据自身的控制逻辑以及子节点的返回结果决定自身的状态。

// 装饰节点都有属性“子节点结束时作用（IsDecorateChildEnds）”可以配置，如果该值配置为真，则仅当子节点结束（成功或失败）的时候，装饰节点的装饰逻辑才起作用。

type Decorator struct {
	common.Node
	IsDecorateChildEnds bool
}

func (d *Decorator) SetChild(child behavior.Noder) {
	if d.Node.GetChildren() == nil {
		d.Node.Children = make([]behavior.Noder, 1)
	}
	d.Node.Children[0] = child
}

func (d *Decorator) GetChild() behavior.Noder {
	if len(d.Node.GetChildren()) > 0 {
		return d.Node.Children[0]
	}
	return nil
}

func (d *Decorator) IsValidate() error {
	if d.Node.GetChildren() == nil {
		return ErrNilChildren
	}
	if len(d.Node.GetChildren()) != 1 {
		return ErrTooManyChild
	}
	return nil
}
