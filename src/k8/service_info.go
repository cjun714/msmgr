package k8

// ServiceInfo is used to store resource defination assigned to service
type ServiceInfo struct {
	Name        string
	ImgURL      string
	RepliCount  int
	CPUQuota    string
	RAMQuota    string
	ClusterPort int
	TargetPort  int
	NodePort    int
	Desc        string
}
