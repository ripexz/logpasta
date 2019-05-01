package main

type Paste struct {
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type PasteData struct {
	Paste `json:"paste"`
}
