package args

import (
	"fmt"
	"os"
	"strconv"
)

var Op = 0 //1 = start 2 = stop
var Delay = -1
var Button = -1 // 0 = left, 1 = middle, 2 = right

var Buttons = map[int]string {
	0: "left",
	1: "center",
	2: "right",
}

func Parse() {
	for k := 0; k < len(os.Args); k++ {
		switch os.Args[k] {
		case "--start":
			Op = 1
		case "--stop":
			Op = 2
		case "--shutdown":
			Op = 3
		case "--delay":
			var err error
			if Delay, err = strconv.Atoi(os.Args[k+1]); err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
		case "--button":
			var err error
			if Button, err = strconv.Atoi(os.Args[k+1]); err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
		case "--help":
			fmt.Println("--start | --stop - start or stop clicking")
			fmt.Println("--shutdown - kill the server")
			fmt.Println("--delay")
			os.Exit(0)
		}
	}
}