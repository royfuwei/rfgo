// Package swag Code generated by swaggo/swag. DO NOT EDIT
package swag

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/app": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Api app name",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/domain.App"
                        }
                    }
                }
            }
        },
        "/jwt/decode": {
            "post": {
                "description": "Decode jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Decode jwt token",
                "parameters": [
                    {
                        "description": "json web token",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ReqJwtToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {}
                    }
                }
            }
        },
        "/jwt/sign": {
            "post": {
                "description": "Sign jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Sign jwt token",
                "parameters": [
                    {
                        "description": "jwt sign",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ReqJwtSign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {}
                    }
                }
            }
        },
        "/jwt/verify": {
            "post": {
                "description": "Verify jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Verify jwt token",
                "parameters": [
                    {
                        "description": "json web token",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ReqJwtToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.App": {
            "type": "object",
            "properties": {
                "app": {
                    "type": "string"
                }
            }
        },
        "domain.ReqJwtSign": {
            "type": "object",
            "properties": {
                "uid": {
                    "type": "string"
                }
            }
        },
        "domain.ReqJwtToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "rfgo open API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
