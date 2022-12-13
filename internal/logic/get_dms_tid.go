package logic

import (
	"fmt"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
)

// GetDmsTid 获取DMS上的租户ID
func (l *ALiClientLogic) GetDmsTid() (tid *int64) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	client, _ := clientS.GetAliDMSToken()
	getUserActiveTenantRequest := &dms_enterprise20181101.GetUserActiveTenantRequest{}
	runtime := &util.RuntimeOptions{}
	getTid, _err := client.GetUserActiveTenantWithOptions(getUserActiveTenantRequest, runtime)
	if _err != nil {
		fmt.Println("get tid failed!")
	}
	tid = getTid.Body.Tenant.Tid
	return tid
}
