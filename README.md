# orderbox

This project serves the purpose of exploring the Go programming language and its capabilities, using the book Let's Go, from Alex Edwards, as a guide.
The main premise is creating orders and saving them to a relational database.
Even though the project has a simple premise, it explores several important web development concepts.

## User Authentication

- Password hashing using the bcrypt algorithm.
- User authentication using a custom middleware.
- User registration and login.

## Stateful HTTP handlers

- Session management and authentication.
- Structured and safe state sharing between HTTP requests for the same user.
- Customizing session behaviour, including timeouts and cookie settings.

## Security

- Using a self-signed TLS certificate, with tweaks to the default settings to improve security and performance.
- Setting up the application so that all requests and responses are served securely over HTTPS.
- Connection timeouts on the server to mitigate slow-client attacks.
- Protection against CSRF attacks using a custom middleware.
- Password authentication using bcrypt.

## Testing

- Writing tests for the application's handlers and services.
- Unit, integration and end-to-end tests.

## Database

- Interacting with a relational database using the `database/sql` package.

## More

- Embedded files, and templates
- Generics and type assertions
- Logging and Error Handling
- UI development using HTML, CSS and JavaScript
