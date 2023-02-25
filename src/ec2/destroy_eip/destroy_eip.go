package destroy_eip

import (
    "fmt"
	"os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func DestroyEIP(eip string) {
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