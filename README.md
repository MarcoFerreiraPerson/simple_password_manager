# simple_password_manager

A lightweight, locally hosted password manager built with **Go (backend)**, **PostgreSQL (database)**, and **React (frontend)**.  
The goal is to create a secure, simple, and private vault for managing passwords — entirely under your control.

---

## Features (Planned)

- Secure account creation with **master password**
- Store, retrieve, update, and delete **encrypted passwords**
- Clean **React + Tailwind** frontend
- Simple **Go REST API** backend
- All password data **encrypted** using AES-256
- Master password never stored or transmitted in plaintext
- Optional **local LLM assistant** with RAG for quick password lookups
- Easy to deploy via **Docker**

---

## Project Roadmap

### Phase 1: Core Backend (Go + PostgreSQL)

Build a REST API with the following endpoints:

**Auth Endpoints:**
- `POST /register`: Create an account with master password
- `POST /login`: Verify master password using HMAC (zero-knowledge)
- `PUT /account/password`: Change master password (re-encrypt vault)
- `DELETE /account`: Delete account and vault

**Vault Endpoints:**
- `GET /vault`: Retrieve all saved credentials
- `POST /vault`: Add a new encrypted password
- `PUT /vault/:id`: Update a stored password
- `DELETE /vault/:id`: Delete a stored password

Security:

- Use **Argon2id** to derive encryption keys from the master password
- Store only **encrypted passwords** (AES-256-GCM)
- Use an HMAC-based **verifier** to authenticate the master password (no plaintext password storage)

---

### Phase 2: Database Schema (PostgreSQL)

#### `users` table

| Field    | Type   | Description                        |
|----------|--------|------------------------------------|
| id       | SERIAL | Primary key                        |
| email    | TEXT   | Unique user email                  |
| salt     | TEXT   | Salt used for key derivation       |
| verifier | TEXT   | HMAC of derived key for login check|

#### `vault` table

| Field              | Type      | Description                              |
|--------------------|-----------|------------------------------------------|
| id                 | SERIAL    | Primary key                              |
| user_id            | INTEGER   | Foreign key to `users`                   |
| name               | TEXT      | Name of the account (e.g. GitHub)        |
| username           | TEXT      | Optional username                        |
| encrypted_password | TEXT      | AES-256 encrypted base64 string          |
| created_at         | TIMESTAMP | Auto timestamp                           |
| updated_at         | TIMESTAMP | Auto timestamp                           |

---

### Phase 3: Frontend (React + TailwindCSS)

- Clean and responsive UI for managing vault entries
- Secure forms for login, registration, and vault actions
- Use browser APIs (e.g. `crypto.subtle`) to support client-side encryption (optional zero-knowledge mode)

---

### Phase 4: Local LLM Assistant (Optional)

- Integrate with a local LLM (e.g. `Ollama`)
- Use RAG to search vault entries using natural language
- Only enable this after vault is unlocked in memory

---

### Phase 5: Deployment

- Create Dockerfiles for backend and frontend
- Add `docker-compose.yml` to run the full stack locally
- Include self-signed TLS certs for HTTPS in development

---

## Security Principles

- Never store or transmit master passwords
- All passwords are encrypted using AES-256-GCM
- Key derivation with **Argon2id + salt**
- Optional client-side zero-knowledge vault access

---
