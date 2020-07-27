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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a oauth hello server.",
    "title": "Swagger OAuth2 Hello Example",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/auth/callback": {
      "get": {
        "security": [],
        "summary": "return access_token",
        "responses": {
          "200": {
            "description": "login",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/hello": {
      "get": {
        "security": [
          {
            "GoogleOauthSecurity": [
              "user"
            ]
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Example"
        ],
        "summary": "hello api",
        "operationId": "hello",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hello"
            }
          }
        }
      }
    },
    "/login": {
      "get": {
        "security": [],
        "summary": "login through oauth2 server",
        "responses": {
          "200": {
            "description": "login",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Hello": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "GoogleOauthSecurity": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "Admin scope",
        "user": "User scope"
      }
    },
    "KeyCloakOauthSecurity": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "Admin scope",
        "user": "User scope"
      }
    }
  },
  "tags": [
    {
      "description": "Example",
      "name": "Example"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a oauth hello server.",
    "title": "Swagger OAuth2 Hello Example",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/auth/callback": {
      "get": {
        "security": [],
        "summary": "return access_token",
        "responses": {
          "200": {
            "description": "login",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/hello": {
      "get": {
        "security": [
          {
            "GoogleOauthSecurity": [
              "user"
            ]
          }
        ],
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Example"
        ],
        "summary": "hello api",
        "operationId": "hello",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Hello"
            }
          }
        }
      }
    },
    "/login": {
      "get": {
        "security": [],
        "summary": "login through oauth2 server",
        "responses": {
          "200": {
            "description": "login",
            "schema": {
              "properties": {
                "access_token": {
                  "type": "string",
                  "format": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Hello": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "GoogleOauthSecurity": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "Admin scope",
        "user": "User scope"
      }
    },
    "KeyCloakOauthSecurity": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "admin": "Admin scope",
        "user": "User scope"
      }
    }
  },
  "tags": [
    {
      "description": "Example",
      "name": "Example"
    }
  ]
}`))
}
