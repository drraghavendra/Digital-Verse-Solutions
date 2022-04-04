package main

type CreateNFTRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Blockchain  string `json:"blockchain"`
}
