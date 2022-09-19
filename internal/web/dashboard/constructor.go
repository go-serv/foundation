package dashboard

import "embed"

//go:embed build/*
var contentFs embed.FS

func NewDashboard() *dashboard {
	d := new(dashboard)
	d.urlPath = "/dashboard"
	d.contentFs = &contentFs
	return d
}
