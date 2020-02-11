package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(u string, form url.Values) ([]byte, error) {
	// header
	httpHeader := http.Header{}
	httpHeader.Add("Content-Type", "application/x-www-form-urlencoded")

	// form
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}
	req.Header = httpHeader

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return d, nil
}
