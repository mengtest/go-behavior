package decorator

import (
	"context"

	"github.com/szyhf/go-behavior"
)

// 类似于逻辑“非”操作，非节点对子节点的返回值执行如下操作：
// 如果子节点失败，那么此节点返回成功。
// 如果子节点成功，那么此节点返回失败。
// 如果子节点返回正在执行，则同样返回正在执行。

type Not struct {
	Decorator
}

func (n *Not) Run(ctx context.Context) behavior.Status {
	child := n.Decorator.GetChild()
	if child == nil {
		panic("Child of NotNode should exist.")
	}
	switch s := n.Decorator.GetChild().Run(ctx); s {
	case behavior.StatusSuccess:
		return behavior.StatusFailure
	case behavior.StatusFailure:
		return behavior.StatusSuccess
	default:
		return s
	}
}
