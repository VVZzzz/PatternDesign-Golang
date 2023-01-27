package abstract_factory

func operateMemberRecord(f MemberDAOFactory) {
	f.CreateMemberMainRecordDAO().SaveMainRecord()
	f.CreateMemberSerialRecordDAO().SaveSerialRecord()
}

func PayMember() {
	f := &PayMemberFactory{}
	operateMemberRecord(f)
}

func NonPayMember() {
	f := &PayMemberFactory{}
	operateMemberRecord(f)
}
