package exegraph

import (
	"encoding/json"
)

// Message object
type Message struct {
	LastID     string          `json:"last_id,omitempty"`
	ExecutedID []string        `json:"executed_ids,omitempty"`
	Data       json.RawMessage `json:"data,omitempty"`
}

// Validate message data
func (m Message) Validate() error {
	if len(m.Data) < 1 {
		return ErrMessageEmpty
	}
	return nil
}

// MarkExecuted task
func (m *Message) MarkExecuted(id string) {
	if !m.Executed(id) {
		m.ExecutedID = append(m.ExecutedID, id)
	}
}

// Executed task
func (m Message) Executed(id string) bool {
	for _, eid := range m.ExecutedID {
		if eid == id {
			return true
		}
	}
	return false
}

// Unmarshal message
func (m *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// UnmarshalData raw data to object
func (m Message) UnmarshalData(l interface{}) error {
	return json.Unmarshal(m.Data, l)
}

// SetData value
func (m *Message) SetData(d interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(d); err == nil {
		m.Data = json.RawMessage(data)
	}
	return
}
