package apiv1

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type ActionFactory struct {
	cpiFactory CPIFactory
}

type Action interface{}

func NewActionFactory(cpiFactory CPIFactory) ActionFactory {
	return ActionFactory{cpiFactory}
}

func (f ActionFactory) Create(method string, context CallContext, apiVersions ApiVersions) (interface{}, error) {
	cpi, err := f.cpiFactory.New(context)
	if err != nil {
		return nil, err
	}

	// binds concrete values to interfaces

	switch method {
	case "info":
		return cpi.Info, nil

	case "create_stemcell":
		return func(imagePath string, props CloudPropsImpl) (StemcellCID, error) {
			return cpi.CreateStemcell(imagePath, props, apiVersions)
		}, nil

	case "delete_stemcell":
		return func(cid StemcellCID) (interface{}, error) {
			return nil, cpi.DeleteStemcell(cid, apiVersions)
		}, nil

	case "create_vm":
		return func(
			agentID AgentID, stemcellCID StemcellCID, props CloudPropsImpl,
			networks Networks, diskCIDs []DiskCID, env VMEnv) (interface{}, error) {

			return cpi.CreateVM(agentID, stemcellCID, props, networks, diskCIDs, env, apiVersions)
		}, nil

	case "delete_vm":
		return func(cid VMCID) (interface{}, error) {
			return nil, cpi.DeleteVM(cid, apiVersions)
		}, nil

	case "calculate_vm_cloud_properties":
		return func(vmResource VMResources) (interface{}, error) {
			return cpi.CalculateVMCloudProperties(vmResource, apiVersions)
		}, nil

	case "set_vm_metadata":
		return func(cid VMCID, metadata VMMeta) (interface{}, error) {
			return nil, cpi.SetVMMetadata(cid, metadata, apiVersions)
		}, nil

	case "has_vm":
		return func(cid VMCID) (interface{}, error) {
			return cpi.HasVM(cid, apiVersions)
		}, nil

	case "reboot_vm":
		return func(cid VMCID) (string, error) {
			return "", cpi.RebootVM(cid, apiVersions)
		}, nil

	case "get_disks":
		return func(cid VMCID) ([]DiskCID, error) {
			diskCIDs, err := cpi.GetDisks(cid, apiVersions)
			if len(diskCIDs) == 0 {
				return []DiskCID{}, err
			}
			return diskCIDs, err
		}, nil

	case "create_disk":
		return func(size int, props CloudPropsImpl, vmCID *VMCID) (DiskCID, error) {
			return cpi.CreateDisk(size, props, vmCID, apiVersions)
		}, nil

	case "delete_disk":
		return func(cid DiskCID) (interface{}, error) {
			return nil, cpi.DeleteDisk(cid, apiVersions)
		}, nil

	case "attach_disk":
		return func(vmCID VMCID, diskCID DiskCID) (interface{}, error) {
			return cpi.AttachDisk(vmCID, diskCID, apiVersions)
		}, nil

	case "detach_disk":
		return func(vmCID VMCID, diskCID DiskCID) (interface{}, error) {
			return nil, cpi.DetachDisk(vmCID, diskCID, apiVersions)
		}, nil

	case "has_disk":
		return func(cid DiskCID) (interface{}, error) {
			return cpi.HasDisk(cid, apiVersions)
		}, nil

	default:
		return nil, bosherr.Errorf("Unknown method '%s'", method)
	}
}
