# API Documentation for Personal Finance App

## Table of Contents
- [Overview](#overview)
- [Getting Started](#getting-started)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [Login](#login)
  - [Logout](#logout)
  - [Payment](#payment)
- [Request and Response Format](#request-and-response-format)
- [Error Handling](#error-handling)
- [Example Requests](#example-requests)
- [License](#license)

## Overview
API ini adalah bagian dari aplikasi manajemen keuangan pribadi yang memungkinkan pengguna untuk mengelola pendapatan, pengeluaran, dan kategori. API ini mendukung autentikasi dan menyediakan endpoint untuk login, logout, dan manajemen pembayaran.

## Getting Started
Ikuti langkah-langkah di bawah ini untuk menjalankan API ini secara lokal.

1. **Clone the repository**
   ```bash
   git clone https://github.com/nrrarnn/api-golang.git
2. ## Navigasi Direktori

Berikut adalah struktur direktori proyek ini:
  cd api-golang

3. ## Install dependencies Mengunduh semua dependensi yang diperlukan oleh proyek:
go mod tidy

4. ## Set environment variable untuk JWT_SECRET Mengatur kunci rahasia untuk JWT agar dapat digunakan untuk autentikasi:
export JWT_SECRET=your_api_key

5. ## Jalankan server
go run main.go
