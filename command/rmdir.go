package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func RmDir(client *hdfs.Client, dir string) {
	if err := client.Remove(dir);err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}