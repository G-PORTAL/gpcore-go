# GPCORE Golang Client

This is the official GPCORE Golang client. Please open an issue if you have found
a problem or having a question.

## Recommendations

- Golang 1.18 or higher

## Documentation

You can find the current API documentation [here](https://buf.build/gportal/gpcore).
Packages for Go and Node.js can be [downloaded](https://buf.build/gportal/gpcore/assets/main)
from the buf.build repository. If you want to get the raw ```.proto``` definition files,
pull it from [this](https://github.com/G-PORTAL/gpcore-proto) repository or use
the [buf](https://docs.buf.build/reference/cli/buf) tool:

```
buf export buf.build/gportal/gpcore -o .
```

## Service Account

You can create your own Service Account in the GPCORE Panel.

Once you have created the Service Account, you can add it to projects as a member.
The Service Account will only have access to Projects that the Service Account is a member of.


## Example usage

```go
package main

import (
	"context"
	"log"

	authv1 "buf.build/gen/go/gportal/gpcore/protocolbuffers/go/gpcore/api/auth/v1"
	cloudv1 "buf.build/gen/go/gportal/gpcore/protocolbuffers/go/gpcore/api/cloud/v1"
	"github.com/G-PORTAL/gpcore-co/pkg/gpcore/client"
	"github.com/G-PORTAL/gpcore-co/pkg/gpcore/client/auth"
)

func main() {
	conn, err := client.NewClient(
		// You can create your own Service Account in the GPCORE Panel.
		// The Service Account can be added to projects as a member.
		&auth.ProviderKeycloakClientAuth{
			ClientID:     "47d838c3-5738-4a67-aafb-48b364fab41b", // Set your Client ID
			ClientSecret: "MdErjhYSP4Nuq3h6qcYCMdErjhYSP4Nu", // Set your Client Secret
		},
	)
	if err != nil {
		log.Fatal("failed to create client:\n", err)
	}

	ctx := context.Background()
	user, err := conn.AuthClient().GetUser(ctx, &authv1.GetUserRequest{})
	if err != nil {
		log.Fatal("failed to fetch user information:\n", err)
	}
	log.Println("User ID: ", user.User.Id)

	projects, err := conn.CloudClient().ListProjects(ctx, &cloudv1.ListProjectsRequest{})
	if err != nil {
		log.Fatal("failed to fetch project list: \n", err)
	}
	for _, project := range projects.Projects {
		log.Println("Project ID: ", project.Id)
	}
}

```