package generator

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type CodeInfo struct {
	PkIFName    string // インターフェイス側パッケージ名称
	PkCTName    string // コンテナ側パッケージ名称
	PkMockName  string // モック側パッケージ名称
	IFPath      string
	CTPath      string
	MockPath    string
	PublicName  string
	PrivateName string
	MockFlg     bool
}

func (info CodeInfo) CreateCode() {
	createInterface(info)
	createContainer(info)
	createMock(info)
}

func createInterface(info CodeInfo) {
	f := jen.NewFile(info.PkIFName)

	f.Type().Id(info.PublicName).Interface()

	fmt.Printf("%#v", f)
}

func createContainer(info CodeInfo) {
	f := jen.NewFile(info.PkCTName)

	f.Type().Id(info.PrivateName).Struct()

	f.Func().
		Id(fmt.Sprintf("New"+info.PublicName)).
		Params().
		Qual(info.IFPath, info.PublicName).
		Block()

	fmt.Printf("%#v", f)
}

func createMock(info CodeInfo) {
	f := jen.NewFile(info.PkMockName)

	f.Type().Id("Mock" + info.PublicName).Struct(
		jen.Qual("github.com/stretchr/testify/mock", "Mock"),
	)

	fmt.Printf("%#v", f)
}
