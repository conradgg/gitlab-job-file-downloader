# GitLab Job File Downloader

[![Build & Deploy](https://github.com/conradgg/gitlab-job-file-downloader/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/conradgg/gitlab-job-file-downloader/actions/workflows/main.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/conradgg/gitlab-job-file-downloader/badge)](https://www.codefactor.io/repository/github/conradgg/gitlab-job-file-downloader)

GitLab Job File Downloader is a simple Job downloader from gitlab with hooks.

## Getting started

### Prerequisites
***
- **[Go](https://go.dev/)**: any one of the **two latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

### How to install GitLab Job File Downloader
***
- Download the latest version
- Install dependencies

```bash
    $ go get . && go mod tidy
```
- Build the project

```bash
    $ go build
```

### Usage
***
1. Download the latest version and install dependencies.
2. Add to gonfig.go:
    - Project ID on the project page. Example: 44066311.
    - Open the file path in Job and paste everything after `/artifacts/file/`
        - Example URL:
    `https://gitlab.com/conradgg/testCI/-/jobs/3876871442/artifacts/file/build/libs/testCI.jar`
        - Paste `build/libs/testCI.jar`
    - Create a new token in the Project Folder -> Settings -> Access Tokens
        - Create a new token with API scope and Maintainer role.
    - Create a new HTTP Webhook with the URL:
        - `http://YourIP(Domain):port/ListenPath`
	    - Add a Job events trigger and disable SSL Verification.
3. Build and run the project.
4. Test webhooks
