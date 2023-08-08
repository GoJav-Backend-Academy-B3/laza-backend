// Code generated by swaggo/swag. DO NOT EDIT.

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
            "email": "soberkoder@swagger.io"
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
        "/address": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Post details of address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Post Details",
                "parameters": [
                    {
                        "description": "create address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.AddressRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Address"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/address/": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get all Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Get All Address",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Address"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/address/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get details of address corresponding is the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Get Details for a given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the address",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Address"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "put details of address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "put Details",
                "parameters": [
                    {
                        "description": "update address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.AddressRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "ID of the address",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Address"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete the address corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Delete address identified by the given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the address",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "data": {
                                            "$ref": "#/definitions/model.Address"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "description": {
                                            "type": "string"
                                        },
                                        "isError": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {},
                "isError": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_primary": {
                    "type": "boolean"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Order"
                    }
                },
                "phone_number": {
                    "type": "string"
                },
                "receiver_name": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Category": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                }
            }
        },
        "model.CreditCard": {
            "type": "object",
            "properties": {
                "card_number": {
                    "type": "string"
                },
                "expired_month": {
                    "type": "integer"
                },
                "expired_year": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Order"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "address_id": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "credit_card_id": {
                    "type": "integer"
                },
                "gopay_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "order_status": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "transaction_bank_id": {
                    "type": "integer"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Transaction"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "cart_users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "categories": {
                    "description": "many 2 many",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Category"
                    }
                },
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Order"
                    }
                },
                "price": {
                    "type": "number"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Review"
                    }
                },
                "size": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Size"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "wishlisted_users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "model.Review": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Size": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "product": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "size": {
                    "type": "string"
                }
            }
        },
        "model.Transaction": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "expiry_time": {
                    "type": "integer"
                },
                "fraud_status": {
                    "type": "string"
                },
                "gross_amount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                },
                "payment_type": {
                    "type": "string"
                },
                "settlement_time": {
                    "type": "integer"
                },
                "signature": {
                    "type": "string"
                },
                "transaction_status": {
                    "type": "string"
                },
                "transaction_time": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Address"
                    }
                },
                "cart_products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "credit_cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CreditCard"
                    }
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "is_verified": {
                    "type": "boolean"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Order"
                    }
                },
                "password": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Review"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verification_codes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VerificationCode"
                    }
                },
                "verification_tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VerificationToken"
                    }
                },
                "wishlisted_product": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                }
            }
        },
        "model.VerificationCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.VerificationToken": {
            "type": "object",
            "properties": {
                "expiry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "requests.AddressRequest": {
            "type": "object",
            "required": [
                "city",
                "country",
                "is_primary",
                "phone_number",
                "receiver_name"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "is_primary": {
                    "type": "boolean"
                },
                "phone_number": {
                    "type": "string",
                    "maxLength": 20
                },
                "receiver_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
            "type": "apiKey",
            "name": "X-Auth-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Laza",
	Description:      "This is a Final Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
