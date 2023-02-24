// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Reboots an Amazon EC2 instance.]
// snippet-keyword:[Amazon Elastic Compute Cloud]
// snippet-keyword:[RebootInstances function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[ec2]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2018-03-16]
/*
   Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package reboot_instance

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

// Usage:
// go run main.go <instance id>
func RebootInstance() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create new EC2 client
    svc := ec2.New(sess)

	//look for a instance which has tag Name=minecraft and get its instance id
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

	// We set DryRun to true to check to see if the instance exists and we have the
    // necessary permissions to monitor the instance.
    input := &ec2.RebootInstancesInput{
        InstanceIds: []*string{
            aws.String(instanceId),
        },
        DryRun: aws.Bool(true),
    }
    result, err := svc.RebootInstances(input)
    awsErr, ok := err.(awserr.Error)

    // If the error code is `DryRunOperation` it means we have the necessary
    // permissions to Start this instance
    if ok && awsErr.Code() == "DryRunOperation" {
        // Let's now set dry run to be false. This will allow us to reboot the instances
        input.DryRun = aws.Bool(false)
        result, err = svc.RebootInstances(input)
        if err != nil {
            fmt.Println("Error", err)
        } else {
            fmt.Println("Success", result)
        }
    } else { // This could be due to a lack of permissions
        fmt.Println("Error", err)
    }
}