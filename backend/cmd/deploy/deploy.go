package deploy

import (
	"fmt"
	"os"

	"v0/internal/data"
	"v0/internal/deployment"
)

func DeployClient(client *data.Client) error {
	// Deploy the client site
	if err := deployment.DeployClient(client); err != nil {
		fmt.Fprintf(os.Stderr, "Error deploying client site: %v\n", err)
	}

	fmt.Println("Client site deployed successfully!")

	return nil
}
