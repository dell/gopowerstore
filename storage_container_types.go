package gopowerstore

type StorageContainerStorageProtocolEnum string

const (
	StorageContainerStorageProtocolEnumSCSI StorageContainerStorageProtocolEnum = "SCSI"
	StorageContainerStorageProtocolEnumNVME StorageContainerStorageProtocolEnum = "NVME"
)

type StorageContainer struct {
	ID              string                              `json:"id,omitempty"`
	Name            string                              `json:"name,omitempty"`
	Quota           int64                               `json:"quota,omitempty"`
	StorageProtocol StorageContainerStorageProtocolEnum `json:"storage_protocol,omitempty"`
	HighWaterMark   int16                               `json:"high_water_mark,omitempty"`
}

func (s StorageContainer) Fields() []string {
	return []string{"id", "name", "quota", "storage_protocol"}
}
