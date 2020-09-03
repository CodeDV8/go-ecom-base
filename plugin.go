package EComBase

// Plugin - The definition of the plugin interface
type Plugin interface {
	Init(interface{})
	Done(interface{})
}
