package gopowerstore

type Alert struct {
	ID                 string  `json:"id"`
	EventCode          string  `json:"event_code"`
	Severity           string  `json:"severity"`
	ResourceType       string  `json:"resource_type"`
	ResourceName       string  `json:"resource_name"`
	Description        string  `json:"description_l10n"`
	GeneratedTimestamp string  `json:"generated_timestamp"`
	RaisedTimestamp    string  `json:"raised_timestamp"`
	ClearedTimestamp   string  `json:"cleared_timestamp"`
	State              string  `json:"state"`
	Events             []Event `json:"events"`
}

func (e *Alert) Fields() []string {
	return []string{"id", "severity", "event_code", "resource_name", "generated_timestamp",
		"cleared_timestamp", "raised_timestamp", "state", "events", "description_l10n", "resource_type"}
}
