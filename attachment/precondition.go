package attachment

import (
	"context"

	behavior "github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

type Precondition struct {
	common.Node
	evaluate func(ctx context.Context) behavior.Status
}

func (p *Precondition) Run(ctx context.Context) behavior.Status {
	if status := p.RunBeforeAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	if status := p.evaluate(ctx); status != behavior.StatusSuccess {
		return status
	}
	if status := p.RunAfterAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	return behavior.StatusSuccess
}
