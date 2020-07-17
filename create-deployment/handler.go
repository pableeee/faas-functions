package function

import (
	"encoding/json"
	"fmt"

	"github.com/pableeee/k8s-simple-wrapper/cmd"
)

const kubeconfig = "/var/openfaas/secrets/config"

type request struct {
	Image  string   `json:"image"`
	Name   string   `json:"name"`
	Params []string `json:"params"`
}

// Handle a serverless request
func Handle(req []byte) string {
	var r request
	var s string

	err := json.Unmarshal(req, &r)

	if err != nil {
		return fmt.Sprintf("Can't unmarshal request: %s", string(req))
	}

	fmt.Printf("Image:%s , Name:%s", r.Image, r.Name)

	deployment := cmd.DeploymentManagerImpl{}
	s, err = deployment.CreateDeployment(kubeconfig, "default", r.Image, r.Name)

	if err != nil {
		return fmt.Sprintf("Error creating deployment: %s", string(req))
	}

	return s
}
