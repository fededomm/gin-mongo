package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets891247be894f5e343bef1f391a410176a427c5e0 = "{{ define \"healthcheck.tmpl\" }}\n<html>\n    <h1>\n    <p>{{ .title }}</P>\n    </h1>\n</html>\n{{ end }}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"tmpl"}, "/tmpl": []string{"healtcheck.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1681291201, 1681291201642942863),
		Data:     nil,
	}, "/tmpl": &assets.File{
		Path:     "/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1681228181, 1681228181459790753),
		Data:     nil,
	}, "/tmpl/healtcheck.tmpl": &assets.File{
		Path:     "/tmpl/healtcheck.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1681204966, 1681204966849519111),
		Data:     []byte(_Assets891247be894f5e343bef1f391a410176a427c5e0),
	}}, "")
