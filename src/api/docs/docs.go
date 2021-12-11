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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/projects": {
            "put": {
                "description": "Modify a project that is already in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Modify a project",
                "parameters": [
                    {
                        "description": "Create a project",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a project to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Add a project",
                "parameters": [
                    {
                        "description": "Create a project",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a project that is already in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Update a project",
                "parameters": [
                    {
                        "description": "Update a project",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "post": {
                "description": "Add a task to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Add a task",
                "parameters": [
                    {
                        "description": "Create a task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Project": {
            "type": "object",
            "required": [
                "description",
                "finish_date",
                "name",
                "start_date"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "finish_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "leader": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "dto.Task": {
            "type": "object",
            "required": [
                "assigned_to",
                "creation_date",
                "description",
                "estimated_hours",
                "name",
                "project_id",
                "start_date",
                "state",
                "worked_hours"
            ],
            "properties": {
                "assigned_to": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "estimated_hours": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "errors.ErrResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "error"
                },
                "message": {
                    "type": "string"
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
	Title:       "PSA Projects API",
	Description: "This API gives access to the projects module.",
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
