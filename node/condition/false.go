package condition

import (
	"context"

	"github.com/szyhf/go-behavior"
)

//False节点总是返回Failure

type False struct {
	Condition
}

func (this *False) Run(ctx context.Context) behavior.Status {
	return behavior.StatusFailure
}
