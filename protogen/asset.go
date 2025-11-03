package protogen

import "embed"

//go:embed openapiv2/openapiv2.swagger.json
var SwaggerFS embed.FS // 'SwaggerFS' diekspor (huruf besar)
