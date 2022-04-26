package gopowerstore

// K8sCluster defines K8s cluster information
type K8sCluster struct {
	Name      string `json:"name"`
	IPAddress string `json:"address"`
	Port      int    `json:"port"`
	Token     string `json:"token"`
}
