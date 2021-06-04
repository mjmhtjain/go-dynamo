package repo

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mjmhtjain/go-dynamo/src/model"
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
		svc: createDynamoSession(),
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

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
	}))

	return dynamodb.New(sess)
}
