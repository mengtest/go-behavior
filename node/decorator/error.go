package decorator

import "errors"

var ErrNilChildren = errors.New("children slice没初始化")
var ErrTooManyChild = errors.New("decorator节点应有且只有一个子节点")
