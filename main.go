package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    html := `
    <html>
    <head>
        <title>Go App</title>
        <style>
            body {
                background-color: #f0f4f8;
                font-family: Arial, sans-serif;
                color: #333;
                text-align: center;
                padding-top: 100px;
            }
            .message {
                font-size: 24px;
                font-weight: bold;
                color: #0066cc;
            }
        </style>
    </head>
    <body>
        <div class="message">Welcome to the Go App!</div>
    </body>
    </html>
    `
    fmt.Fprint(w, html)
}

func main() {
    http.HandleFunc("/", homePage)
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

