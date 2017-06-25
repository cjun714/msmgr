package k8

import (
	"util/log"

	yaml "gopkg.in/yaml.v2"
)

/** SVC sample
* kind: Service
* apiVersion: v1
* metadata:
*   name: eureka-server
*   labels:
*     name: eureka-server
* spec:
*   type: NodePort
*   ports:
*   - nodePort: 32112
*     targetPort: 8080
*     port: 8080
*   selector:
*     name: eureka-server
**/

// SVC is used to define service
type SVC struct {
	Kind       string
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name   string
		Labels struct {
			Name string
		}
	}
	Spec struct {
		Type  string
		Ports []struct {
			NodePort   int `yaml:"nodePort"`
			TargetPort int `yaml:"targetPort"`
			Port       int
		}
		Selector struct {
			Name string
		}
	}
}

// func (*SVC) ToFile(path string) {

// }

// NewSVC new SVC instance from yaml byte[]
func NewSVC(yamlBytes []byte) *SVC {
	svc := new(SVC)

	e := yaml.Unmarshal(yamlBytes, svc)
	if e != nil {
		log.E(e)
		return nil
	}
	return svc
}
