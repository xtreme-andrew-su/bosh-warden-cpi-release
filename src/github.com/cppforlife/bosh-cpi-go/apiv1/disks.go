package apiv1

type Disks interface {
	CreateDisk(int, DiskCloudProps, *VMCID, ApiVersions) (DiskCID, error)
	DeleteDisk(DiskCID, ApiVersions) error

	AttachDisk(VMCID, DiskCID, ApiVersions) (interface{}, error)
	DetachDisk(VMCID, DiskCID, ApiVersions) error

	HasDisk(DiskCID, ApiVersions) (bool, error)
}

type DiskCloudProps interface {
	As(interface{}) error
	_final() // interface unimplementable from outside
}

type DiskCID struct {
	cloudID
}

func NewDiskCID(cid string) DiskCID {
	if cid == "" {
		panic("Internal incosistency: Disk CID must not be empty")
	}
	return DiskCID{cloudID{cid}}
}
