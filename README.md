[![codecov](https://codecov.io/github/hi2gage/groupinary-graph/graph/badge.svg?token=OIPVJ3RZHG)](https://codecov.io/github/hi2gage/groupinary-graph)

# Groupinary GraphQL Service

### What is Groupinary?
Think of it as a private urban dictionary tailored to capture and preserve the unique lexicon and language quirks of your social circles.

It's more than just a dictionary; it's a linguistic time capsule, preserving the inside jokes, memes, and colloquial expressions that define the shared experience within your group.

Groupinary allows users to log and organize words, phrases, and expressions specific to their social circles. Whether it's a hilarious inside joke, a memorable catchphrase, or a unique term coined within the group, Groupinary ensures that nothing is forgotten.

### What is this repo?
This repository contains the backend service for Groupinary. It's written in Go and uses Ent. The backend powers the GraphQL API


### Why?
After my friends moved away after college and I wanted a project to help us keep in touch because I miss them.


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