package openapi

const SpecJSON = `{
  "openapi": "3.0.3",
  "info": {
    "title": "ownp-bg API",
    "version": "1.0.0",
    "description": "Gin + PostgreSQL sample API"
  },
  "paths": {
    "/api/v1/health": {
      "get": {
        "summary": "Health check with PostgreSQL query",
        "responses": {
          "200": {
            "description": "Service healthy",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "status": { "type": "string" },
                    "database_time": { "type": "string", "format": "date-time" }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Database query failed"
          }
        }
      }
    }
  }
}`
