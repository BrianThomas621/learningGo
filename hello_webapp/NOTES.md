# Our First Web Application!

First, make a distinction between web application, web server, and web service: 

* *Web Server*: a program that uses HTTP to serve files that form web pages to users in response to their requests. These files are stored in a directory (such as `docroot`), and are forwarded by their computers' HTTP clients. 

* *Web Application:* not only returns files but processes requests and performs operations by responding to HTTP requests and returns HTML to the client in an HTTP response message. 

* *Web Service:* don't render HTML to a user but instead returns data in any other format to another program. 

Everything revolves around the HTTP protocol, which is a request-response protocol. These requests include GET, HEAD, POST, PUT, and some others. HTTP responses are sent every time there is a request. These are status codes such as 1XX, 2XX, 4XX, etc...

## Parts of a Web Application 

1. *Handler:* receives and processes the HTTP request that is sent from the client (web browser), and bundles data into the HTTP response to be sent back to the client. 
2. *Templates:* code that is convered to HTML that is sent back to the client as an HTTP response message. Templates can be partly HTML or not at all. Template engines generate the final HTML with templates and data. 

## Code Notes: 

A _handler function_ in Go is a call-back function triggered by an event. The hanlder function here is called _handler_ and takes two arguments: 

1. `http.ResponseWriter`: an _interface_ used by an HTTP handler to construct an HTTP response.
2. `*http.Request`: a pointer to a Request _structure_, which is an HTTP request received by a server or to be sent by a client.

The `main()` function calls two functions, also from the `net/http` package: 

* `http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))`: HandleFunc registers the handler function for the given pattern.
* `http.ListenAndServe(addr string, handler Handler) error`: starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux: 

 