package handler

import (
	"Slack_SlashCommand_Sample/reqbody"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
	}

	log.Println(r.Body)

	var reqBody reqbody.UserInput
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Hello!, %s", reqBody.Text)
}
