package create_instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
	"log"
)

var (
	instanceType string
)

func CreateInstance() {
	fmt.Println("Please enter the instance type: ")
	fmt.Scanln(&instanceType)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create EC2 service client
	svc := ec2.New(sess)

	securityGroupId := createSecurityGroup()

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:          aws.String("ami-04f70e24ce4712a3c"),
		InstanceType:     aws.String(instanceType),
		MinCount:         aws.Int64(1),
		MaxCount:         aws.Int64(1),
		SecurityGroupIds: aws.StringSlice([]string{*securityGroupId}),
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return
	}

	fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("Minecraft"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	fmt.Println("Your ip to connect is: ", *runResult.Instances[0].PublicIpAddress )
}
