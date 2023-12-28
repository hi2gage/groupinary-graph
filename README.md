[![codecov](https://codecov.io/github/hi2gage/groupinary-graph/graph/badge.svg?token=OIPVJ3RZHG)](https://codecov.io/github/hi2gage/groupinary-graph)

# Project Name

Brief project description goes here.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
  - [Running Tests](#running-tests)


## Getting Started

Provide a quick overview of what the project is about and what users need to get started.
## Prerequisites
### Prerequisites for Running locally
- Docker

### Prerequisites for Testing locally
- Go 1.21 or later installed locally for Testing 
    - (TODO: Containerize)
- [gotestsum](https://github.com/gotestyourself/gotestsum) for Testing


### Installation

```zsh
# Clone Repo
git clone git@github.com:hi2gage/groupinary-graph.git

# Setup .env file
cp .env.example .env
# Edit the .env to match auth0

```

## Usage

### Running locally
```zsh
# Run Docker
docker-compose up --build

# Open Web GraphQL playground
open https://localhost:8080/

# (optional) Generate key.pem and cert.pem
openssl genrsa -out key.pem 2048
openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 365
```

## Development

### Running Tests

```zsh
# Run tests
./runtests.zsh
```