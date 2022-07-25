# Golang Template

> Template for golang projects @okp4.

[![version](https://img.shields.io/github/v/release/okp4/template-go?style=for-the-badge&logo=github)](https://github.com/okp4/template-go/releases)
[![lint](https://img.shields.io/github/workflow/status/okp4/template-go/Lint?label=lint&style=for-the-badge&logo=github)](https://github.com/okp4/template-go/actions/workflows/lint.yml)
[![build](https://img.shields.io/github/workflow/status/okp4/template-go/build?label=build&style=for-the-badge&logo=github)](https://github.com/okp4/template-go/actions/workflows/build.yml)
[![test](https://img.shields.io/github/workflow/status/okp4/template-go/test?label=test&style=for-the-badge&logo=github)](https://github.com/okp4/template-go/actions/workflows/test.yml)
[![conventional commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge&logo=conventionalcommits)](https://conventionalcommits.org)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg?style=for-the-badge)](https://opensource.org/licenses/BSD-3-Clause)

## Purpose & Philosophy

This repository holds the template for building Golang projects with a consistent set of standards accros all [OKP4](https://github.com/okp4) projects. We are convinced that the quality of the code depends on clear and consistent coding conventions, with an automated enforcement (CI).

This way, the template promotes:

- the use of [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/), [semantic versioning](https://semver.org/) and [semantic releasing](https://github.com/cycjimmy/semantic-release-action) which automates the whole package release workflow including: determining the next version number, generating the release notes, and publishing the artifacts (project tarball, docker images, etc.)
- unit testing
- linting via [golangci-lint](https://github.com/golangci/golangci-lint)
- a uniformed way of building the project for several platforms via a Makefile using a docker image

## How to use

> 🚨 do not fork this repository as it is a [template repository](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template)

1. Click on [Use this template](https://github.com/okp4/template-go/generate)
2. Give a name to your project
3. Wait until the first run of CI finishes
4. Clone your new project and happy coding!

## Prerequisites

- Be sure you have [Golang](https://go.dev/doc/install) installed.
- [Docker](https://docs.docker.com/engine/install/) as well if you want to use the Makefile.

## Build

```sh
make build
```
