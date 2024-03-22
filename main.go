package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lsp/analysis"
	"lsp/lsp"
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

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}

		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, couldn't parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// reply
		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}

		logger.Printf("Opened: %s", request.Params.TextDocument.URI)

		// Sync the state of the 'analysis engine' with the state of the editor
		// (mapping the file to the contents)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		logger.Printf("Changed: %s", request.Params.TextDocument.URI)

		for _, change := range request.Params.ContentChanges {
			// Sync the state of the 'analysis engine' with the state of the editor
			// (mapping the file to the contents)
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		// Create a response
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		// Write the response back
		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: %s", err)
			return
		}

		// Create a response
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		// Write the response back
		writeResponse(writer, response)

	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/codeAction: %s", err)
			return
		}

		// Create a response
		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)

		// Write the response back
		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("that was not a good file :(")
	}

	return log.New(logfile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
