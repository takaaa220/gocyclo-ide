import { ExtensionContext, Uri, window as Window, workspace } from "vscode";
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
} from "vscode-languageclient/node";

let client: LanguageClient;

export async function activate(context: ExtensionContext) {
  const serverModule = Uri.joinPath(context.extensionUri, "lsp-server").fsPath;

  const serverOptions: ServerOptions = {
    run: { command: serverModule, transport: TransportKind.stdio },
    debug: { command: serverModule, transport: TransportKind.stdio },
  };

  const clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: "file", language: "go" }],
    synchronize: {
      fileEvents: [workspace.createFileSystemWatcher("**/*.go")],
    },
    outputChannel: Window.createOutputChannel("LSP Logs"),
    traceOutputChannel: Window.createOutputChannel("LSP Trace"),
  };

  try {
    client = new LanguageClient(
      "gocyclo-lsp-client",
      "Gocyclo LSP Client",
      serverOptions,
      clientOptions
    );
  } catch (error) {
    void Window.showErrorMessage(`Failed to create language client: ${error}`);
  }

  Window.showInformationMessage("Gocyclo LSP Client is now active!");

  client
    .start()
    .catch((error) => client.error(`Failed to start the server: ${error}`));
}

export async function deactivate() {
  await client?.stop();
}
