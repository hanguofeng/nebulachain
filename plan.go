package main

type Plan struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	KeyPrefix string `json:"key_prefix"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
}
