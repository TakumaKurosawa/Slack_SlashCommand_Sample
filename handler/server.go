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
				Accessory: nil,
			},
			{
				Type: "section",
				Text: response.SlackText{
					Type: "mrkdwn",
					Text: "https://techbowl.co.jp/techtrain/mentors/78",
				},
				Accessory: response.SlackAccessory{
					Type:     "image",
					ImageURL: "https://techbowl-production.s3.ap-northeast-1.amazonaws.com/mentor-profile-image/27b892da10905f353cf1a6835440a089.jpg",
					AltText:  "hoge",
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
