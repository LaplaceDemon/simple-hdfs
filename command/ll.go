package command

import (
	"fmt"
	"os"

	"github.com/colinmarc/hdfs/v2"
)

func Ll(client *hdfs.Client, dir string) {
	dirInfo, err := client.Stat(dir)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	if !dirInfo.IsDir() {
		printFileInfoForLL(dirInfo)
		return
	}

	infos, err := client.ReadDir(dir)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	if infos == nil {
		fmt.Printf("infos is nil\n")
		return
	}
	for _, info := range infos {
		printFileInfoForLL(info)
	}
}

func printFileInfoForLL(info os.FileInfo) {
	lab := ""
	if !info.IsDir() {
		lab = "File"
	} else {
		lab = "Dir"
	}

	status := info.Sys().(*hdfs.FileStatus)
	owner := status.GetOwner()
	group := status.GetGroup()

	fmt.Printf("%-4s%+32s    %10.3f MB  %+10s  %+8s  %s\n", lab, info.Name(), float64(info.Size())/(1024.0*1024), owner, group, info.Mode())
}

