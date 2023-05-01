package generator

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/labstack/gommon/log"
)

type CodeInfo struct {
	PublicName   string
	PrivateName  string
	PkIFName     string // インターフェイス側パッケージ名称
	PkCTName     string // コンテナ側パッケージ名称
	PkMockName   string // モック側パッケージ名称
	IFPath       string
	CTPath       string
	MockPath     string
	IFName       string
	CTName       string
	MockName     string
	IFFileName   string
	CTFileName   string
	MockFileName string
	MockFlg      bool
}

func (info *CodeInfo) SetPackageInfo(command string) {
	switch command {
	case "controller":
		info.PkIFName = "handlers"
		info.PkCTName = "controllers"
		info.IFPath = "login/controllers/handlers"
		info.CTPath = "login/controllers"
		info.IFName = info.PublicName + "Handler"
		info.CTName = info.PrivateName + "Controller"
		info.IFFileName = info.PublicName + "Handler.go"
		info.CTFileName = info.PublicName + "Controller.go"
		info.MockFlg = false
	case "usecase":
		info.PkIFName = "interfaces"
		info.PkCTName = "usecases"
		info.IFPath = "login/usecases/interfaces"
		info.CTPath = "login/usecases"
		info.PkMockName = "mockusecases"
		info.MockPath = "login/usecases/mockusecases"
		info.IFName = info.PublicName + "Interface"
		info.CTName = info.PrivateName + "UseCase"
		info.MockName = "Mock" + info.PublicName + "UseCase"
		info.IFFileName = info.PublicName + "Interface.go"
		info.CTFileName = info.PublicName + "UseCase.go"
		info.MockFileName = "Mock" + info.PublicName + "UseCase.go"
		info.MockFlg = true
	case "repository":
		info.PkIFName = "adapters"
		info.PkCTName = "repositories"
		info.IFPath = "login/repositories/adapters"
		info.CTPath = "login/repositories"
		info.PkMockName = "mockrepositories"
		info.MockPath = "login/repositories/mockrepositories"
		info.IFName = info.PublicName + "Adapter"
		info.CTName = info.PrivateName + "Repository"
		info.MockName = "Mock" + info.PublicName + "Repository"
		info.IFFileName = info.PublicName + "Adapter.go"
		info.CTFileName = info.PublicName + "Repository.go"
		info.MockFileName = "Mock" + info.PublicName + "Repository.go"
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

	f.Type().Id(info.IFName).Interface()

	err := f.Save("../" + info.IFPath + "/" + info.IFFileName)
	if err != nil {
		log.Error(err.Error())
	}
}

func createContainer(info CodeInfo) {
	f := jen.NewFile(info.PkCTName)

	f.ImportName(info.IFPath, info.PkIFName)

	f.Type().Id(info.CTName).Struct()

	f.Func().
		Id(fmt.Sprintf("New"+info.PublicName)).
		Params().
		Qual(info.IFPath, info.IFName).
		Block(
			jen.Return(jen.Op("&").Id(info.CTName).Block()))

	err := f.Save("../" + info.CTPath + "/" + info.CTFileName)
	if err != nil {
		log.Error(err.Error())
	}
}

func createMock(info CodeInfo) {
	f := jen.NewFile(info.PkMockName)

	f.ImportName("github.com/stretchr/testify/mock", "mock")

	f.Type().Id(info.MockName).Struct(
		jen.Qual("github.com/stretchr/testify/mock", "Mock"),
	)

	err := f.Save("../" + info.MockPath + "/" + info.MockFileName)
	if err != nil {
		log.Error(err.Error())
	}
}
