package main

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"

)
var uauth Uauthcfg
type Uauthcfg struct {
	cfg ucloud.Config
	auth auth.Credential
}
type IPSet struct {

	Ip string
	Mac string
}
type Host struct {
	UHostId string
	Zone string
	PrivateIPSet []IPSet
	InternetIPSet []IPSet
	VpcID string
	SubNetId string

}
type UniInfo struct {
	UniId string
	Ip string
	Mac string
}
func (host *Host) addUni (uni UniInfo){

}
type ClientInfo struct {
	Region string
	PublicKey string
	PrivateKey string
	Projectid string
}

func GetHostID(eip string) string  {
	var  hostid string
	unetClient:=unet.NewClient(&uauth.cfg,&uauth.auth)
	req:=unetClient.NewDescribeEIPRequest()
	respon,err:=unetClient.DescribeEIP(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return err.Error()
	} else {
		for _, k := range respon.EIPSet {
			for _, ip := range k.EIPAddr {
				if ip.IP == eip {
					hostid= k.Resource.ResourceID
				}
			}
		}
	}
	return hostid
}
func GetHostDescribe(hostid string) (Host,string) {
	var hostdescribe Host
	var internetipset []IPSet
	var privateipset []IPSet
	hostdescribe=Host{}
	uhostclient:=uhost.NewClient(&uauth.cfg,&uauth.auth)
	req:=uhostclient.NewDescribeUHostInstanceRequest()
	req.UHostIds=[]string{hostid}
	respon,err:=uhostclient.DescribeUHostInstance(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return hostdescribe,err.Error()
	} else {
		for _,host:=range respon.UHostSet{
			if host.UHostId==hostid{
				hostdescribe.UHostId=hostid
				hostdescribe.Zone=host.Zone
				for _,v :=range host.IPSet{
					if v.Type =="Private" {
						tmp:=IPSet{
							Ip:  v.IP,
							Mac: v.Mac,
						}
						if v.Default=="true"{
							hostdescribe.SubNetId=v.SubnetId
							hostdescribe.VpcID=v.VPCId
						}

						privateipset=append(privateipset, tmp)
						hostdescribe.PrivateIPSet=privateipset

					}
					if v.Type=="Internation" || v.Type=="BGP" {
						tmp:=IPSet{
							Ip:  v.IP,
							Mac: v.Mac,
						}
						internetipset=append(internetipset, tmp)
						hostdescribe.InternetIPSet=internetipset
					}
				}
			}
		}
	}

	return hostdescribe,""
}

func CreateUNI(uhost Host) string{
	Uvpcclient:=NewInternalVPCClient(&uauth.cfg,&uauth.auth)
	req:=Uvpcclient.NewCreateNetworkInterfaceRequest()
	req.VPCId=ucloud.String(uhost.VpcID)
	req.SubnetId=ucloud.String(uhost.SubNetId)
	req.Remark=ucloud.String("test-jimmy")
	Uvpcclient.config.BaseUrl="http://api.ucloud.cn/"
	response,err :=Uvpcclient.CreateNetworkInterface(req)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
		return err.Error()
	}else {
		fmt.Printf(response.NetworkInterface.Remark)
	}
return ""
}
func main() {
	var host Host
	Uclient :=ClientInfo{
		"cn-bj2",
		"publickey",
		"privatekey",
		"org-im0svb",
	}
	cfg := ucloud.NewConfig()
	cfg.Region = Uclient.Region
	cfg.ProjectId=Uclient.Projectid
	credential := auth.NewCredential()
	credential.PublicKey = Uclient.PublicKey
	credential.PrivateKey = Uclient.PrivateKey
	uauth=Uauthcfg{
		cfg:  cfg,
		auth: credential,
	}
	hostid:=GetHostID("106.75.9.179")
	host.UHostId=hostid
	fmt.Printf("%s\n",host.UHostId)

	HostDescribe,err:=GetHostDescribe(host.UHostId)
	if err !="" {
		fmt.Println(err)
	}
	fmt.Println(HostDescribe.UHostId,HostDescribe.InternetIPSet,HostDescribe.VpcID)
	CreateUNI(HostDescribe)
}