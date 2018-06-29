package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"

	bwcstem "github.com/cppforlife/bosh-warden-cpi/stemcell"
	bwcvm "github.com/cppforlife/bosh-warden-cpi/vm"
)

type CreateVMMethod struct {
	stemcellFinder bwcstem.Finder
	vmCreator      bwcvm.Creator
}

func NewCreateVMMethod(stemcellFinder bwcstem.Finder, vmCreator bwcvm.Creator) CreateVMMethod {
	return CreateVMMethod{
		stemcellFinder: stemcellFinder,
		vmCreator:      vmCreator,
	}
}

func (a CreateVMMethod) CreateVM(
	agentID apiv1.AgentID, stemcellCID apiv1.StemcellCID,
	cloudProps apiv1.VMCloudProps, networks apiv1.Networks,
	associatedDiskCIDs []apiv1.DiskCID, env apiv1.VMEnv, _ apiv1.ApiVersions) (interface{}, error) {

	stemcell, found, err := a.stemcellFinder.Find(stemcellCID)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Finding stemcell '%s'", stemcellCID)
	}

	if !found {
		return apiv1.VMCID{}, bosherr.Errorf("Expected to find stemcell '%s'", stemcellCID)
	}

	var customCloudProps VMCloudProperties

	err = cloudProps.As(&customCloudProps)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Parsing VM cloud properties")
	}

	vmProps, err := customCloudProps.AsVMProps()
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Validating 'ports' configuration")
	}

	vm, err := a.vmCreator.Create(agentID, stemcell, vmProps, networks, env)
	if err != nil {
		return apiv1.VMCID{}, bosherr.WrapErrorf(err, "Creating VM with agent ID '%s'", agentID)
	}

	return vm.ID(), nil
}
