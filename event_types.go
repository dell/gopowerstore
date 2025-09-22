package gopowerstore

type Event struct {
	ID           string `json:"id"`
	Code         string `json:"event_code"`
	Severity     string `json:"severity"`
	ResourceName string `json:"resource_name"`
	Description  string `json:"description_l10n"`
	Timestamp    string `json:"generated_timestamp"`
}

func (e *Event) Fields() []string {
	return []string{"id", "severity", "event_code", "description_l10n", "resource_name", "generated_timestamp"}
}
