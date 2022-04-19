package gopowerstore

// SoftwareInstalled queries the software packages that are known by the cluster.
type SoftwareInstalled struct {
	// Unique identifier of the installed software instance.
	ID string `json:"id,omitempty"`
	// Whether this information represents the common software release version that is supported on all appliances in the cluster.
	IsCluster bool `json:"is_cluster,omitempty"`
	// Version of the installed release software package release.
	ReleaseVersion string `json:"release_version,omitempty"`
	// Build version of the installed software package release.
	BuildVersion string `json:"build_version,omitempty"`
	// Unique identifier of this build.
	BuildID string `json:"build_id,omitempty"`
}

// Fields returns fields which must be requested to fill struct
func (h *SoftwareInstalled) Fields() []string {
	return []string{"id", "is_cluster", "release_version",
		"build_version", "build_id"}
}
