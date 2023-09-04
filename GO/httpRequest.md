
通过用户名密码获取ticket，访问proxmox API json
```go
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ticket struct {
	Data struct {
		Ticket              string `json:"ticket"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Cap                 struct {
			Sdn struct {
				SDNAllocate       int `json:"SDN.Allocate"`
				PermissionsModify int `json:"Permissions.Modify"`
				SDNAudit          int `json:"SDN.Audit"`
			} `json:"sdn"`
			Dc struct {
				SDNAllocate int `json:"SDN.Allocate"`
				SDNAudit    int `json:"SDN.Audit"`
				SysAudit    int `json:"Sys.Audit"`
			} `json:"dc"`
			Access struct {
				PermissionsModify int `json:"Permissions.Modify"`
				GroupAllocate     int `json:"Group.Allocate"`
				UserModify        int `json:"User.Modify"`
			} `json:"access"`
			Storage struct {
				PermissionsModify         int `json:"Permissions.Modify"`
				DatastoreAudit            int `json:"Datastore.Audit"`
				DatastoreAllocate         int `json:"Datastore.Allocate"`
				DatastoreAllocateTemplate int `json:"Datastore.AllocateTemplate"`
				DatastoreAllocateSpace    int `json:"Datastore.AllocateSpace"`
			} `json:"storage"`
			Vms struct {
				VMConfigNetwork    int `json:"VM.Config.Network"`
				VMConfigHWType     int `json:"VM.Config.HWType"`
				VMMigrate          int `json:"VM.Migrate"`
				VMAudit            int `json:"VM.Audit"`
				VMConfigCDROM      int `json:"VM.Config.CDROM"`
				VMConsole          int `json:"VM.Console"`
				VMBackup           int `json:"VM.Backup"`
				PermissionsModify  int `json:"Permissions.Modify"`
				VMPowerMgmt        int `json:"VM.PowerMgmt"`
				VMClone            int `json:"VM.Clone"`
				VMSnapshot         int `json:"VM.Snapshot"`
				VMSnapshotRollback int `json:"VM.Snapshot.Rollback"`
				VMConfigMemory     int `json:"VM.Config.Memory"`
				VMConfigCloudinit  int `json:"VM.Config.Cloudinit"`
				VMConfigOptions    int `json:"VM.Config.Options"`
				VMConfigCPU        int `json:"VM.Config.CPU"`
				VMAllocate         int `json:"VM.Allocate"`
				VMMonitor          int `json:"VM.Monitor"`
				VMConfigDisk       int `json:"VM.Config.Disk"`
			} `json:"vms"`
			Nodes struct {
				SysPowerMgmt      int `json:"Sys.PowerMgmt"`
				SysModify         int `json:"Sys.Modify"`
				PermissionsModify int `json:"Permissions.Modify"`
				SysAudit          int `json:"Sys.Audit"`
				SysSyslog         int `json:"Sys.Syslog"`
				SysConsole        int `json:"Sys.Console"`
			} `json:"nodes"`
		} `json:"cap"`
		Username string `json:"username"`
	} `json:"data"`
}

func GetTicket() (string, string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := http.Client{Transport: tr}
	payload := []byte(`{"username": "root@pam", "password": "Sec197011"}`)
	req, _ := http.NewRequest("POST", "https://192.168.65.115:8006/api2/json/access/ticket", bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")

	rep, _ := c.Do(req)
	if rep.StatusCode != 200 {
		fmt.Println("网络请求失败")
	}

	defer func() {
		rep.Body.Close()
	}()
	body, _ := io.ReadAll(rep.Body)
	var ticket Ticket
	err := json.Unmarshal(body, &ticket)
	if err != nil {
		fmt.Println("error decoding JSON", err)
	}
	return ticket.Data.CSRFPreventionToken, ticket.Data.Ticket
}

func getvmstatus() {
	token, ticket := GetTicket()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "https://192.168.65.115:8006/api2/json/nodes/wego/qemu/102/status/current", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("CSRFPreventionToken", token)
	req.Header.Add("Cookie", "PVEAuthCookie="+ticket+"")
	rep, _ := c.Do(req)
	fmt.Println(rep.StatusCode)
	if rep.StatusCode != 200 {
		fmt.Println("网络请求失败")
	}

	defer func() {
		rep.Body.Close()
	}()
	body, _ := io.ReadAll(rep.Body)
	fmt.Println(string(body))
}

func main() {
	getvmstatus()
}


```