# Interview Intern Dot Backend

**Aplikasi ini adalah backend REST API untuk manajemen user, project, dan task, yang dapat digunakan untuk kebutuhan interview, magang, atau pembelajaran pengembangan aplikasi backend modern.**

- Pengembang: **Bagus Bimawan Sembodo**

Aplikasi ini menyediakan fitur:
- Registrasi dan login user dengan autentikasi JWT
- Manajemen project (CRUD) yang terhubung ke user
- Manajemen task yang terhubung ke project
- Setiap endpoint yang sensitif dilindungi oleh JWT middleware
- Struktur kode rapi dan scalable dengan Layered Architecture Pattern

## Project Structure & Pattern

Project ini menggunakan **Layered Architecture Pattern** (juga dikenal sebagai Clean Architecture/Onion Architecture secara sederhana) yang terdiri dari beberapa layer utama:

- **Model**: Berisi definisi struktur data (entity) yang merepresentasikan tabel di database.
- **Repository**: Berisi kode akses data ke database (CRUD), terpisah dari logic bisnis.
- **Service**: (Opsional, pada beberapa fitur) Berisi logic bisnis, validasi, dan pengolahan data sebelum/selesai dari repository.
- **Controller**: Berisi handler untuk HTTP request, menerima input dari client, memanggil service/repository, dan mengembalikan response.
- **Middleware**: Berisi logic yang berjalan sebelum/selama request diproses (misal: JWT Auth).
- **Config**: Berisi konfigurasi database dan environment.
- **Routes**: Berisi pengaturan endpoint dan middleware.
- **Utils**: Berisi fungsi-fungsi utilitas (misal: JWT, context helper).

### Alasan Penggunaan Layered Pattern
- **Maintainability**: Kode lebih mudah dipelihara karena setiap layer punya tanggung jawab jelas.
- **Testability**: Setiap layer bisa di-test secara terpisah.
- **Scalability**: Mudah menambah fitur baru tanpa mengganggu bagian lain.
- **Separation of Concerns**: Memisahkan logic bisnis, akses data, dan presentasi (API handler).

## API Documentation

### Authentication
- **POST /register**
  - Request: `{ "name": "string", "email": "string", "password": "string" }`
  - Response: 200 OK / 400 Bad Request

- **POST /login**
  - Request: `{ "email": "string", "password": "string" }`
  - Response: `{ "token": "jwt_token", "user": { ... } }`

### User
- **GET /user/:id**
  - Response: `{ "name": "string", "email": "string" }`

### Project
- **POST /users/:user_id/projects** (JWT required)
  - Request: `{ "title": "string", "description": "string" }`
  - Response: `{ "message": "Project created successfully", "data": { ... } }`

- **GET /users/:user_id/projects** (JWT required)
  - Response: `{ "message": "Projects retrieved successfully", "data": [ ... ] }`

- **GET /projects/:id** (JWT required)
  - Response: `{ "message": "Project retrieved successfully", "data": { ... } }`

- **PUT /projects/:id** (JWT required)
  - Request: `{ "title": "string", "description": "string" }`
  - Response: `{ "message": "project updated successfully" }`

### Task
- **PUT /tasks/:id** (JWT required)
  - Request: `{ "title": "string", "description": "string", "is_done": true }`
  - Response: `{ "message": "task updated successfully" }`

> **Catatan:** Semua endpoint yang membutuhkan autentikasi JWT harus mengirimkan header:
> `Authorization: Bearer <jwt_token>`

## API Client Collection
- Dokumentasi dan contoh request dapat diimport ke Postman menggunakan file koleksi (bisa dibuat dari endpoint di atas).

## Cara Menjalankan
1. Copy `.env-example` ke `.env` dan sesuaikan konfigurasi database & JWT_SECRET.
2. Jalankan perintah:
   ```bash
   go run ./cmd/main.go
   ```
3. Gunakan Postman atau API client lain untuk mencoba endpoint di atas.

---

**Kenapa Layered Pattern?**
Karena pattern ini memudahkan pengembangan, testing, dan scaling aplikasi backend, serta sudah terbukti di banyak project production. 