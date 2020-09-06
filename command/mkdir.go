package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func MkDir(client *hdfs.Client, dir string) {
	err := client.Mkdir(dir, 0755)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}