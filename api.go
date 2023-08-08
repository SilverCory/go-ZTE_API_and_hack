package go_ZTE_API_and_hack

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

// Login - ZTE login
func Login(password string) error {
	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":   []string{"false"},
			"goformId": []string{"LOGIN"},
			"password": []string{base64.URLEncoding.EncodeToString([]byte(password))},
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

// Logout - ZTE logout
func Logout() error {
	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":   []string{"false"},
			"goformId": []string{"LOGOFF"},
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

// GetSMSMessages - ZTE get SMS messages
func GetSMSMessages() ([]string, error) {
	// curl -s --header "Referer: http://<modem_ip>/index.html" http://<modem_ip>//goform/goform_get_cmd_process\?isTest\=false\&cmd\=sms_data_total\&page=0\&data_per_page\=500\&mem_store\=1\&tags\=10\&order_by\=order+by+id+desc
	// Build the request
	//req, err := buildSetCMDRequest(
	//	nil,
	//	url.Values{
	//		"isTest":      []string{"false"},
	//		"cmd":         []string{"sms_data_total"},
	//		"notCallback": []string{"true"},
	//	})
	//if err != nil {
	//	return err
	//}
	//
	//// Send the request
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	return err
	//}
	//
	//// Check the response
	//if resp.StatusCode != 200 {
	//	return fmt.Errorf("got status code %d", resp.StatusCode)
	//}
	//
	//return nil
	panic("not implemented")
}

// DeleteSMSMessages - ZTE delete SMS messages
func DeleteSMSMessages(messageIDs []string) error {
	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":      []string{"false"},
			"goformId":    []string{"DELETE_SMS"},
			"msgId":       messageIDs,
			"notCallback": []string{"true"},
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

// SendSMSMessage - ZTE send SMS message
func SendSMSMessage(phoneNumber string, message string) error {

	// URLEncode the phoneNumber
	phoneNumber = url.QueryEscape(phoneNumber)

	// Hex encode the message
	message = fmt.Sprintf("%X", message)

	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":      []string{"false"},
			"goformId":    []string{"SEND_SMS"},
			"Number":      []string{phoneNumber},
			"notCallback": []string{"true"},
			"sms_time":    []string{""},   // TODO: Copilot: sms_time is in the format "2017-01-01 00:00:00"?
			"id":          []string{"-1"}, // TODO: Copilot: id is -1 for new messages and a positive number for existing messages?
			"encode_type": []string{"UNICODE"},
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

// ToggleWifi - ZTE disable wifi
func ToggleWifi(on bool) error {
	var wifiEnabled string = "0"
	if on {
		wifiEnabled = "1"
	}

	// Build the request
	req, err := buildSetCMDRequest(
		nil,
		url.Values{
			"isTest":        []string{"false"},
			"goformId":      []string{"SET_WIFI_INFO"},
			"m_ssid_enable": []string{"0"},
			"wifiEnabled":   []string{wifiEnabled},
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
