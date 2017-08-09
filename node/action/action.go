package action

import (
	"context"

	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

type Action struct {
	common.Node
	runFunc func(ctx context.Context) behavior.Status
}

func NewAction(runFunc func(ctx context.Context) behavior.Status) *Action {
	return &Action{
		runFunc: runFunc,
	}
}

func (a *Action) Run(ctx context.Context) behavior.Status {
	if a.runFunc == nil {
		return behavior.StatusSuccess
	}
	return a.runFunc(ctx)
}
