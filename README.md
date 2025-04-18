# Daylick Backend

This is the backend of **Daylick** — a flexible planning and productivity platform.  
Built with **Go**, **PostgreSQL**, and structured around a template-driven architecture, it supports autosaving, flexible data tracking, and real-time statistics.

## 🧠 Features

- User registration and authentication
- Workspace and role-based access system
- Template and type-driven data modeling
- Daily reports and long-term plans with flexible timelines
- Data aggregation for simple statistics

## 📂 Project Structure

daylick-backend/ 
  ├── controllers # Request handlers 
  ├── db # Database setup and migrations 
  ├── docs # Project documentation 
  ├── middlewares # Session/auth/etc. 
  ├── models # Data models 
  ├── repository # Data access layer(communicate with the database)
  ├── routes # Endpoint registration 
  ├── services # Business logic 
  ├── utils # Helpers

## ▶️ Running the Backend

```bash
go run main.go