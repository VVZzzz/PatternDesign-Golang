package abstract_factory

import "fmt"

// MemberMainRecordDAO 会员记录表
type MemberMainRecordDAO interface {
	SaveMainRecord()
}

// MemberSerialRecordDAO 会员操作流水表
type MemberSerialRecordDAO interface {
	SaveSerialRecord()
}

// MemberDAOFactory 抽象模式工厂接口
type MemberDAOFactory interface {
	CreateMemberMainRecordDAO() MemberMainRecordDAO
	CreateMemberSerialRecordDAO() MemberSerialRecordDAO
}

// PayMemberMainRecordDAO 付费会员记录
type PayMemberMainRecordDAO struct {
}

func (m *PayMemberMainRecordDAO) SaveMainRecord() {
	fmt.Println("save pay member main record")
}

type PayMemberSerialRecordDAO struct {
}

func (m *PayMemberSerialRecordDAO) SaveSerialRecord() {
	fmt.Println("save pay member serial record")
}

// PayMemberFactory 付费会员抽象工厂的具体实现
type PayMemberFactory struct {
}

func (p *PayMemberFactory) CreateMemberMainRecordDAO() MemberMainRecordDAO {
	return &PayMemberMainRecordDAO{}
}

func (p *PayMemberFactory) CreateMemberSerialRecordDAO() MemberSerialRecordDAO {
	return &PayMemberSerialRecordDAO{}
}

// NonPayMemberMainRecordDAO 非付费会员记录
type NonPayMemberMainRecordDAO struct {
}

func (m *NonPayMemberMainRecordDAO) SaveMainRecord() {
	fmt.Println("save non-pay member main record")
}

type NonPayMemberSerialRecordDAO struct {
}

func (m *NonPayMemberSerialRecordDAO) SaveSerialRecord() {
	fmt.Println("save non-pay member serial record")
}

// NonPayMemberFactory 非付费会员抽象工厂的具体实现
type NonPayMemberFactory struct {
}

func (p *NonPayMemberFactory) CreateMemberMainRecordDAO() MemberMainRecordDAO {
	return &NonPayMemberMainRecordDAO{}
}

func (p *NonPayMemberFactory) CreateMemberSerialRecordDAO() MemberSerialRecordDAO {
	return &NonPayMemberSerialRecordDAO{}
}
