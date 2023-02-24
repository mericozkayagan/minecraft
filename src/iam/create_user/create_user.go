package create_user

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iam"
)

func CreateUser(region string) {
    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )

    // Create a IAM service client.
    svc := iam.New(sess)

    _, err = svc.GetUser(&iam.GetUserInput{
        UserName: aws.String("minecraftRole"),
    })

    if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
        result, err := svc.CreateUser(&iam.CreateUserInput{
            UserName: aws.String("minecraftRole"),
        })

        if err != nil {
            fmt.Println("CreateUser Error", err)
            return
        }

        fmt.Println("Success", result)
    } else {
        fmt.Println("GetUser Error", err)
    }
}