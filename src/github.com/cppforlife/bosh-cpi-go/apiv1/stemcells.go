package apiv1

type Stemcells interface {
	CreateStemcell(string, StemcellCloudProps, ApiVersions) (StemcellCID, error)
	DeleteStemcell(StemcellCID, ApiVersions) error
}

type StemcellCloudProps interface {
	As(interface{}) error
}

type StemcellCID struct {
	cloudID
}

func NewStemcellCID(cid string) StemcellCID {
	if cid == "" {
		panic("Internal incosistency: Stemcell CID must not be empty")
	}
	return StemcellCID{cloudID{cid}}
}
