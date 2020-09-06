package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/colinmarc/hdfs/v2"
	"simple-hdfs/command"
)

func flagInit() (h bool, nn0, nn1, user, cmd string, args []string) {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&nn0, "nn0", "", "hdfs name node 0")
	flag.StringVar(&nn1, "nn1", "", "hdfs name node 1")
	flag.StringVar(&user, "user", "", "test configuration and exit")
	flag.StringVar(&cmd, "cmd", "", "command")
	flag.Parse()
	args = flag.Args()
	return
}

func main() {
	_, nn0, nn1, user, cmd, args := flagInit()

	options := hdfs.ClientOptions{
		Addresses: []string{nn0, nn1},
		User:      user,
	}

	client, err := hdfs.NewClient(options)
	if err != nil {
		fmt.Printf("%+w\n", err)
		return
	}
	defer client.Close()

	cmd = args[0]
	switch cmd {
	case "ls":
		command.Ls(client, args[1])
	case "ll":
		command.Ll(client, args[1])
	case "mkdir":
		command.MkDir(client, args[1])
	case "mkdirall":
		command.MkDirAll(client, args[1])
	case "touch":
		command.Touch(client, args[1])
	case "rm":
		if !checkCmd(fmt.Sprintf("Are you sure you want to remove this file?\n[rm %s]\n\nDo you want to continue(Y/N)?\n", args[1])) {
			return
		}
		command.Rm(client, args[1])
	case "rmdir":
		if !checkCmd(fmt.Sprintf("Are you sure you want to remove this directory?\n[rmdir %s]\n\nDo you want to continue(Y/N)?\n", args[1])) {
			return
		}
		command.RmDir(client, args[1])
	case "rmdirall":
		if !checkCmd(fmt.Sprintf("Are you sure you want to remove this directory and all files in it?\n[rmdirall %s]\n\nDo you want to continue(Y/N)?\n", args[1])) {
			return
		}
		command.RmDirAll(client, args[1])
	case "rename":
		command.Rename(client, args[1], args[2])
	case "less":
		size := 32 * 1024
		if len(args) >= 3 {
			size, err = strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("error:%s\n", err.Error())
				return
			}
		}
		command.Less(client, args[1], size)
	case "readAt":
		offsetInt := 0
		size := 32 * 1024
		if len(args) >= 3 {
			offsetInt, err = strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("error:%s\n", err.Error())
				return
			}

			size, err = strconv.Atoi(args[3])
			if err != nil {
				fmt.Printf("error:%s\n", err.Error())
				return
			}
		}
		command.ReadAt(client, args[1], int64(offsetInt), size)
	default:
		fmt.Println("unknow cmd")
		return
	}
}

func checkCmd(info string) bool {
	fmt.Printf(info)
	reader := bufio.NewReader(os.Stdin)
	yes, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error:%s\n", err.Error())
		return false
	}
	if yes != "Y\r\n" && yes != "y\r\n" && yes != "Y\n" && yes != "y\n" {
		return false
	}
	return true
}
