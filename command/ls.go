package command

import (
	"fmt"
	"os"

	"github.com/colinmarc/hdfs/v2"
)

func Ls(client *hdfs.Client, dir string) {
	dirInfo, err := client.Stat(dir)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	if !dirInfo.IsDir() {
		printFileInfoForLs(dirInfo)
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
		printFileInfoForLs(info)
	}
}

func printFileInfoForLs(info os.FileInfo) {
	fmt.Printf("%s\n", info.Name())
}