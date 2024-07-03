# Merge-Solver

Merge-Solver is a toolkit that includes a Go API and a Rust CLI to facilitate automatic merging of code files without manual changes. It prioritizes changes from the newer files while seamlessly integrating content from older versions, ensuring updates are incorporated and compatibility is maintained.

## Features

- **API (Go):** Service to perform code file merging.
- **CLI (Rust):** Interacts with the API to send code files and receive the merged result.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, make sure you have the following installed:
- [Go](https://golang.org/doc/install) (version 1.22 or higher)
- [Rust](https://www.rust-lang.org/tools/install)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Setup

Create a `.env` file in the `api` directory with the following content:

```py
OPENAI_SECRET_KEY="secret_key"
```

Replace with your actual OpenAI API key.

### Running Locally

1. **Setting up the Go API**

   Navigate to the API directory and build the service:

   ```bash
   cd api/
   go run cmd/app/main.go # or just `air` to run with hot reload
   ```

2. **Setting up the Rust CLI**

   Navigate to the CLI directory and build the executable:

   ```bash
   cd cli/
   cargo build --release
   ./target/release/cli --old path/to/old_file.rs --new path/to/new_file.rs
   ```

### Using the CLI

With the API running, you can use the CLI to merge two files:

```bash
./target/release/file_diff --old path/to/old_file.rs --new path/to/new_file.rs
```

## How It Works

- The **API** has a POST `/merge` endpoint that accepts a JSON object with two fields: `old` and `new`, representing the old and new code.
- The **CLI** reads the content of the provided files, sends them to the API, and displays the merged result.
