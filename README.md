# API Documentation

## Table of Contents
- [Overview](#overview)
- [Getting Started](#getting-started)
- [Navigasi Direktori](#navigasi-direktori)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
   - [Login](#login)
     - [Penggunaan Login](#penggunaan-login)
   - [Logout](#logout)
     - [Penggunaan Logout](#penggunaan-logout)
   - [Payment](#payment)
     - [Penggunaan Payment](#penggunaan-payment)

## Overview
API ini dikembangkan untuk memenuhi kebutuhan integrasi antara merchant dan bank. Fungsionalitas utama API ini mencakup login, logout, dan payment bagi pelanggan yang terdaftar. Semua data pelanggan, merchant, dan riwayat transaksi disimpan dalam file JSON sebagai simulasi.

## Getting Started
Ikuti langkah-langkah di bawah ini untuk menjalankan API ini secara lokal.

1. **Clone the repository**
   ```bash
   git clone https://github.com/nrrarnn/api-golang.git

2. **Navigasi Direktori**
  Berikut adalah struktur direktori proyek ini:
    ```
    cd api-golang
    ```

3. **Install dependencies**
    ```
    go mod tidy
    ```

4. **Set environment variable untuk JWT_SECRET**
    ```
    export JWT_SECRET=your_api_key
    ```

5. ## Jalankan server
    ```
    go run main.go
    ```

## Authentication

API ini menggunakan JSON Web Tokens (JWT) untuk autentikasi. Pengguna perlu melakukan login untuk mendapatkan token yang harus disertakan dalam setiap permintaan yang memerlukan autentikasi.

### Login

**Endpoint**: `/login`  
**Method**: `POST`  

**Request Body**:
```json
{
  "name": "your_username",
  "password": "your_password"
}
```

## Endpoints

Berikut adalah daftar endpoint yang tersedia dalam API ini:

### 1. Login

- **Endpoint**: `/login`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
      "name": "your_username",
      "password": "your_password"
    }
    ```
- **Response** (Jika login berhasil):
    ```json
    {
      "message": "Login successful for customer: testuser",
      "token": "your_jwt_token"
    }
    ```
- **Response** (Jika login gagal):
    ```json
    {
      "error": "Invalid credentials"
    }
    ```

### 2. Logout

- **Endpoint**: `/logout`
- **Method**: `POST`
- **Authorization**: Token JWT harus disertakan dalam header.
- **Response** (Jika logout berhasil):
    ```json
    {
      "message": "Logout successful"
    }
    ```

### 3. Payment

- **Endpoint**: `/payment`
- **Method**: `POST`
- **Authorization**: Token JWT harus disertakan dalam header.
- **Request Body**:
    ```json
    {
      "customer_id": "1",
      "amount": 100,
      "merchant_id": "merchants_1"
    }
    ```
- **Response** (Jika pembayaran berhasil):
    ```json
    {
      "message": "Payment successful"
    }
    ```

### Catatan

- Semua endpoint yang memerlukan autentikasi harus menyertakan token JWT dalam header `Authorization`.
- Token memiliki masa berlaku tertentu. Pastikan untuk melakukan login kembali setelah token kadaluarsa.
- Pastikan untuk menggunakan `Content-Type: application/json` untuk semua permintaan yang mengirimkan data JSON.

