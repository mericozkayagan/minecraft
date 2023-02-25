package attach_policy

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iam"
)

func AttachPolicy() {
    // Initialize a session in us-west-2 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create a IAM service client.
    svc := iam.New(sess)

    var pageErr error
    policyName := "AdministratorAccess"
    policyArn := "arn:aws:iam::aws:policy/AdministratorAccess"

    // Paginate through all role policies. If our role exists on any role
    // policy we will set the pageErr and return false. Stopping the
    // pagination.
    err := svc.ListAttachedRolePoliciesPages(
        &iam.ListAttachedRolePoliciesInput{
            RoleName: aws.String("minecraftRole"),
        },
        func(page *iam.ListAttachedRolePoliciesOutput, lastPage bool) bool {
            if page != nil && len(page.AttachedPolicies) > 0 {
                for _, policy := range page.AttachedPolicies {
                    if *policy.PolicyName == policyName {
                        pageErr = fmt.Errorf("%s is already attached to this role", policyName)
                        return false
                    }
                }
                // We should keep paginating because we did not find our role
                return true
            }
            return false
        },
    )

    if pageErr != nil {
        fmt.Println("Error", pageErr)
        return
    }

    if err != nil {
        fmt.Println("Error", err)
        return
    }

    _, err = svc.AttachRolePolicy(&iam.AttachRolePolicyInput{
        PolicyArn: &policyArn,
        RoleName:  aws.String("minecraftRole"),
    })

    if err != nil {
        fmt.Println("Unable to attach role policy to role")
        return
    }
    fmt.Println("Role attached successfully")
}