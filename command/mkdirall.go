package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func MkDirAll(client *hdfs.Client, dir string) {
	err := client.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
