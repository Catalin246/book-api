# Go API

## Set up the project 

## 1. Install go
## 2. Configuring the environment variables

`export GOPATH=$HOME/go`
`export PATH=$PATH:$GOPATH/bin`

## 3. Setting up a workspace

`$ mkdir ~/go-workspace`
`$ cd ~/go-workspace`
`$ mkdir src pkg bin`

## 4. Building the API

### 4.1 Choosing a framework (e.g., Gorilla mux)

### 4.2 Creating a basic server

### 4.2.1 Create Your Project Directory

Navigate to the src directory in your workspace.

`cd ~/go-workspace/src`

reate a new directory for your API project.

`mkdir book-api`
`cd book-api`

### 4.2.2 Initialize a Go Module

`go mod init book-api`


### 4.2.3 Create Your Server Code

`New-Item -ItemType File -Name "main.go" -Force`

### 4.2.4 Run Your Server

Make sure you're in the book-api directory.

`go run main.go`

Open a web browser or a tool like curl or Postman, and navigate to: http://localhost:8000/