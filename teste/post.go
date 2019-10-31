package main
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func newfileUploadRequest(uri string, params map[string]string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("answer", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token=d6186c8cf73335b77666fe45e3171ec2dc56e21d", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func main() {
	path, _ := os.Getwd()
	path += "/answer.json"
	fmt.Println(path)
	extraParams := map[string]string{
		"file":       "answer",
	}
	request, err := newfileUploadRequest("https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token=d6186c8cf73335b77666fe45e3171ec2dc56e21d", extraParams, path)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

    resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}
