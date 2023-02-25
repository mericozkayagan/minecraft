// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[Allocates a static IP address and associates it with an Amazon EC2 instance.]
// snippet-keyword:[Amazon Elastic Compute Cloud]
// snippet-keyword:[AllocateAddress function]
// snippet-keyword:[AssociateAddress function]
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

package destroy_eip

import (
    "fmt"
	"os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

// Attempts to allocate an VPC Elastic IP Address for region. The IP
// address will be associated with the instance ID passed in.
//

func DestroyEIP(eip string) {

    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

    // Create an EC2 service client.
    svc := ec2.New(sess)

    // Attempt to allocate the Elastic IP address.
    _, err := svc.DisassociateAddress(&ec2.DisassociateAddressInput{
		AssociationId: aws.String(eip),
    })
    if err != nil {
        exitErrorf("Unable to allocate IP address, %v", err)
    }
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}