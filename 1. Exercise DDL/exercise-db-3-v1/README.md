# Exercise SQL

## Alter table

### Description

Terdapat tabel `students` di database sebagai berikut:

| id | fullname | address | gender | day_of_birth | month_of_birth | year_of_birth | grade |
|----|----------|---------|--------|--------------|----------------|---------------|--------|
| 1  | Andi     | Jakarta ... | laki-laki | 12 | 8 | 2001 | 90 |
| 2  | Budi     | Bandung ... | laki-laki | 1 | 1 | 2002 | 80 |
| 3  | Caca     | Surabaya ... | perempuan | 2 | 2 | 2003 | 70 |
| 4  | Deni     | Semarang ... | laki-laki | 3 | 3 | 2004 | 60 |

Dengan spesifikasi kolom sebagai berikut:

- `id` bertipe `INTEGER` dan merupakan primary key
- `fullname` bertipe `VARCHAR(255)` yang tidak boleh `null`
- `address` bertipe `TEXT`
- `gender` bertipe `VARCHAR(50)` yang tidak boleh `null`
- `day_of_birth` bertipe `INTEGER`
- `month_of_birth` bertipe `INTEGER`
- `year_of_birth` bertipe `INTEGER`
- `grade` bertipe `INTEGER`

Ada beberapa perbaikan struktur tabel yang diperlukan, yaitu:

- menghapus kolom `day_of_birth`, `month_of_birth`, `year_of_birth` dan menggantinya dengan kolom baru bernama `date_of_birth` yang memiliki tipe `DATE` dan **tidak boleh `null`**
- menghapus kolom `address` dan menambahnya menjadi beberapa kolom yang lebih spesifik, yaitu:
  - `street` bertipe `VARCHAR(255)`
  - `city` bertipe `VARCHAR(100)`
  - `province` bertipe `VARCHAR(100)`
  - `country` bertipe `VARCHAR(100)`
  - `postal_code` bertipe `VARCHAR(50)`
- mengganti `grade` yang awalnya adalah bertipe `INTEGER` menjadi `FLOAT`

> Pastikan urutan penambahan kolom sesuai dengan urutan di atas.

Buatlah perubahan tersebut di database menggunakan _SQL command_ pada file:

- `add_column_table_students.sql` untuk menambah kolom yang diperlukan
- `drop_column_table_students.sql` untuk menghapus kolom yang tidak dibutuhkan
- `modify_column_table_students.sql` untuk mengubah tipe data kolom yang diperlukan

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 109) dan **`main_test.go`** (line 13) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"Username"`, `"Password"` dan `"DatabaseName"`saja.

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

Jika berhasil maka akan terdapat tabel baru bernama `students` yang kolomnya sudah sesuai dengan spesifikasi yang diberikan.
