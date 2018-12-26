package pb

const (
	Swagger = ` 
{
  "swagger": "2.0",
  "info": {
    "title": "package nf;",
    "version": "version not set"
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
    "/v1/SayHello": {
      "post": {
        "summary": "#API 0.SayHello:gRPC testing,Sends a greeting.",
        "operationId": "SayHello",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbHelloReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbHelloRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWAppFilter": {
      "post": {
        "summary": "create nf with App Filter",
        "operationId": "CreateNfWAppFilter",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateNfWAppFilterRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWUserFilter": {
      "post": {
        "summary": "create nf with User Filter",
        "operationId": "CreateNfWUserFilter",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateNfWUserFilterRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/CreateNfWithAddrs": {
      "post": {
        "summary": "#API 1.CreateNfWithAddrs：create notification with addrs(email addrs, phone numbers).",
        "operationId": "CreateNfWithAddrs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateNfResponse"
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
    },
    "/v1/nf/DescribeNfs": {
      "post": {
        "summary": "#API 2.DescribeNfs:describe single Notification with filter.",
        "operationId": "DescribeNfs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/DescribeUserNfs": {
      "post": {
        "summary": "describe User Notification with filter",
        "operationId": "DescribeUserNfs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsRequest"
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
    "pbCreateNfResponse": {
      "type": "object",
      "properties": {
        "nf_post_id": {
          "type": "string"
        }
      }
    },
    "pbCreateNfWAppFilterRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
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
        "app_id": {
          "type": "string"
        },
        "app_versions": {
          "type": "string"
        },
        "cluster_status": {
          "type": "string"
        }
      }
    },
    "pbCreateNfWUserFilterRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
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
        "user_filter_type": {
          "type": "string"
        },
        "user_filter_condition": {
          "type": "string"
        },
        "userid_list": {
          "type": "string"
        }
      }
    },
    "pbCreateNfWithAddrsRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "string"
        },
        "notify_type": {
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
    "pbDescribeNfsRequest": {
      "type": "object",
      "properties": {
        "nf_post_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "notify_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "userids_str": {
          "type": "string"
        },
        "status": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "limit": {
          "type": "integer",
          "format": "int64"
        },
        "offset": {
          "type": "integer",
          "format": "int64"
        },
        "sort_key": {
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
    },
    "pbHelloReply": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "title": "The response message containing the greetings"
    },
    "pbHelloRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      },
      "description": "The request message containing the user's name."
    }
  }
}

`
)