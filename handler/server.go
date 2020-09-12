package handler

import (
	"Slack_SlashCommand_Sample/response"
	"encoding/json"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	res := response.SlackMessage{
		Blocks: []response.SlackBlock{
			{
				Type: "section",
				Text: response.SlackText{
					Type: "mrkdwn",
					Text: values.Get("text"),
				},
			},
		},
	}
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
