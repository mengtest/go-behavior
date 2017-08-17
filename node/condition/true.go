package condition

import (
	"context"

	"github.com/szyhf/go-behavior"
)

//True节点总是返回Success

type True struct {
	Condition
}

func Run(ctx context.Context) behavior.Status {
	return behavior.StatusSuccess
}
