package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/zolunga/go-back-test/awsgo"
	"github.com/zolunga/go-back-test/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var data models.Secret
	fmt.Print("Secret", secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	json.Unmarshal([]byte(*key.SecretString), &data)
	fmt.Println("Lectura ok", data)
	return data, nil
}
