package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func Rm(client *hdfs.Client, path string) {
	info, err := client.Stat(path)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	if info.IsDir() {
		fmt.Printf("path is a dir\n")
		return
	}

	if err := client.Remove(path);err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}