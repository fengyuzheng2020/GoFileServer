# GoFileServer
GoFileServer is a lightweight and customizable file server written in Go, providing a simple and secure way to upload and download files. It includes features such as configurable upload directories, support for various file types, and straightforward file management.

Features:

File Upload: Easily upload files via HTTP POST requests.
File Download: Download files from the server with proper content disposition.
Configurability: Customize upload directories and server port via command-line parameters.
Minimalistic Design: Keep it simple and easy to understand, making it a great starting point for building file-related applications.
Usage:

Upload files with a simple HTTP POST request.
Download files with a GET request to /download/filename.

Getting Started:
```sh
# Clone the repository
git clone https://github.com/yourusername/GoFileServer.git

# Navigate to the project directory
cd GoFileServer

# Run the server with default settings
go run main.go

# Customize upload directory and port (optional)
go run main.go -upload-dir=/path/to/upload/dir/ -port=8080

```
Explore the potential of GoFileServer for your file management needs!
