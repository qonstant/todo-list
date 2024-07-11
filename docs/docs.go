// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/todo-list/tasks": {
            "get": {
                "description": "List all tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "List all tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Error retrieving tasks",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task Input",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createTaskInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/db.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error creating task",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/todo-list/tasks/{id}": {
            "put": {
                "description": "Update an existing task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update an existing task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task Input",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateTaskInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error updating task",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task",
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
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid task ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error deleting task",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/todo-list/tasks/{id}/done": {
            "put": {
                "description": "Mark a task as done",
                "tags": [
                    "Tasks"
                ],
                "summary": "Mark a task as done",
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
                            "$ref": "#/definitions/db.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid task ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error marking task as done",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Task": {
            "type": "object",
            "properties": {
                "active_at": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "http.createTaskInput": {
            "type": "object",
            "properties": {
                "active_at": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "http.updateTaskInput": {
            "type": "object",
            "properties": {
                "active_at": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}