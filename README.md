# GoZix Boilerplate

[documentation-img]: https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff
[documentation-url]: https://pkg.go.dev/github.com/gozix/boilerplate/v3
[license-img]: https://img.shields.io/github/license/gozix/boilerplate.svg?style=for-the-badge
[license-url]: https://github.com/gozix/boilerplate/blob/master/LICENSE
[release-img]: https://img.shields.io/github/tag/gozix/boilerplate.svg?label=release&color=24B898&logo=github&style=for-the-badge
[release-url]: https://github.com/gozix/boilerplate/releases/latest
[build-status-img]: https://img.shields.io/github/actions/workflow/status/gozix/boilerplate/go.yml?logo=github&style=for-the-badge
[build-status-url]: https://github.com/gozix/boilerplate/actions
[go-report-img]: https://img.shields.io/badge/go%20report-A%2B-green?style=for-the-badge
[go-report-url]: https://goreportcard.com/report/github.com/gozix/boilerplate
[code-coverage-img]: https://img.shields.io/codecov/c/github/gozix/boilerplate.svg?style=for-the-badge&logo=codecov
[code-coverage-url]: https://codecov.io/gh/gozix/boilerplate

[![License][license-img]][license-url]
[![Documentation][documentation-img]][documentation-url]

[![Release][release-img]][release-url]
[![Build Status][build-status-img]][build-status-url]
[![Go Report Card][go-report-img]][go-report-url]
[![Code Coverage][code-coverage-img]][code-coverage-url]

Example of simple GoZix application.

## Installation

```shell
git clone https://github.com/gozix/boilerplate.git
```

## Project Structure

```text
.
├── .bin                // Resources and binary components
├── cmd                 // Project components
│   └── app             // Executable component
│       ├── internal    // Component libraries
│       └── main.go     // Component entry point
├── internal            // Project common libraries
├── vendor              // Project dependencies
├── Gopkg.lock          // Dep lock file
├── Gopkg.toml          // Dep manifest 
├── Makefile            // Project make file
└── README.md           // Project readme
```

## Documentation

You can find documentation on [pkg.go.dev][documentation-url] and read source code if needed.

## Questions

If you have any questions, feel free to create an issue.

