package res

import _ "embed"

//go:embed launcher.js
var LauncherScript string

//go:embed global.css
var GlobalStyles string
