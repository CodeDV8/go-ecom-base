package EComBase

type Plugin interface {
	Init(interface{})
	Done(interface{})
}
