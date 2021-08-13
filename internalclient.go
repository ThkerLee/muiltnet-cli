package main

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)
type UCloudInternalClient struct {
	*Client
}
func NewInternalHostClient(config *ucloud.Config, credential *auth.Credential) *UCloudInternalClient {
	meta := ClientMeta{Product: "uhost"}
	client := NewClientWithMeta(config, credential, meta)
	return &UCloudInternalClient{
		client,
	}
}
func NewInternalVPCClient(config *ucloud.Config, credential *auth.Credential) *UCloudInternalClient {
	meta := ClientMeta{Product: "VPC"}
	client := NewClientWithMeta(config, credential, meta)
	return &UCloudInternalClient{
		client,
	}
}
type CreateNetworkInterfaceRequest struct {
	request.CommonBase
	VPCId 		*string `required:"true"`
	SubnetId    *string `required:"true"`
	Name		*string `required:"false"`
	PrivateIp  	*string `required:"false"`
	SecurityGroupId *string `required:"false"`
	Tag 		*string `required:"false"`
	Remark 		*string `required:"false"`

}
func (c *UCloudInternalClient) NewCreateNetworkInterfaceRequest() *CreateNetworkInterfaceRequest {
	req := &CreateNetworkInterfaceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}
type NetworkInterface struct {
	InterfaceId      string
	VPCId            string
	SubnetId         string
	PrivateIpSet     []string
	MacAddress       string
	Status           int
	Name             string
	Netmask          string
	Gateway          string
	AttachInstanceId string
	Default          bool
	CreateTime       int
	Remark           string
	Tag              string
	EIPIdSet         []string
	FirewallIdSet    []string
}

type CreateNetworkInterfaceResponse struct {
	response.CommonBase
	NetworkInterface NetworkInterface
}

func (c *UCloudInternalClient) CreateNetworkInterface(req *CreateNetworkInterfaceRequest) (*CreateNetworkInterfaceResponse, error) {
	var err error
	var res CreateNetworkInterfaceResponse
	reqCopier := *req
	err = c.Client.InvokeActionGet("CreateNetworkInterface", &reqCopier, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}



type IAllocateUdiskRequest struct {
	request.CommonBase
	DiskSize              *int    `required:"true"`
	Region                *string `required:"true"`
	Zone                  *string `required:"true"`
	DiskName              *string `required:"true"`
	Backend               *string `required:"true"`
	channel               *string    `required:"true"`
	top_organization_id   *string    `required:"true"`
	organization_id       *string    `required:"true"`
	OSType                *string `required:"true"`
	OSName                *string `required:"true"`
	DiskType              *string `required:"false"`
	RdmaClusterId         *string `required:"false"`
	ChargeType            *string `required:"false"`
	Quantity              *int    `required:"false"`
	NetCapability         *string `required:"false"`
	HotplugFeature        *bool   `required:"false"`
	CloudInitFeature      *bool   `required:"false"`
	MinimalCpuPlatform    *string `required:"false"`
	RSSDAttachableFeature *bool   `required:"false"`
	NeedExpandPartition   *bool   `required:"false"`
	GPUFeature            *bool   `required:"false"`
	IPv6Feature           *bool   `required:"false"`
}
type IAllocateUdiskResponse struct {
	response.CommonBase
	RetCode int
	Action  string
	UDiskID string
	Message string
}
func (c *UCloudInternalClient) NewInternalIAllocateUdiskRequest() *IAllocateUdiskRequest {
	req := &IAllocateUdiskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}
func (c *UCloudInternalClient) CreateInternalIAllocateUdisk(req *IAllocateUdiskRequest) (*IAllocateUdiskResponse, error) {
	var err error
	var res IAllocateUdiskResponse

	reqCopier := *req
	// reqCopier.Password = ToBase64Query(reqCopier.Password)
	err = c.Client.InvokeAction("IAllocateUdisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
type GetUDiskDetailInfoRequest struct {
	request.CommonBase
	top_organization_id   *string    `required:"true"`
	organization_id       *string    `required:"true"`
	Backend  *string    `required:"true"`
	az_group  *string    `required:"true"`
	zone_id *string    `required:"true"`
	UDiskId *string    `required:"true"`
}
type UDiskDetailInfo struct {
	Status  *string    `required:"true"`
}
type GetUDiskDetailInfoResponse struct {
	response.CommonBase
	RetCode int
	Action  string
	DetailInfo UDiskDetailInfo

}
func (c *UCloudInternalClient) NewInternalGetUDiskDetailInfoRequest() *GetUDiskDetailInfoRequest {
	req := &GetUDiskDetailInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}
func (c *UCloudInternalClient) CreateInternalGetUDiskDetailInfo(req *IAllocateUdiskRequest) (*GetUDiskDetailInfoResponse, error) {
	var err error
	var res GetUDiskDetailInfoResponse

	reqCopier := *req
	// reqCopier.Password = ToBase64Query(reqCopier.Password)
	err = c.Client.InvokeAction("IAllocateUdisk", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
type IGetResourceinfoRequest struct {
	request.CommonBase
	Backend               *string `required:"true"`
	ResourceId              *string    `required:"true"`
	ResourceType              *string    `required:"true"`
	Limit              *string    `required:"true"`
	ZoneId 				*string    `required:"true"`
	Status 				*string    `required:"true"`

}
type ResourceInfo struct {
	Id              *string    `required:"false"`
	ResourceId              *string    `required:"false"`
	RegionId              *int    `required:"false"`
	ResourceType              *int    `required:"false"`
	ZoneId              *int    `required:"false"`
	TopOrganizationId 		*int    `required:"false"`
	OrganizationId 		*int    `required:"false"`
	Updated 		*int    `required:"false"`
	Created 		*int    `required:"false"`
	Status	 		*int    `required:"false"`
	VPCId              *string    `required:"false"`
	SubnetId              *string    `required:"false"`
	BusinessId              *string    `required:"false"`
}
type IGetResourceInfoResponse struct {
	response.CommonBase
	RetCode int
	Action  string
	Infos	[]ResourceInfo
}
func (c *UCloudInternalClient) NewInternaIGetResourceInfoRequest() *IGetResourceinfoRequest {
	req := &IGetResourceinfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}
func (c *UCloudInternalClient) CreateInternaIGetResourceInfo(req *IGetResourceinfoRequest) (*IGetResourceInfoResponse, error) {
	var err error
	var res IGetResourceInfoResponse

	reqCopier := *req
	// reqCopier.Password = ToBase64Query(reqCopier.Password)
	err = c.Client.InvokeAction("IGetResourceList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
type ZoneMap struct {
	ZoneId int
	AzGroup int
	RegionName string
	ZoneName string
	ZoneCName string
}