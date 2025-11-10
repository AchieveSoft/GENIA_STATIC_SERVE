# 🌐 Genia Static File Server

This project is a lightweight static file server written in Go. It serves files from a specified directory and supports custom index and 404 pages via a JSON configuration file.

## 📦 Features

- Serve static files from a configurable directory
- Customizable index and not-found fallback pages
- Simple configuration via `genia-static-serve.json`
- Minimal dependencies, easy to deploy

## 🛠 Configuration

Create a `genia-static-serve.json` file in the root directory with the following structure:

```json
{
  "port": 8080,
  "path": "./public",
  "indexFile": "index.html",
  "notfoundFile": "404.html"
}
