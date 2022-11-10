# Exercise Database

## Exercise Query Join

### Description

Sekolah SMP CAMP sedang mengadakan ujian sekolah. Untuk memudahkan proses penilaian, sekolah tersebut memiliki aplikasi untuk menyimpan data nilai siswa. Aplikasi tersebut memiliki 2 tabel _SQL_, yaitu tabel `students` dan `reports`. Tabel `students` berisi data-data siswa, sedangkan tabel `reports` berisi data-data nilai siswa. Setiap siswa dapat memiliki lebih dari 1 nilai di table `reports` berdasarkan banyak mata pelajaran (`study`) yang diambil .

Berikut adalah list kolom dari tabel `students`:

- `id` bertipe `SERIAL` sebagai _primary key_
- `fullname` bertipe `VARCHAR(255)` berisi nama lengkap siswa
- `date_of_birth` bertipe `DATE` berisi tanggal lahir siswa
- `class` bertipe `VARCHAR(255)` berisi kelas siswa
- `status` bertipe `VARCHAR(50)` berisi status siswa ("active" atau "inactive")

Berikut adalah list kolom dari tabel `reports`:

- `id` bertipe `SERIAL` sebagai _primary key_
- `student_id` bertipe `INT` sebagai penghubung (_foreign key_)dari tabel `students`
- `study` bertipe `VARCHAR(255)` berisi nama mata pelajaran
- `score` bertipe `INT` berisi nilai siswa

Setelah kegiatan ujian sekolah, sekolah tersebut ingin mengetahui data-data siswa yang masih aktif dan harus mengikut remidi karena nilai mata pelajaran tertentu kurang dari 70.

Buatlah _query join_ yang dapat menampilkan data-data siswa yang harus mengikut remidi. Urutkan data berdasarkan score terendah.

hasil _query join_ yang diharapkan:

| id | fullname | class | status | study | score |
|----|----------|-------|--------|-------|-------|
| 1  | ...      | ...   | ...    | ...   | ...   |
| 2  | ...      | ...   | ...    | ...   | ...   |


Dengan keterangan sebagai berikut:

- `id` adalah id `reports`
- `fullname` adalah nama lengkap siswa dari tabel `students`
- `class` adalah kelas siswa dari tabel `students`
- `status` adalah status siswa dari tabel `students`
- `study` adalah nama mata pelajaran dari tabel `reports`
- `score` adalah nilai siswa dari tabel `reports`

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 119) dan **`main_test.go`** (line 12) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"Username"`, `"Password"` dan `"DatabaseName"`saja.

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

Jika tabel `students` yang dimasukkan adalah berikut:

| id | fullname | date_of_birth | class | status |
| -- | -------- | ----- | ------- | ------ |
| 1 | John Doe | 2000-01-01 | 1A | active |
| 2 | Jane Willy | 2003-01-01  | 1A | active |
| 3 | John Smith | 2002-03-01 | 1A | inactive |
| 4 | Bob Abrams | 2003-04-01 | 1B | active |

Dan table `reports` berisi data:

| id | student_id | study | score |
| -- | ---------- | ----- | ----- |
| 1 | 1 | Math | 90 |
| 2 | 1 | English | 80 |
| 3 | 1 | Science | 70 |
| 4 | 1 | Indonesia | 70 |
| 5 | 2 | Math | 55 |
| 6 | 2 | English | 80 |
| 7 | 2 | Science | 61 |
| 8 | 2 | Indonesia | 70 |
| 9 | 3 | Math | 90 |
| 10 | 3 | English | 80 |
| 11 | 3 | Science | 70 |
| 12 | 3 | Indonesia | 70 |
| 13 | 4 | Math | 65 |
| 14 | 4 | English | 30 |
| 15 | 4 | Science | 40 |
| 16 | 4 | Indonesia | 50 |

Maka hasil _query join_ yang diharapkan adalah:

| id | fullname | class | status | study | score |
|----|----------|-------|--------|-------|-------|
| 14  | Bob Abrams | 1B | active | English | 30 |
| 15  | Bob Abrams | 1B | active | Science | 40 |
| 16  | Bob Abrams | 1B | active | Indonesia | 50 |
| 5  | Jane Willy | 1A | active | Math | 55 |
| 7  | Jane Willy | 1A | active | Science | 61 |
| 13 | Bob Abrams | 1B | active | Math | 65 |

Hasil di atas adalah data dengan nilai mata pelajaran (`score`) yang kurang dari 70 dan status siswa `active` yang diurutkan berdasarkan nilai terendah.
