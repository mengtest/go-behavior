package composite

import (
	"context"
	"fmt"

	behavior "github.com/szyhf/go-behavior"
)

// 顺序执行所有子节点，当所有子节点返回成功时，则返回成功。若有任意一个子节点失败，则返回失败。
type And struct {
	Composite
	runFunc   func(s *And, ctx context.Context) behavior.Status
	abortFunc func(ctx context.Context) bool
}

func NewAnd(abortFunc func(ctx context.Context) bool) behavior.And {
	var seq *And
	if abortFunc == nil {
		seq = &And{
			runFunc: runAndWithoutAbort,
		}
	} else {
		seq = &And{
			runFunc:   runAndWithAbort,
			abortFunc: abortFunc,
		}
	}
	return seq
}

func (s *And) Run(ctx context.Context) behavior.Status {
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

func runAndWithAbort(s *And, ctx context.Context) behavior.Status {
	finRes := behavior.StatusSuccess
	for idx, child := range s.Node.GetChildren() {
		if child != nil {
			if s.abortFunc(ctx) {
				return behavior.StatusFailure
			}
			if status := child.Run(ctx); status == behavior.StatusFailure {
				finRes = behavior.StatusFailure
			}
		} else {
			println(fmt.Sprintf("child of id=%d at idx=%d is nil", s.GetID(), idx))
		}
	}
	return finRes
}

func runAndWithoutAbort(s *And, ctx context.Context) behavior.Status {
	finRes := behavior.StatusSuccess
	for idx, child := range s.Node.GetChildren() {
		if child != nil {
			if status := child.Run(ctx); status != behavior.StatusFailure {
				finRes = behavior.StatusFailure
			}
		} else {
			println(fmt.Sprintf("child of id=%d at idx=%d is nil", s.GetID(), idx))
		}
	}
	return finRes
}
