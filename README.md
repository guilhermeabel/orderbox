# orderbox

This project serves the purpose of exploring the Go programming language and its capabilities, using the book Let's Go, from Alex Edwards, as a guide.
The main premise is creating orders and saving them to a relational database.
Even though the project has a simple premise, it explores several important web development concepts, such as:

## Stateful HTTP handlers

- Session management and authentication.
- Structured and safe state sharing between HTTP requests for the same user.
- Customizing session behaviour, including timeouts and cookie settings.

## Security

- Using a self-signed TLS certificate, with tweaks to the default settings to improve security and performance.
- Setting up the application so that all requests and responses are served securely over HTTPS.
- Connection timeouts on the server to mitigate slow-client attacks.

## Logging

- Using structured logging to log information about the application's requests and responses.
