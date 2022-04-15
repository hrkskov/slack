package slack

// Metadata contains metadata information
type Metadata struct {
	EventType    string `json:"event_type,omitempty"`
	EventPayload string `json:"event_payload,omitempty"`
}
