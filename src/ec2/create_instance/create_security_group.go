package create_instance

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func createSecurityGroup() *string {
	namePtr := "Minecraft-sg"
	descPtr := "The ports required for the Minecraft server"

	flag.Parse()

	if namePtr == "" || descPtr == "" {
		flag.PrintDefaults()
		exitErrorf("Group name and description require")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	// If the VPC ID wasn't provided in the CLI retrieve the first in the account.
	// snippet-start:[ec2.go.create_security_group.vpcid]
	// Get a list of VPCs so we can associate the group with the first VPC.
	result, err := svc.DescribeVpcs(nil)
	if err != nil {
		exitErrorf("Unable to describe VPCs, %v", err)
	}
	if len(result.Vpcs) == 0 {
		exitErrorf("No VPCs found to associate security group with.")
	}

	vpcIDPtr := aws.StringValue(result.Vpcs[0].VpcId)

	// snippet-end:[ec2.go.create_security_group.vpcid]

	// Create the security group with the VPC, name and description.
	// snippet-start:[ec2.go.create_security_group.create]
	createRes, err := svc.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(namePtr),
		Description: aws.String(descPtr),
		VpcId:       aws.String(vpcIDPtr),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidVpcID.NotFound":
				exitErrorf("Unable to find VPC with ID %q.", vpcIDPtr)
			case "InvalidGroup.Duplicate":
				exitErrorf("Security group %q already exists.", namePtr)
			}
		}
		exitErrorf("Unable to create security group %q, %v", namePtr, err)
	}

	fmt.Printf("Created security group %s with VPC %s.\n",
		aws.StringValue(createRes.GroupId), vpcIDPtr)
	// snippet-end:[ec2.go.create_security_group.create]

	// Add permissions to the security group
	// snippet-start:[ec2.go.create_security_group.permissions]
	output, err := svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupName: aws.String(namePtr),
		IpPermissions: []*ec2.IpPermission{
			// Can use setters to simplify seting multiple values without the
			// needing to use aws.String or associated helper utilities.
			(&ec2.IpPermission{}).
				SetIpProtocol("tcp").
				SetFromPort(25565).
				SetToPort(25565).
				SetIpRanges([]*ec2.IpRange{
					{CidrIp: aws.String("0.0.0.0/0")},
				}),
			(&ec2.IpPermission{}).
				SetIpProtocol("tcp").
				SetFromPort(22).
				SetToPort(22).
				SetIpRanges([]*ec2.IpRange{
					(&ec2.IpRange{}).
						SetCidrIp("0.0.0.0/0"),
				}),
		},
	})
	if err != nil {
		exitErrorf("Unable to set security group %q ingress, %v", namePtr, err)
	}

	fmt.Println("Successfully set security group ingress")

	return output.SecurityGroupRules[0].GroupId
}

// snippet-start:[ec2.go.create_security_group.exit]
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// snippet-end:[ec2.go.create_security_group.exit]
// snippet-end:[ec2.go.create_security_group.complete]
