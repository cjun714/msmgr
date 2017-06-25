package k8

import (
	"util/log"

	yaml "gopkg.in/yaml.v2"
)

/**
kind: ReplicationController
apiVersion: v1
metadata:
  name: hello
  labels:
    name: hello
spec:
  replicas: 1
  selector:
    name: hello
  template:
    metadata:
      labels:
        name: hello
    spec:
      containers:
      - name: hello
        image: 20.26.33.121:5000/hello
        resources:
          limits:
            cpu: 600m
            memory: 600Mi
          requests:
            cpu: 500m
            memory: 500Mi
        env:
        - name: SPRING_CONFIG_URI
          value: 'http://ms:Microservice123!@20.26.33.122:32011'
        - name: CONFIG_ACTIVE
          value: "test"
        command: ["sh","-c","java -server -Xms256m -Xmx256m  -Djava.io.tmpdir=/var/tmp -Duser.timezone=Asia/Shanghai -jar hello.jar --spring.profiles.active=${CONFIG_ACTIVE} --server.port=8080"]
        ports:
        - containerPort: 8080
**/

// RC is used to define replication
type RC struct {
	Kind       string
	APIVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name   string
		Labels struct {
			Name string
		}
	}
	Spec struct {
		Replicas int
		Selector struct {
			Name string
		}
		Template struct {
			Metadata struct {
				Labels struct {
					Name string
				}
			}
			Spec struct {
				Containers []struct {
					Name      string
					Image     string
					Resources struct {
						Limits struct {
							CPU    string
							Memory string
						}
						Requests struct {
							CPU    string
							Memory string
						}
					}
					Env []struct {
						Name  string
						Value string
					}
					Command []string `yaml:",flow"`
					Ports   []struct {
						ContainerPort int `yaml:"containerPort"`
					}
				}
			}
		}
	}
}

// NewRC new SVC instance from yaml byte[]
func NewRC(yamlBytes []byte) *RC {
	rc := new(RC)

	e := yaml.Unmarshal(yamlBytes, rc)
	if e != nil {
		log.E(e)
		return nil
	}
	return rc
}
