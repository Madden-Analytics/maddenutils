package maddenutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

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
func PutConfig(configName string, region string, value string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	_, err = ssmsvc.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(configName),
		Value:     aws.String(value),
		Overwrite: aws.Bool(true),
		Type:      aws.String("String"),
	})

	return err
}

// DeleteConfig - Deletes config value (JSON) to AWS Parameter Store
func DeleteConfig(configName string, region string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))

	_, err = ssmsvc.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(configName),
	})

	return err
}
