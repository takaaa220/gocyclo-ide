package lsp

import (
	"log"

	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	// Must include a backend implementation
	// See CommonLog for other options: https://github.com/tliron/commonlog
	_ "github.com/tliron/commonlog/simple"
)

const lsName = "Language Server for Gocyclo"

var (
	version string = "0.0.1"
	handler protocol.Handler
)

// StartServer starts the LSP server and listens for client requests
func StartServer() error {
	// This increases logging verbosity (optional)
	commonlog.Configure(2, nil)

	handler := protocol.Handler{
		Initialize:          initialize,
		Initialized:         initialized,
		Shutdown:            shutdown,
		SetTrace:            setTrace,
		TextDocumentDidSave: didSave,
	}

	s := server.NewServer(&handler, lsName, false)
	return s.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	commonlog.NewInfoMessage(0, "Initializing server...")

	log.Println("Initializing server...")

	capabilities := handler.CreateServerCapabilities()

	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		Save: true,
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func didSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	commonlog.NewInfoMessage(0, "Saved!")
	log.Println("Saved!", params.TextDocument.URI)

	return nil
}
