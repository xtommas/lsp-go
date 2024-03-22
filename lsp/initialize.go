package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ...
	// there's tons more parameters, but we won't cover them here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverinfo"`
}

type ServerCapabilities struct {
	TextDocumentSyncKind int            `json:"textDocumentSync"`
	HoverProvider        bool           `json:"hoverProvider"`
	DefinitionProvider   bool           `json:"definitionProvider"`
	CodeActionProvider   bool           `json:"codeActionProvider"`
	CompletionProvider   map[string]any `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSyncKind: 1,
				HoverProvider:        true,
				DefinitionProvider:   true,
				CodeActionProvider:   true,
				CompletionProvider:   map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "lsp_test",
				Version: "0.0.1",
			},
		},
	}
}
