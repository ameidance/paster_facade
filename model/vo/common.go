package vo

type Request interface {
	CheckParams() bool
}
