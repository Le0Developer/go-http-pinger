# go-http-pinger

This is a simple HTTP pinger written in Go. It sends a GET request to a given URL and prints the response status code every 15 seconds.

## Why?

Running this on my Raspberry Pi.
The remote server will open an Instatus incident if it does not receive a request for 30 seconds.

## Usage

Just buiold & run.

## Environment variables

- `URL`: The URL to ping.
- `API_TOKEN`: Your API token


