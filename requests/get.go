package requests

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"codenation/models"
)

type Get struct{}

func NewGet() *Get {
	return &Get{}
}

func (p Get) Json() models.Answer{

	resp, err:= http.Get("https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=d6186c8cf73335b77666fe45e3171ec2dc56e21d")
	if err != nil {
    log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
    log.Fatal(err)
	}

	var answer models.Answer
	err = json.Unmarshal([]byte(string(body)), &answer)
	return answer
}