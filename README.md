# Autobar Services

This repository is a set of microservices that power the Autobar platform. Each service is designed to handle specific functionalities, ensuring modularity and scalability. In order to preserve common functionality and avoid code duplication, shared libraries are maintained in the `shared-libraries` [repository](https://github.com/autobar-dev/shared-libraries).

They were originally built using SemaphoreCI, because at the time it was the cleanest way to built services in monorepos like this one. However, this should be migrated to GitHub Actions in the future.

If you'd like to run any of these services, look in `docker-compose.yml` for the required env files and dependencies.

If you're interested in using code from this repository in your own projects commercially, please reach out to me on LinkedIn or at the email address found on [my website](https://adampisula.tech).

## Services

### Auth (Golang)

The Auth service manages user and module authentication. It handles request authentication and authorization, as well as user management tasks such as registration or login.

### Currency (Rust)

The Currency service is responsible for currency conversion and exchange rate management. It provides APIs to convert amounts between different currencies and fetch current exchange rates.

### Email (Golang)

The Email service handles sending emails to users.

### Email Template (Node.js)

The Email Template service manages email templates used by the Email service. It uses a React-based templating system to render dynamic email content.

### File (Golang)

The File service manages file uploads, storage, and retrieval from a storage backend like S3. It provides APIs for uploading and deleting files, as well as generating pre-signed URLs for secure access.

### Gateway (nginx)

Conf files for the API Gateway using nginx.

### Module (Golang)

The Module service manages communication and registration of Autobar modules. It handles all operationss related to modules, such as preparation, secure activation, deactivation and status monitoring.

### Product (Golang)

The Product service manages product information, including creation, updating, deletion, and retrieval of product details.

### Realtime (Golang)

The Realtime service handles real-time communication and updates between the Autobar platform and its modules. It uses WebSockets to provide instant data synchronization and notifications. In order to scale efficiently, it uses an AMQP message broker (RabbitMQ) to distribute messages across multiple instances.

### User (Golang)

The User service manages user profiles and related operations, such as profile updates and user preferences.

### Wallet (Golang)

The Wallet service manages user wallets, including balance management, transaction history, and fund transfers. In order to ensure data consistency and integrity, the balance is not kept in the database, but is calculated on-the-fly based on the transaction history.
