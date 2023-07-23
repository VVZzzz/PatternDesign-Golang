package builder

func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}

// Director 负责整体算法/流程框架的构建
type Director struct {
	builder IBuilder
}

func (r *Director) Construct() {
	// 流程/算法: A+B+C
	r.builder.ProcessA()
	r.builder.ProcessB()
	r.builder.ProcessC()
}

func (r *Director) GetResult() interface{} {
	return r.builder.OutputResult()
}

// IBuilder 负责具体每个步骤/流程的构建
type IBuilder interface {
	ProcessA()
	ProcessB()
	ProcessC()

	OutputResult() interface{}
}

type BuilderX struct {
	ctxResult string
}

func (r *BuilderX) ProcessA() {
	r.ctxResult += "X_A "
}

func (r *BuilderX) ProcessB() {
	r.ctxResult += "X_B "
}

func (r *BuilderX) ProcessC() {
	r.ctxResult += "X_C "
}

func (r *BuilderX) OutputResult() interface{} {
	return r.ctxResult
}

type BuilderY struct {
	ctxResult string
}

func (r *BuilderY) ProcessA() {
	r.ctxResult += "Y_A "
}

func (r *BuilderY) ProcessB() {
	r.ctxResult += "Y_B "
}

func (r *BuilderY) ProcessC() {
	r.ctxResult += "Y_C "
}

func (r *BuilderY) OutputResult() interface{} {
	return r.ctxResult
}
