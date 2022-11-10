# Exercise DB

## Ecercise Query

### Description

Di dalam database sekolah terdapat tabel `reports` dengan kolom:

- `id` bertipe `interger` yang merupakan _primary key_
- `first_name` bertipe `varchar(255)`, yang berisi nama depan siswa
- `last_name` bertipe `varchar(255)`, yang berisi nama belakang siswa
- `student_class` `varchar(100)`, berisi kelas siswa
- `final_score` bertipe `integer`, berisi nilai akhir siswa
- `absent` bertipe `integer`, berisi jumlah ketidakhadiran siswa

Kalian diminta untuk membuat query untuk menampilkan data murid yang **tidak lulus**. Murid yang tidak lulus adalah murid yang memiliki nilai akhir dibawah 70 atau jumlah ketidak hadirannya lebih dari 5 kali.

Buatlah _query_ untuk menampilkan data murid yang tidak lulus dengan format tabel `reports` seperti berikut:

| id | student_name | student_class | final_score | absent |
|----|--------------|---------------|-------------|--------|
| 1  | John Doe     | 1A            | 60          | 6      |
| 2  | Jane Doe     | 1A            | 65          | 7      |
|... | ...          | ...           | ...         | ...    |

Kolom `student_name` merupakan gabungan dari `first_name` dan `last_name` yang dipisahkan oleh spasi `" "`. Kerjakan _query_ ini di dalam file `select.sql`.

> hint: Untuk menggabungkan 2 kolom bertipe `varchar`, gunakan fungsi `CONCAT` atau memakai _concatenation_ dengan operator `||`

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 111) dan **`main_test.go`** (line 11) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"Username"`, `"Password"` dan `"DatabaseName"`saja.

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

### Test Case Examples

#### Test Case 1

Jika tabel yang dimasukkan adalah berikut:

| id | first_name | last_name | student_class | final_score | absent |
|----|------------|-----------|---------------|-------------|--------|
| 1  | John       | Doe       | 1A            | 60          | 6      |
| 2  | Rendy       | William       | 1A            | 65          | 7      |
| 3  | Abrams       | Smith       | 1A            | 70          | 5      |
| 4  | Wendy       | Doe       | 1A            | 75          | 4      |

Maka hasil _query_ yang diharapkan adalah:

| id | student_name | student_class | final_score | absent |
|----|--------------|---------------|-------------|--------|
| 1  | John Doe     | 1A            | 60          | 6      |
| 2  | Rendy William     | 1A            | 65          | 7      |
