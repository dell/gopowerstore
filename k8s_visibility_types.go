package gopowerstore

type K8sCluster struct {
	Name      string `json:"name"`
	IPAddress string `json:"address"`
	Port      string `json:"port"`
	Token     string `json:"token"`
}
