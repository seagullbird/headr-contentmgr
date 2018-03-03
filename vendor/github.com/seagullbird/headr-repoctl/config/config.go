package config

import "path/filepath"

var (
	PORT         = "unset"
	DEV          = true
	DATADIR      = "/data"
	SITESDIR     = filepath.Join(DATADIR, "sites")
	InitialTheme = "gohugo-theme-ananke"
)
