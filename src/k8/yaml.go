package k8

import (
	"config"
	"util/file"
	"util/yaml"
)

// CreateYaml is used to create k8 .yaml files including: rc.yaml,svc.yaml
func CreateYaml(svcInfo ServiceInfo) error {
	dirPath := config.YamlPath + svcInfo.Name + "/"
	rcPath := dirPath + "rc.yaml"
	svcPath := dirPath + "svc.yaml"

	e := file.MkDir(dirPath)
	if e != nil {
		return e
	}
	rc := NewRC([]byte(RCTmpl))
	svc := NewSVC([]byte(SVCTmpl))

	// TODO update rc/svc according svcInfo
	rcBytes, e := yaml.ToYaml(rc)
	svcBytes, e := yaml.ToYaml(svc)

	file.WriteFile(rcPath, rcBytes)
	file.WriteFile(svcPath, svcBytes)
	return nil
}
