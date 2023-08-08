package go_ZTE_API_and_hack

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ModemIP string = "192.168.0.1"

// buildSetCMDRequest builds a request with the given url and form parameters
func buildSetCMDRequest(urlParams url.Values, formParams url.Values) (*http.Request, error) {
	var body io.Reader = nil
	if formParams != nil {
		body = bytes.NewBufferString(formParams.Encode())
	}

	req, err := http.NewRequest(fmt.Sprintf("http://%s/goform/goform_set_cmd_process%s", ModemIP, urlParams.Encode()), "POST", body)
	if err != nil {
		return req, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Set("Referer", fmt.Sprintf("http://%s/index.html", ModemIP))

	return req, err
}
