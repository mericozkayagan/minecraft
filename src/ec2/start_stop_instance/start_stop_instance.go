package start_stop_instance

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mericozkayagan/minecraft/src/ec2/filter_by_tag"
)

func StartStopInstance(command string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create new EC2 client
	svc := ec2.New(sess)

	instanceId, _ := filter_by_tag.FilterByTag()

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
		_, err := svc.StartInstances(input)
		awsErr, ok := err.(awserr.Error)

		// If the error code is `DryRunOperation` it means we have the necessary
		// permissions to Start this instance
		if ok && awsErr.Code() == "DryRunOperation" {
			// Let's now set dry run to be false. This will allow us to start the instances
			input.DryRun = aws.Bool(false)
			_, err = svc.StartInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				_, publicIp := filter_by_tag.FilterByTag()
				fmt.Println("Successfully started the instance with the config: ", publicIp)
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
		_, err := svc.StopInstances(input)
		awsErr, ok := err.(awserr.Error)
		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			_, err = svc.StopInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Successfully stopped the instance.")
			}
		} else {
			fmt.Println("Error", err)
		}
	}
}
