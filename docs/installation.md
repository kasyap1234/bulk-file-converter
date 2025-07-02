# Bulk File Converter

A tool for bulk file conversions, including PDF to image and image to PDF.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [System Dependencies](#system-dependencies)
  - [Go Dependencies](#go-dependencies)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Convert PDF files to images (PNG, JPEG, etc.)
- Convert images to PDF
- Batch processing

## Installation

### System Dependencies

#### Ubuntu/Debian

```bash
sudo apt update
sudo apt install -y poppler-utils
```

#### macOS

Using Homebrew:

```bash
brew install poppler
```

#### Windows

- Download Poppler for Windows from [http://blog.alivate.com.au/poppler-windows/](http://blog.alivate.com.au/poppler-windows/)
- Extract and add the `bin` directory to your `PATH` environment variable.

### Go Dependencies

Install Go (version 1.24+ recommended): [https://go.dev/dl/](https://go.dev/dl/)

Install project dependencies:

```bash
go mod download
```

#### Go Libraries for Image to PDF Conversion

This project uses the following Go libraries for image and PDF processing:

- [`github.com/signintech/gopdf`](https://github.com/signintech/gopdf) — for creating PDFs from images.
- [`github.com/phpdave11/gofpdf`](https://github.com/phpdave11/gofpdf) — alternative PDF generation library.
- [`github.com/disintegration/imaging`](https://github.com/disintegration/imaging) — for image manipulation.

Install them with:

```bash
go get github.com/signintech/gopdf
go get github.com/phpdave11/gofpdf
go get github.com/disintegration/imaging
```

## Usage

1. Ensure all dependencies are installed.
2. Build the project:

   ```bash
   go build -o bulk-file-converter
   ```

3. Run the converter:

   ```bash
   ./bulk-file-converter
   ```

## Contributing

See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## License

MIT License. See [LICENSE](LICENSE).