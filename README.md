# ⚙️ MHF-API

## 🌐 Overview

This API project provides a robust interface for accessing and managing uncompiled mhfdat.bin data.

Designed with scalability and maintainability in mind, the API is organized following clean architecture principles. It features well-defined controllers for handling requests, middlewares for tasks such as logging and routing, models to represent data structures, and utility functions for seamless integration with tools like New Relic.

This structure ensures efficient request handling, extensibility, and ease of development while maintaining performance and monitoring capabilities.

## 🚀 Getting Started

### ✅ Prerequisites

- 🛠️ Go 1.15+ installed.
- 📊 New Relic account (if you want to use monitoring).
- 💻 Git installed for commit fetching.

### 📥 Installation

1. Clone the repository:

```sh
git clone https://github.com/mhf-utils/mhf-api.git
```

2. Install dependencies:

```sh
go mod download
```

3. [Configure your environment](https://github.com/mhf-utils/mhf-api?tab=readme-ov-file#%EF%B8%8F-configuration-overview) variables for logging and monitoring.

### ▶️ Running the API

Don't forget to change the configuration.

To run the API, execute:

```sh
go build . && ENVIRONMENT=dev go run .
```

### 📡 API Usage

Here are some example endpoints:

- **GET** `/items`: Fetch all items.
- **GET** `/items/{id}`: Fetch a specific item by ID.

## 🗂️ Project Structure

The project is organized as follows:

```
mhf-api/
├── config/                          # 🔧 For application configurations
|   ├── dev/                         # 💻 Contains configurations for the development environment
|   |   ├── base.json                # ⚙️ General configuration (e.g., host, port)
|   |   ├── logger.json              # 📝 Configuration for logging system
|   |   ├── mhfdat.json              # 📂 Path to the mhfdat file
|   |   └── newrelic.json            # 📊 New Relic monitoring parameters
|   |
|   ├── prod/                        # 🏢 Contains configurations for the production environment
|   |   ├── base.json                # ⚙️ General configuration
|   |   ├── logger.json              # 📝 Configuration for logging system
|   |   ├── mhfdat.json              # 📂 Path to the mhfdat file
|   |   └── newrelic.json            # 📊 New Relic monitoring parameters
|   |
|   └── index.go                     # 🗂️ Initializes configurations based on the environment
|
├── core/                            # 💡 Contains the core logic of the API
|   └── index.go                     # ⚙️ Core of the application (main logic)
|
├── server/                          # 🌐 Handles all server-related functionalities (routes, middleware)
|   ├── common/                      # 📦 Shared code for various parts of the server
|   |   └── index.go                 # ⚙️ Shared logic or common code across multiple server components
|   |
|   ├── controllers/                 # 🗄️ Folder for request handlers and business logic
|   |   └── item.go                  # 🛍️ Manages requests related to "items" (CRUD operations)
|   |
|   ├── middlewares/                 # 🛡️ Contains middleware functions (logging, routing, etc.)
|   |   ├── index.go                 # ⚙️ Global middleware functions
|   |   ├── item.go                  # 🔄 Middleware and routing for item-related requests
|   |   └── logger.go                # 📝 Middleware for HTTP request logging
|   |
|   ├── models/                      # 🏗️ Contains data models representing API objects
|   |
|   └── index.go                     # 🚀 Initializes the server, sets up routes, and starts the server
|
├── utils/                           # 🛠️ Folder for utility functions (logging, ASCII art, New Relic)
|   ├── ascii/                       # 🎨 Contains ASCII art template shown when the server starts
|   ├── binary/                      # 🗂️ Utility functions for handling binary files
|   ├── logger/                      # 📝 Initialization and configuration of the logging system
|   ├── newrelic/                    # 📊 Functions for New Relic integration and performance monitoring
|   └── pointers/                    # 📌 List and declaration of pointers to access data
|
└── main.go                          # 🚪 Main entry point of the application where everything is initialized
```

## ⚙️ Configuration Overview

The API configuration is handled through the `config` package. It uses `viper` to manage configuration files for different environments (e.g., `dev`, `prod`) and sets up key application settings, such as the server host, logging, MHF data file paths, and New Relic integration.

### 🗂️ Configuration Files

Within the `config/` directory, there are subdirectories for each environment (e.g., `dev/` and `prod/`), each containing configuration files:

- **`base.json`**: General application settings.
- **`logger.json`**: Configures the logging system (format and file path).
- **`mhfdat.json`**: Specifies the path to the `mhfdat` file that the API interacts with.
- **`newrelic.json`**: Contains settings for New Relic integration, including the license key and app name.

These configuration files allow the application to be easily configured based on the environment in which it's running.

### 🛠️ How the Configuration Works

When the API starts, it loads the appropriate configuration based on the environment (defined by the `ENVIRONMENT` variable). The configuration loader fetches the relevant JSON files, decodes them, and merges their values into the global `Config` struct.

Here's a breakdown of the `Config` struct:

```go
type Config struct {
  Host     string     // The server host (automatically detected if not set)
  Port     string     // The port the server listens on
  Logger   Logger     // Logger configuration
  Mhfdat   Mhfdat     // MHF data file path
  NewRelic NewRelic   // New Relic settings
}

type Logger struct {
  Format   string     // Logging format (e.g., JSON or text)
  FilePath string     // File path for log output
}

type Mhfdat struct {
  FilePath string     // Path to the mhfdat.bin file
}

type NewRelic struct {
  License string      // New Relic license key
  AppName string      // Application name for monitoring in New Relic
  AppLogForwardingEnabled bool // Whether log forwarding is enabled
}
```

### 🧰 Loading Configurations

When the application starts, it uses the `LoadConfig` function to load configurations for the environment:

```go
func LoadConfig(env string) (*Config, error) {
  var config Config

  config_files := []ConfigFile{
    {Name: "base"},
    {Name: "logger"},
    {Name: "mhfdat"},
    {Name: "newrelic"},
  }

  viper.SetConfigType("json")
  path := fmt.Sprintf("./config/%s", env)
  viper.AddConfigPath(path)

  for _, config_file := range config_files {
    viper.SetConfigName(config_file.Name)

    if err := viper.ReadInConfig(); err != nil {
      return nil, err
    }
    if err := viper.Unmarshal(&config); err != nil {
      return nil, err
    }
  }

  if config.Host == "" {
    config.Host = getOutboundIP4().To4().String()
  }
  return &config, nil
}
```

### 📝 Example Configurations

Example of a `base.json` file (for general settings):

```json
{
  "host": "127.0.0.1",
  "port": ":8080"
}
```

Example of a `logger.json` file:

```json
{
  "format": "json",
  "filePath": "./logs/app.log"
}
```

Example of a `mhfdat.json` file:

```json
{
  "filePath": "/path/to/mhfdat.bin"
}
```

Example of a `newrelic.json` file:

```json
{
  "license": "your-new-relic-license-key",
  "appName": "MHF-API",
  "appLogForwardingEnabled": true
}
```

## ⚙️ How It Works

### 🔑 Main Application

The entry point is in `main.go`, where the application initializes the logger and New Relic monitoring, and starts the server:

1. **Logger Initialization**: Initializes the logger using a configuration file.
2. **New Relic Monitoring**: Sets up application monitoring with New Relic.
3. **Commit Retrieval**: The application retrieves the latest Git commit (short hash) to display with the server's ASCII title.
4. **Server Startup**: The server is started using the `Init()` function from the `server` package, which binds the router and applies middleware.

### 🌐 Server Initialization

The `server.Init()` function handles:

- **Router Setup**: The Gorilla Mux router is set up to manage routes.
- **Middleware Chaining**: Middlewares, such as logging, are applied to every request.
- **Listening**: The server listens for incoming HTTP requests

 on the configured port.


## 🛠️ Example

### Controllers:

This file manages item-related endpoints:

- **`List`**: Returns a list of all data.
- **`Read`**: Fetches a specific data by its ID.

### Middlewares:

- **GET** `/items`: Returns all items.
- **GET** `/items/{id}`: Returns a specific item by ID.


Middlewares are defined in `middlewares/`:

- **Router Middleware**: `GetRouterItem()` defines the routes for item-related API requests.
- **Logging Middleware**: Logs incoming requests and their responses for debugging and monitoring purposes.
- **Custom Middlewares**: Additional middlewares (like `user.go`) can be added as needed for other features.

### 🏗️ Models

The `models/item.go` file defines the `Item` struct:

```go
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

This model is used to represent items across the API.


### 🛠️ Utilities

1. **ASCII Art**: Displayed when the server starts.
2. **Logger**: Configured with the New Relic app to log events and errors.
3. **New Relic**: Used for application performance monitoring.

## 🤝 Contributing

Feel free to submit issues or pull requests to improve the API or add new features. Contributions are always welcome!

## 📝 License

This project is licensed under the MIT License.
