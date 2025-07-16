# HCM 03 Go Tung - Food & Drinks App

A Go-based web application for food and drinks management with MySQL database.

## Project Structure

```
hcm_03_go_tung/
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Go dependencies checksum
├── .env                 # Environment variables (create this file)
├── .gitignore           # Git ignore rules
├── README.md            # Project documentation
├── config/              # Configuration files
│   └── database.go      # Database connection setup
├── models/              # Data models (GORM)
│   ├── user.go
│   ├── product.go
│   ├── cart.go
│   ├── order.go
│   ├── order_item.go
│   ├── product_review.go
│   ├── product_suggestion.go
│   └── migrate.go       # Database migration
├── routes/              # API routes definition
├── controller/          # HTTP handlers
├── services/            # Business logic layer
├── repository/          # Data access layer
├── migrations/          # Database migration files
└── assets/              # Static assets
```

## Database Models

- **User**: User management (id, name, email, role, created_time)
- **Product**: Product catalog (id, name, type, category, price, image_url, rating, created_at)
- **Cart**: Shopping cart (id, user_id, product_id, quantity)
- **Order**: Order management (id, user_id, status, created_at)
- **OrderItem**: Order details (order_id, product_id, quantity, price)
- **ProductReview**: Product reviews (id, user_id, product_id, rating, comment, created_at)
- **ProductSuggestion**: Product suggestions (id, user_id, product_id, created_at)

## Setup Instructions

### 1. Prerequisites
- Go 1.24.4 or higher
- MySQL 8.0 or higher

### 2. Database Setup
```sql
CREATE DATABASE hcm_03_go_tung CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'go_user'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON hcm_03_go_tung.* TO 'go_user'@'localhost';
FLUSH PRIVILEGES;
```

### 3. Environment Configuration
Create `.env` file in the root directory:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=go_user
DB_PASSWORD=your_password
DB_NAME=hcm_03_go_tung
```

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Run Application
```bash
go run main.go
```

The application will automatically create database tables using GORM migrations.

## Technology Stack

- **Backend**: Go (Gin framework)
- **Database**: MySQL
- **ORM**: GORM
- **Authentication**: JWT (planned)
- **API**: RESTful

## Development

### Project Architecture
- **Controller Layer**: HTTP request/response handling
- **Service Layer**: Business logic
- **Repository Layer**: Data access
- **Model Layer**: Data structures

### Adding New Features
1. Create model in `models/`
2. Add repository in `repository/`
3. Implement service in `services/`
4. Create controller in `controller/`
5. Define routes in `routes/`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the MIT License.
