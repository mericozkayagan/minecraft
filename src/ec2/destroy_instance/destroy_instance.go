package destroy_instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
)

func DestroyInstance(instanceId *string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create EC2 service client
	svc := ec2.New(sess)

	describeResult, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", describeResult)
	}

	deleteResult, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: aws.StringSlice([]string{*instanceId}),
	})

	if err != nil {
		fmt.Println("Could not delete instance", err)
		return
	}

	fmt.Println("Deleting instance: ", *deleteResult.TerminatingInstances[0].InstanceId)

}
