package utils_test

type PublicStruct struct {
	obj privateStruct
}

func (p *PublicStruct) PublicFunc() {
	p.privateFunc()
}

func (p PublicStruct) PublicFunc2() {
	p.privateFunc()
}

func (p *PublicStruct) privateFunc() {
	p.obj.Test()
}

type privateStruct struct{}

func (p *privateStruct) Test() {}
