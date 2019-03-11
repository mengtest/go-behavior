package composite

import (
	"context"
	"fmt"

	behavior "github.com/szyhf/go-behavior"
)

// 顺序执行所有子节点，当所有子节点返回失败时，则返回失败。若有一个子节点成功，则返回成功。
type Or struct {
	Composite
	runFunc   func(s *Or, ctx context.Context) behavior.Status
	abortFunc func(ctx context.Context) bool
}

func NewOr(abortFunc func(ctx context.Context) bool) behavior.Or {
	var seq *Or
	if abortFunc == nil {
		seq = &Or{
			runFunc: runOrWithoutAbort,
		}
	} else {
		seq = &Or{
			runFunc:   runOrWithAbort,
			abortFunc: abortFunc,
		}
	}
	return seq
}

func (s *Or) Run(ctx context.Context) behavior.Status {
	if status := s.RunBeforeAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	if status := s.runFunc(s, ctx); status != behavior.StatusSuccess {
		return status
	}
	if status := s.RunAfterAttachments(ctx); status != behavior.StatusSuccess {
		return status
	}
	return behavior.StatusSuccess
}

func runOrWithAbort(s *Or, ctx context.Context) behavior.Status {
	finRes := behavior.StatusFailure
	for idx, child := range s.Node.GetChildren() {
		if child != nil {
			if s.abortFunc(ctx) {
				return behavior.StatusFailure
			}
			if status := child.Run(ctx); status == behavior.StatusSuccess {
				finRes = behavior.StatusSuccess
			}
		} else {
			println(fmt.Sprintf("child of id=%d at idx=%d is nil", s.GetID(), idx))
		}
	}
	return finRes
}

func runOrWithoutAbort(s *Or, ctx context.Context) behavior.Status {
	finRes := behavior.StatusFailure
	for idx, child := range s.Node.GetChildren() {
		if child != nil {
			if status := child.Run(ctx); status != behavior.StatusSuccess {
				finRes = behavior.StatusSuccess
			}
		} else {
			println(fmt.Sprintf("child of id=%d at idx=%d is nil", s.GetID(), idx))
		}
	}
	return finRes
}
