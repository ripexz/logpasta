package main

type Paste struct {
	UUID    string `json:"uuid,omitempty"`
	Content string `json:"content,omitempty"`
}

type PasteData struct {
	Paste `json:"paste"`
}
