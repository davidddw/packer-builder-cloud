package client

import (
	"errors"
	"fmt"
	"github.com/nilshell/xmlrpc"
	"log"
	"regexp"
)

type XenAPIClient struct {
	Session  interface{}
	Host     string
	Url      string
	Username string
	Password string
	RPC      *xmlrpc.Client
}

type APIResult struct {
	Status           string
	Value            interface{}
	ErrorDescription string
}

type XenAPIObject struct {
	Ref    string
	Client *XenAPIClient
}

type Host XenAPIObject
type VM XenAPIObject
type SR XenAPIObject
type VDI XenAPIObject
type Network XenAPIObject
type VBD XenAPIObject
type VIF XenAPIObject
type PIF XenAPIObject
type Pool XenAPIObject
type Task XenAPIObject

type VDIType int

const (
	_ VDIType = iota
	Disk
	CD
	Floppy
)

type TaskStatusType int

const (
	_ TaskStatusType = iota
	Pending
	Success
	Failure
	Cancelling
	Cancelled
)

func (c *XenAPIClient) RPCCall(result interface{}, method string, params []interface{}) (err error) {
	fmt.Println(params)
	p := new(xmlrpc.Params)
	p.Params = params
	err = c.RPC.Call(method, *p, result)
	return err
}

func (client *XenAPIClient) Login() (err error) {
	//Do loging call
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)
	params[0] = client.Username
	params[1] = client.Password

	err = client.RPCCall(&result, "session.login_with_password", params)
	if err == nil {
		// err might not be set properly, so check the reference
		if result["Value"] == nil {
			return errors.New ("Invalid credentials supplied")
		}
	}	
	client.Session = result["Value"]
	return err
}

func (client *XenAPIClient) APICall(result *APIResult, method string, params ...interface{}) (err error) {
	if client.Session == nil {
		fmt.Println("Error: no session")
		return fmt.Errorf("No session. Unable to make call")
	}

	//Make a params slice which will include the session
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session

	if params != nil {
		for idx, element := range params {
			p[idx+1] = element
		}
	}

	res := xmlrpc.Struct{}

	err = client.RPCCall(&res, method, p)

	if err != nil {
		return err
	}

	result.Status = res["Status"].(string)

	if result.Status != "Success" {
		fmt.Println("Encountered an API error: ", result.Status)
		fmt.Println(res["ErrorDescription"])
		return fmt.Errorf("API Error: %s", res["ErrorDescription"])
	} else {
		result.Value = res["Value"]
	}
	return
}

func (client *XenAPIClient) GetHosts() (err error) {
	result := APIResult{}
	_ = client.APICall(&result, "host.get_all")
	hosts := result.Value
	fmt.Println(hosts)
	return nil
}

func (client *XenAPIClient) GetPools() (pools []*Pool, err error) {
	pools = make([]*Pool, 0)
	result := APIResult{}
	err = client.APICall(&result, "pool.get_all")
	if err != nil {
		return pools, err
	}

	for _, elem := range result.Value.([]interface{}) {
		pool := new(Pool)
		pool.Ref = elem.(string)
		pool.Client = client
		pools = append(pools, pool)
	}

	return pools, nil
}

func (client *XenAPIClient) GetDefaultSR() (sr *SR, err error) {
	pools, err := client.GetPools()

	if err != nil {
		return nil, err
	}

	pool_rec, err := pools[0].GetRecord()

	if err != nil {
		return nil, err
	}

	if pool_rec["default_SR"] == "" {
		return nil, errors.New("No default_SR specified for the pool.")
	}

	sr = new(SR)
	sr.Ref = pool_rec["default_SR"].(string)
	sr.Client = client

	return sr, nil
}

func (client *XenAPIClient) GetVMByUuid(vm_uuid string) (vm *VM, err error) {
	vm = new(VM)
	result := APIResult{}
	err = client.APICall(&result, "VM.get_by_uuid", vm_uuid)
	if err != nil {
		return nil, err
	}
	vm.Ref = result.Value.(string)
	vm.Client = client
	return
}

func (client *XenAPIClient) GetVMByNameLabel(name_label string) (vms []*VM, err error) {
	vms = make([]*VM, 0)
	result := APIResult{}
	err = client.APICall(&result, "VM.get_by_name_label", name_label)
	if err != nil {
		return vms, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vm := new(VM)
		vm.Ref = elem.(string)
		vm.Client = client
		vms = append(vms, vm)
	}

	return vms, nil
}

func (client *XenAPIClient) GetSRByNameLabel(name_label string) (srs []*SR, err error) {
	srs = make([]*SR, 0)
	result := APIResult{}
	err = client.APICall(&result, "SR.get_by_name_label", name_label)
	if err != nil {
		return srs, err
	}

	for _, elem := range result.Value.([]interface{}) {
		sr := new(SR)
		sr.Ref = elem.(string)
		sr.Client = client
		srs = append(srs, sr)
	}

	return srs, nil
}

func (client *XenAPIClient) GetNetworkByUuid(network_uuid string) (network *Network, err error) {
	network = new(Network)
	result := APIResult{}
	err = client.APICall(&result, "network.get_by_uuid", network_uuid)
	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = client
	return
}

func (client *XenAPIClient) GetNetworkByNameLabel(name_label string) (networks []*Network, err error) {
	networks = make([]*Network, 0)
	result := APIResult{}
	err = client.APICall(&result, "network.get_by_name_label", name_label)
	if err != nil {
		return networks, err
	}

	for _, elem := range result.Value.([]interface{}) {
		network := new(Network)
		network.Ref = elem.(string)
		network.Client = client
		networks = append(networks, network)
	}

	return networks, nil
}

func (client *XenAPIClient) GetVdiByNameLabel(name_label string) (vdis []*VDI, err error) {
	vdis = make([]*VDI, 0)
	result := APIResult{}
	err = client.APICall(&result, "VDI.get_by_name_label", name_label)
	if err != nil {
		return vdis, err
	}

	for _, elem := range result.Value.([]interface{}) {
		vdi := new(VDI)
		vdi.Ref = elem.(string)
		vdi.Client = client
		vdis = append(vdis, vdi)
	}

	return vdis, nil
}

func (client *XenAPIClient) GetSRByUuid(sr_uuid string) (sr *SR, err error) {
	sr = new(SR)
	result := APIResult{}
	err = client.APICall(&result, "SR.get_by_uuid", sr_uuid)
	if err != nil {
		return nil, err
	}
	sr.Ref = result.Value.(string)
	sr.Client = client
	return
}

func (client *XenAPIClient) GetVdiByUuid(vdi_uuid string) (vdi *VDI, err error) {
	vdi = new(VDI)
	result := APIResult{}
	err = client.APICall(&result, "VDI.get_by_uuid", vdi_uuid)
	if err != nil {
		return nil, err
	}
	vdi.Ref = result.Value.(string)
	vdi.Client = client
	return
}

func (client *XenAPIClient) GetPIFs() (pifs []*PIF, err error) {
	pifs = make([]*PIF, 0)
	result := APIResult{}
	err = client.APICall(&result, "PIF.get_all")
	if err != nil {
		return pifs, err
	}
	for _, elem := range result.Value.([]interface{}) {
		pif := new(PIF)
		pif.Ref = elem.(string)
		pif.Client = client
		pifs = append(pifs, pif)
	}

	return pifs, nil
}

func (client *XenAPIClient) CreateTask() (task *Task, err error) {
	result := APIResult{}
	err = client.APICall(&result, "task.create", "packer-task", "Packer task")

	if err != nil {
		return
	}

	task = new(Task)
	task.Ref = result.Value.(string)
	task.Client = client
	return
}

func (client *XenAPIClient) CreateNetwork(name_label string, name_description string, bridge string) (network *Network, err error) {
	network = new(Network)

	net_rec := make(xmlrpc.Struct)
	net_rec["name_label"] = name_label
	net_rec["name_description"] = name_description 
	net_rec["bridge"] = bridge
	net_rec["other_config"] = make(xmlrpc.Struct)

	result := APIResult{}
	err = client.APICall(&result, "network.create", net_rec)
	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = client

	return network, nil
}

// Host associated function

func (self *Host) GetAddress() (address string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "host.get_address", self.Ref)
	if err != nil {
		return "", err
	}
	address = result.Value.(string)
	return address, nil
}


// VM associated functions

func (self *VM) Clone(label string) (new_instance *VM, err error) {
	new_instance = new(VM)

	result := APIResult{}
	err = self.Client.APICall(&result, "VM.clone", self.Ref, label)
	if err != nil {
		return nil, err
	}
	new_instance.Ref = result.Value.(string)
	new_instance.Client = self.Client
	return
}

func (self *VM) Provision() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.provision", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VM) Copy (newName string, targetSr *SR) (new_instance *VM, err error) {
	new_instance = new(VM)

	result := APIResult{}
	err = self.Client.APICall(&result, "VM.copy", self.Ref, newName, targetSr.Ref)
	if err != nil {
		return nil, err
	}
	new_instance.Ref = result.Value.(string)
	new_instance.Client = self.Client
	return
}

func (self *VM) Snapshot(label string) (snapshot *VM, err error) {
	snapshot = new(VM)

	result := APIResult{}
	err = self.Client.APICall(&result, "VM.snapshot", self.Ref, label)
	if err != nil {
		return nil, err
	}
	snapshot.Ref = result.Value.(string)
	snapshot.Client = self.Client
	return
}

func (self *VM) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.destroy", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VM) Start(paused, force bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.start", self.Ref, paused, force)
	if err != nil {
		return err
	}
	return
}

func (self *VM) CleanShutdown() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.clean_shutdown", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VM) HardShutdown() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.hard_shutdown", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VM) Unpause() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.unpause", self.Ref)
	if err != nil {
		return err
	}
	return
}

func (self *VM) Resume(paused, force bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.resume", self.Ref, paused, force)
	if err != nil {
		return err
	}
	return
}

func (self *VM) GetHVMBootPolicy() (bootOrder string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_HVM_boot_policy", self.Ref)
	if err != nil {
		return "", err
	}
	bootOrder = ""
	if result.Value != nil {
		bootOrder = result.Value.(string)
	}

	return bootOrder, nil
}


func (self *VM) SetHVMBoot(policy, bootOrder string) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.set_HVM_boot_policy", self.Ref, policy)
	if err != nil {
		return err
	}
	result = APIResult{}
	params := make(xmlrpc.Struct)
	params["order"] = bootOrder
	err = self.Client.APICall(&result, "VM.set_HVM_boot_params", self.Ref, params)
	if err != nil {
		return err
	}
	return
}

func (self *VM) SetPVBootloader(pv_bootloader, pv_args string) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.set_PV_bootloader", self.Ref, pv_bootloader)
	if err != nil {
		return err
	}
	result = APIResult{}
	err = self.Client.APICall(&result, "VM.set_PV_bootloader_args", self.Ref, pv_args)
	if err != nil {
		return err
	}
	return
}

func (self *VM) GetDomainId() (domid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_domid", self.Ref)
	if err != nil {
		return "", err
	}
	domid = result.Value.(string)
	return domid, nil
}

func (self *VM) GetResidentOn() (host *Host, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_resident_on", self.Ref)
	if err != nil {
		return nil, err
	}

	host = new(Host)
	host.Ref = result.Value.(string)
	host.Client = self.Client

	return host, nil
}


func (self *VM) GetPowerState() (state string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_power_state", self.Ref)
	if err != nil {
		return "", err
	}
	state = result.Value.(string)
	return state, nil
}

func (self *VM) GetUuid() (uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	uuid = result.Value.(string)
	return uuid, nil
}

func (self *VM) GetVBDs() (vbds []VBD, err error) {
	vbds = make([]VBD, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_VBDs", self.Ref)
	if err != nil {
		return vbds, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vbd := VBD{}
		vbd.Ref = elem.(string)
		vbd.Client = self.Client
		vbds = append(vbds, vbd)
	}

	return vbds, nil
}

func (self *VM) GetVIFs() (vifs []VIF, err error) {
	vifs = make([]VIF, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_VIFs", self.Ref)
	if err != nil {
		return vifs, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vif := VIF{}
		vif.Ref = elem.(string)
		vif.Client = self.Client
		vifs = append(vifs, vif)
	}

	return vifs, nil
}

func (self *VM) GetDisks() (vdis []*VDI, err error) {
	// Return just data disks (non-isos)
	vdis = make([]*VDI, 0)
	vbds, err := self.GetVBDs()
	if err != nil {
		return nil, err
	}

	for _, vbd := range vbds {
		rec, err := vbd.GetRecord()
		if err != nil {
			return nil, err
		}
		if rec["type"] == "Disk" {

			vdi, err := vbd.GetVDI()
			if err != nil {
				return nil, err
			}
			vdis = append(vdis, vdi)

		}
	}
	return vdis, nil
}

func (self *VM) GetGuestMetricsRef() (ref string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.get_guest_metrics", self.Ref)
	if err != nil {
		return "", nil
	}
	ref = result.Value.(string)
	return ref, err
}

func (self *VM) GetGuestMetrics() (metrics map[string]interface{}, err error) {
	metrics_ref, err := self.GetGuestMetricsRef()
	if err != nil {
		return nil, err
	}
	if metrics_ref == "OpaqueRef:NULL" {
		return nil, nil
	}

	result := APIResult{}
	err = self.Client.APICall(&result, "VM_guest_metrics.get_record", metrics_ref)
	if err != nil {
		return nil, err
	}
	return result.Value.(xmlrpc.Struct), nil
}

func (self *VM) SetStaticMemoryRange(min, max uint) (err error) {
	result := APIResult{}
	strMin := fmt.Sprintf("%d", min)
	strMax := fmt.Sprintf("%d", max)
	err = self.Client.APICall(&result, "VM.set_memory_limits", self.Ref, strMin, strMax, strMin, strMax)
	if err != nil {
		return err
	}
	return
}

func (self *VM) ConnectVdi(vdi *VDI, vdiType VDIType) (err error) {

	// 1. Create a VBD

	vbd_rec := make(xmlrpc.Struct)
	vbd_rec["VM"] = self.Ref
	vbd_rec["VDI"] = vdi.Ref
	vbd_rec["userdevice"] = "autodetect"
	vbd_rec["empty"] = false
	vbd_rec["other_config"] = make(xmlrpc.Struct)
	vbd_rec["qos_algorithm_type"] = ""
	vbd_rec["qos_algorithm_params"] = make(xmlrpc.Struct)

	switch vdiType {
	case CD:
		vbd_rec["mode"] = "RO"
		vbd_rec["bootable"] = true
		vbd_rec["unpluggable"] = false
		vbd_rec["type"] = "CD"
	case Disk:
		vbd_rec["mode"] = "RW"
		vbd_rec["bootable"] = false
		vbd_rec["unpluggable"] = false
		vbd_rec["type"] = "Disk"
	case Floppy:
		vbd_rec["mode"] = "RW"
		vbd_rec["bootable"] = false
		vbd_rec["unpluggable"] = true
		vbd_rec["type"] = "Floppy"
	}

	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.create", vbd_rec)

	if err != nil {
		return err
	}

	vbd_ref := result.Value.(string)
	fmt.Println("VBD Ref:", vbd_ref)

	result = APIResult{}
	err = self.Client.APICall(&result, "VBD.get_uuid", vbd_ref)

	fmt.Println("VBD UUID: ", result.Value.(string))
	/*
	   // 2. Plug VBD (Non need - the VM hasn't booted.
	   // @todo - check VM state
	   result = APIResult{}
	   err = self.Client.APICall(&result, "VBD.plug", vbd_ref)

	   if err != nil {
	       return err
	   }
	*/
	return
}

func (self *VM) DisconnectVdi(vdi *VDI) error {
	vbds, err := self.GetVBDs()
	if err != nil {
		return fmt.Errorf("Unable to get VM VBDs: %s", err.Error())
	}

	for _, vbd := range vbds {
		rec, err := vbd.GetRecord()
		if err != nil {
			return fmt.Errorf("Could not get record for VBD '%s': %s", vbd.Ref, err.Error())
		}

		if recVdi, ok := rec["VDI"].(string); ok {
			if recVdi == vdi.Ref {
				_ = vbd.Unplug()
				err = vbd.Destroy()
				if err != nil {
					return fmt.Errorf("Could not destroy VBD '%s': %s", vbd.Ref, err.Error())
				}

				return nil
			}
		} else {
			log.Printf("Could not find VDI record in VBD '%s'", vbd.Ref)
		}
	}

	return fmt.Errorf("Could not find VBD for VDI '%s'", vdi.Ref)
}

func (self *VM) SetPlatform(params map[string]string) (err error) {
	result := APIResult{}
	platform_rec := make(xmlrpc.Struct)
	for key, value := range params {
		platform_rec[key] = value
	}

	err = self.Client.APICall(&result, "VM.set_platform", self.Ref, platform_rec)

	if err != nil {
		return err
	}
	return
}

func (self *VM) SetVCpuMax( vcpus uint) (err error) {
	result := APIResult{}
	strVcpu := fmt.Sprintf("%d", vcpus)

	err = self.Client.APICall(&result, "VM.set_VCPUs_max", self.Ref, strVcpu)

	if err != nil {
		return err
	}
	return
}

func (self *VM) SetVCpuAtStartup( vcpus uint) (err error) {
	result := APIResult{}
	strVcpu := fmt.Sprintf("%d", vcpus)

	err = self.Client.APICall(&result, "VM.set_VCPUs_at_startup", self.Ref, strVcpu)

	if err != nil {
		return err
	}
	return
}

func (self *VM) ConnectNetwork(network *Network, device string) (vif *VIF, err error) {
	// Create the VIF

	vif_rec := make(xmlrpc.Struct)
	vif_rec["network"] = network.Ref
	vif_rec["VM"] = self.Ref
	vif_rec["MAC"] = ""
	vif_rec["device"] = device
	vif_rec["MTU"] = "1504"
	vif_rec["other_config"] = make(xmlrpc.Struct)
	vif_rec["qos_algorithm_type"] = ""
	vif_rec["qos_algorithm_params"] = make(xmlrpc.Struct)

	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.create", vif_rec)

	if err != nil {
		return nil, err
	}

	vif = new(VIF)
	vif.Ref = result.Value.(string)
	vif.Client = self.Client

	return vif, nil
}

//      Setters

func (self *VM) SetIsATemplate(is_a_template bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.set_is_a_template", self.Ref, is_a_template)
	if err != nil {
		return err
	}
	return
}

func (self *VM) SetHaAlwaysRun(ha_always_run bool) (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VM.set_ha_always_run", self.Ref, ha_always_run)
	if err != nil {
		return err
	}
	return
}

// SR associated functions

func (self *SR) CreateVdi(name_label string, size int64) (vdi *VDI, err error) {
	vdi = new(VDI)

	vdi_rec := make(xmlrpc.Struct)
	vdi_rec["name_label"] = name_label
	vdi_rec["SR"] = self.Ref
	vdi_rec["virtual_size"] = fmt.Sprintf("%d", size)
	vdi_rec["type"] = "user"
	vdi_rec["sharable"] = false
	vdi_rec["read_only"] = false

	oc := make(xmlrpc.Struct)
	oc["temp"] = "temp"
	vdi_rec["other_config"] = oc

	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.create", vdi_rec)
	if err != nil {
		return nil, err
	}

	vdi.Ref = result.Value.(string)
	vdi.Client = self.Client

	return
}

func (self *SR) GetUuid() (sr_uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "SR.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	sr_uuid = result.Value.(string)
	return sr_uuid, nil
}

// Network associated functions

func (self *Network) GetAssignedIPs() (ip_map map[string]string, err error) {
	ip_map = make(map[string]string, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "network.get_assigned_ips", self.Ref)
	if err != nil {
		return ip_map, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		ip_map[k] = v.(string)
	}
	return ip_map, nil
}

func (self *Network) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "network.destroy", self.Ref)
	if err != nil {
		return err
	}
	return
}


// PIF associated functions

func (self *PIF) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "PIF.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}

// Pool associated functions

func (self *Pool) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "pool.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}

// VBD associated functions
func (self *VBD) GetRecord() (record map[string]interface{}, err error) {
	record = make(map[string]interface{})
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.get_record", self.Ref)
	if err != nil {
		return record, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		record[k] = v
	}
	return record, nil
}

func (self *VBD) GetVDI() (vdi *VDI, err error) {
	vbd_rec, err := self.GetRecord()
	if err != nil {
		return nil, err
	}

	vdi = new(VDI)
	vdi.Ref = vbd_rec["VDI"].(string)
	vdi.Client = self.Client

	return vdi, nil
}

func (self *VBD) Eject() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.eject", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) Unplug() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.unplug", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VBD) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VBD.destroy", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

// VIF associated functions

func (self *VIF) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.destroy", self.Ref)
	if err != nil {
		return err
	}
	return nil
}

func (self *VIF) GetNetwork() (network *Network, err error) {

	network = new(Network)
	result := APIResult{}
	err = self.Client.APICall(&result, "VIF.get_network", self.Ref)

	if err != nil {
		return nil, err
	}
	network.Ref = result.Value.(string)
	network.Client = self.Client
	return

}

// VDI associated functions

func (self *VDI) GetUuid() (vdi_uuid string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_uuid", self.Ref)
	if err != nil {
		return "", err
	}
	vdi_uuid = result.Value.(string)
	return vdi_uuid, nil
}

func (self *VDI) GetVBDs() (vbds []VBD, err error) {
	vbds = make([]VBD, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_VBDs", self.Ref)
	if err != nil {
		return vbds, err
	}
	for _, elem := range result.Value.([]interface{}) {
		vbd := VBD{}
		vbd.Ref = elem.(string)
		vbd.Client = self.Client
		vbds = append(vbds, vbd)
	}

	return vbds, nil
}

func (self *VDI) GetVirtualSize() (virtual_size string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.get_virtual_size", self.Ref)
	if err != nil {
		return "", err
	}
	virtual_size = result.Value.(string)  
	return virtual_size, nil
}


func (self *VDI) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "VDI.destroy", self.Ref)
	if err != nil {
		return err
	}
	return
}

// Task associated functions

func (self *Task) GetStatus() (status TaskStatusType, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "task.get_status", self.Ref)
	if err != nil {
		return
	}
	rawStatus := result.Value.(string)
	switch rawStatus {
	case "pending":
		status = Pending
	case "success":
		status = Success
	case "failure":
		status = Failure
	case "cancelling":
		status = Cancelling
	case "cancelled":
		status = Cancelled
	default:
		panic(fmt.Sprintf("Task.get_status: Unknown status '%s'", rawStatus))
	}
	return
}

func (self *Task) GetProgress() (progress float64, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "task.get_progress", self.Ref)
	if err != nil {
		return
	}
	progress = result.Value.(float64)
	return
}

func (self *Task) GetResult() (object *XenAPIObject, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "task.get_result", self.Ref)
	if err != nil {
		return
	}
	switch ref := result.Value.(type) {
	case string:
		// @fixme: xapi currently sends us an xmlrpc-encoded string via xmlrpc.
		// This seems to be a bug in xapi. Remove this workaround when it's fixed
		re := regexp.MustCompile("^<value><array><data><value>([^<]*)</value>.*</data></array></value>$")
		match := re.FindStringSubmatch(ref)
		if match == nil {
			object = nil
		} else {
			object = &XenAPIObject{
				Ref:    match[1],
				Client: self.Client,
			}
		}
	case nil:
		object = nil
	default:
		err = fmt.Errorf("task.get_result: unknown value type %T (expected string or nil)", ref)
	}
	return
}

func (self *Task) GetErrorInfo() (errorInfo []string, err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "task.get_error_info", self.Ref)
	if err != nil {
		return
	}
	errorInfo = make([]string, 0)
	for _, infoRaw := range result.Value.([]interface{}) {
		errorInfo = append(errorInfo, fmt.Sprintf("%v", infoRaw))
	}
	return
}

func (self *Task) Destroy() (err error) {
	result := APIResult{}
	err = self.Client.APICall(&result, "task.destroy", self.Ref)
	return
}

// Client Initiator

func NewXenAPIClient(host, username, password string) (client XenAPIClient) {
	client.Host = host
	client.Url = "http://" + host
	client.Username = username
	client.Password = password
	client.RPC, _ = xmlrpc.NewClient(client.Url, nil)
	return
}
