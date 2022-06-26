package slack

// Metadata contains metadata information
type Metadata struct {
	EventType    string `json:"event_type"`
	EventPayload map[string]string `json:"event_payload"`
}
