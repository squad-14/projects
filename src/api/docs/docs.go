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
            "get": {
                "description": "Get all the projects in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get all projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.Project"
                            }
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
                "summary": "Create a project",
                "parameters": [
                    {
                        "description": "Create a project",
                        "name": "project",
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
                            "$ref": "#/definitions/utils.Project"
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
        },
        "/projects/:id": {
            "get": {
                "description": "Get a specific project based on it's ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Project"
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
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a project",
                        "name": "project",
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
                            "$ref": "#/definitions/utils.Project"
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
            },
            "delete": {
                "description": "Delete a project that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Delete a project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
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
        },
        "/projects/:id/state": {
            "patch": {
                "description": "Update a project's state that is already in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Update a project's state",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a project's state",
                        "name": "state",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/utils.StateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Project"
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
        },
        "/tasks": {
            "get": {
                "description": "Gat all the tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get all tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project's ID to filter tasks",
                        "name": "project_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.Task"
                            }
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
                "summary": "Create a task",
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
                            "$ref": "#/definitions/utils.Task"
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
        },
        "/tasks/:id": {
            "get": {
                "description": "Gat a specific task based on it's ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Task"
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
            "put": {
                "description": "Update a task that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a task",
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
                            "$ref": "#/definitions/utils.Task"
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
            },
            "delete": {
                "description": "Delete a task that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
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
                    "type": "string",
                    "example": "Project's description"
                },
                "finish_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "leader": {
                    "type": "string",
                    "example": "Project's leader"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Project's members"
                    ]
                },
                "name": {
                    "type": "string",
                    "example": "Project's name"
                },
                "start_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "TODO",
                        "IN_PROGRESS",
                        "DONE"
                    ],
                    "example": "Project's state"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "dto.Task": {
            "type": "object",
            "required": [
                "description",
                "name",
                "project_id"
            ],
            "properties": {
                "assigned_to": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "description": {
                    "type": "string",
                    "example": "task's description"
                },
                "estimated_hours": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "task's name"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "TODO",
                        "IN_PROGRESS",
                        "DONE"
                    ]
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "errors.ErrResponse": {
            "type": "object",
            "properties": {
                "err_code": {
                    "type": "string",
                    "example": "error.Put.validateState.projects"
                },
                "message": {
                    "type": "string",
                    "example": "Invalid state sarcasm"
                }
            }
        },
        "utils.Project": {
            "type": "object",
            "required": [
                "description",
                "finish_date",
                "name",
                "start_date"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Project's description"
                },
                "finish_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "id": {
                    "type": "integer"
                },
                "leader": {
                    "type": "string",
                    "example": "Project's leader"
                },
                "name": {
                    "type": "string",
                    "example": "Project's name"
                },
                "start_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "TODO",
                        "IN_PROGRESS",
                        "DONE"
                    ],
                    "example": "Project's state"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Deleted ok"
                }
            }
        },
        "utils.StateDTO": {
            "type": "object",
            "properties": {
                "state": {
                    "type": "string",
                    "enum": [
                        "TODO",
                        "IN_PROGRESS",
                        "DONE"
                    ]
                }
            }
        },
        "utils.Task": {
            "type": "object",
            "required": [
                "description",
                "name",
                "project_id"
            ],
            "properties": {
                "assigned_to": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "description": {
                    "type": "string",
                    "example": "task's description"
                },
                "estimated_hours": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "task's name"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string",
                    "example": "2021-12-14T12:41:09.993-04:00"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "TODO",
                        "IN_PROGRESS",
                        "DONE"
                    ]
                },
                "worked_hours": {
                    "type": "integer"
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
	Host:        "https://squad14-2c-2021.herokuapp.com/",
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
