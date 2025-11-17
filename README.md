# go-gist

`go-gist` is a lightweight command-line utility for temporary file and directory sharing using object storage (such as Amazon S3 or any S3-compatible backend). It allows users to upload any file or directory for a specified duration and retrieve it later using a short, human-readable code.

The tool automatically compresses directories, uploads the resulting archive to object storage, sets an expiration based on the user-provided duration, and provides a simple installation mechanism for clients using the fetch command.

---

## Features

- Upload any file or directory to object storage.
- Automatic ZIP compression for directories.
- Time-limited availability with automatic expiry.
- Human-readable access codes for downloads.
- Simple CLI interface with symmetric push/fetch operations.
- Extensible design with pluggable storage backends.

---

## Installation

Build the executable from source:

```bash
git clone https://github.com/yourname/go-gist
cd go-gist
go build -o go-gist ./cmd/go-gist
```

Move it into your PATH if desired:

```bash
mv go-gist /usr/local/bin/
```

---

## Usage

`go-gist` provides two primary subcommands: `push` and `fetch`.

### Push

Uploads a file or directory for a specified duration.

```bash
go-gist push <path> <duration>
```

**Example**

```bash
go-gist push ./project 5h
```

This command:

1. Detects whether the path is a file or directory.
2. Compresses the content into a ZIP archive.
3. Uploads the archive to the configured object storage backend.
4. Stores metadata such as expiration time.
5. Outputs a code such as `simple-fish-69` that identifies the upload.

---

### Fetch

Downloads and extracts a previously uploaded file or directory using its code.

```bash
go-gist fetch <code>
```

**Example**

```bash
go-gist fetch simple-fish-69
```

This command:

1. Retrieves metadata by the code.
2. Downloads the associated ZIP archive from object storage.
3. Extracts it into the current working directory (or a user-specified location).

---

## Configuration

The project is designed to be storage-agnostic. Object storage settings such as credentials, bucket names, and endpoints will be configurable through environment variables or a configuration file. Support for S3, MinIO, and Cloudflare R2 will be included.

A typical configuration will include:

- Access key and secret key
- Default bucket name
- Region or endpoint
- Default expiration policy
- Local paths for temporary ZIP files
- Metadata storage backend (BoltDB or remote storage)

Detailed configuration documentation will be provided as storage backends are implemented.

---

## Project Structure

```
go-gist/
  cmd/
    go-gist/
      main.go
  internal/
    cli/
      cli.go
    storage/
      ...
    compression/
      ...
    codegen/
      ...
    metadata/
      ...
  go.mod
  README.md
```

- `cmd/go-gist`: Entry point for the command-line binary.
- `internal/cli`: Command implementations for push/fetch.
- `internal/compression`: ZIP and extraction utilities.
- `internal/storage`: Backend implementations for S3 or compatible services.
- `internal/metadata`: Local or remote metadata storage.
- `internal/codegen`: Human-readable code generator module.

---

## Roadmap

- Implement ZIP compression for files and directories.
- Add S3-compatible storage backend.
- Implement metadata persistence using BoltDB.
- Add code generation utilities.
- Add extraction logic for fetch command.
- Add configuration support.
- Add automated tests and integration tests.
- Provide precompiled binaries and installation scripts.

---

## License

MIT License. See `LICENSE` file for details.

---

## Contributing

Contributions are welcome. Please open an issue for discussion before submitting a pull request.
