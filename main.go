package main

import (
	"fmt"
	"fuckclicking/args"
	"github.com/go-vgo/robotgo"
	"net"
	"net/rpc"
	"os"
	"sync"
	"time"
)

var spamming = false
var delay = time.Millisecond * 25 //ms

type Request struct {
	Op int
	Delay int
	Button int
}
type Response struct {}
type Handler struct {}

func (h *Handler) Execute(req Request, res *Response) (err error) {
	switch req.Op {
	case 1:
		if !spamming {
			spamming = true
			go Spam()
		}
	case 2:
		spamming = false
	case 3:
		os.Exit(0)
	}

	if req.Button != -1 {
		args.Button = req.Button
	}

	if req.Delay != -1 {
		args.Delay = req.Delay
		delay = time.Millisecond * time.Duration(args.Delay)
	}

	return nil
}

func main() {
	args.Parse()

	if client, err := rpc.Dial("tcp", "127.0.0.1:2938"); err == nil {
		//be the client
		_ = client.Call("Handler.Execute", &Request{args.Op, args.Delay, args.Button}, &Response{})
		return
	} else {
		//become the server

		_ = rpc.Register(&Handler{})
		var listener net.Listener
		if listener, err = net.Listen("tcp","127.0.0.1:2938"); err != nil {
			fmt.Println(err.Error())
			return
		}

		if args.Button == -1 {
			args.Button = 1
		}

		if args.Delay == -1 {
			args.Delay = 25
		}

		if args.Op == 1 {
			spamming = true
			go Spam()
		} else if args.Op == 3 {
			os.Exit(0)
		}

		delay = time.Millisecond * time.Duration(args.Delay)

		go func() {
			for {
				rpc.Accept(listener)
			}
		}()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func Spam() {
	for spamming {
		time.Sleep(delay)
		fmt.Println("clicking")
		robotgo.Click(args.Buttons[args.Button], false)
	}
}