package tencentcam

import (
	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/util"
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func DetachPolicyFromUser() {
	UserUin := GetUserUin()
	request := cam.NewDetachUserPolicyRequest()
	request.PolicyId = common.Uint64Ptr(1)
	request.DetachUin = common.Uint64Ptr(UserUin)
	_, err := CAMClient().DetachUserPolicy(request)
	util.HandleErrNoExit(err)
	if err == nil {
		log.Debugln("成功移除 crossfire 用户的权限 (Successfully removed the privileges of the crossfire user)")
	}
}

func DeleteUser() {
	request := cam.NewDeleteUserRequest()
	request.Name = common.StringPtr("crossfire")
	request.Force = common.Uint64Ptr(1)
	_, err := CAMClient().DeleteUser(request)
	util.HandleErrNoExit(err)
	if err == nil {
		log.Debugln("删除 crossfire 用户成功 (Delete crossfire user successfully)")
	}
}

func CancelTakeoverConsole() {
	DetachPolicyFromUser()
	DeleteUser()
	log.Infoln("成功删除 crossfire 用户，已取消控制台接管 (Successful deletion of crossfire user, console takeover cancelled)")
}
