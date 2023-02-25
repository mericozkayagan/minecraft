package assoicate_eip

import (
    "fmt"
	"os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func AssociateEIP(instanceId string) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

    // Create an EC2 service client.
    svc := ec2.New(sess)

	// Attempt to allocate the Elastic IP address.
    allocRes, err := svc.AllocateAddress(&ec2.AllocateAddressInput{
        Domain: aws.String("vpc"),
    })
    if err != nil {
        exitErrorf("Unable to allocate IP address, %v", err)
    }

    // Associate the new Elastic IP address with an existing EC2 instance.
    assocRes, err := svc.AssociateAddress(&ec2.AssociateAddressInput{
        AllocationId: allocRes.AllocationId,
        InstanceId:   aws.String(instanceId),
    })
    if err != nil {
        exitErrorf("Unable to associate IP address with %s, %v",
            instanceId, err)
    }

    fmt.Printf("Successfully allocated %s with instance %s.\n\tallocation id: %s, association id: %s\n",
        *allocRes.PublicIp, instanceId, *allocRes.AllocationId, *assocRes.AssociationId)
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}