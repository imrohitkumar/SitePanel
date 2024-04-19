# SitePanel: A FreeBSD Hosting Panel
=====================================

## Features

* User management with password hashing and salting
* Domain management with support for subdomains
* FTP account management with password hashing and salting
* MySQL database management with support for multiple databases
* Backup management with support for scheduling backups
* SSL certificate management with support for Let's Encrypt

### Getting Started

To get started with SitePanel, follow these steps:

1. Clone the repository: `git clone https://github.com/imrohitkumar/SitePanel.git`
2. Install Go and PostgreSQL on your system
3. Create a PostgreSQL database and user
4. Set the `DATABASE_URL` environment variable
5. Build and run the application: `go build main.go &&./main`
6. Access the application at `http://localhost:8080`
