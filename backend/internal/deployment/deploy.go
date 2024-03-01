package deployment

import (
	"fmt"
	"os"
	"os/exec"
	"v0/internal/data"
)

// DeployClient deploys a whitelabel client site using Docker
func DeployClient(clientData *data.Client) error {

	// Get the current environment variables
	env := os.Environ()

	// Add or modify environment variables as needed
	// For example, to set a custom environment variable:
	env = append(env, "VITE_CLIENT_ID="+clientData.ID.String())
	env = append(env, "VITE_API_URL=http://localhost:8080")

	// pass arguments to the Docker Compose command
	cmd := exec.Command("docker-compose", "-f", "../../docker-compose.yml", "-p", clientData.ID.String(), "up", "-d")
	cmd.Env = env

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return nil
}
