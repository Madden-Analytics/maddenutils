// Package maddenutils contains functions to interact with Madden Analytics API v1
package maddenutils

import (
	"bytes"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Request - Function to make general request to the Madden API
func Request(requestType string, endpoint string, auth string, json []byte) (int, []byte) {

	req, _ := http.NewRequest(requestType, endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+auth)
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": requestType,
		"endpoint":    endpoint,
		"body":        string(json),
	}).Debug("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": requestType,
			"endpoint":    endpoint,
			"body":        string(json),
			"response":    err,
		}).Error("ERROR")
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"requesttype":  requestType,
			"endpoint":     endpoint,
			"errorMessage": err,
		}).Error("ERROR Reading Response")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		log.WithFields(log.Fields{
			"requesttype":  requestType,
			"endpoint":     endpoint,
			"errorMessage": err,
		}).Fatal("Unauthorized")
	}

	return resp.StatusCode, response
}
