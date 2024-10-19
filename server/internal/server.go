package internal

import (
	"fmt"
	"log"
	"net/url"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const lsName = "Language Server for Gocyclo"

var (
	version string = "0.0.1"
	handler protocol.Handler
)

func StartServer() error {

	handler := protocol.Handler{
		Initialize:                     initialize,
		Initialized:                    initialized,
		Shutdown:                       shutdown,
		SetTrace:                       setTrace,
		TextDocumentDidSave:            didSave,
		TextDocumentCodeLens:           codeLens,
		WorkspaceDidChangeWatchedFiles: didChangeWatchedFiles,
	}

	s := server.NewServer(&handler, lsName, false)
	return s.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	log.Println("Initializing Language Server")

	capabilities := handler.CreateServerCapabilities()

	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		Save: true,
	}
	capabilities.CodeLensProvider = &protocol.CodeLensOptions{}

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
	log.Println("Saved", params.TextDocument.URI)

	go context.Notify(protocol.ServerWorkspaceCodeLensRefresh, nil)

	return nil
}

func didChangeWatchedFiles(context *glsp.Context, params *protocol.DidChangeWatchedFilesParams) error {
	log.Println("Changed watched files")

	return nil
}

func codeLens(context *glsp.Context, params *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	log.Println("CodeLens requested")

	parsedURL, err := url.Parse(params.TextDocument.URI)
	if err != nil {
		log.Fatalf("Error parsing URI: %v", err)
		return nil, err
	}

	complexities := calculateFunctionComplexities(parsedURL.Path)
	lensList := make([]protocol.CodeLens, len(complexities))
	for i, complexity := range complexities {
		lensList[i] = protocol.CodeLens{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      uint32(complexity.Pos.Line - 1),
					Character: uint32(complexity.Pos.Column - 1),
				},
				End: protocol.Position{
					Line:      uint32(complexity.Pos.Line - 1),
					Character: uint32(complexity.Pos.Column - 1),
				},
			},
			Command: &protocol.Command{
				Title:     fmt.Sprintf("Cyclomatic Complexity: %d", complexity.Complexity),
				Command:   "",
				Arguments: nil,
			},
		}
	}

	return lensList, nil
}
