package run

type XpFunc func(uint) uint

var XpFuncs = map[uint]XpFunc{
	1: XpFuncSlow,
	2: XpFuncFast,
}

func XpFuncSlow(level uint) uint {
	return 100 + (level * 10)
}

func XpFuncFast(level uint) uint {
	return 50 + (level * 10)
}
