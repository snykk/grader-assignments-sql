# Exercise DB

## Ecercise Update Database

### Description

Terdapat table `students` di database dengan kolom:

- `id` (Integer)
- `first_name` (varchar(100))
- `last_name` (varchar(100))
- `date_of_birth` (date)
- `address` (varchar(255))
- `class` (varchar(100))
- `status` (varchar(100))

Sebelumnya kalian sudah memasukkan data ke dalam table tersebut. Berikut data dalam bentuk tabel:

| id | first_name | last_name | date_of_birth | address | class | status |
|----|------------|-----------|---------------|---------|-------|--------|
| 1  | Abdi       | Doe       | 2003-12-01    | Jakarta | 1A    | active |
| 2  | Jane       | Doe       | 2004-02-01    | Jakarta | 1A    | active |
| 3  | Bernard       | Smith     | 2004-02-01    | Jakarta | 1A    | active |
| 4  | Jane       | Smith     | 2003-12-02    | Jakarta | 1B    | active |
| 5  | Andrew       | Doe       | 2004-07-04    | Jakarta | 1B    | inactive |
| 6  | Rendy       | Doe       | 2004-06-10    | Jakarta | 1B    | inactive |
| 7  | John       | Smith     | 2004-05-11    | Jakarta | 1B    | inactive |
| 8  | Herry       | Smith     | 2004-04-12    | - | 1B    | active |
| 9  | John       | William       | 2004-03-20    | - | 1B    | active |
| 10 | Wendy      | Doe       | 2004-02-21    | - | 1B    | active |

Sekarang, kalian diminta untuk melakukan update data pada table `students` yang belum mengisi kolom `address` (data dalam bentuk `null`) dan `status` berisi `active`. Sementara kolom tersebut diisi dengan `Bandung`.

Buatlah _query_ yang dapat melakukan _update_ data tersebut ke dalam file `update.sql`. Pastikan teks yang dimasukkan sama persis seperti tabel di atas (tidak terdapat perbedaan huruf kecil dan besar)

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 75) dan **`main_test.go`** (line 40) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"Username"`, `"Password"` dan `"DatabaseName"`saja.

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

Jika berhasil maka data di kolom `address` yang memiliki kondisi di atas akan berubah menjadi `Jakarta`.
