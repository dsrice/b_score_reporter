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

func (info *CodeInfo) SetPackageInfo(command string) {
	switch command {
	case "controller":
		info.PkIFName = "handlers"
		info.PkCTName = "controllers"
		info.IFPath = "app/controllers/handlers"
		info.CTPath = "app/controllers"
		info.MockFlg = false
	case "usecase":
		info.PkIFName = "interfaces"
		info.PkCTName = "usecases"
		info.IFPath = "app/usecases/interfaces"
		info.CTPath = "app/usecases"
		info.PkMockName = "mockusecases"
		info.MockPath = "app/usecases/mockusecases"
		info.MockFlg = true
	case "repository":
		info.PkIFName = "adapters"
		info.PkCTName = "repositories"
		info.IFPath = "app/repositories/adapters"
		info.CTPath = "app/repositories"
		info.PkMockName = "mockrepositories"
		info.MockPath = "app/repositories/mockrepositories"
		info.MockFlg = true
	}
}

func (info CodeInfo) CreateCode() {
	createInterface(info)
	createContainer(info)

	if info.MockFlg {
		createMock(info)
	}
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
