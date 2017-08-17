package behavior

type Status int

const (
	StatusFailure Status = iota
	StatusSuccess
	StatusRunning
)

func (s Status) String() string {
	switch s {
	case StatusSuccess:
		return "成功"
	case StatusFailure:
		return "失败"
	case StatusRunning:
		return "运行中"
	default:
		return "未知状态"
	}
}
