package theStd

import (
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	shellops "github.com/therif/gotherif-shellops"
)

//go:embed conf/*
var FEmbedFs embed.FS

func exists(pathstr string) (bool, error) {
	_, err := os.Stat(pathstr)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAllFilenames(filterext string, dirnya string, recursive bool, showpath bool, showext bool) (out []string, err error) {
	if len(dirnya) == 0 {
		dirnya = "."
	}

	entries, err := ioutil.ReadDir(dirnya)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var fp string
		if filterext != "" {
			if filterext == path.Ext(entry.Name()) {
				if showpath {
					fp = path.Join(dirnya, entry.Name())
				} else {
					fp = entry.Name()
					if !showext {
						fp = fp[:len(entry.Name())-len(path.Ext(entry.Name()))]
					}
				}
			}
		} else {
			if showpath {
				fp = path.Join(dirnya, entry.Name())
			} else {
				fp = entry.Name()
				if !showext {
					fp = fp[:len(entry.Name())-len(path.Ext(entry.Name()))]
				}
			}
		}

		if recursive && entry.IsDir() {
			res, err := GetAllFilenames(filterext, fp, recursive, showpath, showext)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		if fp != "" {
			out = append(out, fp)
		}
	}
	return
}

func GetAllFilenamesEmbeded(fsnya *embed.FS, dirnya string, recursive bool, showpath bool, showext bool) (out []string, err error) {
	if len(dirnya) == 0 {
		dirnya = "."
	}

	entries, err := fsnya.ReadDir(dirnya)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var fp string
		if showpath {
			fp = path.Join(dirnya, entry.Name())
		} else {
			fp = entry.Name()
			if !showext {
				fp = fp[:len(entry.Name())-len(path.Ext(entry.Name()))]
				//fileName[:len(fileName)-len(filepath.Ext(fileName))]
			}
		}

		if recursive && entry.IsDir() {
			res, err := GetAllFilenamesEmbeded(fsnya, fp, recursive, showpath, showext)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}

func TulisContentFile(targetfilenya string, srcfile string, chmodfile string, backuptarget bool, isRealFile bool, skipexec bool) {

	stopError := false
	if skipexec {
		if backuptarget {
			fmt.Println("rm -f " + targetfilenya + ".back")
			fmt.Println("cp -f " + targetfilenya + " " + targetfilenya + ".back")
		}

		if _, err := os.Stat(targetfilenya); err == nil {
			fmt.Println("chmod 777 " + targetfilenya)
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Println("touch " + targetfilenya)
			fmt.Println("chmod 777 " + targetfilenya)
		} else {
			stopError = true
		}

		if !stopError {
			fmt.Println("Emulate Opening File as RW " + targetfilenya)
			fmt.Println("Here Emulate Writing to file...")
			fmt.Printf("\nLength: %d bytes", 0)
			fmt.Printf("\nFile Name: %s", targetfilenya)
			fmt.Println("")

			if chmodfile != "" {
				fmt.Println("chmod " + chmodfile + "" + targetfilenya)
			} else {
				fmt.Println("chmod 644 " + targetfilenya)
			}
		}

	} else {

		if backuptarget {
			shellops.AsyncCmdBashSudo("rm -f "+targetfilenya+".back", skipexec)
			shellops.AsyncCmdBashSudo("cp -f "+targetfilenya+" "+targetfilenya+".back", skipexec)
		}

		if _, err := os.Stat(targetfilenya); err == nil {
			// path/to/whatever exists
			shellops.AsyncCmdBashSudo("chmod 777 "+targetfilenya, skipexec)

		} else if errors.Is(err, os.ErrNotExist) {
			shellops.AsyncCmdBashSudo("touch "+targetfilenya, skipexec)
			shellops.AsyncCmdBashSudo("chmod 777 "+targetfilenya, skipexec)
		} else {
			stopError = true
		}

		if !stopError {
			//file, err := os.OpenFile(targetfilenya, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			file, err := os.OpenFile(targetfilenya, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
				TulisContentFile(targetfilenya, srcfile, chmodfile, false, isRealFile, skipexec)
			} else {
				defer file.Close()

				var data []byte
				if isRealFile {
					ex, err := os.Executable()
					if err != nil {
						panic(err)
					}
					exPath := path.Dir(ex)

					if _, err := os.Stat(exPath + "/" + srcfile); err == nil {
						data, err = os.ReadFile(exPath + "/" + srcfile)
						if err != nil {
							stopError = true
							log.Fatalf("failed readfile: %s", err)
						}
					} else {
						stopError = true
						log.Fatalf("failed file: %s", err)
					}

				} else {
					data, err = FEmbedFs.ReadFile(srcfile)
					if err != nil {
						stopError = true
						log.Fatalf("failed readfile: %s", err)
					}
				}

				err = file.Truncate(0)
				//len, err := file.WriteString(string(data))
				len, err := file.WriteAt(data, 0) // Write at 0 beginning
				if err != nil {
					log.Fatalf("failed writing to file: %s", err)
				}

				fmt.Printf("\nLength: %d bytes", len)
				fmt.Printf("\nFile Name: %s", file.Name())
				fmt.Println("")

				if chmodfile != "" {
					shellops.AsyncCmdBashSudo("chmod "+chmodfile+" "+targetfilenya, skipexec)
				} else {
					shellops.AsyncCmdBashSudo("chmod 644 "+targetfilenya, skipexec)
				}
			}
		}

	}
}
