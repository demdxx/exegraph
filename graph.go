package exegraph

// Base object
type Graph struct {
	Items []*Item
}

// SetItem executer
func (g *Graph) SetItem(id string, ex Executer, relations ...*Item) (*Item, error) {
	for _, it := range g.Items {
		if it.ID == id {
			it.Executer = ex
			it.Relations = relations
			return it, nil
		}
	}
	it := &Item{ID: id, Executer: ex, Relations: relations}
	g.Items = append(g.Items, it)
	return it, nil
}

// Item by ID
func (g *Graph) Item(id string) *Item {
	for _, it := range g.Items {
		if it.ID == id {
			return it
		}
	}
	return nil
}

// Execute message event
func (g *Graph) Execute(msg *Message) error {
	for _, it := range g.Items {
		if it.CanGoNext(msg) {
			if data, err := it.Executer.Exec(msg); err != nil {
				return err
			} else {
				msg.LastID = it.ID
				msg.MarkExecuted(it.ID)
				return msg.SetData(data)
			}
		}
	}
	return nil
}
