// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.19
// +build go1.19

package main

// prop combines the name of a property with the name of the structure it is in.
type prop [2]string

const (
	nothing = iota
	wantStar
	wantOpt
	wantOptStar
)

// goplsStar records the optionality of each field in the protocol.
var goplsStar = map[prop]int{
	{"ClientCapabilities", "textDocument"}:                       wantOpt,
	{"ClientCapabilities", "window"}:                             wantOpt,
	{"ClientCapabilities", "workspace"}:                          wantOpt,
	{"CodeAction", "edit"}:                                       wantOpt,
	{"CodeAction", "kind"}:                                       wantOpt,
	{"CodeActionClientCapabilities", "codeActionLiteralSupport"}: wantOpt,
	{"CodeActionContext", "triggerKind"}:                         wantOpt,
	{"CodeLens", "command"}:                                      wantOpt,
	{"CompletionClientCapabilities", "completionItem"}:           wantOpt,
	{"CompletionClientCapabilities", "insertTextMode"}:           wantOpt,
	{"CompletionItem", "insertTextFormat"}:                       wantOpt,
	{"CompletionItem", "insertTextMode"}:                         wantOpt,
	{"CompletionItem", "kind"}:                                   wantOpt,
	{"CompletionParams", "context"}:                              wantOpt,
	{"Diagnostic", "severity"}:                                   wantOpt,
	{"DidSaveTextDocumentParams", "text"}:                        wantOptStar,
	{"DocumentHighlight", "kind"}:                                wantOpt,
	{"FileOperationPattern", "matches"}:                          wantOpt,
	{"FileSystemWatcher", "kind"}:                                wantOpt,
	{"Hover", "range"}:                                           wantOpt,
	{"InitializeResult", "serverInfo"}:                           wantOpt,
	{"InlayHint", "kind"}:                                        wantOpt,
	{"InlayHint", "position"}:                                    wantStar,

	{"Lit_CompletionClientCapabilities_completionItem", "commitCharactersSupport"}: nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "deprecatedSupport"}:       nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "documentationFormat"}:     nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "insertReplaceSupport"}:    nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "insertTextModeSupport"}:   nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "labelDetailsSupport"}:     nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "preselectSupport"}:        nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "resolveSupport"}:          nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "snippetSupport"}:          nothing,
	{"Lit_CompletionClientCapabilities_completionItem", "tagSupport"}:              nothing,
	{"Lit_CompletionClientCapabilities_completionItemKind", "valueSet"}:            nothing,
	{"Lit_CompletionClientCapabilities_completionList", "itemDefaults"}:            nothing,
	{"Lit_CompletionList_itemDefaults", "commitCharacters"}:                        nothing,
	{"Lit_CompletionList_itemDefaults", "data"}:                                    nothing,
	{"Lit_CompletionList_itemDefaults", "editRange"}:                               nothing,
	{"Lit_CompletionList_itemDefaults", "insertTextFormat"}:                        nothing,
	{"Lit_CompletionList_itemDefaults", "insertTextMode"}:                          nothing,
	{"Lit_CompletionOptions_completionItem", "labelDetailsSupport"}:                nothing,
	{"Lit_DocumentSymbolClientCapabilities_symbolKind", "valueSet"}:                nothing,
	{"Lit_FoldingRangeClientCapabilities_foldingRange", "collapsedText"}:           nothing,
	{"Lit_FoldingRangeClientCapabilities_foldingRangeKind", "valueSet"}:            nothing,
	{"Lit_InitializeResult_serverInfo", "version"}:                                 nothing,
	{"Lit_NotebookDocumentChangeEvent_cells", "data"}:                              nothing,
	{"Lit_NotebookDocumentChangeEvent_cells", "structure"}:                         nothing,
	{"Lit_NotebookDocumentChangeEvent_cells", "textContent"}:                       nothing,
	{"Lit_NotebookDocumentChangeEvent_cells_structure", "didClose"}:                nothing,
	{"Lit_NotebookDocumentChangeEvent_cells_structure", "didOpen"}:                 nothing,
	{"Lit_NotebookDocumentFilter_Item0", "pattern"}:                                nothing,
	{"Lit_NotebookDocumentFilter_Item0", "scheme"}:                                 nothing,
	{"Lit_NotebookDocumentSyncOptions_notebookSelector_Elem_Item0", "cells"}:       nothing,
	{"Lit_SemanticTokensClientCapabilities_requests", "full"}:                      nothing,
	{"Lit_SemanticTokensClientCapabilities_requests", "range"}:                     nothing,
	{"Lit_SemanticTokensClientCapabilities_requests_full_Item1", "delta"}:          nothing,
	{"Lit_SemanticTokensOptions_full_Item1", "delta"}:                              nothing,
	{"Lit_ServerCapabilities_workspace", "fileOperations"}:                         nothing,
	{"Lit_ServerCapabilities_workspace", "workspaceFolders"}:                       nothing,

	{"Lit_ShowMessageRequestClientCapabilities_messageActionItem", "additionalPropertiesSupport"}:           nothing,
	{"Lit_SignatureHelpClientCapabilities_signatureInformation", "activeParameterSupport"}:                  nothing,
	{"Lit_SignatureHelpClientCapabilities_signatureInformation", "documentationFormat"}:                     nothing,
	{"Lit_SignatureHelpClientCapabilities_signatureInformation", "parameterInformation"}:                    nothing,
	{"Lit_SignatureHelpClientCapabilities_signatureInformation_parameterInformation", "labelOffsetSupport"}: nothing,

	{"Lit_TextDocumentContentChangeEvent_Item0", "range"}:       wantStar,
	{"Lit_TextDocumentContentChangeEvent_Item0", "rangeLength"}: nothing,
	{"Lit_TextDocumentFilter_Item0", "pattern"}:                 nothing,
	{"Lit_TextDocumentFilter_Item0", "scheme"}:                  nothing,
	{"Lit_TextDocumentFilter_Item1", "language"}:                nothing,
	{"Lit_TextDocumentFilter_Item1", "pattern"}:                 nothing,

	{"Lit_WorkspaceEditClientCapabilities_changeAnnotationSupport", "groupsOnLabel"}: nothing,
	{"Lit_WorkspaceSymbolClientCapabilities_symbolKind", "valueSet"}:                 nothing,
	{"Lit__InitializeParams_clientInfo", "version"}:                                  nothing,

	{"Moniker", "kind"}:                                       wantOpt,
	{"PartialResultParams", "partialResultToken"}:             wantOpt,
	{"ResourceOperation", "annotationId"}:                     wantOpt,
	{"ServerCapabilities", "completionProvider"}:              wantOpt,
	{"ServerCapabilities", "documentLinkProvider"}:            wantOpt,
	{"ServerCapabilities", "executeCommandProvider"}:          wantOpt,
	{"ServerCapabilities", "positionEncoding"}:                wantOpt,
	{"ServerCapabilities", "signatureHelpProvider"}:           wantOpt,
	{"ServerCapabilities", "workspace"}:                       wantOpt,
	{"TextDocumentClientCapabilities", "codeAction"}:          wantOpt,
	{"TextDocumentClientCapabilities", "completion"}:          wantOpt,
	{"TextDocumentClientCapabilities", "documentSymbol"}:      wantOpt,
	{"TextDocumentClientCapabilities", "foldingRange"}:        wantOpt,
	{"TextDocumentClientCapabilities", "hover"}:               wantOpt,
	{"TextDocumentClientCapabilities", "publishDiagnostics"}:  wantOpt,
	{"TextDocumentClientCapabilities", "rename"}:              wantOpt,
	{"TextDocumentClientCapabilities", "semanticTokens"}:      wantOpt,
	{"TextDocumentSyncOptions", "change"}:                     wantOpt,
	{"TextDocumentSyncOptions", "save"}:                       wantOpt,
	{"WorkDoneProgressParams", "workDoneToken"}:               wantOpt,
	{"WorkspaceClientCapabilities", "didChangeConfiguration"}: wantOpt,
	{"WorkspaceClientCapabilities", "didChangeWatchedFiles"}:  wantOpt,
	{"WorkspaceEditClientCapabilities", "failureHandling"}:    wantOpt,
	{"XInitializeParams", "clientInfo"}:                       wantOpt,
}

// keep track of which entries in goplsStar are used
var usedGoplsStar = make(map[prop]bool)

// For gopls compatibility, use a different, typically more restrictive, type for some fields.
var renameProp = map[prop]string{
	{"CancelParams", "id"}:         "interface{}",
	{"Command", "arguments"}:       "[]json.RawMessage",
	{"CompletionItem", "textEdit"}: "TextEdit",
	{"Diagnostic", "code"}:         "interface{}",

	{"DocumentDiagnosticReportPartialResult", "relatedDocuments"}: "map[DocumentURI]interface{}",

	{"ExecuteCommandParams", "arguments"}: "[]json.RawMessage",
	{"FoldingRange", "kind"}:              "string",
	{"Hover", "contents"}:                 "MarkupContent",
	{"InlayHint", "label"}:                "[]InlayHintLabelPart",

	// removing this causes the test in json_test.go to fail
	// First, the custom unmarshaler returns the 'wrong' error type
	// Second, cmp.Diff reports too many errors.
	// TODO(pjw): fix json_test.go and maybe the generated code
	{"Lit_SemanticTokensClientCapabilities_requests", "full"}:        "interface{}",
	{"Lit_SemanticTokensClientCapabilities_requests", "range"}:       "bool",
	{"NotebookCellTextDocumentFilter", "notebook"}:                   "NotebookDocumentFilter",
	{"RelatedFullDocumentDiagnosticReport", "relatedDocuments"}:      "map[DocumentURI]interface{}",
	{"RelatedUnchangedDocumentDiagnosticReport", "relatedDocuments"}: "map[DocumentURI]interface{}",

	// this one also has the json_test.go problem
	{"RenameClientCapabilities", "prepareSupportDefaultBehavior"}: "interface{}",

	{"SemanticTokensOptions", "full"}:                         "bool",
	{"SemanticTokensOptions", "range"}:                        "interface{}",
	{"ServerCapabilities", "callHierarchyProvider"}:           "interface{}",
	{"ServerCapabilities", "codeActionProvider"}:              "interface{}",
	{"ServerCapabilities", "colorProvider"}:                   "interface{}",
	{"ServerCapabilities", "declarationProvider"}:             "bool",
	{"ServerCapabilities", "definitionProvider"}:              "bool",
	{"ServerCapabilities", "diagnosticProvider"}:              "interface{}",
	{"ServerCapabilities", "documentFormattingProvider"}:      "bool",
	{"ServerCapabilities", "documentHighlightProvider"}:       "bool",
	{"ServerCapabilities", "documentRangeFormattingProvider"}: "bool",
	{"ServerCapabilities", "documentSymbolProvider"}:          "bool",
	{"ServerCapabilities", "foldingRangeProvider"}:            "interface{}",
	{"ServerCapabilities", "hoverProvider"}:                   "bool",
	{"ServerCapabilities", "implementationProvider"}:          "interface{}",
	{"ServerCapabilities", "inlayHintProvider"}:               "interface{}",
	{"ServerCapabilities", "inlineValueProvider"}:             "interface{}",
	{"ServerCapabilities", "linkedEditingRangeProvider"}:      "interface{}",
	{"ServerCapabilities", "monikerProvider"}:                 "interface{}",
	{"ServerCapabilities", "notebookDocumentSync"}:            "interface{}",
	{"ServerCapabilities", "referencesProvider"}:              "bool",
	{"ServerCapabilities", "renameProvider"}:                  "interface{}",
	{"ServerCapabilities", "selectionRangeProvider"}:          "interface{}",
	{"ServerCapabilities", "semanticTokensProvider"}:          "interface{}",
	{"ServerCapabilities", "textDocumentSync"}:                "interface{}",
	{"ServerCapabilities", "typeDefinitionProvider"}:          "interface{}",
	{"ServerCapabilities", "typeHierarchyProvider"}:           "interface{}",
	{"ServerCapabilities", "workspaceSymbolProvider"}:         "bool",
	{"TextDocumentEdit", "edits"}:                             "[]TextEdit",
	{"TextDocumentSyncOptions", "save"}:                       "SaveOptions",
	{"WorkspaceEdit", "documentChanges"}:                      "[]DocumentChanges",
}

// which entries of renameProp were used
var usedRenameProp = make(map[prop]bool)

type adjust struct {
	prefix, suffix string
}

// disambiguate specifies prefixes or suffixes to add to all values of
// some enum types to avoid name conflicts
var disambiguate = map[string]adjust{
	"CodeActionTriggerKind":        {"CodeAction", ""},
	"CompletionItemKind":           {"", "Completion"},
	"CompletionItemTag":            {"Compl", ""},
	"DiagnosticSeverity":           {"Severity", ""},
	"DocumentDiagnosticReportKind": {"Diagnostic", ""},
	"FileOperationPatternKind":     {"", "Pattern"},
	"InsertTextFormat":             {"", "TextFormat"},
	"SemanticTokenModifiers":       {"Mod", ""},
	"SemanticTokenTypes":           {"", "Type"},
	"SignatureHelpTriggerKind":     {"Sig", ""},
	"SymbolTag":                    {"", "Symbol"},
	"WatchKind":                    {"Watch", ""},
}

// which entries of disambiguate got used
var usedDisambiguate = make(map[string]bool)

// for gopls compatibility, replace generated type names with existing ones
var goplsType = map[string]string{
	"And_RegOpt_textDocument_colorPresentation": "WorkDoneProgressOptionsAndTextDocumentRegistrationOptions",
	"ConfigurationParams":                       "ParamConfiguration",
	"DocumentDiagnosticParams":                  "string",
	"DocumentDiagnosticReport":                  "string",
	"DocumentUri":                               "DocumentURI",
	"InitializeParams":                          "ParamInitialize",
	"LSPAny":                                    "interface{}",

	"Lit_CodeActionClientCapabilities_codeActionLiteralSupport":                "PCodeActionLiteralSupportPCodeAction",
	"Lit_CodeActionClientCapabilities_codeActionLiteralSupport_codeActionKind": "FCodeActionKindPCodeActionLiteralSupport",

	"Lit_CodeActionClientCapabilities_resolveSupport":     "PResolveSupportPCodeAction",
	"Lit_CodeAction_disabled":                             "PDisabledMsg_textDocument_codeAction",
	"Lit_CompletionClientCapabilities_completionItem":     "PCompletionItemPCompletion",
	"Lit_CompletionClientCapabilities_completionItemKind": "PCompletionItemKindPCompletion",

	"Lit_CompletionClientCapabilities_completionItem_insertTextModeSupport": "FInsertTextModeSupportPCompletionItem",

	"Lit_CompletionClientCapabilities_completionItem_resolveSupport": "FResolveSupportPCompletionItem",
	"Lit_CompletionClientCapabilities_completionItem_tagSupport":     "FTagSupportPCompletionItem",

	"Lit_CompletionClientCapabilities_completionList":     "PCompletionListPCompletion",
	"Lit_CompletionList_itemDefaults":                     "PItemDefaultsMsg_textDocument_completion",
	"Lit_CompletionList_itemDefaults_editRange_Item1":     "FEditRangePItemDefaults",
	"Lit_CompletionOptions_completionItem":                "PCompletionItemPCompletionProvider",
	"Lit_DocumentSymbolClientCapabilities_symbolKind":     "PSymbolKindPDocumentSymbol",
	"Lit_DocumentSymbolClientCapabilities_tagSupport":     "PTagSupportPDocumentSymbol",
	"Lit_FoldingRangeClientCapabilities_foldingRange":     "PFoldingRangePFoldingRange",
	"Lit_FoldingRangeClientCapabilities_foldingRangeKind": "PFoldingRangeKindPFoldingRange",
	"Lit_GeneralClientCapabilities_staleRequestSupport":   "PStaleRequestSupportPGeneral",
	"Lit_InitializeResult_serverInfo":                     "PServerInfoMsg_initialize",
	"Lit_InlayHintClientCapabilities_resolveSupport":      "PResolveSupportPInlayHint",
	"Lit_MarkedString_Item1":                              "Msg_MarkedString",
	"Lit_NotebookDocumentChangeEvent_cells":               "PCellsPChange",
	"Lit_NotebookDocumentChangeEvent_cells_structure":     "FStructurePCells",
	"Lit_NotebookDocumentFilter_Item0":                    "Msg_NotebookDocumentFilter",

	"Lit_NotebookDocumentSyncOptions_notebookSelector_Elem_Item0": "PNotebookSelectorPNotebookDocumentSync",

	"Lit_PrepareRenameResult_Item1": "Msg_PrepareRename2Gn",

	"Lit_PublishDiagnosticsClientCapabilities_tagSupport":       "PTagSupportPPublishDiagnostics",
	"Lit_SemanticTokensClientCapabilities_requests":             "PRequestsPSemanticTokens",
	"Lit_SemanticTokensClientCapabilities_requests_full_Item1":  "FFullPRequests",
	"Lit_SemanticTokensClientCapabilities_requests_range_Item1": "FRangePRequests",

	"Lit_SemanticTokensOptions_full_Item1":  "PFullESemanticTokensOptions",
	"Lit_SemanticTokensOptions_range_Item1": "PRangeESemanticTokensOptions",
	"Lit_ServerCapabilities_workspace":      "Workspace6Gn",

	"Lit_ShowMessageRequestClientCapabilities_messageActionItem": "PMessageActionItemPShowMessage",
	"Lit_SignatureHelpClientCapabilities_signatureInformation":   "PSignatureInformationPSignatureHelp",

	"Lit_SignatureHelpClientCapabilities_signatureInformation_parameterInformation": "FParameterInformationPSignatureInformation",

	"Lit_TextDocumentContentChangeEvent_Item0":                    "Msg_TextDocumentContentChangeEvent",
	"Lit_TextDocumentFilter_Item0":                                "Msg_TextDocumentFilter",
	"Lit_TextDocumentFilter_Item1":                                "Msg_TextDocumentFilter",
	"Lit_WorkspaceEditClientCapabilities_changeAnnotationSupport": "PChangeAnnotationSupportPWorkspaceEdit",
	"Lit_WorkspaceSymbolClientCapabilities_resolveSupport":        "PResolveSupportPSymbol",
	"Lit_WorkspaceSymbolClientCapabilities_symbolKind":            "PSymbolKindPSymbol",
	"Lit_WorkspaceSymbolClientCapabilities_tagSupport":            "PTagSupportPSymbol",
	"Lit_WorkspaceSymbol_location_Item1":                          "PLocationMsg_workspace_symbol",
	"Lit__InitializeParams_clientInfo":                            "Msg_XInitializeParams_clientInfo",
	"Or_CompletionList_itemDefaults_editRange":                    "OrFEditRangePItemDefaults",
	"Or_Declaration": "[]Location",
	"Or_DidChangeConfigurationRegistrationOptions_section": "OrPSection_workspace_didChangeConfiguration",
	"Or_GlobPattern":                "string",
	"Or_InlayHintLabelPart_tooltip": "OrPTooltipPLabel",
	"Or_InlayHint_tooltip":          "OrPTooltip_textDocument_inlayHint",
	"Or_LSPAny":                     "interface{}",
	"Or_NotebookDocumentFilter":     "Msg_NotebookDocumentFilter",
	"Or_NotebookDocumentSyncOptions_notebookSelector_Elem": "PNotebookSelectorPNotebookDocumentSync",

	"Or_NotebookDocumentSyncOptions_notebookSelector_Elem_Item0_notebook": "OrFNotebookPNotebookSelector",

	"Or_ParameterInformation_documentation":                     "string",
	"Or_ParameterInformation_label":                             "string",
	"Or_PrepareRenameResult":                                    "Msg_PrepareRename2Gn",
	"Or_ProgressToken":                                          "interface{}",
	"Or_Result_textDocument_completion":                         "CompletionList",
	"Or_Result_textDocument_declaration":                        "Or_textDocument_declaration",
	"Or_Result_textDocument_definition":                         "[]Location",
	"Or_Result_textDocument_documentSymbol":                     "[]interface{}",
	"Or_Result_textDocument_implementation":                     "[]Location",
	"Or_Result_textDocument_semanticTokens_full_delta":          "interface{}",
	"Or_Result_textDocument_typeDefinition":                     "[]Location",
	"Or_Result_workspace_symbol":                                "[]SymbolInformation",
	"Or_TextDocumentContentChangeEvent":                         "Msg_TextDocumentContentChangeEvent",
	"Or_TextDocumentFilter":                                     "Msg_TextDocumentFilter",
	"Or_WorkspaceFoldersServerCapabilities_changeNotifications": "string",
	"Or_WorkspaceSymbol_location":                               "OrPLocation_workspace_symbol",
	"PrepareRenameResult":                                       "PrepareRename2Gn",
	"Tuple_ParameterInformation_label_Item1":                    "UIntCommaUInt",
	"WorkspaceFoldersServerCapabilities":                        "WorkspaceFolders5Gn",
	"[]LSPAny":                                                  "[]interface{}",
	"[]Or_NotebookDocumentSyncOptions_notebookSelector_Elem":    "[]PNotebookSelectorPNotebookDocumentSync",
	"[]Or_Result_textDocument_codeAction_Item0_Elem":            "[]CodeAction",
	"[]PreviousResultId":                                        "[]PreviousResultID",
	"[]uinteger":                                                "[]uint32",
	"boolean":                                                   "bool",
	"decimal":                                                   "float64",
	"integer":                                                   "int32",
	"map[DocumentUri][]TextEdit":                                "map[DocumentURI][]TextEdit",
	"uinteger":                                                  "uint32",
}

var usedGoplsType = make(map[string]bool)

// methodNames is a map from the method to the name of the function that handles it
var methodNames = map[string]string{
	"$/cancelRequest":                        "CancelRequest",
	"$/logTrace":                             "LogTrace",
	"$/progress":                             "Progress",
	"$/setTrace":                             "SetTrace",
	"callHierarchy/incomingCalls":            "IncomingCalls",
	"callHierarchy/outgoingCalls":            "OutgoingCalls",
	"client/registerCapability":              "RegisterCapability",
	"client/unregisterCapability":            "UnregisterCapability",
	"codeAction/resolve":                     "ResolveCodeAction",
	"codeLens/resolve":                       "ResolveCodeLens",
	"completionItem/resolve":                 "ResolveCompletionItem",
	"documentLink/resolve":                   "ResolveDocumentLink",
	"exit":                                   "Exit",
	"initialize":                             "Initialize",
	"initialized":                            "Initialized",
	"inlayHint/resolve":                      "Resolve",
	"notebookDocument/didChange":             "DidChangeNotebookDocument",
	"notebookDocument/didClose":              "DidCloseNotebookDocument",
	"notebookDocument/didOpen":               "DidOpenNotebookDocument",
	"notebookDocument/didSave":               "DidSaveNotebookDocument",
	"shutdown":                               "Shutdown",
	"telemetry/event":                        "Event",
	"textDocument/codeAction":                "CodeAction",
	"textDocument/codeLens":                  "CodeLens",
	"textDocument/colorPresentation":         "ColorPresentation",
	"textDocument/completion":                "Completion",
	"textDocument/declaration":               "Declaration",
	"textDocument/definition":                "Definition",
	"textDocument/diagnostic":                "Diagnostic",
	"textDocument/didChange":                 "DidChange",
	"textDocument/didClose":                  "DidClose",
	"textDocument/didOpen":                   "DidOpen",
	"textDocument/didSave":                   "DidSave",
	"textDocument/documentColor":             "DocumentColor",
	"textDocument/documentHighlight":         "DocumentHighlight",
	"textDocument/documentLink":              "DocumentLink",
	"textDocument/documentSymbol":            "DocumentSymbol",
	"textDocument/foldingRange":              "FoldingRange",
	"textDocument/formatting":                "Formatting",
	"textDocument/hover":                     "Hover",
	"textDocument/implementation":            "Implementation",
	"textDocument/inlayHint":                 "InlayHint",
	"textDocument/inlineValue":               "InlineValue",
	"textDocument/linkedEditingRange":        "LinkedEditingRange",
	"textDocument/moniker":                   "Moniker",
	"textDocument/onTypeFormatting":          "OnTypeFormatting",
	"textDocument/prepareCallHierarchy":      "PrepareCallHierarchy",
	"textDocument/prepareRename":             "PrepareRename",
	"textDocument/prepareTypeHierarchy":      "PrepareTypeHierarchy",
	"textDocument/publishDiagnostics":        "PublishDiagnostics",
	"textDocument/rangeFormatting":           "RangeFormatting",
	"textDocument/references":                "References",
	"textDocument/rename":                    "Rename",
	"textDocument/selectionRange":            "SelectionRange",
	"textDocument/semanticTokens/full":       "SemanticTokensFull",
	"textDocument/semanticTokens/full/delta": "SemanticTokensFullDelta",
	"textDocument/semanticTokens/range":      "SemanticTokensRange",
	"textDocument/signatureHelp":             "SignatureHelp",
	"textDocument/typeDefinition":            "TypeDefinition",
	"textDocument/willSave":                  "WillSave",
	"textDocument/willSaveWaitUntil":         "WillSaveWaitUntil",
	"typeHierarchy/subtypes":                 "Subtypes",
	"typeHierarchy/supertypes":               "Supertypes",
	"window/logMessage":                      "LogMessage",
	"window/showDocument":                    "ShowDocument",
	"window/showMessage":                     "ShowMessage",
	"window/showMessageRequest":              "ShowMessageRequest",
	"window/workDoneProgress/cancel":         "WorkDoneProgressCancel",
	"window/workDoneProgress/create":         "WorkDoneProgressCreate",
	"workspace/applyEdit":                    "ApplyEdit",
	"workspace/codeLens/refresh":             "CodeLensRefresh",
	"workspace/configuration":                "Configuration",
	"workspace/diagnostic":                   "DiagnosticWorkspace",
	"workspace/diagnostic/refresh":           "DiagnosticRefresh",
	"workspace/didChangeConfiguration":       "DidChangeConfiguration",
	"workspace/didChangeWatchedFiles":        "DidChangeWatchedFiles",
	"workspace/didChangeWorkspaceFolders":    "DidChangeWorkspaceFolders",
	"workspace/didCreateFiles":               "DidCreateFiles",
	"workspace/didDeleteFiles":               "DidDeleteFiles",
	"workspace/didRenameFiles":               "DidRenameFiles",
	"workspace/executeCommand":               "ExecuteCommand",
	"workspace/inlayHint/refresh":            "InlayHintRefresh",
	"workspace/inlineValue/refresh":          "InlineValueRefresh",
	"workspace/semanticTokens/refresh":       "SemanticTokensRefresh",
	"workspace/symbol":                       "Symbol",
	"workspace/willCreateFiles":              "WillCreateFiles",
	"workspace/willDeleteFiles":              "WillDeleteFiles",
	"workspace/willRenameFiles":              "WillRenameFiles",
	"workspace/workspaceFolders":             "WorkspaceFolders",
	"workspaceSymbol/resolve":                "ResolveWorkspaceSymbol",
}
