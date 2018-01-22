package exegraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	graph := &Graph{}
	t1, _ := graph.SetItem("task1", task("task1"))
	t11, _ := graph.SetItem("task1.1", task("task1.1"), t1) // After task1
	t12, _ := graph.SetItem("task1.2", task("task1.2"), t1) // After task1
	graph.SetItem("task1.1.1", task("task1.1.1"), t11, t12) // After 1.1 & 1.2

	msg := Message{}
	assert.NoError(t, graph.Execute(&msg))
	assert.NoError(t, graph.Execute(&msg))
	assert.NoError(t, graph.Execute(&msg))
	assert.NoError(t, graph.Execute(&msg))
}

func task(code string) ExecuterFnk {
	return func(m *Message) (interface{}, error) {
		var list []string
		m.UnmarshalData(&list)
		fmt.Println("Execute", code)
		return append(list, code), nil
	}
}
