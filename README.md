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
