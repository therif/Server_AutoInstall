package install

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"autoinstaller/theStd"
	"autoinstaller/theStd/confStd"

	shellops "github.com/therif/gotherif-shellops"
)

func Start(skipexec bool) {

	defer fmt.Println("\nInstall Process End!")

	if skipexec {
		fmt.Println("\nSkip Execution, will show command only!")
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)

	//get list config available
	var listmenunya []string
	listMenuAvEx, err := theStd.GetAllFilenames(".the", exPath, false, false, false)
	if err != nil {
		log.Fatalf("No external config files: %s", err)
	} else {
		for _, ve := range listMenuAvEx {
			listmenunya = append(listmenunya, strings.ReplaceAll(ve, "_", " ")+" (EXT)")
		}
	}

	listMenuAv, err := theStd.GetAllFilenamesEmbeded(&theStd.FEmbedFs, "conf/main", false, false, false)
	if err != nil {
		log.Fatalf("No external config files: %s", err)
	} else {
		for _, vi := range listMenuAv {
			listmenunya = append(listmenunya, strings.ReplaceAll(vi, "_", " "))
		}
	}

	fmt.Println("")
	fmt.Println("-------- AUTO INSTALLER --------")

	for k, mn := range listmenunya {
		fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToUpper(mn))
	}

	fmt.Println("--")
	fmt.Println("0. EXIT")
	fmt.Print("Fill according to choice [0-" + strconv.Itoa(len(listmenunya)) + "] : ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	jawaban, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Fatalf("Answer error : %s", err)
	} else {

		if jawaban == 0 {
			//exit
		} else if jawaban > 0 && jawaban < len(listmenunya) {
			sval := listmenunya[jawaban-1]
			blExt := strings.Contains(sval, " (EXT)")
			sval = strings.Replace(sval, " (EXT)", "", 1)
			sval = strings.ReplaceAll(sval, " ", "_")
			fmt.Println("Processing " + sval)

			if blExt {
				install(exPath+"/"+sval+".the", true, skipexec)
			} else {
				install("conf/main/"+sval+".the", false, skipexec)
			}

		} else {
			fmt.Println("Not a option! ")
			//Start()
		}

	}

}

func install(targetCfgFile string, isRealFile bool, skipexec bool) {

	if skipexec {
		fmt.Println("\nSkip Execution, Emulation will show command only!")
	}

	fmt.Println("loading configs.....")

	if isRealFile {
		confStd.ConfigsKu.ReadConfig(targetCfgFile)
	} else {
		confStd.ConfigsKu.ReadFsConfig(targetCfgFile)
	}

	fmt.Println("")
	fmt.Println("")

	if confStd.ConfigsKu.Name == "" {
		fmt.Println("Error, No Configuration File!")
		os.Exit(1)
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)
	dtLogoAda := false
	var dtLogo []byte

	if _, err := os.Stat(exPath + "/logo.conf"); err == nil {
		if dtLogo, err = os.ReadFile(exPath + "/logo.conf"); err == nil {
			dtLogoAda = true
			fmt.Println(string(dtLogo))
		}
	}

	fmt.Println("")
	fmt.Println("Config Name  : " + confStd.ConfigsKu.Name)
	fmt.Println("Config Ver   : " + confStd.ConfigsKu.Version)
	fmt.Println("Config by    : " + confStd.ConfigsKu.By)
	fmt.Println("--this config design for :")
	fmt.Println("OS    : " + confStd.ConfigsKu.Os.Name)
	fmt.Println("Ver   : " + confStd.ConfigsKu.Os.Ver)
	fmt.Println("Type  : " + confStd.ConfigsKu.Os.Type)
	fmt.Println("Arch  : " + confStd.ConfigsKu.Os.Arch)

	//system package manager
	fmt.Println("")
	fmt.Print(" Check Update Operating System ? [Y/N] : ")
	input00 := bufio.NewScanner(os.Stdin)
	input00.Scan()
	jawaban00 := input00.Text()
	if jawaban00 == "y" || jawaban00 == "Y" {
		fmt.Println("")
		fmt.Println("=============")
		log.Println("-- Check System Update --")
		shellops.AsyncCmdBashSudo(confStd.ConfigsKu.PkgMan.Update, skipexec)

		if dtLogoAda {
			fmt.Println(string(dtLogo))
		}
	}

	fmt.Println("")
	fmt.Println(" Upgrade Operating System ? [Y/N]")
	input0 := bufio.NewScanner(os.Stdin)
	input0.Scan()
	jawaban0 := input0.Text()
	if jawaban0 == "y" || jawaban0 == "Y" {
		fmt.Println("")
		fmt.Println("=============")
		log.Println("-- Upgrade " + confStd.ConfigsKu.Name + " System to Latest --")

		shellops.AsyncCmdBashSudo(confStd.ConfigsKu.PkgMan.UpgradeList, skipexec)
		shellops.AsyncCmdBashSudo(confStd.ConfigsKu.PkgMan.Upgrade, skipexec)
	}

	//Install OS / PKG Redefault requirement
	if confStd.ConfigsKu.PkgReqInstall.Install != "" {
		fmt.Println("")
		fmt.Println("")
		fmt.Print(" Install Requirement Package ? [Y/N] : ")
		input1 := bufio.NewScanner(os.Stdin)
		input1.Scan()
		jawaban1 := input1.Text()
		if jawaban1 == "y" || jawaban1 == "Y" {
			fmt.Println("")
			log.Println("-- Install Requirement Packages --")

			shellops.AsyncCmdBashSudo(confStd.ConfigsKu.PkgReqInstall.Install, skipexec)
		}
	}

	//proses install
	for _, pkgnya := range confStd.ConfigsKu.Pkg {
		if pkgnya.Name != "" && pkgnya.Install != "" {

			fmt.Println("")
			fmt.Println("")
			fmt.Println("======== INSTALL " + strings.ToUpper(pkgnya.Name) + " ? ========")
			fmt.Print("Fill according to choice : [Y/N] ")
			input99 := bufio.NewScanner(os.Stdin)
			input99.Scan()
			jawaban99 := input99.Text()
			if jawaban99 == "y" || jawaban99 == "Y" {
				fmt.Println("")
				fmt.Println("=============")
				log.Println("-- Install " + pkgnya.Name + " --")

				shellops.AsyncCmdBashSudo(pkgnya.Install, skipexec)

				if pkgnya.InstallReq != "" {
					log.Println("-- Install " + pkgnya.Name + " Requirement --")
					shellops.AsyncCmdBashSudo(pkgnya.InstallReq, skipexec)
				}

				if pkgnya.Service.Restart != "" {

				}
			}
		}

	}

	if confStd.ConfigsKu.CustomInstall != nil && len(confStd.ConfigsKu.CustomInstall) > 0 {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("======== RUN OTHER CUSTOM INSTALLATION ? ========")
		fmt.Print("Fill according to choice [Y/N] : ")
		inputZ := bufio.NewScanner(os.Stdin)
		inputZ.Scan()
		jawabanZ := inputZ.Text()
		if jawabanZ == "y" || jawabanZ == "Y" {
			fmt.Println("")

			iJmlProses := len(confStd.ConfigsKu.CustomInstall)

			if iJmlProses > 0 {
				var lastinputnya string
				iProses := 0

				for _, prosesnya := range confStd.ConfigsKu.CustomInstall {
					fmt.Println("Processing Custom Install [" + strconv.Itoa(iProses+1) + "/" + strconv.Itoa(iJmlProses) + "] ")

					if prosesnya.Act != "" {

						if strings.ToLower(prosesnya.Act) == "cmd" {
							if strings.ToLower(prosesnya.Msg) != "" {
								if strings.ToLower(prosesnya.TextPre) != "" {
									newTextPre := strings.Replace(prosesnya.TextPre, "${input}", lastinputnya, -1)
									fmt.Println(newTextPre)
								}

								newMsg := strings.Replace(prosesnya.Msg, "${input}", lastinputnya, -1)
								if prosesnya.Sudo {
									shellops.AsyncCmdBashSudo(newMsg, skipexec)
								} else {
									shellops.AsyncCmdBashSudo(newMsg, skipexec)
								}

								if strings.ToLower(prosesnya.TextAfter) != "" {
									newTextAfter := strings.Replace(prosesnya.TextAfter, "${input}", lastinputnya, -1)
									fmt.Println(newTextAfter)
								}
							}

						} else if strings.ToLower(prosesnya.Act) == "input" {
							fmt.Println("")
							if strings.ToLower(prosesnya.TextPre) != "" {
								newTextPre := strings.Replace(prosesnya.TextPre, "${input}", lastinputnya, -1)
								fmt.Println(newTextPre)
							}

							input901 := bufio.NewScanner(os.Stdin)
							input901.Scan()
							lastinputnya = input901.Text()

							if strings.ToLower(prosesnya.TextAfter) != "" {
								newTextAfter := strings.Replace(prosesnya.TextAfter, "${input}", lastinputnya, -1)
								fmt.Println(newTextAfter)
							}

						} else if strings.ToLower(prosesnya.Act) == "output" {
							if strings.ToLower(prosesnya.Msg) != "" {
								newMsg := strings.Replace(prosesnya.Msg, "${input}", lastinputnya, -1)
								fmt.Println(newMsg)
							}

						} else if strings.ToLower(prosesnya.Act) == "replacefile" {
							if prosesnya.Dest != "" && prosesnya.Src != "" {
								if prosesnya.Sudo {
									theStd.TulisContentFile(prosesnya.Dest, prosesnya.Src, prosesnya.Chmod, prosesnya.Backup, isRealFile, skipexec)
								} else {
									theStd.TulisContentFile(prosesnya.Dest, prosesnya.Src, prosesnya.Chmod, prosesnya.Backup, isRealFile, skipexec)
								}
							}

						}
					}
					iProses++
				}
			}

		}
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("-- Completed !!! --")
	fmt.Println("")
	fmt.Println("")

}
