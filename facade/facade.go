package facade

func NewDoAPI() DoAPI {
	return &doApiImpl{
		a: NewModuleAAPI(),
		b: NewModuleBAPI(),
	}

}

type DoAPI interface {
	DoFunc() string
}

type doApiImpl struct {
	a ModuleAAPI
	b ModuleBAPI
}

func (d *doApiImpl) DoFunc() string {
	aStr := d.a.DoFuncA()
	bStr := d.b.DoFuncB()
	return aStr + bStr
}

// ModuleA API
func NewModuleAAPI() ModuleAAPI {
	return &moduleAAPIImpl{}

}

type ModuleAAPI interface {
	DoFuncA() string
}

type moduleAAPIImpl struct {
}

func (m *moduleAAPIImpl) DoFuncA() string {
	return "Module A do funcA"
}

// ModuleA API
func NewModuleBAPI() ModuleBAPI {
	return &moduleBAPIImpl{}

}

type ModuleBAPI interface {
	DoFuncB() string
}

type moduleBAPIImpl struct {
}

func (d *moduleBAPIImpl) DoFuncB() string {
	return "Module B do funcB"
}
