package maddenutils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	zerowidth "github.com/trubitsyn/go-zero-width"
)

// GetAuth generates bearer token from Madden
func GetAuth(baseURL string, accountID string, key string) string {

	var maddenAuth Bearer
	maddenKey := APIKey{
		AccountID: accountID,
		APIKey:    key,
	}

	maddenKeyJSON, _ := json.Marshal(maddenKey)

	// Get Madden Bearer Token
	statusCode, response := Request(
		"POST",
		baseURL+"/auth",
		"",
		maddenKeyJSON,
	)
	if statusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"errorMessage": string(response),
		}).Fatal("error getting auth")
	} else {
		json.Unmarshal(response, &maddenAuth)
	}

	return maddenAuth.AccessToken
}

// GetBasicToken - Hashes a basic auth token
func GetBasicToken(user string, password string) string {

	auth := user + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))

}

// BatchTransactions batches transaction struct into given size
func BatchTransactions(batchSize int, batchSlice []Transaction) [][]Transaction {

	batches := make([][]Transaction, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// BatchStocks batches stock struct into given size
func BatchStocks(batchSize int, batchSlice []Stock) [][]Stock {

	batches := make([][]Stock, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// BatchPIM batches pim struct into given size
func BatchProducts(batchSize int, batchSlice []Product) [][]Product {

	batches := make([][]Product, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// StrToInt converts decimal string to int
func StrToInt(str string) int {
	nonFractionalPart := strings.Split(str, ".")
	theInt, _ := strconv.Atoi(nonFractionalPart[0])
	return theInt
}

// GetMD5Hash - Hashes 1/2 strings to md5
func GetMD5Hash(string1 string, string2 string) string {
	hasher := md5.New()
	hash := ""
	if len(string2) == 0 {
		hasher.Write([]byte(string1))
		hash = hex.EncodeToString(hasher.Sum(nil))
	} else {
		hasher.Write([]byte(string1 + "-" + string2))
		hash = hex.EncodeToString(hasher.Sum(nil))
	}
	return hash
}

func CleanString(s string) string {
	s = strings.TrimSpace(s)
	s = zerowidth.RemoveZeroWidthCharacters(s)
	return s
}
