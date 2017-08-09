package composite

import (
	"context"

	"github.com/szyhf/go-behavior"
)

// 该节点以给定的顺序依次执行其子节点，直到所有子节点成功返回，该节点也返回成功。只要其中某个子节点失败，那么该节点也失败。

// Sequence 实现了与（&&）的功能。我们知道表达式 R=A&&B&&C&&D 执行的时候首先执行 A，如果 A 是 false 则返回 false，
// 如果 A 是 true 则执行 B，如果 B 是 false 则返回 false，否则如果 B 是 true 则执行 C，如果 C 是 false 则返回 false，
// 否则如果 C 是 true 则执行 D，并且返回D 的值。

// 最一般的意义上， Sequence 节点实现了一个序列。 实际上， Sequence 节点不仅可以管理‘动作’子节点，也可以管理‘条件’子节点。
// 如上图的应用中， 如果1和2号节点是条件节点的话，这两个条件节点实际上用作进入下面其他节点的 precondition，
// 只有这两个条件是 true，下面的其他节点才有可能执行。

// 此外， Sequence 上还可以添加‘中断条件’作为终止执行的条件。字段AbortFunc就是可选的‘中断条件’。
// 该‘中断条件’在每处理下一个子节点的时候被检查，当为true时，则不再继续，返回失败（Failure）。

type Sequence struct {
	Composite
	runFunc   func(s *Sequence, ctx context.Context) behavior.Status
	abortFunc func(ctx context.Context) bool
}

func NewSequence(abortFunc func(ctx context.Context) bool) behavior.Sequencer {
	var seq *Sequence
	if abortFunc == nil {
		seq = &Sequence{
			runFunc: runSequenceWithoutAbort,
		}
	} else {
		seq = &Sequence{
			runFunc:   runSequenceWithAbort,
			abortFunc: abortFunc,
		}
	}
	return seq
}

func (s *Sequence) Run(ctx context.Context) behavior.Status {
	return s.runFunc(s, ctx)
}

func runSequenceWithAbort(s *Sequence, ctx context.Context) behavior.Status {
	for _, child := range s.Node.GetChildren() {
		if s.abortFunc(ctx) {
			return behavior.StatusFailure
		}
		if status := child.Run(ctx); status != behavior.StatusSuccess {
			// 待定：是否缓存Sequence为Running的Node。
			return status
		}
	}
	return behavior.StatusSuccess
}

func runSequenceWithoutAbort(s *Sequence, ctx context.Context) behavior.Status {
	for _, child := range s.Node.GetChildren() {
		if status := child.Run(ctx); status != behavior.StatusSuccess {
			// 待定：是否缓存Sequence为Running的Node。
			return status
		}
	}
	return behavior.StatusSuccess
}
