package main

import (
	"flag"
	"fmt"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var User string
var Whitelist map[string]bool

func main() {
	var user_param = flag.String("user", "", "username to search for")
	var whitelist_param = flag.String("whitelist", "supervisord", "comma separated list of allowed process names")
	count := 0
	flag.Parse()

	if *user_param == "" {
		user_struct, _ := user.Current()
		User = user_struct.Name
	} else {
		User = *user_param
	}

	Whitelist = make(map[string]bool)
	for _, k := range strings.Split(*whitelist_param, ",") {
		Whitelist[k] = true
	}

	proclist, _ := process.Processes()
	for _, x := range proclist {
		puser, _ := x.Username()
		if puser == User {
			process_name, _ := x.Name()
			ppid, _ := x.Ppid()
			if _, ok := Whitelist[process_name]; !ok && ppid == 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
