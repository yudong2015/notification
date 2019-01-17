// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated by "makestatic"; DO NOT EDIT.

package static

var Files = map[string]string{
	"api.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "Notification Project",
    "version": "0.0.1",
    "contact": {
      "name": "Notification Project",
      "url": "https://github.com/openpitrix/notification"
    }
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/notifications": {
      "get": {
        "summary": "#API 2.DescribeNfs:describe single Notification with filter.",
        "operationId": "DescribeNfs",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "content_type",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "sent_type",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "title",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "content",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "owner",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "userids_str",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "status",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "limit",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "offset",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sort_key",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "notification"
        ]
      },
      "post": {
        "summary": "#API 1.CreateNfWithAddrs：create notification with addrs(email addrs, phone numbers).",
        "operationId": "CreateNfWithAddrs",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/pbCreateNfWithAddrsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateNfWithAddrsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    }
  },
  "definitions": {
    "pbCreateNfWithAddrsRequest": {
      "type": "object",
      "properties": {
        "content_type": {
          "type": "string"
        },
        "sent_type": {
          "type": "string"
        },
        "addrs_str": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "short_content": {
          "type": "string"
        },
        "expired_days": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "pbCreateNfWithAddrsResponse": {
      "type": "object",
      "properties": {
        "notification_id": {
          "type": "string"
        }
      }
    },
    "pbDescribeNfsResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "description": "The Authorization header must be set to Bearer followed by a space and a token. For example, 'Bearer vHUabiBEIKi8n1RdvWOjGFulGSM6zunb'.",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "BearerAuth": []
    }
  ]
}
`,
}