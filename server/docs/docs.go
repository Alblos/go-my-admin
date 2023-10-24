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
                "description": "Gets all the connections that a user has added",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Get all connections",
                "responses": {
                    "200": {
                        "description": "Returns all connections",
                        "schema": {
                            "$ref": "#/definitions/connections.HandleGetAllConectionsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/connections/create": {
            "post": {
                "description": "Creates a connection with the given parameters",
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
                            "$ref": "#/definitions/connections.HandleCreateConnectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The connection was created successfully",
                        "schema": {
                            "$ref": "#/definitions/connections.HandleCreateConnectionResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is invalid or that some required fields are missing",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Error creating connection",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/connections/delete/{id}": {
            "delete": {
                "description": "Delete a connection with the given ID",
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
                        "description": "The connection was deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/connections.HandleDeleteConnectionResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid connection ID",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Connection with the given ID does not exist",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/connections/test": {
            "post": {
                "description": "Test a connection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "connections"
                ],
                "summary": "Test a connection",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/connections.HandleTestConnectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The connection was tested successfully",
                        "schema": {
                            "$ref": "#/definitions/connections.HandleTestConnectionResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is invalid or that some required fields are missing",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "402": {
                        "description": "Connection could not be established",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
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
                            "$ref": "#/definitions/connections.HandleUpdateConnectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The connection was updated successfully",
                        "schema": {
                            "$ref": "#/definitions/connections.HandleUpdateConnectionResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is invalid or that some required fields are missing",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Connection with the given ID does not exist",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/connections/{id}": {
            "get": {
                "description": "Get the data of a connection by its ID",
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
                            "$ref": "#/definitions/connections.HandleGetConnectionByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid connection ID",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Connection with the given ID does not exist",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
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
        "connections.HandleCreateConnectionRequest": {
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
        "connections.HandleCreateConnectionResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "connections.HandleDeleteConnectionResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "connections.HandleGetAllConectionsResponse": {
            "type": "object",
            "properties": {
                "connections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Connection"
                    }
                },
                "error": {
                    "type": "boolean"
                }
            }
        },
        "connections.HandleGetConnectionByIdResponse": {
            "type": "object",
            "properties": {
                "connection": {
                    "$ref": "#/definitions/types.Connection"
                },
                "error": {
                    "type": "boolean"
                }
            }
        },
        "connections.HandleTestConnectionRequest": {
            "type": "object",
            "properties": {
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
        "connections.HandleTestConnectionResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "connections.HandleUpdateConnectionRequest": {
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
        "connections.HandleUpdateConnectionResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
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
        },
        "types.Connection": {
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
        "types.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean"
                },
                "message": {
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
