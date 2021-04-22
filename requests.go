// Package maddenutils contains functions to interact with Madden Analytics API v1
package maddenutils

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//Get ~ Function to GET data from Madden API
func Get(endpoint string, auth string) (int, []byte) {

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+auth)
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "GET",
		"endpoint":    endpoint,
		"body":        nil,
	}).Debug("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "GET/Error",
			"endpoint":    endpoint,
			"body":        nil,
			"response":    err,
		}).Error("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"requesttype": "GET/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Error("ERROR")

	} else {
		log.WithFields(log.Fields{
			"requesttype": "GET/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    "",
		}).Debug("OK")
		return resp.StatusCode, response
	}
	return resp.StatusCode, response
}

//Post - Function to POST data to Madden API
func Post(endpoint string, auth string, json []byte) (int, []byte) {

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+auth)
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "POST",
		"endpoint":    endpoint,
		"body":        string(json),
	}).Debug("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "POST/Error",
			"endpoint":    endpoint,
			"body":        string(json),
			"response":    err,
		}).Error("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"requesttype": "POST/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Error("ERROR")

	} else {
		log.WithFields(log.Fields{
			"requesttype": "POST/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Debug("OK")
		return resp.StatusCode, response
	}
	return resp.StatusCode, response
}

//Put - Function to PUT data to Madden API
func Put(endpoint string, auth string, json []byte) (int, []byte) {

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+auth)
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "PUT",
		"endpoint":    endpoint,
		"body":        string(json),
	}).Debug("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Error",
			"endpoint":    endpoint,
			"body":        string(json),
			"response":    err,
		}).Error("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Error("ERROR")

	} else {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Debug("OK")
		return resp.StatusCode, response
	}
	return resp.StatusCode, response
}

//Patch - Function to Patch data to Madden API
func Patch(endpoint string, auth string, json []byte) []byte {

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+auth)
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "PATCH",
		"endpoint":    endpoint,
		"body":        string(json),
	}).Debug("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "PATCH/Error",
			"endpoint":    endpoint,
			"body":        string(json),
			"response":    err,
		}).Error("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		log.WithFields(log.Fields{
			"requesttype": "PATCH/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Error("ERROR")

	} else {
		log.WithFields(log.Fields{
			"requesttype": "PATCH/Response",
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Debug("OK")
		return response
	}
	return response
}

//Request - Function to make general request to the Madden API
func Request(requestType string, endpoint string, auth string, json []byte) (int, []byte) {

	req, err := http.NewRequest(requestType, endpoint, bytes.NewBuffer(json))
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

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

		log.WithFields(log.Fields{
			"requesttype": requestType,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Debug("OK")

	} else {

		log.WithFields(log.Fields{
			"requesttype": requestType,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Error("ERROR")

	}
	return resp.StatusCode, response
}
