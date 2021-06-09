package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func killPidServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		//log.Printf("ERROR: %+v", err)
		return errors.Wrap(err, "Exception: can't open the file or doesn't exist!")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning can't remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		log.Printf("ERROR: while converting pid int to string: %+v", err)
		return errors.Wrap(err, "Exception : can't convert PID!")
	}

	fmt.Printf("Killing server with PID -> %d\n", pid)
	return nil
}

func setupLoggin() {

	out, err := os.OpenFile("server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer out.Close()
	log.SetOutput(out)
}

func main() {

	if err := killPidServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Fatal(err)
	}

}
