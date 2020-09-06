package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func Touch(client *hdfs.Client, path string) {
	err := client.CreateEmptyFile(path)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
