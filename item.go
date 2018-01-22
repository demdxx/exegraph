package exegraph

type Executer interface {
	Exec(msg *Message) (interface{}, error)
}

// ExecuterFnk function
type ExecuterFnk func(msg *Message) (interface{}, error)

// Exec the function
func (f ExecuterFnk) Exec(msg *Message) (interface{}, error) {
	return f(msg)
}

// Item graph object
type Item struct {
	ID        string
	Relations []*Item
	Executer  Executer
}

// CanGoNext execute
func (it *Item) CanGoNext(msg *Message) bool {
	if msg.Executed(it.ID) {
		return false
	}
	for _, r := range it.Relations {
		if !msg.Executed(r.ID) {
			return false
		}
	}
	return true
}
