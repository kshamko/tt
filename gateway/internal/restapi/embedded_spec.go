// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Data GW Swagger",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/data/{id}": {
      "get": {
        "tags": [
          "Data"
        ],
        "operationId": "Data",
        "parameters": [
          {
            "type": "string",
            "description": "id of a record",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Data"
            }
          },
          "404": {
            "description": "data not found",
            "schema": {
              "$ref": "#/definitions/ApiInvalidResponse"
            }
          },
          "500": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/ApiInvalidResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiInvalidResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Data": {
      "type": "object",
      "properties": {
        "alias": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "city": {
          "type": "string"
        },
        "code": {
          "type": "string"
        },
        "coordinates": {
          "type": "array",
          "items": {
            "type": "number",
            "format": "double"
          }
        },
        "country": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "province": {
          "type": "string"
        },
        "regions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timezone": {
          "type": "string"
        },
        "unlocs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Data GW Swagger",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/data/{id}": {
      "get": {
        "tags": [
          "Data"
        ],
        "operationId": "Data",
        "parameters": [
          {
            "type": "string",
            "description": "id of a record",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/Data"
            }
          },
          "404": {
            "description": "data not found",
            "schema": {
              "$ref": "#/definitions/ApiInvalidResponse"
            }
          },
          "500": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/ApiInvalidResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiInvalidResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Data": {
      "type": "object",
      "properties": {
        "alias": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "city": {
          "type": "string"
        },
        "code": {
          "type": "string"
        },
        "coordinates": {
          "type": "array",
          "items": {
            "type": "number",
            "format": "double"
          }
        },
        "country": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "province": {
          "type": "string"
        },
        "regions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timezone": {
          "type": "string"
        },
        "unlocs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
}