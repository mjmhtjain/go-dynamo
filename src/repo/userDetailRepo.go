package repo

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mjmhtjain/go-dynamo/src/model"
)

var (
	accessKeyID        string = ""
	secretAccessKey    string = ""
	awsRegion          string = ""
	awsGatewayEndpoint string = ""
)

type UserDetailRepo interface {
	FindById(id string) (*model.UserDetail, error)
	Save(*model.UserDetail)
}

type UserDetailRepoImpl struct {
	svc *dynamodb.DynamoDB
}

func NewUserDetailRepo() UserDetailRepo {
	return &UserDetailRepoImpl{
		svc: initDynamoSession(),
	}
}

func (u *UserDetailRepoImpl) FindById(id string) (*model.UserDetail, error) {
	tableName := "UserDetail"

	result, err := u.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find the record"
		return nil, errors.New(msg)
	}

	userDetail := model.UserDetail{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &userDetail)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return &userDetail, nil
}

func (u *UserDetailRepoImpl) Save(*model.UserDetail) {

}

func initDynamoSession() *dynamodb.DynamoDB {

	accessKeyID = os.Getenv("aws.access-key")
	secretAccessKey = os.Getenv("aws.secret-key")
	awsRegion = os.Getenv("aws.dynamodb.region")
	awsGatewayEndpoint = os.Getenv("aws.dynamodb.endpoint")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Endpoint:    aws.String(awsGatewayEndpoint),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	}))

	return dynamodb.New(sess)
}
