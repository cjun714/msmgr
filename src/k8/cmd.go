package k8

import (
	"util/exec"
	"util/file"
	"util/log"
)

// Create is used to create pod using cmd: kubectl create -f XXX.yaml
func Create(yamlPath string) bool {
	log.I("Create POD by: " + yamlPath)
	if e := exec.RunCmd("kubectl", "create", "-f", yamlPath); e != nil {
		return false
	}
	return true
}

// Stop is used to stop pod using cmd: kubectl delete -f XXX.yaml
func Stop(yamlPath string) bool {
	log.I("Delete POD by: " + yamlPath)
	if e := exec.RunCmd("kubectl", "delete", "-f", yamlPath); e != nil {
		return false
	}
	return true
}

// Delete is used to stop pod and delete .ymal
func Delete(yamlPath string) bool {
	Stop(yamlPath)
	file.Delete(yamlPath)
	return true
}

// GetPod is kubectl create -f
func GetPod(pod string) bool {
	log.I("Get POD : " + pod)
	if e := exec.RunCmd("kubectl", "get", "pod", pod); e != nil {
		return false
	}
	return true
}
