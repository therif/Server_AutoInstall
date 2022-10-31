package confStd

import (
	"autoinstaller/theStd"
	"encoding/json"
	"log"
	"os"
)

type theCfgService struct {
	Stop    string
	Start   string
	Restart string
	Status  string
}

type thePkgMan struct {
	Name        string
	Update      string
	Upgrade     string
	UpgradeList string
}

type thePkg struct {
	Name        string
	Install     string
	Conf        string
	ConfDefault string
	DataFolder  string
	InstallReq  string
	Service     theCfgService
	Configure   []theConfProses `json:"configure"`
}

type theReq struct {
	Install string
}

type theConfProses struct {
	Act       string
	Msg       string
	TextPre   string
	TextAfter string
	Sudo      bool

	Dest   string
	Src    string
	Chmod  string
	Backup bool
}

type theOs struct {
	Name   string
	Ver    string
	Type   string
	Arch   string
	Distro string
}

type the struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	By      string `json:"by"`

	Os            theOs
	PkgMan        thePkgMan       `json:"pkg_manager"`
	PkgReqInstall theReq          `json:"pkgreqinstall"`
	Pkg           []thePkg        `json:"pkg"`
	CustomInstall []theConfProses `json:"custominstallpkg"`
}

var ConfigsKu the

func (conf *the) ReadFsConfig(filefsnya string) *the {

	jsonCfgFile, err := theStd.FEmbedFs.ReadFile(filefsnya)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(jsonCfgFile, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

func (conf *the) ReadConfig(filenya string) *the {

	jsonCfgFile, err := os.ReadFile(filenya)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(jsonCfgFile, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
