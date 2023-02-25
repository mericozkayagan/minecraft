package start_stop_instance

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func StartStopInstance(command string, instanceId *string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
	// Create new EC2 client
	svc := ec2.New(sess)

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}

	// Turn monitoring on
	if strings.TrimSpace(strings.ToLower(command)) == "start" {
		// We set DryRun to true to check to see if the instance exists and we have the
		// necessary permissions to monitor the instance.
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(*instanceId),
			},
			DryRun: aws.Bool(true),
		}
		result, err := svc.StartInstances(input)
		awsErr, ok := err.(awserr.Error)

		// If the error code is `DryRunOperation` it means we have the necessary
		// permissions to Start this instance
		if ok && awsErr.Code() == "DryRunOperation" {
			// Let's now set dry run to be false. This will allow us to start the instances
			input.DryRun = aws.Bool(false)
			result, err = svc.StartInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Successfully started the instance: ", result.StartingInstances[0].InstanceId)
			}
		} else { // This could be due to a lack of permissions
			fmt.Println("Error", err)
		}
	} else if strings.TrimSpace(strings.ToLower(command)) == "stop" { // Turn instances off
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(*instanceId),
			},
			DryRun: aws.Bool(true),
		}
		result, err := svc.StopInstances(input)
		awsErr, ok := err.(awserr.Error)
		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			result, err = svc.StopInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Successfully stopped the instance: ", result.StoppingInstances[0].InstanceId)
			}
		} else {
			fmt.Println("Error", err)
		}
	}
}
