# Flutter App Upgrade Local Server

# Steps

1. Clone the repository to your local machine.

2. Make sure you have Go installed on your machine. You can download it from the official website: https://golang.org/dl/.

3. Navigate to the project directory in your terminal.

4. Run go run main.go to start the server.

The server will be running on http://localhost:8000.
You can use the following endpoints:

* GET /appcast/:version to retrieve a JSON file based on the version parameter.
* POST /upload to upload one or more files to the server.
* GET /download/:filename to download a file from the server based on the filename parameter.

To stop the server, press Ctrl+C in your terminal.

## Summary

| URL                  | Method | Description |
|:---|---|---|
| /appcast/:version    | Get    | fetch the appcast json file which named {version}, you should put the config file in the src/files/ first
| /upload              | Post   | upload file to server, storage in src/files/
| /download/:filename  | Get    | download file named {filename}, the file should storage in src/files/
