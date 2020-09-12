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
					Text: fmt.Sprintf("*input: %s*へのおすすめメンターランキング", values.Get("text")),
				},
			},
			{
				Type: "section",
				Text: response.SlackText{
					Type: "mrkdwn",
					Text: "<https://techbowl.co.jp/techtrain/mentors/78|黒澤拓磨>\n :star::star::star::star::star:",
				},
				Accessory: &response.SlackAccessory{
					Type:     "image",
					ImageURL: "https://techbowl-production.s3.ap-northeast-1.amazonaws.com/mentor-profile-image/27b892da10905f353cf1a6835440a089.jpg",
					AltText:  "hoge",
				},
			},
			{
				Type: "section",
				Text: response.SlackText{
					Type: "mrkdwn",
					Text: "<https://techbowl.co.jp/techtrain/mentors/79|石塚 崇寛>\n :star::star::star::star:",
				},
				Accessory: &response.SlackAccessory{
					Type:     "image",
					ImageURL: "https://techbowl-production.s3.ap-northeast-1.amazonaws.com/mentor-profile-image/9b6428a47821e88eded480ce2377669b.jpg",
					AltText:  "hoge",
				},
			},
			{
				Type: "section",
				Text: response.SlackText{
					Type: "mrkdwn",
					Text: "<https://techbowl.co.jp/techtrain/mentors/67|白川 みちる>\n :star::star::star:",
				},
				Accessory: &response.SlackAccessory{
					Type:     "image",
					ImageURL: "https://techbowl-production.s3.ap-northeast-1.amazonaws.com/mentor-profile-image/81f80b447359317aadad05102fc39759.jpg",
					AltText:  "hoge",
				},
			},
		},
	}
	data, err := json.Marshal(&res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
