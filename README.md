# Dylets-Survey-Form


For handling the form submission, we utilise the (Gorilla Mux package)[https://github.com/gorilla/mux] web framework for routing and net/http for handling HTTP requests. A briefe searche through the internet led to the conclution that it is a powerful HTTP router and URL matcher for building Go web servers.

## Error Handling

We check the HTTP method and return an error if it's not a POST request.
* Errors during form parsing and JSON encoding are logged and appropriate HTTP error responses are returned.
* The server is started using http.ListenAndServe, and log.Fatal is used to log fatal errors and exit the application.