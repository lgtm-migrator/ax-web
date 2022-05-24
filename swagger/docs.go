// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "axiangcoding",
            "email": "axiangcoding@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/demo/get": {
            "get": {
                "tags": [
                    "Demo API"
                ],
                "summary": "Demo for Get",
                "parameters": [
                    {
                        "maxLength": 255,
                        "minLength": 10,
                        "type": "string",
                        "description": "param1, max 255 words",
                        "name": "param1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "param2, required",
                        "name": "param2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "param3, if it's null, validate nothing. if it's not null, must match email regex",
                        "name": "paramEmail",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/demo/post": {
            "post": {
                "tags": [
                    "Demo API"
                ],
                "summary": "Demo for Post",
                "parameters": [
                    {
                        "description": "getParam",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CommonParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/system/info": {
            "get": {
                "tags": [
                    "System API"
                ],
                "summary": "System Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ApiJson": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "v1.CommonParam": {
            "type": "object",
            "required": [
                "param2"
            ],
            "properties": {
                "param1": {
                    "description": "param1, max 255 words",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "param2": {
                    "description": "param2, required",
                    "type": "string"
                },
                "paramEmail": {
                    "description": "param3, if it's null, validate nothing. if it's not null, must match email regex",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "安东星",
	Description: "安东星接口文档",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
