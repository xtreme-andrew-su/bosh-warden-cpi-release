package action

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type CalculateVMCloudPropertiesMethod struct{}

func NewCalculateVMCloudPropertiesMethod() CalculateVMCloudPropertiesMethod {
	return CalculateVMCloudPropertiesMethod{}
}

func (a CalculateVMCloudPropertiesMethod) CalculateVMCloudProperties(res apiv1.VMResources, _ apiv1.ApiVersions) (apiv1.VMCloudProps, error) {
	return apiv1.NewVMCloudPropsFromMap(map[string]interface{}{}), nil
}
