package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var ap string
var dir string = "./"

func main() {
	username := flag.String("cd", "", "file")
	flag.Parse()
	temp := strings.TrimSuffix(*username, ".")
	fmt.Println(temp)
	if len(temp) > 0 {
		if len(path.Ext(temp)) > 0 {
			ap = temp
		} else {
			dir = temp
			ap = getLastApkname()
		}
	} else {
		args := os.Args
		if len(args) > 1 {
			a := args[1]
			if len(path.Ext(a)) > 0 {
				ap = a
			} else {
				dir = a
				ap = getLastApkname()
			}
		} else {
			ap = getLastApkname()
		}
	}

	command := "adb"
	params := []string{"install", "-r", ap}

	execCommand(command, params)
}

func execCommand(commandName string, params []string) {
	cmd := exec.Command(commandName, params...)
	fmt.Printf("exec: %s\n", strings.Join(cmd.Args[1:], " "))
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	fmt.Println(string(stdout))
	time.Sleep(2 * time.Second)
	cmd.Wait()
}

func getLastApkname() string {
	var temp int64
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().Unix() >= temp && ".apk" == path.Ext(info.Name()) {
			ap = p
		}
		return nil
	})
	return ap
}
