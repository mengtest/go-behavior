package action

import (
	"context"

	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

type Action struct {
	common.Node
	ActFunc func(ctx context.Context) behavior.Status
}

func NewAction(actFunc func(ctx context.Context) behavior.Status) *Action {
	return &Action{
		ActFunc: actFunc,
	}
}

func (a *Action) Run(ctx context.Context) behavior.Status {
	if a.ActFunc == nil {
		return behavior.StatusSuccess
	}
	return a.ActFunc(ctx)
}
