package validate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Email(email string) error {
	url := "https://isitarealemail.com/api/email/validate?email=" + url.QueryEscape(email)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	EmailStatus := struct {
		Status string `json:"status,omitempty"`
	}{}

	json.Unmarshal(body, &EmailStatus)
	
	if EmailStatus.Status == "valid" {
		return nil
	}else {
		return fmt.Errorf("invalid email")
	}
}
