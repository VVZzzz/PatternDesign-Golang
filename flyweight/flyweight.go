package flyweight

import (
	"fmt"
	"strings"
)

// FlyWeight 享元interface
type FlyWeight interface {
	/*
	*判断传入的安全实体和权限,是否和享元对象内部状态匹配配
	*@param securityEntity安全实体
	*@param permit 权限
	*@return true表示匹配,false表示不匹配
	 */
	Match(securityEntity string, permit string) bool
}

// AuthorizationFlyWeight 封装授权数据中重复出现部分的享元对象
type AuthorizationFlyWeight struct {
	securityEntity string
	permit         string
}

func NewAuthorizationFlyWeight(key string) *AuthorizationFlyWeight {
	keySlice := strings.Split(key, ",")
	if len(keySlice) < 2 {
		return nil
	}
	return &AuthorizationFlyWeight{securityEntity: keySlice[0], permit: keySlice[1]}
}

func (w *AuthorizationFlyWeight) GetSecurityEntity() string {
	return w.securityEntity
}

func (w *AuthorizationFlyWeight) Match(securityEntity string, permit string) bool {
	return w.securityEntity == securityEntity && w.permit == permit
}

// 单例模式的享元工厂-----------------------------------------------
var flyWeightFactory *FlyWeightFactory

type FlyWeightFactory struct {
	maps map[string]FlyWeight
}

func (f *FlyWeightFactory) GetFlyWeight(key string) FlyWeight {
	val := f.maps[key]
	if val == nil {
		val = NewAuthorizationFlyWeight(key)
		f.maps[key] = val
	}
	return val
}

func GetFlyWeightFactory() *FlyWeightFactory {
	if flyWeightFactory == nil {
		flyWeightFactory = &FlyWeightFactory{make(map[string]FlyWeight)}
	}
	return flyWeightFactory
}

// 测试用 DB-----------------------------------------------
/**
模拟用来测试的数据库: 名字,安全实体,权限
"张三,人员列表,查看"
"李四,人员列表,查看"
"李四,薪资数据,查看"
"李四,薪资数据,修改"
*/

type TestDB struct {
	RawData []string
}

var testDB *TestDB

func GetTestDB() *TestDB {
	return testDB
}

func InitTestDB() {
	testDB = &TestDB{RawData: []string{
		"张三,人员列表,查看",
		"李四,人员列表,查看",
		"李四,薪资数据,查看",
		"李四,薪资数据,修改",
	}}

}

// 测试用安全权限管理器-----------------------------------------------
type SecurityPermitMgr struct {
	/*
		在运行期间,用来存放登录人员对应的权限
		-在Web应用中,这些数据通常会存放到session中
	*/
	Data map[string][]FlyWeight
}

/**
 * 模拟登录的功能
 * 用户登陆后,将所用的权限从数据库取出并放入内存/缓存中
 */
func (mgr *SecurityPermitMgr) Login(user string) {
	secPermits := mgr.QueryAuthorizationFlyWeightByUser(user)
	mgr.Data[user] = secPermits
}

func (mgr *SecurityPermitMgr) QueryAuthorizationFlyWeightByUser(user string) []FlyWeight {
	var permits []FlyWeight
	// query db
	db := GetTestDB()
	for _, data := range db.RawData {
		slice := strings.Split(data, ",")
		if len(slice) < 2 {
			continue
		}
		if user != slice[0] {
			continue
		}
		permitFlyWeight := GetFlyWeightFactory().GetFlyWeight(slice[1] + "," + slice[2])
		permits = append(permits, permitFlyWeight)
	}

	return permits
}

func (mgr *SecurityPermitMgr) HasPermit(user string, securityEntity string, permit string) bool {
	permits := mgr.QueryAuthorizationFlyWeightByUser(user)
	if len(permits) <= 0 {
		fmt.Printf("%s 未登录或者没有分配任何权限\n", user)
		return false
	}
	var ret bool
	//注意观察实例地址,确定是否是一个实例
	fmt.Printf("HasPermit %+v\n", permits)
	for _, p := range permits {
		if p.Match(securityEntity, permit) {
			fmt.Println(true)
			ret = true
		}
	}
	if !ret {
		fmt.Println(false)
	}
	return ret
}

var securityPermitMgr *SecurityPermitMgr

func GetSecurityPermitMgr() *SecurityPermitMgr {
	if securityPermitMgr == nil {
		securityPermitMgr = &SecurityPermitMgr{map[string][]FlyWeight{}}
	}
	return securityPermitMgr
}
