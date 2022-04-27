package khl

import "encoding/json"

// In this file is all weird solutions for silly KHL official developers.

// UnmarshalJSON deals with untyped mention in custom message in event.
func (m *EventCustomMessage) UnmarshalJSON(bytes []byte) error {
	type tempECM EventCustomMessage
	t := struct {
		tempECM
		Mention []json.RawMessage `json:"mention"`
	}{}
	err := json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}
	t.tempECM.Mention = make([]string, len(t.Mention))
	for index, item := range t.Mention {
		if item[0] == '"' {
			t.tempECM.Mention[index] = string(item[1 : len(item)-1])
		} else {
			t.tempECM.Mention[index] = string(item)
		}
	}
	*m = EventCustomMessage(t.tempECM)
	return nil
}
