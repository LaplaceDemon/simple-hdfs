package command

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func Less(client *hdfs.Client, path string, size int) {
	reader, err := client.Open(path)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	bs := make([]byte, size, size)
	len, err := reader.Read(bs)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Println(string(bs[0:len]))
}
