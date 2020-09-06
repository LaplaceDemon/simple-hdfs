package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func RmDirAll(client *hdfs.Client, dir string) {
	if err := client.RemoveAll(dir);err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
