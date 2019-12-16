package common

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

// GetAwsCredentials - Common Utility Function to Get AWS Credentials from Local Workstation
func GetAwsCredentials() (*credentials.Credentials, error) {
	DebugMessage("[common] func GetAwsCredentials()")
	// Retrieve the credentials from the Environment
	fmt.Println("GetAwsCredentials() --> Getting AWS Credentials from Environment")
	envCreds := credentials.NewEnvCredentials()
	envCredsValue, err := envCreds.Get()
	if err != nil {
		// Error
		fmt.Printf("GetAwsCredentials() --> ERROR: Unable to retrieve AWS Credentials from Environment --> %v\n", err)
		// Example:
		// ERROR: EnvAccessKeyNotFound: AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment

		// Retrieve the credentials from the Shared Credentials File Default Profile
		fmt.Println("GetAwsCredentials() --> Getting AWS Credentials from Shared File (e.g. ~/.aws/credentials")
		homedir := GetUserHomeDir()
		fileCreds := credentials.NewSharedCredentials(homedir+"/.aws/credentials", "default")
		fileCredsValue, err := fileCreds.Get()
		if err != nil {
			// Error
			fmt.Printf("GetAwsCredentials() --> ERROR: Unable to retrieve AWS Credentials from Shared File --> %v\n", err)
			fmt.Printf("GetAwsCredentials() --> ERROR: Unable to retrieve AWS Credentials from Environment or Shared File\n")
			return nil, err
			// Examples:
			// ERROR: EnvAccessKeyNotFound: AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment
			// ERROR: EnvSecretNotFound: AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment
		} else {
			// Success: Return Shared File Credentials
			fmt.Println("GetAwsCredentials() --> SUCCESS: Returning Credentials from Shared File")
			fmt.Printf("\nFILE CREDS:\n\n%v\n", fileCredsValue)
			return fileCreds, nil
		}

	} else {
		// Success: Return Environment Credentials
		fmt.Println("GetAwsCredentials() --> SUCCESS: Returning Credentials from Environment")
		fmt.Printf("\nENV CREDS:\n\n%v\n", envCredsValue)
		return envCreds, nil
		// Example:
		// CREDS: {<ACCESS_KEY> <SECRET_KEY> <SESSION_TOKEN> EnvProvider}

	}
}
