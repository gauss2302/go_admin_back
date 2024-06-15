# Administation System Implemented in Go


# General

This Golang-based admin application offers a user-friendly interface for managing administrative tasks. Designed to be modular, extensible, and easy to deploy, it includes features such as user management, role-based access control, and various other admin functionalities.

## Features

- **User Management**: Easily manage user accounts, including creating, deleting, and modifying them.
- **Role-Based Access Control (RBAC)**: Define roles and permissions to control access to different features and sections of the admin application.
- **Modular Design**: Built with modularity in mind, allowing for easy extension and customization of functionality.
- **RESTful API**: Exposes a RESTful API for seamless integration with other services or front-end applications.
- **Authentication and Authorization**: Provides secure mechanisms for user authentication and authorization to ensure data integrity and confidentiality.

## ⚙️ Build Instructions

**📂 Clone the repository:**
   ```sh
   git clone https://github.com/gauss2302/go_admin_back.git
   cd your-repo

   go get -v ./...

   go run . 
  ```
**🛠️ Build:**
   ```sh
   go build -o go_admin_back
  ```

**▶️ Run the built application:**
   ```sh
   go build -o go_admin_back
  ```

**🧪 Testing:**
   ```sh
   go test ./...
  ```