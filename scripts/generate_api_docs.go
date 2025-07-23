package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// API endpoint documentation structure
type APIEndpoint struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Description string            `json:"description"`
	Parameters  []Parameter       `json:"parameters"`
	Response    ResponseSchema    `json:"response"`
	Examples    []Example         `json:"examples"`
	Module      string            `json:"module"`
}

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Example     string `json:"example"`
}

type ResponseSchema struct {
	Type        string            `json:"type"`
	Description string            `json:"description"`
	Properties  map[string]string `json:"properties"`
}

type Example struct {
	Title    string `json:"title"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type APIDocumentation struct {
	Version     string        `json:"version"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	BaseURL     string        `json:"base_url"`
	Endpoints   []APIEndpoint `json:"endpoints"`
	GeneratedAt time.Time     `json:"generated_at"`
}

// Stateset Core API endpoints
func getAPIEndpoints() []APIEndpoint {
	return []APIEndpoint{
		// Bank module endpoints
		{
			Method:      "GET",
			Path:        "/cosmos/bank/v1beta1/balances/{address}",
			Description: "Query account balances",
			Module:      "bank",
			Parameters: []Parameter{
				{
					Name:        "address",
					Type:        "string",
					Required:    true,
					Description: "Account address to query",
					Example:     "stateset1abc123def456ghi789",
				},
			},
			Response: ResponseSchema{
				Type:        "object",
				Description: "Account balance information",
				Properties: map[string]string{
					"balances": "array of coin objects",
					"pagination": "pagination info",
				},
			},
			Examples: []Example{
				{
					Title:   "Get account balance",
					Request: "GET /cosmos/bank/v1beta1/balances/stateset1abc123def456ghi789",
					Response: `{
  "balances": [
    {
      "denom": "ustate",
      "amount": "1000000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}`,
				},
			},
		},
		
		// Security module endpoints
		{
			Method:      "POST",
			Path:        "/stateset/security/v1/rules",
			Description: "Create a new security rule",
			Module:      "security",
			Parameters: []Parameter{
				{
					Name:        "rule",
					Type:        "object",
					Required:    true,
					Description: "Security rule configuration",
					Example:     `{"name": "velocity_rule", "type": "velocity", "threshold": 10}`,
				},
			},
			Response: ResponseSchema{
				Type:        "object",
				Description: "Created security rule",
				Properties: map[string]string{
					"id": "string - Rule ID",
					"name": "string - Rule name",
					"is_active": "boolean - Rule status",
				},
			},
			Examples: []Example{
				{
					Title:   "Create velocity rule",
					Request: `POST /stateset/security/v1/rules
{
  "name": "High Velocity Detection",
  "type": "velocity",
  "threshold": 10.0,
  "time_window": 60,
  "action": "alert"
}`,
					Response: `{
  "id": "rule_001",
  "name": "High Velocity Detection",
  "type": "velocity",
  "threshold": 10.0,
  "is_active": true
}`,
				},
			},
		},
		
		// Analytics module endpoints
		{
			Method:      "GET",
			Path:        "/stateset/analytics/v1/metrics",
			Description: "Query blockchain performance metrics",
			Module:      "analytics",
			Parameters: []Parameter{
				{
					Name:        "start_time",
					Type:        "string",
					Required:    false,
					Description: "Start time for metrics query (RFC3339 format)",
					Example:     "2024-01-01T00:00:00Z",
				},
				{
					Name:        "end_time",
					Type:        "string",
					Required:    false,
					Description: "End time for metrics query (RFC3339 format)",
					Example:     "2024-01-02T00:00:00Z",
				},
				{
					Name:        "metric_type",
					Type:        "string",
					Required:    false,
					Description: "Type of metrics to retrieve",
					Example:     "block_time",
				},
			},
			Response: ResponseSchema{
				Type:        "object",
				Description: "Performance metrics data",
				Properties: map[string]string{
					"metrics": "array of metric objects",
					"count": "number of metrics returned",
				},
			},
			Examples: []Example{
				{
					Title:   "Get block time metrics",
					Request: "GET /stateset/analytics/v1/metrics?metric_type=block_time",
					Response: `{
  "metrics": [
    {
      "name": "block_time",
      "value": 2.5,
      "timestamp": "2024-01-01T12:00:00Z",
      "block_height": 12345
    }
  ],
  "count": 1
}`,
				},
			},
		},
		
		// Invoice module endpoints
		{
			Method:      "POST",
			Path:        "/stateset/invoice/v1/invoices",
			Description: "Create a new enhanced invoice",
			Module:      "invoice",
			Parameters: []Parameter{
				{
					Name:        "invoice",
					Type:        "object",
					Required:    true,
					Description: "Enhanced invoice data",
					Example:     `{"amount": "50000", "due_date": "2024-02-01T00:00:00Z"}`,
				},
			},
			Response: ResponseSchema{
				Type:        "object",
				Description: "Created invoice information",
				Properties: map[string]string{
					"id": "string - Invoice ID",
					"status": "string - Invoice status",
					"payment_schedule": "array - Payment schedule",
				},
			},
			Examples: []Example{
				{
					Title:   "Create invoice with payment terms",
					Request: `POST /stateset/invoice/v1/invoices
{
  "amount": {
    "denom": "ustate",
    "amount": "50000"
  },
  "due_date": "2024-02-01T00:00:00Z",
  "payment_terms": {
    "due_days": 30,
    "early_pay_discount": 0.02,
    "late_fee_percentage": 0.015
  }
}`,
					Response: `{
  "id": "INV-2024-001",
  "status": "pending",
  "amount": {
    "denom": "ustate",
    "amount": "50000"
  },
  "payment_schedule": [
    {
      "due_date": "2024-02-01T00:00:00Z",
      "amount": "50000",
      "type": "full"
    }
  ]
}`,
				},
			},
		},
		
		// Agreement module endpoints
		{
			Method:      "GET",
			Path:        "/stateset/agreement/v1/agreements/{id}",
			Description: "Get agreement details by ID",
			Module:      "agreement",
			Parameters: []Parameter{
				{
					Name:        "id",
					Type:        "string",
					Required:    true,
					Description: "Agreement ID",
					Example:     "AGR-2024-001",
				},
			},
			Response: ResponseSchema{
				Type:        "object",
				Description: "Agreement details",
				Properties: map[string]string{
					"id": "string - Agreement ID",
					"parties": "array - Agreement parties",
					"status": "string - Agreement status",
					"terms": "object - Agreement terms",
				},
			},
			Examples: []Example{
				{
					Title:   "Get agreement details",
					Request: "GET /stateset/agreement/v1/agreements/AGR-2024-001",
					Response: `{
  "id": "AGR-2024-001",
  "parties": [
    "stateset1party1...",
    "stateset1party2..."
  ],
  "status": "active",
  "terms": {
    "duration": "12 months",
    "value": "100000"
  }
}`,
				},
			},
		},
	}
}

// HTML template for API documentation
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - API Documentation</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f8f9fa;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 40px;
            border-radius: 10px;
            margin-bottom: 30px;
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 2.5em;
        }
        .header p {
            margin: 10px 0 0 0;
            opacity: 0.9;
        }
        .info-box {
            background: white;
            padding: 20px;
            border-radius: 8px;
            margin-bottom: 20px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .endpoint {
            background: white;
            border-radius: 8px;
            margin-bottom: 30px;
            overflow: hidden;
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
        }
        .endpoint-header {
            background: #f8f9fa;
            padding: 20px;
            border-bottom: 1px solid #e9ecef;
        }
        .method {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 4px;
            font-weight: bold;
            font-size: 0.9em;
            margin-right: 10px;
        }
        .method.GET { background: #d4edda; color: #155724; }
        .method.POST { background: #cce5ff; color: #004085; }
        .method.PUT { background: #fff3cd; color: #856404; }
        .method.DELETE { background: #f8d7da; color: #721c24; }
        .path {
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            background: #f8f9fa;
            padding: 2px 6px;
            border-radius: 4px;
            font-size: 1.1em;
        }
        .endpoint-body {
            padding: 20px;
        }
        .section {
            margin-bottom: 20px;
        }
        .section h4 {
            margin-top: 0;
            color: #495057;
            border-bottom: 2px solid #e9ecef;
            padding-bottom: 5px;
        }
        .parameter {
            background: #f8f9fa;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 10px;
        }
        .parameter-name {
            font-weight: bold;
            color: #495057;
        }
        .parameter-type {
            color: #6c757d;
            font-style: italic;
        }
        .required {
            color: #dc3545;
            font-size: 0.8em;
        }
        .code-block {
            background: #f8f9fa;
            border: 1px solid #e9ecef;
            border-radius: 4px;
            padding: 15px;
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            font-size: 0.9em;
            overflow-x: auto;
            margin: 10px 0;
        }
        .example {
            background: #f0f7ff;
            border-left: 4px solid #007bff;
            padding: 15px;
            margin: 10px 0;
        }
        .toc {
            background: white;
            padding: 20px;
            border-radius: 8px;
            margin-bottom: 20px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .toc ul {
            list-style-type: none;
            padding-left: 0;
        }
        .toc li {
            margin: 8px 0;
        }
        .toc a {
            text-decoration: none;
            color: #007bff;
        }
        .toc a:hover {
            text-decoration: underline;
        }
        .module-badge {
            background: #6c757d;
            color: white;
            padding: 2px 8px;
            border-radius: 12px;
            font-size: 0.8em;
            margin-left: 10px;
        }
        .footer {
            text-align: center;
            margin-top: 40px;
            padding: 20px;
            color: #6c757d;
            border-top: 1px solid #e9ecef;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>{{.Title}}</h1>
        <p>{{.Description}}</p>
        <p><strong>Base URL:</strong> {{.BaseURL}} | <strong>Version:</strong> {{.Version}}</p>
        <p><em>Generated on {{.GeneratedAt.Format "January 2, 2006 at 3:04 PM"}}</em></p>
    </div>

    <div class="info-box">
        <h3>📚 Quick Start Guide</h3>
        <p>This documentation covers all available API endpoints for the Stateset Core blockchain. 
        The blockchain provides enhanced security, analytics, and business functionality for supply chain finance.</p>
        
        <h4>🔑 Key Features:</h4>
        <ul>
            <li><strong>Security Module:</strong> Real-time fraud detection and compliance monitoring</li>
            <li><strong>Analytics Module:</strong> Performance metrics and business intelligence</li>
            <li><strong>Enhanced Invoicing:</strong> Multi-currency support and automated payment scheduling</li>
            <li><strong>Agreement Management:</strong> Smart contract-based business agreements</li>
        </ul>
    </div>

    <div class="toc">
        <h3>📋 Table of Contents</h3>
        <ul>
            {{range .Endpoints}}
            <li>
                <a href="#{{.Method}}-{{.Path | replace "/" "-" | replace "{" "" | replace "}" ""}}">
                    <span class="method {{.Method}}">{{.Method}}</span>
                    {{.Path}}
                    <span class="module-badge">{{.Module}}</span>
                </a>
            </li>
            {{end}}
        </ul>
    </div>

    {{range .Endpoints}}
    <div class="endpoint" id="{{.Method}}-{{.Path | replace "/" "-" | replace "{" "" | replace "}" ""}}">
        <div class="endpoint-header">
            <span class="method {{.Method}}">{{.Method}}</span>
            <span class="path">{{.Path}}</span>
            <span class="module-badge">{{.Module}}</span>
            <h3 style="margin: 10px 0 0 0;">{{.Description}}</h3>
        </div>
        
        <div class="endpoint-body">
            {{if .Parameters}}
            <div class="section">
                <h4>📝 Parameters</h4>
                {{range .Parameters}}
                <div class="parameter">
                    <div class="parameter-name">
                        {{.Name}} 
                        <span class="parameter-type">({{.Type}})</span>
                        {{if .Required}}<span class="required">*required</span>{{end}}
                    </div>
                    <div>{{.Description}}</div>
                    {{if .Example}}<div><strong>Example:</strong> <code>{{.Example}}</code></div>{{end}}
                </div>
                {{end}}
            </div>
            {{end}}

            <div class="section">
                <h4>📤 Response</h4>
                <p><strong>Type:</strong> {{.Response.Type}}</p>
                <p>{{.Response.Description}}</p>
                {{if .Response.Properties}}
                <div class="code-block">
                    {{range $key, $value := .Response.Properties}}
                    {{$key}}: {{$value}}<br>
                    {{end}}
                </div>
                {{end}}
            </div>

            {{if .Examples}}
            <div class="section">
                <h4>💡 Examples</h4>
                {{range .Examples}}
                <div class="example">
                    <h5>{{.Title}}</h5>
                    <p><strong>Request:</strong></p>
                    <div class="code-block">{{.Request}}</div>
                    <p><strong>Response:</strong></p>
                    <div class="code-block">{{.Response}}</div>
                </div>
                {{end}}
            </div>
            {{end}}
        </div>
    </div>
    {{end}}

    <div class="footer">
        <p>🚀 Stateset Core Blockchain API Documentation</p>
        <p>For more information, visit our <a href="https://github.com/stateset/core">GitHub repository</a></p>
    </div>
</body>
</html>
`

func main() {
	log.Println("Generating Stateset Core API Documentation...")

	// Create API documentation structure
	apiDoc := APIDocumentation{
		Version:     "v1.0.0",
		Title:       "Stateset Core Blockchain API",
		Description: "Enhanced blockchain API with security, analytics, and business functionality",
		BaseURL:     "https://api.stateset.io",
		Endpoints:   getAPIEndpoints(),
		GeneratedAt: time.Now(),
	}

	// Generate JSON documentation
	jsonData, err := json.MarshalIndent(apiDoc, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create docs directory
	docsDir := "docs/api"
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		log.Fatalf("Failed to create docs directory: %v", err)
	}

	// Write JSON documentation
	jsonFile := filepath.Join(docsDir, "api.json")
	if err := ioutil.WriteFile(jsonFile, jsonData, 0644); err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	// Generate HTML documentation
	tmpl := template.New("api-docs")
	tmpl = tmpl.Funcs(template.FuncMap{
		"replace": func(input, old, new string) string {
			return strings.Replace(input, old, new, -1)
		},
	})
	
	tmpl, err = tmpl.Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	// Create HTML file
	htmlFile := filepath.Join(docsDir, "index.html")
	f, err := os.Create(htmlFile)
	if err != nil {
		log.Fatalf("Failed to create HTML file: %v", err)
	}
	defer f.Close()

	// Execute template
	if err := tmpl.Execute(f, apiDoc); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	// Generate OpenAPI/Swagger specification
	generateOpenAPISpec(apiDoc, docsDir)

	log.Println("✅ API Documentation generated successfully!")
	log.Printf("📄 JSON: %s", jsonFile)
	log.Printf("🌐 HTML: %s", htmlFile)
	log.Printf("📋 OpenAPI: %s", filepath.Join(docsDir, "openapi.yaml"))
}

func generateOpenAPISpec(apiDoc APIDocumentation, docsDir string) {
	openAPIContent := fmt.Sprintf(`openapi: 3.0.3
info:
  title: %s
  description: %s
  version: %s
servers:
  - url: %s
    description: Production server
  - url: http://localhost:1317
    description: Local development server

paths:`, apiDoc.Title, apiDoc.Description, apiDoc.Version, apiDoc.BaseURL)

	for _, endpoint := range apiDoc.Endpoints {
		openAPIContent += fmt.Sprintf(`
  %s:
    %s:
      summary: %s
      tags:
        - %s
      parameters:`, endpoint.Path, strings.ToLower(endpoint.Method), endpoint.Description, endpoint.Module)

		for _, param := range endpoint.Parameters {
			location := "query"
			if strings.Contains(endpoint.Path, "{"+param.Name+"}") {
				location = "path"
			}
			
			openAPIContent += fmt.Sprintf(`
        - name: %s
          in: %s
          required: %t
          schema:
            type: %s
          description: %s`, param.Name, location, param.Required, param.Type, param.Description)
		}

		openAPIContent += fmt.Sprintf(`
      responses:
        '200':
          description: Successful response
          content:
            application/json:
              schema:
                type: %s
                description: %s`, endpoint.Response.Type, endpoint.Response.Description)
	}

	// Write OpenAPI specification
	openAPIFile := filepath.Join(docsDir, "openapi.yaml")
	if err := ioutil.WriteFile(openAPIFile, []byte(openAPIContent), 0644); err != nil {
		log.Printf("Warning: Failed to write OpenAPI spec: %v", err)
	}
}