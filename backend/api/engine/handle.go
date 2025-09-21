package api_engine

import (
	"net/http"
	"strings"

	"github.com/digiconvent/d9t/api/context"
	"github.com/digiconvent/d9t/meta/services"
	"github.com/google/uuid"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) < 2 {
		http.NotFound(w, r)
		return
	}

	pkg := parts[0]
	entity := parts[1]

	if pkgMap, exists := routes[pkg]; exists {
		if entityMap, exists := pkgMap[entity]; exists {

			var operation string
			var id uuid.UUID
			var err error

			if len(parts) == 2 {
				switch r.Method {
				case "POST":
					operation = "create"
				case "GET":
					operation = "list"
				}
				if operation == "" {
					http.Error(w, "", http.StatusMethodNotAllowed)
					return
				}
			} else if len(parts) == 4 {
				rawId := parts[2]
				id, err = uuid.Parse(rawId)
				if err != nil {
					http.NotFound(w, r)
					return
				}
				operation = parts[3]
			}

			if operation != "" {
				if op, exists := entityMap[operation]; exists {
					if op.M != Method(r.Method) {
						http.Error(w, "", http.StatusMethodNotAllowed)
						return
					}
					// do the permission
					ctx := &context.Context{
						Id:       &id,
						Request:  r,
						Response: w,
						Services: services.Ref,
					}
					op.H(ctx)
					return
				} else {
					http.Error(w, "", http.StatusMethodNotAllowed)
					return
				}
			}
		}
	}

	http.NotFound(w, r)
}
