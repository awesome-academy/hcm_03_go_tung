# Hướng dẫn cấu hình database và khởi tạo project

1. Tạo file `.env` ở thư mục gốc với nội dung:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=go_user
DB_PASSWORD=your_password
DB_NAME=hcm_03_go_tung
```

2. Cài đặt các package cần thiết:

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
```

3. Đảm bảo đã tạo database và user trên MySQL:

```
CREATE DATABASE hcm_03_go_tung CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'go_user'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON hcm_03_go_tung.* TO 'go_user'@'localhost';
FLUSH PRIVILEGES;
```

4. Hàm `ConnectDatabase()` đã được tạo ở `config/database.go`. Hãy gọi hàm này ở hàm `main()` để khởi tạo kết nối database.
