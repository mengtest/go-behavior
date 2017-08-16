package condition

import (
	"context"

	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/common"
)

// 条件节点根据比较结果返回成功或失败，但永远不会返回正在执行（Running）

type ConditionCompare int

const (
	ConditionEqual ConditionCompare = iota
	ConditionLarger
	ConditionLesser
	ConditionLargerOrEqual
	ConditionLesserOrEqual

	ConditionContains // 仅限于Slice或者Map
)

// ShouldEqual          = assertions.ShouldEqual
// ShouldNotEqual       = assertions.ShouldNotEqual
// ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
// ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
// ShouldBeLessThan             = assertions.ShouldBeLessThan
// ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
// ShouldContain       = assertions.ShouldContain
// ShouldNotContain    = assertions.ShouldNotContain
// ShouldContainKey    = assertions.ShouldContainKey
// ShouldNotContainKey = assertions.ShouldNotContainKey

type Condition struct {
	//key              string
	//expect           int
	//ConditionCompare ConditionCompare
	common.Node
	EvaluateFunc func(ctx context.Context) bool
}

func NewCondition(evaluateFunc func(ctx context.Context) bool) behavior.Conditioner {
	cond := &Condition{
		EvaluateFunc: evaluateFunc,
	}
	return cond
}

func (c *Condition) Run(ctx context.Context) behavior.Status {
	if c.EvaluateFunc == nil {
		return behavior.StatusSuccess
	}
	if c.EvaluateFunc(ctx) {
		return behavior.StatusSuccess
	} else {
		return behavior.StatusFailure
	}
}

// ShouldBeNil          = assertions.ShouldBeNil
// ShouldNotBeNil       = assertions.ShouldNotBeNil
// ShouldBeTrue         = assertions.ShouldBeTrue
// ShouldBeFalse        = assertions.ShouldBeFalse
// ShouldBeZeroValue    = assertions.ShouldBeZeroValue
// ShouldBeIn          = assertions.ShouldBeIn
// ShouldNotBeIn       = assertions.ShouldNotBeIn
// ShouldBeEmpty       = assertions.ShouldBeEmpty
// ShouldNotBeEmpty    = assertions.ShouldNotBeEmpty
// ShouldHaveLength    = assertions.ShouldHaveLength
