// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Jonathan Brenman"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/contacts": {
            "post": {
                "description": "Create new contact",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add new contact",
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/contacts/id/:id": {
            "get": {
                "description": "Retrieve a contact using his ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get contact by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contact search by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "404": {},
                    "500": {}
                }
            },
            "put": {
                "description": "update a contact using his ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update contact by id",
                "responses": {
                    "200": {},
                    "404": {},
                    "500": {}
                }
            },
            "delete": {
                "description": "delete a contact using his ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete contact by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contact delete by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/contacts/search": {
            "get": {
                "description": "search a contact using his ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search contact by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contact search by query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {},
                    "404": {},
                    "500": {}
                }
            }
        },
        "/contactss": {
            "get": {
                "description": "Get all contacts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all contacts",
                "responses": {
                    "200": {}
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger go-contacts",
	Description: "This is a sample of best practices with golang.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}