package maddenutils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// GetAuth generates bearer token from Madden
func GetAuth(baseURL string, accountID string, key string) string {
	maddenKey := APIKey{
		AccountID: accountID,
		APIKey:    key,
	}

	maddenKeyJSON, _ := json.Marshal(maddenKey)

	// Get Madden Bearer Token
	_, maddenAuthResponse := Post(
		baseURL+"/auth",
		"",
		maddenKeyJSON,
	)
	var maddenAuth MaddenBearer
	json.Unmarshal(maddenAuthResponse, &maddenAuth)

	return maddenAuth.AccessToken
}

// BatchTransactions batches transaction struct into given size
func BatchTransactions(batchSize int, batchSlice []MaddenTransaction) [][]MaddenTransaction {

	batches := make([][]MaddenTransaction, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// BatchStocks batches stock struct into given size
func BatchStocks(batchSize int, batchSlice []MaddenStock) [][]MaddenStock {

	batches := make([][]MaddenStock, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// BatchPIM batches pim struct into given size
func BatchPIM(batchSize int, batchSlice []MaddenPIM) [][]MaddenPIM {

	batches := make([][]MaddenPIM, 0, (len(batchSlice)+batchSize-1)/batchSize)
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

// GetConfig - Returns config values (JSON) from AWS Parameter Store
func GetConfig(configName string, region string) []byte {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(configName),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	value := []byte(*param.Parameter.Value)
	return value
}

// PutConfig - Updates config value (JSON) to AWS Parameter Store
func PutConfig(configName string, region string, value string) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	param, err := ssmsvc.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(configName),
		Value:     aws.String(value),
		Overwrite: aws.Bool(true),
		Type:      aws.String("Śtring"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(param)
}
