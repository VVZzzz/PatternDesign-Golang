package flyweight

import (
	"fmt"
	"testing"
)

/*
"张三,人员列表,查看",
"李四,人员列表,查看",
"李四,薪资数据,查看",
"李四,薪资数据,修改",
*/
func TestFlyWeight(t *testing.T) {
	//init db
	InitTestDB()
	mgr := GetSecurityPermitMgr()
	mgr.Login("张三")
	mgr.Login("李四")

	// debug
	fmt.Printf("data %+v\n", mgr.Data)
	mgr.HasPermit("张三", "薪资数据", "查看")
	mgr.HasPermit("李四", "薪资数据", "查看")

}
