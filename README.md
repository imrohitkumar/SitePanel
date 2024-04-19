#SitePanel: A Web Hosting Panel for FreeBSD
Work in Progress

SitePanel is a web hosting panel built using Go, PostgreSQL, and Gorilla web toolkit. It's designed to provide a simple and intuitive interface for managing websites, domains, subdomains, FTP accounts, MySQL databases, backups, and SSL certificates.

Features

User management with password hashing and salting
Domain management with support for subdomains
FTP account management with password hashing and salting
MySQL database management with support for multiple databases
Backup management with support for scheduling backups
SSL certificate management with support for Let's Encrypt
Screenshots

Coming soon!

Getting Started

To get started with SitePanel, follow these steps:

Clone the repository: git clone https://github.com/imrohitkumar/SitePanel.git
Install Go and PostgreSQL on your system
Create a PostgreSQL database and user
Set the DATABASE_URL environment variable
Build and run the application: go build main.go &&./main
Access the application at http://localhost:8080
Contributing

SitePanel is a work in progress, and we welcome contributions from the community. If you're interested in helping out, please fork the repository and submit a pull request.

License

SitePanel is licensed under the MIT License.

Acknowledgments

We'd like to thank the Go and PostgreSQL communities for their support and resources.

Note

SitePanel is a work in progress, and some features may not be fully implemented or tested. We appreciate your patience and feedback as we continue to develop and improve the application.
