package reboot_instance

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mericozkayagan/minecraft/src/ec2/filter_by_tag"
)

func RebootInstance() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	svc := ec2.New(sess)

	instanceId, _ := filter_by_tag.FilterByTag()

	// We set DryRun to true to check to see if the instance exists and we have the
	// necessary permissions to monitor the instance.
	input := &ec2.RebootInstancesInput{
		InstanceIds: []*string{
			aws.String(*instanceId),
		},
		DryRun: aws.Bool(true),
	}
	_, err := svc.RebootInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is `DryRunOperation` it means we have the necessary
	// permissions to Start this instance
	if ok && awsErr.Code() == "DryRunOperation" {
		// Let's now set dry run to be false. This will allow us to reboot the instances
		input.DryRun = aws.Bool(false)
		_, err = svc.RebootInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			_, publicIp := filter_by_tag.FilterByTag()
			fmt.Println("Successfully started the instance with the IP: ", publicIp)
		}
	} else { // This could be due to a lack of permissions
		fmt.Println("Error", err)
	}
}
