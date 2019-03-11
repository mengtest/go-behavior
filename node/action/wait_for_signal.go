package action

// 等待信号（WaitForSignal）节点模拟了等待某个条件的“阻塞”过程。

// 等待信号节点返回Running，直到它上面附加的条件是true的时候：
// 1、如果有子节点，则执行其子节点，并当子节点结束时，返回该子节点的返回值。
// 2、如果没有子节点，则直接返回成功。

type WaitForSignal struct {
	Action
}
