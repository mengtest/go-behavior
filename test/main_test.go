package test

import (
	"testing"

	"context"

	"github.com/szyhf/go-behavior"
	"github.com/szyhf/go-behavior/node/action"
	"github.com/szyhf/go-behavior/node/composite"
	"github.com/szyhf/go-behavior/node/condition"
	"gitlab.shouyouqianxian.com/new-legend/server/log"
)

var (
	isTheDugeon      = true
	isRoleHasMission = false
)

var DugeonIDEqualCondition = condition.NewCondition(func(ctx context.Context) bool {
	// 跳过验证，假设符合条件
	log.Alertf("CS在副本：%v\n", isTheDugeon)
	return isTheDugeon
})

var RoleContainsMissionCondition = condition.NewCondition(func(ctx context.Context) bool {
	// 跳过验证，假设符合条件
	log.Alertf("CS有任务：%v\n", isRoleHasMission)
	return isRoleHasMission
})

var DugeonCreateNPCAction = action.NewAction(func(ctx context.Context) behavior.Status {
	log.Alertf("在坐标(250,250)创建NPC=FU")
	return behavior.StatusSuccess
})

var BaseSeq = composite.NewSequence(nil)

func TestBase(t *testing.T) {
	// 如果玩家cs进入副本(id=xxx,名字=M记的Restroom)时，身上有任务(id=xxxxx,名字=寻找吖傅的故事)，就在副本的某坐标（M记的从左开始第3间厕所）创建npc（吖傅）
	// Loop(1)
	//      -> Sequence
	//               -> Condition -> Dugeon.ID=xxx
	//               -> Condition -> Role.Missions.Contains(id=xxxx)
	//               -> Action -> CreateNPC(fufufu,x,y)
	BaseSeq.Push(DugeonIDEqualCondition, RoleContainsMissionCondition, DugeonCreateNPCAction)
	BaseSeq.Run(context.TODO())
}

func TestPush(t *testing.T) {
	// 如果玩家cs进入某格，就 推送副本事件（事件名字=吖傅掉坑，事件进度计数=2）
}
