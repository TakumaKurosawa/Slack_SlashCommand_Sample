package handler

import (
	"Slack_SlashCommand_Sample/response"
	"encoding/json"
	"fmt"
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
					Text: fmt.Sprintf("*%s*", values.Get("text")),
				},
			},
			{
				Type: "section",
				Text: response.SlackText{
					Type:        "mrkdwn",
					Text:        "https://techbowl.co.jp/techtrain/mentors/78",
					UnfurlLinks: true,
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
