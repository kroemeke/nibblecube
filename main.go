package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"bytes"
)

/* map of return http codes, where key is the full string
   taken from first line of servers response */
var returnMap = make(map[string]int)
var statsMap = make(map[string]int)
var mutex = &sync.Mutex{}

func showStatistics(delay time.Duration) {
	var oldTotal int
	var delta int
	for {
		print("\033[H\033[2J")
		var newTotal int
		for _, v := range returnMap {
			newTotal = newTotal + v
		}
		delta = newTotal - oldTotal
		oldTotal = newTotal
		mutex.Lock()
		fmt.Printf("Total: %v Errors: %v Rate: %v /s\n", newTotal, statsMap["ERROR"], delta)
		mutex.Unlock()
		showHistogram()
		//fmt.Println(randomPot.multipleChar)
		time.Sleep(1000*time.Millisecond)
	}
}

/* Regenerate the struct that holds random stuff, done every X miliseconds so that 
   we don't have to burn time on this each time */
func regenerateRandomPot(delay time.Duration) {
	for {
		time.Sleep(delay)
		rubik()
	}
}


/* Walk through the histogram of response codes */
func showHistogram() {
	for k, v := range returnMap {
		fmt.Printf("%v \t%v\n", v, k)
	}
}


func getWorker(id int, ip string, port int, host string, socket string) {
	fmt.Printf("%d started\n", id)
	x := strconv.Itoa(port)
	connectionString := ip + ":" + x
	connectionType := ""
	if (socket != "") {
	  connectionType = "unix"
	} else {
	  connectionType = "tcp"
	}
	for {
		get(connectionType, connectionString, host)
	}
}

func get(connectionType string, connectionString string, host string) {
	conn, err := net.Dial(connectionType, connectionString)
	if err != nil {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		statsMap["ERROR"]++
		mutex.Unlock()
	} else {
		var buffer bytes.Buffer
		buffer.WriteString("GET / HTTP/1.1\r\nHost: ")
		buffer.WriteString(host)
		buffer.WriteString("\n")
		buffer.WriteString(http())
		buffer.WriteString("\r\n\r\n")
		conn.Write(buffer.Bytes())
		conn.SetReadDeadline(time.Now().Add(100*time.Millisecond))
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			time.Sleep(time.Second)
			mutex.Lock()
			statsMap["ERROR1"]++
			//statsMap[string(err)]++
			mutex.Unlock()
		} else {
		        status = strings.TrimRight(status, "\r\n")
			mutex.Lock()
			returnMap[status]++
			mutex.Unlock()
		}
	}
}



func main() {
	// Argument parsing
	workers := flag.Int("workers", 64, "number of concurrent workers")
	statsinterval := flag.Duration("statsinterval", 1*time.Second, "seconds between statistics update")
	host := flag.String("host", "localhost", "Host: header to be sent")
	ip := flag.String("ip", "", "ip address")
	socket := flag.String("socket", "", "unix socket path")
	port := flag.Int("port", 80, "port")
	workerdelay := 10*time.Millisecond
	flag.Parse()

	if (*ip != "" ) || (*socket != "" ) {
		for i := 0; i <= *workers; i++ {
			time.Sleep(workerdelay)
			go getWorker(i, *ip, *port, *host, *socket)
		}
		go regenerateRandomPot(50*time.Millisecond)
		go showStatistics(*statsinterval)
		for {
			time.Sleep(60 * time.Second)
		}
	} else {
		// ERROR
		flag.Usage()
		fmt.Printf("\n(!) You need to specify either ip or unix socket.\n")
		os.Exit(1)
	}
}
