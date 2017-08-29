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
	if status := a.RunBeforeAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	if a.ActFunc == nil {
		return behavior.StatusSuccess
	}
	if status := a.ActFunc(ctx); status != behavior.StatusSuccess {
		return status
	}
	if status := a.RunAfterAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	return behavior.StatusSuccess
}

func (this *Action) SetChildren(children ...behavior.Noder) {
	panic("go-behavior: action node should not have child.")
}

func (this *Action) SetChildAtIndex(i int, child behavior.Noder) {
	panic("go-behavior: action node should not have child.")
}
