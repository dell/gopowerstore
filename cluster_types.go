package gopowerstore

// RemoteSystem details about a remote system
type RemoteSystem struct {
	// Unique identifier of the remote system instance.
	ID string `json:"id,omitempty"`
	// User-specified name of the remote system instance.
	// This property supports case-insensitive filtering
	Name string `json:"name,omitempty"`
	// User-specified description of the remote system instance.
	Description string `json:"description,omitempty"`
	// Serial number of the remote system instance
	SerialNumber string `json:"serial_number,omitempty"`
	// Management IP address of the remote system instance
	ManagementAddress string `json:"management_address,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (r *RemoteSystem) Fields() []string {
	return []string{"id", "name", "description", "serial_number", "management_address"}
}

// Cluster details about the cluster
type Cluster struct {
	// Unique identifier of the cluster.
	ID string `json:"id,omitempty"`
	// User-specified name of the cluster
	Name string `json:"name,omitempty"`
	// Management IP address of the remote system instance
	ManagementAddress string `json:"management_address,omitempty"`
	// Current state of the cluster
	State string `json:"state,omitempty"`
	// NVMe Subsystem NQN for cluster
	NVMeNQN string `json:"nvm_subsystem_nqn,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (r *Cluster) Fields() []string {
	return []string{"id", "name", "management_address", "state"}
}
