package gopowerstore

// LimitIDEnum - ID of limits returned by the /limit endpoint
type LimitIDEnum string

const (
	// MaxVolumeSize - Max size of a volume
	MaxVolumeSize LimitIDEnum = "Max_Volume_Size"
	// Max_VirtualVolume_Size - Max size of a virtual volume
	MaxVirtualVolumeSize LimitIDEnum = "Max_VirtualVolume_Size"
	// Max_Folder_Size - Max size of a folder
	MaxFolderSize LimitIDEnum = "Max_Folder_Size"
)

// Limit - Response /limit endpoint
type Limit struct {
	ID    string `json:"id"`
	Limit int64  `json:"limit"`
}

// Fields - Returns fields which must be requested to fill struct
func (l *Limit) Fields() []string {
	return []string{"id", "limit"}
}
