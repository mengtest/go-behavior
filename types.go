package behavior

import (
	"context"
)

// 行为树的基本概念：

// 执行每个节点都会有一个结果（成功，失败或运行）
// 子节点的执行结果由其父节点控制和管理
// 返回运行结果的节点被视作处于运行状态，处于运行状态的节点将被持续执行一直到其返回结束（成功或失败）。在其结束前，其父节点不会把控制转移到后续节点。

type Noder interface {
	Run(ctx context.Context) Status
	GetID() int
	GetParent() Noder
	GetChildren() []Noder
	// GetChildrenCount() int
	// GetChildByIndex(index int) Noder
	// GetChildByID(id int) Noder
	// String() string
	// Clear()
	// // 检测先决条件
	// CheckPreconditions(isAlive bool) bool
	// ApplyEffects(status Status)
	// CheckEvents(eventName string, params map[int]interface{}) bool
	// Attach(attachment Noder, isPrecondition, isEffector, isTransition bool)
	// IsManagingChildrenAsSubTrees() bool
	SetParent(n Noder)
}

// 顺序执行所有子节点返回成功，如果某个子节点失败返回失败。
type Loopper interface {
}

// 顺序执行所有子节点返回成功，如果某个子节点失败返回失败。
type Conditioner interface {
	Noder
}

// 当指定的周期过去后返回成功。
type Waiter interface {
}

// 顺序执行所有子节点返回成功，如果某个子节点失败返回失败。
type Sequencer interface {
	Noder
	Push(node ...Noder)
}

type Selector interface {
	Noder
	Push(node ...Noder)
}

// 修饰节点
type Decorator interface {
}

type Actioner interface {
	Noder
}

// 行为树
type Tree interface {
	GetID() int
	Run(ctx context.Context) Status
}

// 附件
type Attachment interface {
	Run(ctx context.Context) bool
}

type ConditionEvaluater interface {
	Evaluate(ctx context.Context) bool
}
