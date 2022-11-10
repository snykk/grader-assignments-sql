# Exercise DB

## Ecercise Query 2

### Description

Terdapat table `people` yang dimiliki oleh dinas daerah yang memiliki kolom sebagai berikut:

- `id` bertipe `INT` yang merupakan _primary key_
- `NIK` bertipe `VARCHAR(255)` yang berisi nomer induk kependudukan (data unik)
- `first_name` bertipe `VARCHAR(150)` yang berisi nama depan penduduk
- `last_name` bertipe `VARCHAR(150)` yang berisi nama belakang penduduk
- `gender` bertipe `VARCHAR(50)` yang berisi jenis kelamin penduduk (laki-laki atau perempuan)
- `date_of_birth` bertipe `DATE` yang berisi tanggal lahir dengan format `YYYY-MM-DD`
- `height` bertipe `INT` yang berisi tinggi badan penduduk dalam satuan centimeter
- `weight` bertipe `INT` yang berisi berat badan penduduk dalam satuan kilogram
- `address` bertipe `VARCHAR(255)` yang berisi alamat penduduk

Dengan contoh data:

| id |      NIK      | first_name | last_name | gender | date_of_birth | height | weight | address |
|----|---------------|------------|-----------|--------|---------------|--------|--------|---------|
| 1  | 1234567890123 | Andi       | Sukirna       | laki-laki | 1990-01-01    | 170    | 70     | Jl. Abc |
| 2  | 1234567890124 | Sulis       | Indahwati       | perempuan | 1990-01-02    | 160    | 50     | Jl. Abc |
| 3  | 1234567890125 | Andre       | William     | laki-laki | 1990-01-03    | 180    | 80     | Jl. Abc |
| 4  | 1234567890126 | Henny       | Welas     | perempuan | 1990-01-04    | 150    | 40     | Jl. Abc |
| ...| ...          | ...        | ...        | ...    | ...           | ...    | ...    | ...     |

Dinas tersebut sedang mengadakan lomba tarik tambang, dan membutuhkan data penduduk dengan berat badan paling berat.

Ambillah 5 data dengan berat badan paling berat dimana jenis kelaminnya adalah **laki-laki**. Data yang diambil harus diurutkan berdasarkan berat badan dari yang terbesar ke terkecil. Jika hasilnya kurang dari 5, maka ambil semua data yang memenuhi syarat tersebut.

Tampilkan kolom `id`, `NIK`, `fullname`, `date_of_birth`, `weight`, `address`, dan kolom `fullname` merupakan gabungan dari `first_name` dan `last_name` yang dipisahkan oleh spasi.

Contoh hasil tabel _query_ :

| id | NIK | fullname | date_of_birth | weight | address |
|----|-----|----------|---------------|--------|---------|
| 1  | ... | ...      | ...           | ...    | ...     |

Kerjakan _query_ ini di dalam file `select.sql`.

> hint: Untuk menggabungkan 2 kolom bertipe `varchar`, gunakan fungsi `CONCAT` atau memakai _concatenation_ dengan operator `||`

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 110) dan **`main_test.go`** (line 12) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"Username"`, `"Password"` dan `"DatabaseName"`saja.

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

| id |      NIK      | first_name | last_name | gender | date_of_birth | height | weight | address |
|----|---------------|------------|-----------|--------|---------------|--------|--------|---------|
| 1  | 1234567890123 | Andi       | Sukirna       | laki-laki | 1990-01-01    | 170    | 70     | Jl. Abc |
| 2  | 1234567890124 | Sulis       | Indahwati       | perempuan | 1990-01-02    | 160    | 50     | Jl. Abc |
| 3  | 1234567890125 | Andre       | William     | laki-laki | 1990-01-03    | 180    | 80     | Jl. Abc |
| 4  | 1234567890126 | Henny       | Welas     | perempuan | 1990-01-04    | 150    | 40     | Jl. Abc |

Maka hasil _query_ yang diharapkan adalah:

| id | NIK | fullname | date_of_birth | weight | address |
|----|-----|----------|---------------|--------|---------|
| 3  | 1234567890125 | Andre William      |   1990-01-03          | 80    | Jl. Abc     |
| 1  | 1234567890123 | Andi Sukirna      |   1990-01-01          | 70    | Jl. Abc     |

#### Test Case 2

Jika tabel yang dimasukkan adalah berikut:

| id |      NIK      | first_name | last_name | gender | date_of_birth | height | weight | address |
|----|---------------|------------|-----------|--------|---------------|--------|--------|---------|
| 1  | 1234567890123 | Andi       | Sukirna       | laki-laki | 1990-01-01    | 170    | 70     | Jl. Jakarta |
| 2  | 1234567890124 | Sulis       | Indahwati       | perempuan | 1990-01-02    | 160    | 50     | Jl. Jakarta |
| 3  | 1234567890125 | Andre       | William     | laki-laki | 1990-01-03    | 180    | 80     | Jl. Jakarta |
| 4  | 1234567890126 | Henny       | Welas     | perempuan | 1990-01-04    | 150    | 40     | Jl. Jakarta |
| 5 | 1234567890127 | Wendy       | Sukirna       | laki-laki | 1990-01-25    | 170    | 71     | Jl. Jakarta |
| 6 | 1234567890128 | Rendy       | Santoso       | laki-laki | 1990-01-21    | 170    | 75     | Jl. Jakarta |
| 7 | 1234567890129 | Rina       | Santoso       | perempuan | 1990-01-12    | 170    | 75     | Jl. Jakarta |
| 8 | 1234567890130 | Johan       | Roger       | laki-laki | 1990-01-10    | 170    | 69     | Jl. Jakarta |
| 9 | 1234567890131 | Albert       | Sunardi       | laki-laki | 1990-01-11    | 170    | 73     | Jl. Jakarta |
| 10 | 1234567890132 | Firman       | Hardi       | laki-laki | 1990-01-05    | 170    | 74     | Jl. Jakarta |

Maka hasil _query_ yang diharapkan adalah:

| id | NIK | fullname | date_of_birth | weight | address |
|----|-----|----------|---------------|--------|---------|
| 3  | 1234567890125 | Andre William      |   1990-01-03   | 80    | Jl. jakarta     |
| 6 | 1234567890128 | Rendy Santoso       |  1990-01-21     | 75     | Jl. Jakarta |
| 10 | 1234567890132 | Firman Hardi       | 1990-01-05       | 74     | Jl. Jakarta |
| 9 | 1234567890131 | Albert Sunardi       | 1990-01-11       | 73     | Jl. Jakarta |
| 5 | 1234567890127 | Wendy Sukirna       | 1990-01-25       | 71     | Jl. Jakarta |
