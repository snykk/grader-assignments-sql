# Assignment DB 1

## Assignment Data Definition Language

### Description

Perusahaan PT CAMP adalah perusahaan jasa yang bergerak di bidang IT. PT CAMP membuat database untuk kebutuhan internal yang dipakai untuk mencatat kehadiran karyawan. Untuk saat ini terdapat table `users` yang berisi data karyawan dan akses _login_ mereka dan juga terdapat tabel `attendances` yang berisi data kehadiran karyawan.

Tabel `users` berisi kolom sebagai berikut:

- `id` bertipe `int` tidak boleh `null`
- `fullname` bertipe `varchar(255)` tidak boleh `null`
- `email` bertipe `varchar(255)` tidak boleh `null`
- `password` bertipe `varchar(255)` tidak boleh `null`
- `role` bertipe `varchar(100)`

Tabel `attendances` berisi kolom sebagai berikut:

- `id` bertipe `int` tidak boleh `null`
- `user_id` bertipe `int` tidak boleh `null`
- `status` bertipe `varchar(100)` tidak boleh `null`

Perusahaan ini mengubah beberapa bentuk tabel agar lebih informatif dan mudah untuk diolah. Kalian diminta untuk :

1. Menghapus kolom `role` di tabel `users` dan menambah kolom sebagai berikut:

    - `phone` bertipe `varchar(50)` boleh diisi data kosong / `null`
    - `address` bertipe `varchar(255)` boleh diisi data kosong / `null`
    - `department` bertipe `varchar(255)` boleh diisi data kosong / `null`
    - `division` bertipe `varchar(255)` boleh diisi data kosong / `null`
    - `position` bertipe `varchar(255)` boleh diisi data kosong / `null`

2. Menghapus tabel `attendances`

3. Membuat tabel `presences` sebagai ganti dari tabel `attendances` dengan kolom sebagai berikut:

    - `id` bertipe `INT` tidak boleh `null`
    - `user_id` bertipe `INT` tidak boleh `null`
    - `presence_date` bertipe `DATE` tidak boleh `null`
    - `status` bertipe `VARCHAR(50)` tidak boleh `null`
    - `location` bertipe `VARCHAR(255)` boleh diisi data kosong / `null`
    - `description` bertipe `VARCHAR(255)` boleh diisi data kosong / `null`
    - `image_presence` bertipe `VARCHAR(255)` boleh diisi data kosong / `null`
    - `image_location` bertipe `VARCHAR(255)` boleh diisi data kosong / `null`

Kerjakan tugas nomor 1 di file `add_column_users.sql` untuk menambah kolom dan file `drop_column_users.sql` untuk menghapus kolom. Kerjakan tugas nomor 2 di file `drop_table_attendances.sql` dan tugas nomor 3 di file `create_table_presences.sql`.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 118) dan **`main_test.go`** (line 11) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

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

Jika berhasil maka akan terdapat tabel `users` dengan kolom sebagai berikut:

- `id` bertipe `int` tidak boleh kosong
- `fullname` bertipe `varchar(255)` tidak boleh kosong
- `email` bertipe `varchar(255)` tidak boleh kosong
- `password` bertipe `varchar(255)` tidak boleh kosong
- `phone` bertipe `varchar(50)` boleh diisi data kosong
- `address` bertipe `varchar(255)` boleh diisi data kosong
- `department` bertipe `varchar(255)` boleh diisi data kosong
- `division` bertipe `varchar(255)` boleh diisi data kosong
- `position` bertipe `varchar(255)` boleh diisi data kosong

Table `attendance` tidak ada, dan sudah diganti dengan tabel `presence`.
