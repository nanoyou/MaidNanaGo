// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/about": {
            "get": {
                "description": "获取 Maid Nana 调试信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "about"
                ],
                "summary": "调试信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.DebugInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.DebugInfo": {
            "type": "object",
            "properties": {
                "goVersion": {
                    "type": "string"
                },
                "qq": {
                    "type": "object",
                    "properties": {
                        "account": {
                            "type": "integer"
                        },
                        "online": {
                            "type": "boolean"
                        }
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0-alpha",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Main Nana API 文档",
	Description:      "Maid Nana 的 Web API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}