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
        "/connections": {
            "get": {
                "description": "Get all connections",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Get all connections",
                "responses": {
                    "200": {
                        "description": "Returns the connections",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/connections/create": {
            "post": {
                "description": "Create a connection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Create a connection",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/connections.CreateConnectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the ID of the created connection",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns that the request body is invalid or that some required fields are missing",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "If the connection could not be created",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/connections/delete/{id}": {
            "delete": {
                "description": "Delete a connection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Delete a connection",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Connection ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the success message",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns that the connection ID is invalid",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Returns that the connection does not exist",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/connections/update": {
            "put": {
                "description": "Update a connection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Update a connection",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/connections.UpdateConnectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns that the connection was updated successfully",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns that the request body is invalid or that some required fields are missing",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Returns that the connection with the given ID does not exist",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/connections/{id}": {
            "get": {
                "description": "Get a connection by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Get a connection by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Connection ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the connection data",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns that the connection ID is invalid",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Returns that the connection does not exist",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Ping the server",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "Ping the server",
                "responses": {
                    "200": {
                        "description": "Returns pong",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/query/{connectionId}": {
            "post": {
                "description": "Query a database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interactDatabases"
                ],
                "summary": "Query a database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Connection ID",
                        "name": "connectionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interactDatabases.requestQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the rows",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "If the request body is invalid",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "501": {
                        "description": "If the connection could not be established",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "If the query could not be run",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/schema/refresh/{connectionId}": {
            "get": {
                "description": "Refresh the schema of a database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interactDatabases"
                ],
                "summary": "Refresh the schema of a database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Connection ID",
                        "name": "connectionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the schema",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns the error",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "501": {
                        "description": "If the connection could not be established or the schema could not be retrieved",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/schema/{connectionId}": {
            "get": {
                "description": "Get the schema of a database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "interactDatabases"
                ],
                "summary": "Get the schema of a database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Connection ID",
                        "name": "connectionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the schema",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Returns the error",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "501": {
                        "description": "If the connection could not be established or the schema could not be retrieved",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "connections.CreateConnectionRequest": {
            "type": "object",
            "properties": {
                "common_name": {
                    "type": "string"
                },
                "db_name": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "ssl_mode": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "connections.UpdateConnectionRequest": {
            "type": "object",
            "properties": {
                "common_name": {
                    "type": "string"
                },
                "database_name": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "port": {
                    "type": "integer"
                },
                "ssl_mode": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "interactDatabases.requestQuery": {
            "type": "object",
            "properties": {
                "query": {
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
