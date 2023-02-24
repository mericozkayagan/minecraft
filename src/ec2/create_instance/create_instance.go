package create_instance

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
    "log"
)

var (
	instanceType    string
)

func init() {
    fmt.Println("Please enter the region you want to create your server: ")
	fmt.Scanln(&instanceType)

}


func CreateInstance(region string) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )

    // Create EC2 service client
    svc := ec2.New(sess)

    // Specify the details of the instance that you want to create.
    runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
        // An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
        ImageId:      aws.String("ami-e7527ed7"),
        InstanceType: aws.String(instanceType),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    })

    if err != nil {
        fmt.Println("Could not create instance", err)
        return
    }

    fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

    // Add tags to the created instance
    _, errtag := svc.CreateTags(&ec2.CreateTagsInput{
        Resources: []*string{runResult.Instances[0].InstanceId},
        Tags: []*ec2.Tag{
            {
                Key:   aws.String("Name"),
                Value: aws.String("Minecraft"),
            },
        },
    })
    if errtag != nil {
        log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
        return
    }

    fmt.Println("Successfully tagged instance")
}