package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func Rename(client *hdfs.Client, oldpath, newpath string) {
	info, err := client.Stat(oldpath)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	if info.IsDir() {
		fmt.Printf("path is a dir\n")
		return
	}

	if err := client.Rename(oldpath, newpath); err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
