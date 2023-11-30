// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about resources in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.ResourceResponce"
                        }
                    }
                }
            }
        },
        "/fluxhelmreleases": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about FluxCD Kustomizations in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.FluxHelmreleasesResponse"
                        }
                    }
                }
            }
        },
        "/fluxkustomizations": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about FluxCD Kustomizations in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.FluxKustomizationsResponse"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Liveness and readness probe enpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.NodeRespose"
                        }
                    }
                }
            }
        },
        "/ingresses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about ingresses in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.IngressResponse"
                        }
                    }
                }
            }
        },
        "/namespaces": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about namespaces in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.NamespaceResponse"
                        }
                    }
                }
            }
        },
        "/nodes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about nodes in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.NodeRespose"
                        }
                    }
                }
            }
        },
        "/pods": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get info about pods in cluster",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubeApiResponseStruct.NamespaceResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "kubeApiResponseStruct.FluxHelmreleasesResponse": {
            "type": "object",
            "properties": {
                "fluxhelmreleasesnumber": {
                    "type": "integer"
                },
                "fluxhelmreleasesnumberbystatus": {
                    "$ref": "#/definitions/kubeApiResponseStruct.FluxKustomizationsStatus"
                }
            }
        },
        "kubeApiResponseStruct.FluxKustomizationsResponse": {
            "type": "object",
            "properties": {
                "fluxkustomizationsnumber": {
                    "type": "integer"
                },
                "fluxkustomizationsnumberbystatus": {
                    "$ref": "#/definitions/kubeApiResponseStruct.FluxKustomizationsStatus"
                }
            }
        },
        "kubeApiResponseStruct.FluxKustomizationsStatus": {
            "type": "object",
            "properties": {
                "notready": {
                    "type": "integer"
                },
                "ready": {
                    "type": "integer"
                },
                "unknown": {
                    "type": "integer"
                }
            }
        },
        "kubeApiResponseStruct.IngressResponse": {
            "type": "object",
            "properties": {
                "ingressnumber": {
                    "type": "integer"
                }
            }
        },
        "kubeApiResponseStruct.NamespaceResponse": {
            "type": "object",
            "properties": {
                "NamespaceNumber": {
                    "type": "integer"
                }
            }
        },
        "kubeApiResponseStruct.NodeRespose": {
            "type": "object",
            "properties": {
                "contolplanenumber": {
                    "type": "integer"
                },
                "kubernetesversion": {
                    "type": "string"
                },
                "nodenumber": {
                    "type": "integer"
                },
                "osimage": {
                    "type": "string"
                },
                "workernumber": {
                    "type": "integer"
                }
            }
        },
        "kubeApiResponseStruct.PodsResponse": {
            "type": "object",
            "properties": {
                "podsnumber": {
                    "type": "integer"
                },
                "podsnumberbystatus": {
                    "$ref": "#/definitions/kubeApiResponseStruct.PodsStatus"
                }
            }
        },
        "kubeApiResponseStruct.PodsStatus": {
            "type": "object",
            "properties": {
                "failednumber": {
                    "type": "integer"
                },
                "nodenumber": {
                    "type": "integer"
                },
                "runningnumber": {
                    "type": "integer"
                },
                "succeedednumber": {
                    "type": "integer"
                }
            }
        },
        "kubeApiResponseStruct.ResourceResponce": {
            "type": "object",
            "properties": {
                "fluxhelmreleases": {
                    "$ref": "#/definitions/kubeApiResponseStruct.FluxHelmreleasesResponse"
                },
                "fluxkustomizations": {
                    "$ref": "#/definitions/kubeApiResponseStruct.FluxKustomizationsResponse"
                },
                "ingresses": {
                    "$ref": "#/definitions/kubeApiResponseStruct.IngressResponse"
                },
                "namespaces": {
                    "$ref": "#/definitions/kubeApiResponseStruct.NamespaceResponse"
                },
                "nodes": {
                    "$ref": "#/definitions/kubeApiResponseStruct.NodeRespose"
                },
                "pods": {
                    "$ref": "#/definitions/kubeApiResponseStruct.PodsResponse"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "kubeinfo-kubeinfo-backend.kubeinfo.svc",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Kubeinfo API",
	Description:      "Kubeinfo used to get information about k8s cluster component",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}