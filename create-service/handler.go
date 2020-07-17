package function

import (
	"encoding/json"
	"fmt"

	"github.com/pableeee/k8s-simple-wrapper/cmd"
)

const kubeconfig = "/var/openfaas/secrets/config"

type request struct {
	Name string `json:"name"`
	Port unit16 `json:"port"`
}

// Handle a serverless request
func Handle(req []byte) string {
	var r request
	var s string

	err := json.Unmarshal(req, &r)

	if err != nil {
		return fmt.Sprintf("Can't unmarshal request: %s", string(req))
	}

	fmt.Printf("Name:%s , Port:%d", r.Name, r.Port)

	service := cmd.ServiceManagerImpl{}
	s, err = service.CreateService(kubeconfig, "default", r.Name, r.Port)

	if err != nil {
		return fmt.Sprintf("Error creating service: %s", string(req))
	}

	return s
}
