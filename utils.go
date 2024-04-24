package maddenutils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"

	zerowidth "github.com/trubitsyn/go-zero-width"
)

// timeSlice - Functions for sorting times
type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// GetAuth generates bearer token from Madden
func GetAuth(baseURL string, accountID string, key string) (*string, error) {

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
		return nil, errors.New("failed to get Madden Auth")
	} else {
		json.Unmarshal(response, &maddenAuth)
	}

	return &maddenAuth.AccessToken, nil
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

// Read the config file from the current directory and marshal
// into the conf config struct.
func GetConf[T any](configName string) (*T, error) {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	//Set Config File setting
	viper.SetConfigType("json")

	// Get config from AWS
	configFile, err := GetConfig(configName, "eu-north-1")
	if err != nil {
		return nil, err
	}

	// Searches for config file in given paths and read it
	if err := viper.ReadConfig(bytes.NewBuffer(configFile)); err != nil {
		return nil, err
	}

	conf := new(T)
	err = viper.Unmarshal(conf)

	return conf, err
}

func GetLatestDate(dateSlice timeSlice) time.Time {
	latest := time.Time{}
	if len(dateSlice) > 0 {
		sort.Sort(sort.Reverse(dateSlice))
		latest = dateSlice[0]
	}
	return latest
}
