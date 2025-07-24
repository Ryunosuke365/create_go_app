package generator

import _ "embed"

//go:embed templates/main.tpl
var MainTemplate string

//go:embed templates/router.tpl
var RouterTemplate string

//go:embed templates/cors.tpl
var CorsTemplate string
