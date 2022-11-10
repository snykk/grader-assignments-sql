# Exercise SQL

## Create Table

### Description

Setelah kalian belajar tentang `CREATE TABLE`, sekarang saatnya untuk membuat tabel baru. Tabel yang akan kalian buat adalah tabel bernama `persons`, yang memiliki kolom sebagai herikut:

- `id` bertipe `INTEGER` dan merupakan _primary key_
- `NIK` bertipe `VARCHAR(255)` yang tidak boleh `null` dan harus unik, kalian bisa menggunakan _constraints_ `UNIQUE`.
- `fullname` bertipe `VARCHAR(255)` yang tidak boleh `null`
- `gender` bertipe `VARCHAR(50)` yang tidak boleh `null`
- `birth_date` bertipe `DATE` yang tidak boleh `null`
- `is_married` bertipe `BOOLEAN`
- `height` beripe `FLOAT`
- `weight` bertipe `FLOAT`
- `address` bertipe `TEXT`

Buatlah tabel `persons` dengan spesifikasi di atas secara berurutan ke database pada file `create-table.sql` dengan menggunakan _SQL command_ yang sudah dijelaskan di materi.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 58) dan **`main_test.go`** (line 24) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

Contoh:

```go
dbCredentials = Credential{
    Host:         "localhost",
    Username:     "postgres", // <- ubah ini
    Password:     "postgres", // <- ubah ini
    DatabaseName: "kampusmerdeka", // <- ubah ini
    Port:         5432,
}
```

### Expected Result

Jalankan perintah `grader-cli test` untuk mengecek apakah jawaban kalian sudah benar atau belum.

Jika berhasil maka akan terdapat tabel baru bernama `persons` yang kolomnya sudah sesuai dengan spesifikasi di atas.
