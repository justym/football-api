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
    "description": "This is a simple Football API",
    "title": "Football API",
    "contact": {
      "email": "justym8132@gmail.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/football-api/1.0.0",
  "paths": {
    "/rounds/{seasonCode}": {
      "get": {
        "description": "Rounds when {seasonCode}\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "rounds"
        ],
        "summary": "Teams when the season",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up rounds",
            "name": "seasonCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success finding rounds in the seasons",
            "schema": {
              "$ref": "#/definitions/SeasonRounds"
            }
          },
          "400": {
            "description": "bad request"
          }
        }
      }
    },
    "/teams/{seasonCode}": {
      "get": {
        "description": "Teams when {seasonCode}\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "teams"
        ],
        "summary": "Teams when the season",
        "operationId": "findTeams",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up teams",
            "name": "seasonCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success finding teams in the seasons",
            "schema": {
              "$ref": "#/definitions/SeasonTeams"
            }
          },
          "400": {
            "description": "bad request"
          }
        }
      }
    }
  },
  "definitions": {
    "Club": {
      "type": "object",
      "required": [
        "key",
        "name",
        "code"
      ],
      "properties": {
        "code": {
          "type": "string",
          "example": "MUN"
        },
        "key": {
          "type": "string",
          "example": "manutd"
        },
        "name": {
          "type": "string",
          "example": "Manchester United FC"
        }
      }
    },
    "Match": {
      "type": "object",
      "required": [
        "date",
        "team1",
        "team2",
        "score1",
        "score2"
      ],
      "properties": {
        "date": {
          "type": "string",
          "example": "2019-08-10"
        },
        "score1": {
          "$ref": "#/definitions/Score"
        },
        "score2": {
          "$ref": "#/definitions/Score"
        },
        "team1": {
          "$ref": "#/definitions/Club"
        },
        "team2": {
          "$ref": "#/definitions/Club"
        }
      }
    },
    "Round": {
      "type": "object",
      "required": [
        "name",
        "matches"
      ],
      "properties": {
        "matches": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Match"
          }
        },
        "name": {
          "type": "string",
          "example": "Matchday 1"
        }
      }
    },
    "Score": {
      "type": "integer",
      "format": "int64",
      "example": 1
    },
    "SeasonName": {
      "type": "string",
      "example": "Premier League 2010/11"
    },
    "SeasonRounds": {
      "type": "object",
      "required": [
        "name",
        "rounds"
      ],
      "properties": {
        "name": {
          "$ref": "#/definitions/SeasonName"
        },
        "rounds": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        }
      }
    },
    "SeasonTeams": {
      "type": "object",
      "required": [
        "name",
        "club"
      ],
      "properties": {
        "club": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Club"
          }
        },
        "name": {
          "$ref": "#/definitions/SeasonName"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Teams in a Particurarly Seasons",
      "name": "teams"
    },
    {
      "description": "Rounds in a Particurarly Seasons",
      "name": "rounds"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple Football API",
    "title": "Football API",
    "contact": {
      "email": "justym8132@gmail.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/football-api/1.0.0",
  "paths": {
    "/rounds/{seasonCode}": {
      "get": {
        "description": "Rounds when {seasonCode}\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "rounds"
        ],
        "summary": "Teams when the season",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up rounds",
            "name": "seasonCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success finding rounds in the seasons",
            "schema": {
              "$ref": "#/definitions/SeasonRounds"
            }
          },
          "400": {
            "description": "bad request"
          }
        }
      }
    },
    "/teams/{seasonCode}": {
      "get": {
        "description": "Teams when {seasonCode}\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "teams"
        ],
        "summary": "Teams when the season",
        "operationId": "findTeams",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up teams",
            "name": "seasonCode",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "success finding teams in the seasons",
            "schema": {
              "$ref": "#/definitions/SeasonTeams"
            }
          },
          "400": {
            "description": "bad request"
          }
        }
      }
    }
  },
  "definitions": {
    "Club": {
      "type": "object",
      "required": [
        "key",
        "name",
        "code"
      ],
      "properties": {
        "code": {
          "type": "string",
          "example": "MUN"
        },
        "key": {
          "type": "string",
          "example": "manutd"
        },
        "name": {
          "type": "string",
          "example": "Manchester United FC"
        }
      }
    },
    "Match": {
      "type": "object",
      "required": [
        "date",
        "team1",
        "team2",
        "score1",
        "score2"
      ],
      "properties": {
        "date": {
          "type": "string",
          "example": "2019-08-10"
        },
        "score1": {
          "$ref": "#/definitions/Score"
        },
        "score2": {
          "$ref": "#/definitions/Score"
        },
        "team1": {
          "$ref": "#/definitions/Club"
        },
        "team2": {
          "$ref": "#/definitions/Club"
        }
      }
    },
    "Round": {
      "type": "object",
      "required": [
        "name",
        "matches"
      ],
      "properties": {
        "matches": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Match"
          }
        },
        "name": {
          "type": "string",
          "example": "Matchday 1"
        }
      }
    },
    "Score": {
      "type": "integer",
      "format": "int64",
      "example": 1
    },
    "SeasonName": {
      "type": "string",
      "example": "Premier League 2010/11"
    },
    "SeasonRounds": {
      "type": "object",
      "required": [
        "name",
        "rounds"
      ],
      "properties": {
        "name": {
          "$ref": "#/definitions/SeasonName"
        },
        "rounds": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Round"
          }
        }
      }
    },
    "SeasonTeams": {
      "type": "object",
      "required": [
        "name",
        "club"
      ],
      "properties": {
        "club": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Club"
          }
        },
        "name": {
          "$ref": "#/definitions/SeasonName"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Teams in a Particurarly Seasons",
      "name": "teams"
    },
    {
      "description": "Rounds in a Particurarly Seasons",
      "name": "rounds"
    }
  ]
}`))
}
