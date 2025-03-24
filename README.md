# go-iconhash

A powerful tool for calculating favicon hashes used in cybersecurity reconnaissance. This tool is an improved implementation of [Becivells/iconhash](https://github.com/Becivells/iconhash).

## Features

- Calculate MMH3 (MurmurHash3) hash of favicons
- Multiple input sources:
  - Direct URL (e.g., https://example.com/favicon.ico)
  - Local file (e.g., favicon.ico)
  - Base64 encoded data
- Multiple output formats:
  - Plain hash number
  - Fofa search format (`icon_hash="123456789"`)
  - Shodan search format (`http.favicon.hash:123456789`)
- Supports both int32 (default) and uint32 hash outputs
- HTTP API server with authentication
- Model Context Protocol (MCP) support for AI integration
- Enhanced error handling and debugging
- Modern CLI interface with help documentation
- Docker support for containerized usage

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/cyberspacesec/go-iconhash.git
cd go-iconhash

# Build the binary
make build
```

### Using Go

```bash
go install github.com/cyberspacesec/go-iconhash/cmd/iconhash@latest
```

### Using Docker

```bash
# Pull the image
docker pull cyberspacesec/iconhash:latest

# Run with Docker
docker run --rm cyberspacesec/iconhash:latest -u https://example.com/favicon.ico
```

## Usage

### Basic Usage

```bash
# Hash from URL
iconhash https://www.example.com/favicon.ico

# Hash from file
iconhash favicon.ico

# Hash from base64 file
iconhash -b64 encoded.txt

# Start the API server
iconhash server -p 8080
```

### Options

```
Usage:
  iconhash [flags] [url or file]

Flags:
  -b, --b64 string        Path to file containing base64 encoded favicon
  -d, --debug             Enable debug output
  -f, --file string       Path to favicon file
      --fofa              Output in Fofa search format (default true)
  -h, --help              Help for iconhash
  -k, --skip-verify       Skip HTTPS certificate verification (default true)
  -s, --shodan            Output in Shodan search format
  -t, --timeout int       HTTP request timeout in seconds (default 10)
      --uint32            Use uint32 format for hash output (default is int32)
  -u, --url string        URL to download favicon from
  -a, --user-agent string Custom User-Agent for HTTP requests
  -v, --version           Version for iconhash
```

### Examples

#### Hash from URL with Debug Output

```bash
iconhash -d -u https://www.example.com/favicon.ico
```

#### Hash from File with Shodan Format

```bash
iconhash -f favicon.ico --shodan --fofa=false
```

#### Hash from Base64 with Custom Timeout

```bash
iconhash -b64 encoded.txt -t 30
```

### API Server

The iconhash tool includes an HTTP API server that allows you to calculate favicon hashes via HTTP requests. This is useful for integrating with other tools or services.

#### Starting the Server

```bash
# Start the server on default port (8080)
iconhash server

# Start the server on a custom port with debug output
iconhash server -p 3000 -d

# Start the server with authentication
iconhash server -a "your-secret-token"
```

#### API Server Options

```
Server Flags:
  -a, --auth-token string  Authentication token (empty for no auth)
  -d, --debug              Enable debug logging
  -H, --host string        Host address to bind to (default "127.0.0.1")
  -k, --insecure           Skip TLS verification for outbound requests (default true)
  -p, --port int           Port to listen on (default 8080)
  -t, --timeout int        Timeout for outbound requests in seconds (default 10)
```

#### API Endpoints

| Endpoint        | Method     | Description                              |
|-----------------|------------|------------------------------------------|
| `/health`       | GET        | Health check                             |
| `/hash/url`     | GET, POST  | Calculate hash from URL                  |
| `/hash/file`    | POST       | Calculate hash from uploaded file        |
| `/hash/base64`  | POST       | Calculate hash from base64 encoded data  |
| `/mcp`          | POST       | Model Context Protocol interaction       |

#### Authentication

If an authentication token is set, all requests (except `/health`) must include the token in one of these ways:

* In the Authorization header: `Authorization: Bearer your-token`
* As a query parameter: `?token=your-token`

#### Example API Requests

**Hash from URL (GET):**
```bash
curl -X GET "http://localhost:8080/hash/url?url=https://example.com/favicon.ico&format=shodan"
```

**Hash from URL (POST):**
```bash
curl -X POST -d "url=https://example.com/favicon.ico" http://localhost:8080/hash/url
```

**Hash from File:**
```bash
curl -X POST -F "file=@favicon.ico" http://localhost:8080/hash/file
```

**Hash from Base64:**
```bash
curl -X POST -d "data=$(base64 -i favicon.ico)" http://localhost:8080/hash/base64
```

**With Authentication:**
```bash
curl -X GET -H "Authorization: Bearer your-token" "http://localhost:8080/hash/url?url=https://example.com/favicon.ico"
```

### Model Context Protocol (MCP)

The iconhash tool supports the Model Context Protocol, which allows integration with AI services. This enables conversational interaction with the tool to calculate favicon hashes.

#### MCP Endpoint

The MCP endpoint is available at `/mcp` and accepts POST requests with JSON payloads following the Model Context Protocol specification.

#### Example MCP Request:

```json
{
  "version": "1.0",
  "protocol": "Model Context Protocol",
  "context": {
    "messages": [
      {
        "role": "user",
        "content": "Calculate the hash for https://example.com/favicon.ico"
      }
    ]
  }
}
```

#### Example MCP Response:

```json
{
  "version": "1.0",
  "protocol": "Model Context Protocol",
  "message": {
    "role": "assistant",
    "content": "Favicon Hash for https://example.com/favicon.ico:\n\nPlain hash: -1424097501\nFofa format: icon_hash=\"-1424097501\"\nShodan format: http.favicon.hash:-1424097501"
  },
  "usage": {
    "prompt_tokens": 14,
    "completion_tokens": 14,
    "total_tokens": 28,
    "processing_time": 0.325,
    "started_at": "2023-09-28T10:15:30Z",
    "completed_at": "2023-09-28T10:15:30.325Z"
  }
}
```

#### Using MCP with curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "version": "1.0",
  "protocol": "Model Context Protocol",
  "context": {
    "messages": [
      {
        "role": "user",
        "content": "Calculate the hash for https://example.com/favicon.ico"
      }
    ]
  }
}' http://localhost:8080/mcp
```

### Docker Usage

#### Building the Docker Image Locally

```bash
# Build the Docker image
git clone https://github.com/cyberspacesec/go-iconhash.git
cd go-iconhash
docker build -t iconhash:latest .
```

#### Using Docker for CLI Operations

```bash
# Hash from URL
docker run --rm iconhash:latest -u https://example.com/favicon.ico

# Hash from local file (mounting a volume)
docker run --rm -v $(pwd)/samples:/app/samples iconhash:latest -f /app/samples/favicon.ico

# Using Shodan format
docker run --rm iconhash:latest -s -u https://example.com/favicon.ico

# Enable debug output
docker run --rm iconhash:latest -d -u https://example.com/favicon.ico
```

#### Running the API Server with Docker

```bash
# Start the API server on port 8080
docker run -d -p 8080:8080 --name iconhash-server iconhash:latest server -a 0.0.0.0 -p 8080

# Start with authentication
docker run -d -p 8080:8080 --name iconhash-server iconhash:latest server -a 0.0.0.0 -p 8080 --auth-token "your-secret-token"

# Check the logs
docker logs iconhash-server

# Stop the server
docker stop iconhash-server
```

#### Using Docker Compose

The project includes a `docker-compose.yml` file that provides configurations for both server and CLI modes:

```bash
# Start the API server
docker-compose up server

# Start the API server in detached mode
docker-compose up -d server

# Run a CLI command
docker-compose run --rm cli -u https://example.com/favicon.ico

# Run with a custom command
docker-compose run --rm cli -f /app/samples/favicon.ico

# Stop all containers
docker-compose down
```

### Website

The project includes a React-based website for interacting with the IconHash tool. See the [website README](./website/README.md) for more information.

## Development

```bash
# Build the binary
make build

# Run tests
make test

# Build Docker image
make docker-build

# Test with a URL
make test-url

# Test with a sample favicon
make test-sample
```

## Use Cases

### Cybersecurity Reconnaissance

The hash values generated by this tool can be used to search for websites with the same favicon in services like:

- [FOFA](https://fofa.info/) - Use the Fofa output format: `icon_hash="-151231234"`
- [Shodan](https://www.shodan.io/) - Use the Shodan output format: `http.favicon.hash:-151231234`

### Integration with Other Tools

With the API server, you can:
- Integrate with web applications
- Build automation scripts
- Create middleware for processing favicons
- Develop security scanning tools

### AI Integration

With the Model Context Protocol support, you can:
- Create conversational interfaces to the tool
- Integrate with AI assistants
- Build intelligent security tools that process natural language
- Automate reconnaissance tasks using conversational AI

## Technical Background

This tool calculates the MMH3 hash of a favicon icon by:

1. Obtaining the favicon.ico file from a URL, local file, or base64 encoded data
2. Encoding the data in base64 format with line breaks every 76 characters (as per RFC 822)
3. Calculating the 32-bit MMH3 hash of the encoded data

This approach is compatible with the hash calculation methods used by Fofa and Shodan search engines.

## License

[MIT License](LICENSE) 