# DbAdmin

**DbAdmin** is an interactive command-line tool for managing database users and privileges with fine-grained control. It’s built with Go and designed for DevOps engineers, DBAs, and developers who need a fast, secure, and scriptable interface for handling database access or those who are too lazy to run SQL queries 😅😅😅.

## Features

- ✅ Create and delete users and roles
- 🔒 Grant readonly or read/write privileges
- 🚫 Revoke all privileges from a user
- 🎭 Assign roles to users
- 🧵 Interactive CLI mode with prompt-based input
- 🔍 Query user or table privileges
- 🎯 Restrict access to a single database
- 🛡 Secure password generation and validation

## Installation

```bash
   git clone https://github.com/gentcod/dbadmin.git
   cd dbadmin
   make build
```

## Usage

```bash
   dbadmin -d postgres -D postgres://user:password@hostname:port/db_name
```

> NB: Connected database user should be a user with superadmin authorization.