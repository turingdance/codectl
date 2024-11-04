
//don't modify !!!!
// create at 2024-11-04 16:12:02
// creeate by winlion
//go:generate  codectl router -a winlion -s . -d . -n router.go
package codectl

import (
	"github.com/turingdance/infra/restkit"
)

var DefaultRouter *restkit.Router = restkit.NewRouter().PathPrefix("/")
// 初始化路由
func InitRouter(router *restkit.Router) {
	
	// 创建默认用户
	userinfoCtrl := &Userinfo{}
	userinforouter := router.Subrouter().PathPrefix("userinfo")
	//搜索
	userinforouter.HandleFunc("/search", userinfoCtrl.Search).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/count", userinfoCtrl.Count).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/updatedeptandname", userinfoCtrl.Updatedeptandname).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/updateUserinfo", userinfoCtrl.UpdateUserinfo).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/defaultRefId", userinfoCtrl.DefaultRefId).Methods("post","get",)
	
	//创建默认用户
	userinforouter.HandleFunc("/create", userinfoCtrl.Create).Methods("post","put",)
	
	//更新
	userinforouter.HandleFunc("/update", userinfoCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	userinforouter.HandleFunc("/delete", userinfoCtrl.Delete).Methods("post","delete",)
	
	//获取
	userinforouter.HandleFunc("/getOne", userinfoCtrl.GetOne).Methods("post","get",)
	
	//获取
	userinforouter.HandleFunc("/snsinfo", userinfoCtrl.Snsinfo).Methods("post","get",)
	

	
	
	// 获取岗位
	flowCtrl := &Flow{}
	flowrouter := router.Subrouter().PathPrefix("flow")
	//搜索岗位
	flowrouter.HandleFunc("/search", flowCtrl.Search).Methods("post","get",)
	
	//创建岗位
	flowrouter.HandleFunc("/create", flowCtrl.Create).Methods("post","put",)
	
	//更新岗位
	flowrouter.HandleFunc("/update", flowCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	flowrouter.HandleFunc("/delete", flowCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	flowrouter.HandleFunc("/getOne", flowCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新点位
	pointCtrl := &Point{}
	pointrouter := router.Subrouter().PathPrefix("point")
	//搜索点位
	pointrouter.HandleFunc("/search", pointCtrl.Search).Methods("post","get",)
	
	//搜索点位
	pointrouter.HandleFunc("/count", pointCtrl.Count).Methods("post","get",)
	
	//创建点位
	pointrouter.HandleFunc("/create", pointCtrl.Create).Methods("post","put",)
	
	//更新点位
	pointrouter.HandleFunc("/update", pointCtrl.Update).Methods("post","put",)
	
	//删除点位,系统默认都是逻辑删除
	pointrouter.HandleFunc("/delete", pointCtrl.Delete).Methods("post","delete",)
	
	//获取点位
	pointrouter.HandleFunc("/getOne", pointCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除资源管理,系统默认都是逻辑删除
	resourceCtrl := &Resource{}
	resourcerouter := router.Subrouter().PathPrefix("resource")
	//搜索资源管理
	resourcerouter.HandleFunc("/search", resourceCtrl.Search).Methods("post","get",)
	
	//创建资源管理
	resourcerouter.HandleFunc("/create", resourceCtrl.Create).Methods("post","put",)
	
	//更新资源管理
	resourcerouter.HandleFunc("/update", resourceCtrl.Update).Methods("post","put",)
	
	//删除资源管理,系统默认都是逻辑删除
	resourcerouter.HandleFunc("/delete", resourceCtrl.Delete).Methods("post","delete",)
	
	//获取资源管理
	resourcerouter.HandleFunc("/getOne", resourceCtrl.GetOne).Methods("post","get",)
	

	
	
	// patern/module/action
	restHandlerCtrl := &RestHandler{}
	restHandlerrouter := router.Subrouter().PathPrefix("restHandler")
	//patern/module/action
	restHandlerrouter.HandleFunc("/serveHTTP", restHandlerCtrl.ServeHTTP).Methods("post","get",)
	

	
	
	// 搜索
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 管理员专用
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 获取岗位
	deptCtrl := &Dept{}
	deptrouter := router.Subrouter().PathPrefix("dept")
	//搜索岗位
	deptrouter.HandleFunc("/search", deptCtrl.Search).Methods("post","get",)
	
	//创建岗位
	deptrouter.HandleFunc("/create", deptCtrl.Create).Methods("post","put",)
	
	//更新岗位
	deptrouter.HandleFunc("/update", deptCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/delete", deptCtrl.Delete).Methods("post","delete",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/tree", deptCtrl.Tree).Methods("post","get",)
	
	//获取岗位
	deptrouter.HandleFunc("/getOne", deptCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取字典
	dictCtrl := &Dict{}
	dictrouter := router.Subrouter().PathPrefix("dict")
	//搜索字典
	dictrouter.HandleFunc("/search", dictCtrl.Search).Methods("post","get",)
	
	//创建字典
	dictrouter.HandleFunc("/create", dictCtrl.Create).Methods("post","put",)
	
	//更新字典
	dictrouter.HandleFunc("/update", dictCtrl.Update).Methods("post","put",)
	
	//删除字典,系统默认都是逻辑删除
	dictrouter.HandleFunc("/delete", dictCtrl.Delete).Methods("post","delete",)
	
	//获取字典
	dictrouter.HandleFunc("/getOne", dictCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除,系统默认都是逻辑删除
	userinfoCtrl := &Userinfo{}
	userinforouter := router.Subrouter().PathPrefix("userinfo")
	//搜索
	userinforouter.HandleFunc("/search", userinfoCtrl.Search).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/count", userinfoCtrl.Count).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/updatedeptandname", userinfoCtrl.Updatedeptandname).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/updateUserinfo", userinfoCtrl.UpdateUserinfo).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/defaultRefId", userinfoCtrl.DefaultRefId).Methods("post","get",)
	
	//创建默认用户
	userinforouter.HandleFunc("/create", userinfoCtrl.Create).Methods("post","put",)
	
	//更新
	userinforouter.HandleFunc("/update", userinfoCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	userinforouter.HandleFunc("/delete", userinfoCtrl.Delete).Methods("post","delete",)
	
	//获取
	userinforouter.HandleFunc("/getOne", userinfoCtrl.GetOne).Methods("post","get",)
	
	//获取
	userinforouter.HandleFunc("/snsinfo", userinfoCtrl.Snsinfo).Methods("post","get",)
	

	
	
	// 更新
	rightsCtrl := &Rights{}
	rightsrouter := router.Subrouter().PathPrefix("rights")
	//搜索
	rightsrouter.HandleFunc("/search", rightsCtrl.Search).Methods("post","get",)
	
	//搜索
	rightsrouter.HandleFunc("/tree", rightsCtrl.Tree).Methods("post","get",)
	
	//创建
	rightsrouter.HandleFunc("/create", rightsCtrl.Create).Methods("post","put",)
	
	//更新
	rightsrouter.HandleFunc("/update", rightsCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rightsrouter.HandleFunc("/delete", rightsCtrl.Delete).Methods("post","delete",)
	
	//获取
	rightsrouter.HandleFunc("/getOne", rightsCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新短信发送记录
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// 修改当前账号手机号
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 创建岗位
	deptCtrl := &Dept{}
	deptrouter := router.Subrouter().PathPrefix("dept")
	//搜索岗位
	deptrouter.HandleFunc("/search", deptCtrl.Search).Methods("post","get",)
	
	//创建岗位
	deptrouter.HandleFunc("/create", deptCtrl.Create).Methods("post","put",)
	
	//更新岗位
	deptrouter.HandleFunc("/update", deptCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/delete", deptCtrl.Delete).Methods("post","delete",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/tree", deptCtrl.Tree).Methods("post","get",)
	
	//获取岗位
	deptrouter.HandleFunc("/getOne", deptCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新岗位
	flowCtrl := &Flow{}
	flowrouter := router.Subrouter().PathPrefix("flow")
	//搜索岗位
	flowrouter.HandleFunc("/search", flowCtrl.Search).Methods("post","get",)
	
	//创建岗位
	flowrouter.HandleFunc("/create", flowCtrl.Create).Methods("post","put",)
	
	//更新岗位
	flowrouter.HandleFunc("/update", flowCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	flowrouter.HandleFunc("/delete", flowCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	flowrouter.HandleFunc("/getOne", flowCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建岗位
	microappCtrl := &Microapp{}
	microapprouter := router.Subrouter().PathPrefix("microapp")
	//搜索岗位
	microapprouter.HandleFunc("/search", microappCtrl.Search).Methods("post","get",)
	
	//创建岗位
	microapprouter.HandleFunc("/create", microappCtrl.Create).Methods("post","put",)
	
	//更新岗位
	microapprouter.HandleFunc("/update", microappCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	microapprouter.HandleFunc("/delete", microappCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	microapprouter.HandleFunc("/getOne", microappCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除岗位,系统默认都是逻辑删除
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 删除问题采集,系统默认都是逻辑删除
	problemCtrl := &Problem{}
	problemrouter := router.Subrouter().PathPrefix("problem")
	//搜索问题采集
	problemrouter.HandleFunc("/count", problemCtrl.Count).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/countMine", problemCtrl.CountMine).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/search", problemCtrl.Search).Methods("post","get",)
	
	//创建问题采集
	problemrouter.HandleFunc("/create", problemCtrl.Create).Methods("post","put",)
	
	//更新问题采集
	problemrouter.HandleFunc("/update", problemCtrl.Update).Methods("post","put",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/delete", problemCtrl.Delete).Methods("post","delete",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/confirm", problemCtrl.Confirm).Methods("post","get",)
	
	//获取问题采集
	problemrouter.HandleFunc("/getOne", problemCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 删除区域,系统默认都是逻辑删除
	areaCtrl := &Area{}
	arearouter := router.Subrouter().PathPrefix("area")
	//搜索区域
	arearouter.HandleFunc("/search", areaCtrl.Search).Methods("post","get",)
	
	//创建区域
	arearouter.HandleFunc("/create", areaCtrl.Create).Methods("post","put",)
	
	//更新区域
	arearouter.HandleFunc("/update", areaCtrl.Update).Methods("post","put",)
	
	//删除区域,系统默认都是逻辑删除
	arearouter.HandleFunc("/delete", areaCtrl.Delete).Methods("post","delete",)
	
	//获取区域
	arearouter.HandleFunc("/getOne", areaCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新
	configCtrl := &Config{}
	configrouter := router.Subrouter().PathPrefix("config")
	//搜索
	configrouter.HandleFunc("/search", configCtrl.Search).Methods("post","get",)
	
	//创建
	configrouter.HandleFunc("/create", configCtrl.Create).Methods("post","put",)
	
	//创建
	configrouter.HandleFunc("/save", configCtrl.Save).Methods("post","get",)
	
	//更新
	configrouter.HandleFunc("/update", configCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	configrouter.HandleFunc("/delete", configCtrl.Delete).Methods("post","delete",)
	
	//获取
	configrouter.HandleFunc("/getOne", configCtrl.GetOne).Methods("post","get",)
	
	//获取
	configrouter.HandleFunc("/value", configCtrl.Value).Methods("post","get",)
	

	
	
	// 删除岗位,系统默认都是逻辑删除
	deptCtrl := &Dept{}
	deptrouter := router.Subrouter().PathPrefix("dept")
	//搜索岗位
	deptrouter.HandleFunc("/search", deptCtrl.Search).Methods("post","get",)
	
	//创建岗位
	deptrouter.HandleFunc("/create", deptCtrl.Create).Methods("post","put",)
	
	//更新岗位
	deptrouter.HandleFunc("/update", deptCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/delete", deptCtrl.Delete).Methods("post","delete",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/tree", deptCtrl.Tree).Methods("post","get",)
	
	//获取岗位
	deptrouter.HandleFunc("/getOne", deptCtrl.GetOne).Methods("post","get",)
	

	
	
	// 初始化上传策略
	ossCtrl := &Oss{}
	ossrouter := router.Subrouter().PathPrefix("oss")
	//初始化上传策略
	ossrouter.HandleFunc("/policy", ossCtrl.Policy).Methods("post","get",)
	
	//初始化上传策略
	ossrouter.HandleFunc("/callback", ossCtrl.Callback).Methods("post","get",)
	

	
	
	// 搜索
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索岗位
	flowCtrl := &Flow{}
	flowrouter := router.Subrouter().PathPrefix("flow")
	//搜索岗位
	flowrouter.HandleFunc("/search", flowCtrl.Search).Methods("post","get",)
	
	//创建岗位
	flowrouter.HandleFunc("/create", flowCtrl.Create).Methods("post","put",)
	
	//更新岗位
	flowrouter.HandleFunc("/update", flowCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	flowrouter.HandleFunc("/delete", flowCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	flowrouter.HandleFunc("/getOne", flowCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建岗位
	flowCtrl := &Flow{}
	flowrouter := router.Subrouter().PathPrefix("flow")
	//搜索岗位
	flowrouter.HandleFunc("/search", flowCtrl.Search).Methods("post","get",)
	
	//创建岗位
	flowrouter.HandleFunc("/create", flowCtrl.Create).Methods("post","put",)
	
	//更新岗位
	flowrouter.HandleFunc("/update", flowCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	flowrouter.HandleFunc("/delete", flowCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	flowrouter.HandleFunc("/getOne", flowCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取我发起的流程
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 更新机构信息
	orgCtrl := &Org{}
	orgrouter := router.Subrouter().PathPrefix("org")
	//搜索机构信息
	orgrouter.HandleFunc("/search", orgCtrl.Search).Methods("post","get",)
	
	//搜索机构信息
	orgrouter.HandleFunc("/mine", orgCtrl.Mine).Methods("post","get",)
	
	//创建机构信息
	orgrouter.HandleFunc("/create", orgCtrl.Create).Methods("post","put",)
	
	//更新机构信息
	orgrouter.HandleFunc("/update", orgCtrl.Update).Methods("post","put",)
	
	//删除机构信息,系统默认都是逻辑删除
	orgrouter.HandleFunc("/delete", orgCtrl.Delete).Methods("post","delete",)
	
	//获取机构信息
	orgrouter.HandleFunc("/getOne", orgCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建区域
	areaCtrl := &Area{}
	arearouter := router.Subrouter().PathPrefix("area")
	//搜索区域
	arearouter.HandleFunc("/search", areaCtrl.Search).Methods("post","get",)
	
	//创建区域
	arearouter.HandleFunc("/create", areaCtrl.Create).Methods("post","put",)
	
	//更新区域
	arearouter.HandleFunc("/update", areaCtrl.Update).Methods("post","put",)
	
	//删除区域,系统默认都是逻辑删除
	arearouter.HandleFunc("/delete", areaCtrl.Delete).Methods("post","delete",)
	
	//获取区域
	arearouter.HandleFunc("/getOne", areaCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除,系统默认都是逻辑删除
	articleCtrl := &Article{}
	articlerouter := router.Subrouter().PathPrefix("article")
	//搜索
	articlerouter.HandleFunc("/search", articleCtrl.Search).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/create", articleCtrl.Create).Methods("post","put",)
	
	//搜索
	articlerouter.HandleFunc("/count", articleCtrl.Count).Methods("post","get",)
	
	//搜索
	articlerouter.HandleFunc("/totalread", articleCtrl.Totalread).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/publish", articleCtrl.Publish).Methods("post","get",)
	
	//更新
	articlerouter.HandleFunc("/update", articleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	articlerouter.HandleFunc("/delete", articleCtrl.Delete).Methods("post","delete",)
	
	//获取
	articlerouter.HandleFunc("/getOne", articleCtrl.GetOne).Methods("post","get",)
	
	//获取
	articlerouter.HandleFunc("/addReadNum", articleCtrl.AddReadNum).Methods("post","get",)
	

	
	
	// 创建
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索岗位
	deptCtrl := &Dept{}
	deptrouter := router.Subrouter().PathPrefix("dept")
	//搜索岗位
	deptrouter.HandleFunc("/search", deptCtrl.Search).Methods("post","get",)
	
	//创建岗位
	deptrouter.HandleFunc("/create", deptCtrl.Create).Methods("post","put",)
	
	//更新岗位
	deptrouter.HandleFunc("/update", deptCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/delete", deptCtrl.Delete).Methods("post","delete",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/tree", deptCtrl.Tree).Methods("post","get",)
	
	//获取岗位
	deptrouter.HandleFunc("/getOne", deptCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取问题采集
	problemCtrl := &Problem{}
	problemrouter := router.Subrouter().PathPrefix("problem")
	//搜索问题采集
	problemrouter.HandleFunc("/count", problemCtrl.Count).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/countMine", problemCtrl.CountMine).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/search", problemCtrl.Search).Methods("post","get",)
	
	//创建问题采集
	problemrouter.HandleFunc("/create", problemCtrl.Create).Methods("post","put",)
	
	//更新问题采集
	problemrouter.HandleFunc("/update", problemCtrl.Update).Methods("post","put",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/delete", problemCtrl.Delete).Methods("post","delete",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/confirm", problemCtrl.Confirm).Methods("post","get",)
	
	//获取问题采集
	problemrouter.HandleFunc("/getOne", problemCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索短信发送记录
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新
	userinfoCtrl := &Userinfo{}
	userinforouter := router.Subrouter().PathPrefix("userinfo")
	//搜索
	userinforouter.HandleFunc("/search", userinfoCtrl.Search).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/count", userinfoCtrl.Count).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/updatedeptandname", userinfoCtrl.Updatedeptandname).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/updateUserinfo", userinfoCtrl.UpdateUserinfo).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/defaultRefId", userinfoCtrl.DefaultRefId).Methods("post","get",)
	
	//创建默认用户
	userinforouter.HandleFunc("/create", userinfoCtrl.Create).Methods("post","put",)
	
	//更新
	userinforouter.HandleFunc("/update", userinfoCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	userinforouter.HandleFunc("/delete", userinfoCtrl.Delete).Methods("post","delete",)
	
	//获取
	userinforouter.HandleFunc("/getOne", userinfoCtrl.GetOne).Methods("post","get",)
	
	//获取
	userinforouter.HandleFunc("/snsinfo", userinfoCtrl.Snsinfo).Methods("post","get",)
	

	
	
	// 用户专用
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// token 续期
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 搜索
	articleCtrl := &Article{}
	articlerouter := router.Subrouter().PathPrefix("article")
	//搜索
	articlerouter.HandleFunc("/search", articleCtrl.Search).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/create", articleCtrl.Create).Methods("post","put",)
	
	//搜索
	articlerouter.HandleFunc("/count", articleCtrl.Count).Methods("post","get",)
	
	//搜索
	articlerouter.HandleFunc("/totalread", articleCtrl.Totalread).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/publish", articleCtrl.Publish).Methods("post","get",)
	
	//更新
	articlerouter.HandleFunc("/update", articleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	articlerouter.HandleFunc("/delete", articleCtrl.Delete).Methods("post","delete",)
	
	//获取
	articlerouter.HandleFunc("/getOne", articleCtrl.GetOne).Methods("post","get",)
	
	//获取
	articlerouter.HandleFunc("/addReadNum", articleCtrl.AddReadNum).Methods("post","get",)
	

	
	
	// 获取
	articleCtrl := &Article{}
	articlerouter := router.Subrouter().PathPrefix("article")
	//搜索
	articlerouter.HandleFunc("/search", articleCtrl.Search).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/create", articleCtrl.Create).Methods("post","put",)
	
	//搜索
	articlerouter.HandleFunc("/count", articleCtrl.Count).Methods("post","get",)
	
	//搜索
	articlerouter.HandleFunc("/totalread", articleCtrl.Totalread).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/publish", articleCtrl.Publish).Methods("post","get",)
	
	//更新
	articlerouter.HandleFunc("/update", articleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	articlerouter.HandleFunc("/delete", articleCtrl.Delete).Methods("post","delete",)
	
	//获取
	articlerouter.HandleFunc("/getOne", articleCtrl.GetOne).Methods("post","get",)
	
	//获取
	articlerouter.HandleFunc("/addReadNum", articleCtrl.AddReadNum).Methods("post","get",)
	

	
	
	// 等待我审核的
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 获取
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索区域
	areaCtrl := &Area{}
	arearouter := router.Subrouter().PathPrefix("area")
	//搜索区域
	arearouter.HandleFunc("/search", areaCtrl.Search).Methods("post","get",)
	
	//创建区域
	arearouter.HandleFunc("/create", areaCtrl.Create).Methods("post","put",)
	
	//更新区域
	arearouter.HandleFunc("/update", areaCtrl.Update).Methods("post","put",)
	
	//删除区域,系统默认都是逻辑删除
	arearouter.HandleFunc("/delete", areaCtrl.Delete).Methods("post","delete",)
	
	//获取区域
	arearouter.HandleFunc("/getOne", areaCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新区域
	areaCtrl := &Area{}
	arearouter := router.Subrouter().PathPrefix("area")
	//搜索区域
	arearouter.HandleFunc("/search", areaCtrl.Search).Methods("post","get",)
	
	//创建区域
	arearouter.HandleFunc("/create", areaCtrl.Create).Methods("post","put",)
	
	//更新区域
	arearouter.HandleFunc("/update", areaCtrl.Update).Methods("post","put",)
	
	//删除区域,系统默认都是逻辑删除
	arearouter.HandleFunc("/delete", areaCtrl.Delete).Methods("post","delete",)
	
	//获取区域
	arearouter.HandleFunc("/getOne", areaCtrl.GetOne).Methods("post","get",)
	

	
	
	// 处理/upload 逻辑
	attachCtrl := &Attach{}
	attachrouter := router.Subrouter().PathPrefix("attach")
	///attach/render?key=filekey
	attachrouter.HandleFunc("/render", attachCtrl.Render).Methods("post","get",)
	
	//处理/upload 逻辑
	attachrouter.HandleFunc("/upload", attachCtrl.Upload).Methods("post","get",)
	

	
	
	// 创建
	configCtrl := &Config{}
	configrouter := router.Subrouter().PathPrefix("config")
	//搜索
	configrouter.HandleFunc("/search", configCtrl.Search).Methods("post","get",)
	
	//创建
	configrouter.HandleFunc("/create", configCtrl.Create).Methods("post","put",)
	
	//创建
	configrouter.HandleFunc("/save", configCtrl.Save).Methods("post","get",)
	
	//更新
	configrouter.HandleFunc("/update", configCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	configrouter.HandleFunc("/delete", configCtrl.Delete).Methods("post","delete",)
	
	//获取
	configrouter.HandleFunc("/getOne", configCtrl.GetOne).Methods("post","get",)
	
	//获取
	configrouter.HandleFunc("/value", configCtrl.Value).Methods("post","get",)
	

	
	
	// 创建短信发送记录
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// /attach/render?key=filekey
	attachCtrl := &Attach{}
	attachrouter := router.Subrouter().PathPrefix("attach")
	///attach/render?key=filekey
	attachrouter.HandleFunc("/render", attachCtrl.Render).Methods("post","get",)
	
	//处理/upload 逻辑
	attachrouter.HandleFunc("/upload", attachCtrl.Upload).Methods("post","get",)
	

	
	
	// 搜索岗位
	microappCtrl := &Microapp{}
	microapprouter := router.Subrouter().PathPrefix("microapp")
	//搜索岗位
	microapprouter.HandleFunc("/search", microappCtrl.Search).Methods("post","get",)
	
	//创建岗位
	microapprouter.HandleFunc("/create", microappCtrl.Create).Methods("post","put",)
	
	//更新岗位
	microapprouter.HandleFunc("/update", microappCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	microapprouter.HandleFunc("/delete", microappCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	microapprouter.HandleFunc("/getOne", microappCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索点位
	pointCtrl := &Point{}
	pointrouter := router.Subrouter().PathPrefix("point")
	//搜索点位
	pointrouter.HandleFunc("/search", pointCtrl.Search).Methods("post","get",)
	
	//搜索点位
	pointrouter.HandleFunc("/count", pointCtrl.Count).Methods("post","get",)
	
	//创建点位
	pointrouter.HandleFunc("/create", pointCtrl.Create).Methods("post","put",)
	
	//更新点位
	pointrouter.HandleFunc("/update", pointCtrl.Update).Methods("post","put",)
	
	//删除点位,系统默认都是逻辑删除
	pointrouter.HandleFunc("/delete", pointCtrl.Delete).Methods("post","delete",)
	
	//获取点位
	pointrouter.HandleFunc("/getOne", pointCtrl.GetOne).Methods("post","get",)
	

	
	
	// 授权
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索字典
	dictCtrl := &Dict{}
	dictrouter := router.Subrouter().PathPrefix("dict")
	//搜索字典
	dictrouter.HandleFunc("/search", dictCtrl.Search).Methods("post","get",)
	
	//创建字典
	dictrouter.HandleFunc("/create", dictCtrl.Create).Methods("post","put",)
	
	//更新字典
	dictrouter.HandleFunc("/update", dictCtrl.Update).Methods("post","put",)
	
	//删除字典,系统默认都是逻辑删除
	dictrouter.HandleFunc("/delete", dictCtrl.Delete).Methods("post","delete",)
	
	//获取字典
	dictrouter.HandleFunc("/getOne", dictCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索岗位
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 删除机构信息,系统默认都是逻辑删除
	orgCtrl := &Org{}
	orgrouter := router.Subrouter().PathPrefix("org")
	//搜索机构信息
	orgrouter.HandleFunc("/search", orgCtrl.Search).Methods("post","get",)
	
	//搜索机构信息
	orgrouter.HandleFunc("/mine", orgCtrl.Mine).Methods("post","get",)
	
	//创建机构信息
	orgrouter.HandleFunc("/create", orgCtrl.Create).Methods("post","put",)
	
	//更新机构信息
	orgrouter.HandleFunc("/update", orgCtrl.Update).Methods("post","put",)
	
	//删除机构信息,系统默认都是逻辑删除
	orgrouter.HandleFunc("/delete", orgCtrl.Delete).Methods("post","delete",)
	
	//获取机构信息
	orgrouter.HandleFunc("/getOne", orgCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新资源管理
	resourceCtrl := &Resource{}
	resourcerouter := router.Subrouter().PathPrefix("resource")
	//搜索资源管理
	resourcerouter.HandleFunc("/search", resourceCtrl.Search).Methods("post","get",)
	
	//创建资源管理
	resourcerouter.HandleFunc("/create", resourceCtrl.Create).Methods("post","put",)
	
	//更新资源管理
	resourcerouter.HandleFunc("/update", resourceCtrl.Update).Methods("post","put",)
	
	//删除资源管理,系统默认都是逻辑删除
	resourcerouter.HandleFunc("/delete", resourceCtrl.Delete).Methods("post","delete",)
	
	//获取资源管理
	resourcerouter.HandleFunc("/getOne", resourceCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取岗位
	microappCtrl := &Microapp{}
	microapprouter := router.Subrouter().PathPrefix("microapp")
	//搜索岗位
	microapprouter.HandleFunc("/search", microappCtrl.Search).Methods("post","get",)
	
	//创建岗位
	microapprouter.HandleFunc("/create", microappCtrl.Create).Methods("post","put",)
	
	//更新岗位
	microapprouter.HandleFunc("/update", microappCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	microapprouter.HandleFunc("/delete", microappCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	microapprouter.HandleFunc("/getOne", microappCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索机构信息
	orgCtrl := &Org{}
	orgrouter := router.Subrouter().PathPrefix("org")
	//搜索机构信息
	orgrouter.HandleFunc("/search", orgCtrl.Search).Methods("post","get",)
	
	//搜索机构信息
	orgrouter.HandleFunc("/mine", orgCtrl.Mine).Methods("post","get",)
	
	//创建机构信息
	orgrouter.HandleFunc("/create", orgCtrl.Create).Methods("post","put",)
	
	//更新机构信息
	orgrouter.HandleFunc("/update", orgCtrl.Update).Methods("post","put",)
	
	//删除机构信息,系统默认都是逻辑删除
	orgrouter.HandleFunc("/delete", orgCtrl.Delete).Methods("post","delete",)
	
	//获取机构信息
	orgrouter.HandleFunc("/getOne", orgCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建点位
	pointCtrl := &Point{}
	pointrouter := router.Subrouter().PathPrefix("point")
	//搜索点位
	pointrouter.HandleFunc("/search", pointCtrl.Search).Methods("post","get",)
	
	//搜索点位
	pointrouter.HandleFunc("/count", pointCtrl.Count).Methods("post","get",)
	
	//创建点位
	pointrouter.HandleFunc("/create", pointCtrl.Create).Methods("post","put",)
	
	//更新点位
	pointrouter.HandleFunc("/update", pointCtrl.Update).Methods("post","put",)
	
	//删除点位,系统默认都是逻辑删除
	pointrouter.HandleFunc("/delete", pointCtrl.Delete).Methods("post","delete",)
	
	//获取点位
	pointrouter.HandleFunc("/getOne", pointCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取点位
	pointCtrl := &Point{}
	pointrouter := router.Subrouter().PathPrefix("point")
	//搜索点位
	pointrouter.HandleFunc("/search", pointCtrl.Search).Methods("post","get",)
	
	//搜索点位
	pointrouter.HandleFunc("/count", pointCtrl.Count).Methods("post","get",)
	
	//创建点位
	pointrouter.HandleFunc("/create", pointCtrl.Create).Methods("post","put",)
	
	//更新点位
	pointrouter.HandleFunc("/update", pointCtrl.Update).Methods("post","put",)
	
	//删除点位,系统默认都是逻辑删除
	pointrouter.HandleFunc("/delete", pointCtrl.Delete).Methods("post","delete",)
	
	//获取点位
	pointrouter.HandleFunc("/getOne", pointCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新
	articleCtrl := &Article{}
	articlerouter := router.Subrouter().PathPrefix("article")
	//搜索
	articlerouter.HandleFunc("/search", articleCtrl.Search).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/create", articleCtrl.Create).Methods("post","put",)
	
	//搜索
	articlerouter.HandleFunc("/count", articleCtrl.Count).Methods("post","get",)
	
	//搜索
	articlerouter.HandleFunc("/totalread", articleCtrl.Totalread).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/publish", articleCtrl.Publish).Methods("post","get",)
	
	//更新
	articlerouter.HandleFunc("/update", articleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	articlerouter.HandleFunc("/delete", articleCtrl.Delete).Methods("post","delete",)
	
	//获取
	articlerouter.HandleFunc("/getOne", articleCtrl.GetOne).Methods("post","get",)
	
	//获取
	articlerouter.HandleFunc("/addReadNum", articleCtrl.AddReadNum).Methods("post","get",)
	

	
	
	// 删除,系统默认都是逻辑删除
	configCtrl := &Config{}
	configrouter := router.Subrouter().PathPrefix("config")
	//搜索
	configrouter.HandleFunc("/search", configCtrl.Search).Methods("post","get",)
	
	//创建
	configrouter.HandleFunc("/create", configCtrl.Create).Methods("post","put",)
	
	//创建
	configrouter.HandleFunc("/save", configCtrl.Save).Methods("post","get",)
	
	//更新
	configrouter.HandleFunc("/update", configCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	configrouter.HandleFunc("/delete", configCtrl.Delete).Methods("post","delete",)
	
	//获取
	configrouter.HandleFunc("/getOne", configCtrl.GetOne).Methods("post","get",)
	
	//获取
	configrouter.HandleFunc("/value", configCtrl.Value).Methods("post","get",)
	

	
	
	// 删除字典,系统默认都是逻辑删除
	dictCtrl := &Dict{}
	dictrouter := router.Subrouter().PathPrefix("dict")
	//搜索字典
	dictrouter.HandleFunc("/search", dictCtrl.Search).Methods("post","get",)
	
	//创建字典
	dictrouter.HandleFunc("/create", dictCtrl.Create).Methods("post","put",)
	
	//更新字典
	dictrouter.HandleFunc("/update", dictCtrl.Update).Methods("post","put",)
	
	//删除字典,系统默认都是逻辑删除
	dictrouter.HandleFunc("/delete", dictCtrl.Delete).Methods("post","delete",)
	
	//获取字典
	dictrouter.HandleFunc("/getOne", dictCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除岗位,系统默认都是逻辑删除
	flowCtrl := &Flow{}
	flowrouter := router.Subrouter().PathPrefix("flow")
	//搜索岗位
	flowrouter.HandleFunc("/search", flowCtrl.Search).Methods("post","get",)
	
	//创建岗位
	flowrouter.HandleFunc("/create", flowCtrl.Create).Methods("post","put",)
	
	//更新岗位
	flowrouter.HandleFunc("/update", flowCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	flowrouter.HandleFunc("/delete", flowCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	flowrouter.HandleFunc("/getOne", flowCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索资源管理
	resourceCtrl := &Resource{}
	resourcerouter := router.Subrouter().PathPrefix("resource")
	//搜索资源管理
	resourcerouter.HandleFunc("/search", resourceCtrl.Search).Methods("post","get",)
	
	//创建资源管理
	resourcerouter.HandleFunc("/create", resourceCtrl.Create).Methods("post","put",)
	
	//更新资源管理
	resourcerouter.HandleFunc("/update", resourceCtrl.Update).Methods("post","put",)
	
	//删除资源管理,系统默认都是逻辑删除
	resourcerouter.HandleFunc("/delete", resourceCtrl.Delete).Methods("post","delete",)
	
	//获取资源管理
	resourcerouter.HandleFunc("/getOne", resourceCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建资源管理
	resourceCtrl := &Resource{}
	resourcerouter := router.Subrouter().PathPrefix("resource")
	//搜索资源管理
	resourcerouter.HandleFunc("/search", resourceCtrl.Search).Methods("post","get",)
	
	//创建资源管理
	resourcerouter.HandleFunc("/create", resourceCtrl.Create).Methods("post","put",)
	
	//更新资源管理
	resourcerouter.HandleFunc("/update", resourceCtrl.Update).Methods("post","put",)
	
	//删除资源管理,系统默认都是逻辑删除
	resourcerouter.HandleFunc("/delete", resourceCtrl.Delete).Methods("post","delete",)
	
	//获取资源管理
	resourcerouter.HandleFunc("/getOne", resourceCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取机构信息
	orgCtrl := &Org{}
	orgrouter := router.Subrouter().PathPrefix("org")
	//搜索机构信息
	orgrouter.HandleFunc("/search", orgCtrl.Search).Methods("post","get",)
	
	//搜索机构信息
	orgrouter.HandleFunc("/mine", orgCtrl.Mine).Methods("post","get",)
	
	//创建机构信息
	orgrouter.HandleFunc("/create", orgCtrl.Create).Methods("post","put",)
	
	//更新机构信息
	orgrouter.HandleFunc("/update", orgCtrl.Update).Methods("post","put",)
	
	//删除机构信息,系统默认都是逻辑删除
	orgrouter.HandleFunc("/delete", orgCtrl.Delete).Methods("post","delete",)
	
	//获取机构信息
	orgrouter.HandleFunc("/getOne", orgCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除点位,系统默认都是逻辑删除
	pointCtrl := &Point{}
	pointrouter := router.Subrouter().PathPrefix("point")
	//搜索点位
	pointrouter.HandleFunc("/search", pointCtrl.Search).Methods("post","get",)
	
	//搜索点位
	pointrouter.HandleFunc("/count", pointCtrl.Count).Methods("post","get",)
	
	//创建点位
	pointrouter.HandleFunc("/create", pointCtrl.Create).Methods("post","put",)
	
	//更新点位
	pointrouter.HandleFunc("/update", pointCtrl.Update).Methods("post","put",)
	
	//删除点位,系统默认都是逻辑删除
	pointrouter.HandleFunc("/delete", pointCtrl.Delete).Methods("post","delete",)
	
	//获取点位
	pointrouter.HandleFunc("/getOne", pointCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取资源管理
	resourceCtrl := &Resource{}
	resourcerouter := router.Subrouter().PathPrefix("resource")
	//搜索资源管理
	resourcerouter.HandleFunc("/search", resourceCtrl.Search).Methods("post","get",)
	
	//创建资源管理
	resourcerouter.HandleFunc("/create", resourceCtrl.Create).Methods("post","put",)
	
	//更新资源管理
	resourcerouter.HandleFunc("/update", resourceCtrl.Update).Methods("post","put",)
	
	//删除资源管理,系统默认都是逻辑删除
	resourcerouter.HandleFunc("/delete", resourceCtrl.Delete).Methods("post","delete",)
	
	//获取资源管理
	resourcerouter.HandleFunc("/getOne", resourceCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除,系统默认都是逻辑删除
	rightsCtrl := &Rights{}
	rightsrouter := router.Subrouter().PathPrefix("rights")
	//搜索
	rightsrouter.HandleFunc("/search", rightsCtrl.Search).Methods("post","get",)
	
	//搜索
	rightsrouter.HandleFunc("/tree", rightsCtrl.Tree).Methods("post","get",)
	
	//创建
	rightsrouter.HandleFunc("/create", rightsCtrl.Create).Methods("post","put",)
	
	//更新
	rightsrouter.HandleFunc("/update", rightsCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rightsrouter.HandleFunc("/delete", rightsCtrl.Delete).Methods("post","delete",)
	
	//获取
	rightsrouter.HandleFunc("/getOne", rightsCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取区域
	areaCtrl := &Area{}
	arearouter := router.Subrouter().PathPrefix("area")
	//搜索区域
	arearouter.HandleFunc("/search", areaCtrl.Search).Methods("post","get",)
	
	//创建区域
	arearouter.HandleFunc("/create", areaCtrl.Create).Methods("post","put",)
	
	//更新区域
	arearouter.HandleFunc("/update", areaCtrl.Update).Methods("post","put",)
	
	//删除区域,系统默认都是逻辑删除
	arearouter.HandleFunc("/delete", areaCtrl.Delete).Methods("post","delete",)
	
	//获取区域
	arearouter.HandleFunc("/getOne", areaCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索
	configCtrl := &Config{}
	configrouter := router.Subrouter().PathPrefix("config")
	//搜索
	configrouter.HandleFunc("/search", configCtrl.Search).Methods("post","get",)
	
	//创建
	configrouter.HandleFunc("/create", configCtrl.Create).Methods("post","put",)
	
	//创建
	configrouter.HandleFunc("/save", configCtrl.Save).Methods("post","get",)
	
	//更新
	configrouter.HandleFunc("/update", configCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	configrouter.HandleFunc("/delete", configCtrl.Delete).Methods("post","delete",)
	
	//获取
	configrouter.HandleFunc("/getOne", configCtrl.GetOne).Methods("post","get",)
	
	//获取
	configrouter.HandleFunc("/value", configCtrl.Value).Methods("post","get",)
	

	
	
	// 创建岗位
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 更新岗位
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 获取
	userinfoCtrl := &Userinfo{}
	userinforouter := router.Subrouter().PathPrefix("userinfo")
	//搜索
	userinforouter.HandleFunc("/search", userinfoCtrl.Search).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/count", userinfoCtrl.Count).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/updatedeptandname", userinfoCtrl.Updatedeptandname).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/updateUserinfo", userinfoCtrl.UpdateUserinfo).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/defaultRefId", userinfoCtrl.DefaultRefId).Methods("post","get",)
	
	//创建默认用户
	userinforouter.HandleFunc("/create", userinfoCtrl.Create).Methods("post","put",)
	
	//更新
	userinforouter.HandleFunc("/update", userinfoCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	userinforouter.HandleFunc("/delete", userinfoCtrl.Delete).Methods("post","delete",)
	
	//获取
	userinforouter.HandleFunc("/getOne", userinfoCtrl.GetOne).Methods("post","get",)
	
	//获取
	userinforouter.HandleFunc("/snsinfo", userinfoCtrl.Snsinfo).Methods("post","get",)
	

	
	
	// 更新问题采集
	problemCtrl := &Problem{}
	problemrouter := router.Subrouter().PathPrefix("problem")
	//搜索问题采集
	problemrouter.HandleFunc("/count", problemCtrl.Count).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/countMine", problemCtrl.CountMine).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/search", problemCtrl.Search).Methods("post","get",)
	
	//创建问题采集
	problemrouter.HandleFunc("/create", problemCtrl.Create).Methods("post","put",)
	
	//更新问题采集
	problemrouter.HandleFunc("/update", problemCtrl.Update).Methods("post","put",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/delete", problemCtrl.Delete).Methods("post","delete",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/confirm", problemCtrl.Confirm).Methods("post","get",)
	
	//获取问题采集
	problemrouter.HandleFunc("/getOne", problemCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取
	rightsCtrl := &Rights{}
	rightsrouter := router.Subrouter().PathPrefix("rights")
	//搜索
	rightsrouter.HandleFunc("/search", rightsCtrl.Search).Methods("post","get",)
	
	//搜索
	rightsrouter.HandleFunc("/tree", rightsCtrl.Tree).Methods("post","get",)
	
	//创建
	rightsrouter.HandleFunc("/create", rightsCtrl.Create).Methods("post","put",)
	
	//更新
	rightsrouter.HandleFunc("/update", rightsCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rightsrouter.HandleFunc("/delete", rightsCtrl.Delete).Methods("post","delete",)
	
	//获取
	rightsrouter.HandleFunc("/getOne", rightsCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除,系统默认都是逻辑删除
	roleCtrl := &Role{}
	rolerouter := router.Subrouter().PathPrefix("role")
	//搜索
	rolerouter.HandleFunc("/search", roleCtrl.Search).Methods("post","get",)
	
	//创建
	rolerouter.HandleFunc("/create", roleCtrl.Create).Methods("post","put",)
	
	//授权
	rolerouter.HandleFunc("/grant", roleCtrl.Grant).Methods("post","get",)
	
	//更新
	rolerouter.HandleFunc("/update", roleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rolerouter.HandleFunc("/delete", roleCtrl.Delete).Methods("post","delete",)
	
	//获取
	rolerouter.HandleFunc("/getOne", roleCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建问题采集
	problemCtrl := &Problem{}
	problemrouter := router.Subrouter().PathPrefix("problem")
	//搜索问题采集
	problemrouter.HandleFunc("/count", problemCtrl.Count).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/countMine", problemCtrl.CountMine).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/search", problemCtrl.Search).Methods("post","get",)
	
	//创建问题采集
	problemrouter.HandleFunc("/create", problemCtrl.Create).Methods("post","put",)
	
	//更新问题采集
	problemrouter.HandleFunc("/update", problemCtrl.Update).Methods("post","put",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/delete", problemCtrl.Delete).Methods("post","delete",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/confirm", problemCtrl.Confirm).Methods("post","get",)
	
	//获取问题采集
	problemrouter.HandleFunc("/getOne", problemCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索
	rightsCtrl := &Rights{}
	rightsrouter := router.Subrouter().PathPrefix("rights")
	//搜索
	rightsrouter.HandleFunc("/search", rightsCtrl.Search).Methods("post","get",)
	
	//搜索
	rightsrouter.HandleFunc("/tree", rightsCtrl.Tree).Methods("post","get",)
	
	//创建
	rightsrouter.HandleFunc("/create", rightsCtrl.Create).Methods("post","put",)
	
	//更新
	rightsrouter.HandleFunc("/update", rightsCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rightsrouter.HandleFunc("/delete", rightsCtrl.Delete).Methods("post","delete",)
	
	//获取
	rightsrouter.HandleFunc("/getOne", rightsCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取短信发送记录
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建
	articleCtrl := &Article{}
	articlerouter := router.Subrouter().PathPrefix("article")
	//搜索
	articlerouter.HandleFunc("/search", articleCtrl.Search).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/create", articleCtrl.Create).Methods("post","put",)
	
	//搜索
	articlerouter.HandleFunc("/count", articleCtrl.Count).Methods("post","get",)
	
	//搜索
	articlerouter.HandleFunc("/totalread", articleCtrl.Totalread).Methods("post","get",)
	
	//创建
	articlerouter.HandleFunc("/publish", articleCtrl.Publish).Methods("post","get",)
	
	//更新
	articlerouter.HandleFunc("/update", articleCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	articlerouter.HandleFunc("/delete", articleCtrl.Delete).Methods("post","delete",)
	
	//获取
	articlerouter.HandleFunc("/getOne", articleCtrl.GetOne).Methods("post","get",)
	
	//获取
	articlerouter.HandleFunc("/addReadNum", articleCtrl.AddReadNum).Methods("post","get",)
	

	
	
	// 获取
	configCtrl := &Config{}
	configrouter := router.Subrouter().PathPrefix("config")
	//搜索
	configrouter.HandleFunc("/search", configCtrl.Search).Methods("post","get",)
	
	//创建
	configrouter.HandleFunc("/create", configCtrl.Create).Methods("post","put",)
	
	//创建
	configrouter.HandleFunc("/save", configCtrl.Save).Methods("post","get",)
	
	//更新
	configrouter.HandleFunc("/update", configCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	configrouter.HandleFunc("/delete", configCtrl.Delete).Methods("post","delete",)
	
	//获取
	configrouter.HandleFunc("/getOne", configCtrl.GetOne).Methods("post","get",)
	
	//获取
	configrouter.HandleFunc("/value", configCtrl.Value).Methods("post","get",)
	

	
	
	// 创建字典
	dictCtrl := &Dict{}
	dictrouter := router.Subrouter().PathPrefix("dict")
	//搜索字典
	dictrouter.HandleFunc("/search", dictCtrl.Search).Methods("post","get",)
	
	//创建字典
	dictrouter.HandleFunc("/create", dictCtrl.Create).Methods("post","put",)
	
	//更新字典
	dictrouter.HandleFunc("/update", dictCtrl.Update).Methods("post","put",)
	
	//删除字典,系统默认都是逻辑删除
	dictrouter.HandleFunc("/delete", dictCtrl.Delete).Methods("post","delete",)
	
	//获取字典
	dictrouter.HandleFunc("/getOne", dictCtrl.GetOne).Methods("post","get",)
	

	
	
	// 获取岗位
	instanceCtrl := &Instance{}
	instancerouter := router.Subrouter().PathPrefix("instance")
	//搜索岗位
	instancerouter.HandleFunc("/search", instanceCtrl.Search).Methods("post","get",)
	
	//创建岗位
	instancerouter.HandleFunc("/create", instanceCtrl.Create).Methods("post","put",)
	
	//更新岗位
	instancerouter.HandleFunc("/update", instanceCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	instancerouter.HandleFunc("/delete", instanceCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	instancerouter.HandleFunc("/getOne", instanceCtrl.GetOne).Methods("post","get",)
	
	//获取我发起的流程
	instancerouter.HandleFunc("/whatiinit", instanceCtrl.Whatiinit).Methods("post","get",)
	
	//等待我审核的
	instancerouter.HandleFunc("/whatneedapprove", instanceCtrl.Whatneedapprove).Methods("post","get",)
	

	
	
	// 搜索问题采集
	problemCtrl := &Problem{}
	problemrouter := router.Subrouter().PathPrefix("problem")
	//搜索问题采集
	problemrouter.HandleFunc("/count", problemCtrl.Count).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/countMine", problemCtrl.CountMine).Methods("post","get",)
	
	//搜索问题采集
	problemrouter.HandleFunc("/search", problemCtrl.Search).Methods("post","get",)
	
	//创建问题采集
	problemrouter.HandleFunc("/create", problemCtrl.Create).Methods("post","put",)
	
	//更新问题采集
	problemrouter.HandleFunc("/update", problemCtrl.Update).Methods("post","put",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/delete", problemCtrl.Delete).Methods("post","delete",)
	
	//删除问题采集,系统默认都是逻辑删除
	problemrouter.HandleFunc("/confirm", problemCtrl.Confirm).Methods("post","get",)
	
	//获取问题采集
	problemrouter.HandleFunc("/getOne", problemCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建
	rightsCtrl := &Rights{}
	rightsrouter := router.Subrouter().PathPrefix("rights")
	//搜索
	rightsrouter.HandleFunc("/search", rightsCtrl.Search).Methods("post","get",)
	
	//搜索
	rightsrouter.HandleFunc("/tree", rightsCtrl.Tree).Methods("post","get",)
	
	//创建
	rightsrouter.HandleFunc("/create", rightsCtrl.Create).Methods("post","put",)
	
	//更新
	rightsrouter.HandleFunc("/update", rightsCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	rightsrouter.HandleFunc("/delete", rightsCtrl.Delete).Methods("post","delete",)
	
	//获取
	rightsrouter.HandleFunc("/getOne", rightsCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除短信发送记录,系统默认都是逻辑删除
	smstaskCtrl := &Smstask{}
	smstaskrouter := router.Subrouter().PathPrefix("smstask")
	//搜索
	smstaskrouter.HandleFunc("/send", smstaskCtrl.Send).Methods("post","get",)
	
	//搜索短信发送记录
	smstaskrouter.HandleFunc("/search", smstaskCtrl.Search).Methods("post","get",)
	
	//创建短信发送记录
	smstaskrouter.HandleFunc("/create", smstaskCtrl.Create).Methods("post","put",)
	
	//更新短信发送记录
	smstaskrouter.HandleFunc("/update", smstaskCtrl.Update).Methods("post","put",)
	
	//删除短信发送记录,系统默认都是逻辑删除
	smstaskrouter.HandleFunc("/delete", smstaskCtrl.Delete).Methods("post","delete",)
	
	//获取短信发送记录
	smstaskrouter.HandleFunc("/getOne", smstaskCtrl.GetOne).Methods("post","get",)
	

	
	
	// 这是router
	captchaCtrl := &Captcha{}
	captcharouter := router.Subrouter().PathPrefix("captcha")
	//这是router
	captcharouter.HandleFunc("/get", captchaCtrl.Get).Methods("post","get",)
	

	
	
	// 更新岗位
	microappCtrl := &Microapp{}
	microapprouter := router.Subrouter().PathPrefix("microapp")
	//搜索岗位
	microapprouter.HandleFunc("/search", microappCtrl.Search).Methods("post","get",)
	
	//创建岗位
	microapprouter.HandleFunc("/create", microappCtrl.Create).Methods("post","put",)
	
	//更新岗位
	microapprouter.HandleFunc("/update", microappCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	microapprouter.HandleFunc("/delete", microappCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	microapprouter.HandleFunc("/getOne", microappCtrl.GetOne).Methods("post","get",)
	

	
	
	// 删除岗位,系统默认都是逻辑删除
	microappCtrl := &Microapp{}
	microapprouter := router.Subrouter().PathPrefix("microapp")
	//搜索岗位
	microapprouter.HandleFunc("/search", microappCtrl.Search).Methods("post","get",)
	
	//创建岗位
	microapprouter.HandleFunc("/create", microappCtrl.Create).Methods("post","put",)
	
	//更新岗位
	microapprouter.HandleFunc("/update", microappCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	microapprouter.HandleFunc("/delete", microappCtrl.Delete).Methods("post","delete",)
	
	//获取岗位
	microapprouter.HandleFunc("/getOne", microappCtrl.GetOne).Methods("post","get",)
	

	
	
	// 创建机构信息
	orgCtrl := &Org{}
	orgrouter := router.Subrouter().PathPrefix("org")
	//搜索机构信息
	orgrouter.HandleFunc("/search", orgCtrl.Search).Methods("post","get",)
	
	//搜索机构信息
	orgrouter.HandleFunc("/mine", orgCtrl.Mine).Methods("post","get",)
	
	//创建机构信息
	orgrouter.HandleFunc("/create", orgCtrl.Create).Methods("post","put",)
	
	//更新机构信息
	orgrouter.HandleFunc("/update", orgCtrl.Update).Methods("post","put",)
	
	//删除机构信息,系统默认都是逻辑删除
	orgrouter.HandleFunc("/delete", orgCtrl.Delete).Methods("post","delete",)
	
	//获取机构信息
	orgrouter.HandleFunc("/getOne", orgCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新
	accountCtrl := &Account{}
	accountrouter := router.Subrouter().PathPrefix("account")
	//搜索
	accountrouter.HandleFunc("/register", accountCtrl.Register).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/bindwxuser", accountCtrl.Bindwxuser).Methods("post","get",)
	
	//搜索
	accountrouter.HandleFunc("/updateprofile", accountCtrl.Updateprofile).Methods("post","put",)
	
	//创建
	accountrouter.HandleFunc("/login", accountCtrl.Login).Methods("post","get",)
	
	//创建
	accountrouter.HandleFunc("/resetPwd", accountCtrl.ResetPwd).Methods("post","get",)
	
	//用户专用
	accountrouter.HandleFunc("/updateMyPwd", accountCtrl.UpdateMyPwd).Methods("post","put",)
	
	//管理员专用
	accountrouter.HandleFunc("/updatePwd", accountCtrl.UpdatePwd).Methods("post","put",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/resetmobile", accountCtrl.Resetmobile).Methods("post","get",)
	
	//修改当前账号手机号
	accountrouter.HandleFunc("/updateUserName", accountCtrl.UpdateUserName).Methods("post","put",)
	
	//更新
	accountrouter.HandleFunc("/getInfo", accountCtrl.GetInfo).Methods("post","get",)
	
	//token 续期
	accountrouter.HandleFunc("/renewal", accountCtrl.Renewal).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/enable", accountCtrl.Enable).Methods("post","get",)
	
	//更新
	accountrouter.HandleFunc("/disable", accountCtrl.Disable).Methods("post","get",)
	

	
	
	// 更新岗位
	deptCtrl := &Dept{}
	deptrouter := router.Subrouter().PathPrefix("dept")
	//搜索岗位
	deptrouter.HandleFunc("/search", deptCtrl.Search).Methods("post","get",)
	
	//创建岗位
	deptrouter.HandleFunc("/create", deptCtrl.Create).Methods("post","put",)
	
	//更新岗位
	deptrouter.HandleFunc("/update", deptCtrl.Update).Methods("post","put",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/delete", deptCtrl.Delete).Methods("post","delete",)
	
	//删除岗位,系统默认都是逻辑删除
	deptrouter.HandleFunc("/tree", deptCtrl.Tree).Methods("post","get",)
	
	//获取岗位
	deptrouter.HandleFunc("/getOne", deptCtrl.GetOne).Methods("post","get",)
	

	
	
	// 更新字典
	dictCtrl := &Dict{}
	dictrouter := router.Subrouter().PathPrefix("dict")
	//搜索字典
	dictrouter.HandleFunc("/search", dictCtrl.Search).Methods("post","get",)
	
	//创建字典
	dictrouter.HandleFunc("/create", dictCtrl.Create).Methods("post","put",)
	
	//更新字典
	dictrouter.HandleFunc("/update", dictCtrl.Update).Methods("post","put",)
	
	//删除字典,系统默认都是逻辑删除
	dictrouter.HandleFunc("/delete", dictCtrl.Delete).Methods("post","delete",)
	
	//获取字典
	dictrouter.HandleFunc("/getOne", dictCtrl.GetOne).Methods("post","get",)
	

	
	
	// 搜索
	userinfoCtrl := &Userinfo{}
	userinforouter := router.Subrouter().PathPrefix("userinfo")
	//搜索
	userinforouter.HandleFunc("/search", userinfoCtrl.Search).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/count", userinfoCtrl.Count).Methods("post","get",)
	
	//搜索
	userinforouter.HandleFunc("/updatedeptandname", userinfoCtrl.Updatedeptandname).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/updateUserinfo", userinfoCtrl.UpdateUserinfo).Methods("post","put",)
	
	//搜索
	userinforouter.HandleFunc("/defaultRefId", userinfoCtrl.DefaultRefId).Methods("post","get",)
	
	//创建默认用户
	userinforouter.HandleFunc("/create", userinfoCtrl.Create).Methods("post","put",)
	
	//更新
	userinforouter.HandleFunc("/update", userinfoCtrl.Update).Methods("post","put",)
	
	//删除,系统默认都是逻辑删除
	userinforouter.HandleFunc("/delete", userinfoCtrl.Delete).Methods("post","delete",)
	
	//获取
	userinforouter.HandleFunc("/getOne", userinfoCtrl.GetOne).Methods("post","get",)
	
	//获取
	userinforouter.HandleFunc("/snsinfo", userinfoCtrl.Snsinfo).Methods("post","get",)
	

	
}
func init() {
	InitRouter(DefaultRouter)
}
