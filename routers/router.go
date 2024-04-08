package routers

import (
	"multi_version_swagger_demo/controllers/test_api_v1"
	"multi_version_swagger_demo/controllers/test_api_v2"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	initAPIV1()
	initAPIV2()
}

// @NamespacePrefix /v1
// @APIVersion 1.0.0
// @Title API v1
// @Description API v1 description
// @Schemes http
func initAPIV1() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&test_api_v1.UserController{})),
	)
	beego.AddNamespace(ns)
}

// @NamespacePrefix /v2
// @APIVersion 2.0.0
// @Title API v1
// @Description API v2 description
// @Schemes http
func initAPIV2() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/user", beego.NSInclude(&test_api_v2.UserController{})),
	)
	beego.AddNamespace(ns)
}
