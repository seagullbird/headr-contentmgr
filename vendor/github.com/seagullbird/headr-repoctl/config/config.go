package config

import "path/filepath"

var (
	// PORT is the serving port, this value is set during compile time
	PORT = "unset"
	// DATADIR is the root directory for persistent volume
	DATADIR = "/data"
	// SITESDIR contains all user sites
	SITESDIR = filepath.Join(DATADIR, "sites")
	// InitialTheme currently is set dead to gohugo-theme-ananke
	// TODO: InitialTheme should be set in sitemgr and passed to repoctl
	InitialTheme = "gohugo-theme-ananke"
)
