package response

type SlackMessage struct {
	Blocks []SlackBlock `json:"blocks"`
}

type SlackBlock struct {
	Type string    `json:"type"`
	Text SlackText `json:"text"`
}

type SlackText struct {
	Type      string         `json:"type"`
	Text      string         `json:"text"`
	Accessory SlackAccessory `json:"accessory"`
}

type SlackAccessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}
