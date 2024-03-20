package main

import (
	"bufio"
	"log"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("/mnt/c/Users/tomas/Documents/Proyectos/Go/lsp/log.txt")
	logger.Println("Logger started!")
	// we comunicate to the client through st io
	// so we need to listen for the incoming messages
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("that was not a good file :(")
	}

	return log.New(logfile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
