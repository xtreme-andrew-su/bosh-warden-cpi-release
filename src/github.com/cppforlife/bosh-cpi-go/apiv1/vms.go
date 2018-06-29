package apiv1

type VMs interface {
	CreateVM(AgentID, StemcellCID, VMCloudProps, Networks, []DiskCID, VMEnv, ApiVersions) (interface{}, error)
	DeleteVM(VMCID, ApiVersions) error

	CalculateVMCloudProperties(VMResources, ApiVersions) (VMCloudProps, error)

	SetVMMetadata(VMCID, VMMeta, ApiVersions) error
	HasVM(VMCID, ApiVersions) (bool, error)

	RebootVM(VMCID, ApiVersions) error
	GetDisks(VMCID, ApiVersions) ([]DiskCID, error)
}

type VMCloudProps interface {
	As(interface{}) error
	_final() // interface unimplementable from outside
}

type VMResources struct {
	RAM               int `json:"ram"`
	CPU               int `json:"cpu"`
	EphemeralDiskSize int `json:"ephemeral_disk_size"`
}

type VMCID struct {
	cloudID
}

type AgentID struct {
	cloudID
}

type VMMeta struct {
	cloudKVs
}

type VMEnv struct {
	cloudKVs
}

func NewVMCID(cid string) VMCID {
	if cid == "" {
		panic("Internal incosistency: VM CID must not be empty")
	}
	return VMCID{cloudID{cid}}
}

func NewAgentID(id string) AgentID {
	if id == "" {
		panic("Internal incosistency: Agent ID must not be empty")
	}
	return AgentID{cloudID{id}}
}

func NewVMMeta(meta map[string]interface{}) VMMeta {
	return VMMeta{cloudKVs{meta}}
}

func NewVMEnv(env map[string]interface{}) VMEnv {
	return VMEnv{cloudKVs{env}}
}
