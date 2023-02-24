package destroy_instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
)

func DestroyInstance(region string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create EC2 service client
	svc := ec2.New(sess)

	describeResult, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", describeResult)
	}

	//look for a instance which has tag Name=minecraft and get its instance id
	var instanceId string

	for _, reservation := range describeResult.Reservations {
		for _, instance := range reservation.Instances {
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" && *tag.Value == "Minecraft" {
					instanceId = *instance.InstanceId
					break
				}
			}
		}
	}

	deleteResult, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		InstanceIds: aws.StringSlice([]string{instanceId}),
	})

    if err != nil {
        fmt.Println("Could not delete instance", err)
        return
    }

    fmt.Println("Deleting instance", *deleteResult.TerminatingInstances[0].InstanceId)

}
