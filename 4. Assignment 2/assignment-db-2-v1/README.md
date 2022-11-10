# Assignment DB 2

## Assignment Data Manipulation Language

### Description

Sekolah SMP CAMP memiliki sistem pencatatan data untuk siswa dan guru. Data siswa dan guru disimpan dalam database tabel `students` dan `teachers`. SMP Camp saat ini sedang mengadakan penerimaan siswa baru dan melakukan _update_ data guru di database.

Tabel `students` berisi kolom sebagai berikut:

- `id` bertipe `int` tidak boleh `null` sebagai _primary key_
- `first_name` bertipe `varchar(255)` tidak boleh `null`, yang berisi nama depan siswa
- `last_name` bertipe `varchar(255)` tidak boleh `null`, yang berisi nama belakang siswa
- `gender` bertipe `VARCHAR(50)` tidak boleh `null`, yang berisi jenis kelamin siswa (laki-laki atau perempuan)
- `date_of_birth` bertipe `date` tidak boleh `null`, yang berisi tanggal lahir siswa
- `address` bertipe `varchar(255)` dapat diisi data kosong / `null`, yang berisi alamat siswa
- `class` bertipe `varchar(10)` tidak boleh `null`, yang berisi kelas siswa
- `status` bertipe `varchar(50)` tidak boleh `null` (active atau inactive)

Tabel `teachers` berisi kolom sebagai berikut:

- `id` bertipe `int` tidak boleh `null` sebagai _primary key_
- `nip` bertipe `int` tidak boleh `null`, yang berisi nomer induk pegawai
- `fullname` bertipe `varchar(255)` tidak boleh `null`, yang berisi nama lengkap
- `address` bertipe `varchar(255)` dapat diisi data kosong / `null`, yang berisi alamat lengkap
- `groups` bertipe `varchar(10)` tidak boleh `null`, yang berisi golongan jabatan (A, B, C, D)
- `status` bertipe `varchar(50)` tidak boleh `null` (active atau inactive)

kalian diminta untuk:

1. Tambahkan data berikut ke tabel `students`:

    | first_name | last_name |gender |date_of_birth | address | class | status |
    |------------|-----------|-----|----------|---------|-------|--------|
    | Imam       | Rendi      | laki-laki | 2002-02-02    |  Jl Jakarta       | 1A    | active |
    | Andi       | Sukirna      | laki-laki | 2002-02-03    |  Jl Jakarta       | 1A    | active |
    | Achmad       | Fadjar    | laki-laki  | 2002-02-03    |  Jl Depok       | 1A    | active |
    | Achmad       | Kalla    | laki-laki   | 2002-02-03    |  -       | 1A    | active |
    | Aida       | Ishak     | perempuan  | 2002-01-01    |  Jl Depok       | 1A    | active |
    | Alice | Haryono | perempuan  | 2002-01-01    |  Jl Depok       | 1A    | active |
    | Calvin | Lukmantara | laki-laki | 2002-01-05    |  Jl Jakarta Utara       | 1A    | active |
    | Chris | Fong |  laki-laki  | 2002-01-07    |  Jl Jakarta Utara       | 1A    | active |
    | Citra | Andini | perempuan | 2002-01-08    | -       | 1B    | active |
    | Darwin | Leo | laki-laki  | 2003-04-01    |  Jl Jakarta Utara       | 1B    | active |
    | Dewi | Nilka Sari | perempuan | 2003-05-01    |  Jl Jakarta Utara       | 1B    | active |
    | Edy |  Kosasih | laki-laki  | 2003-06-01    |  Jl Sukarno Hatta       | 1B    | active |
    | Fabian |  Gelael | laki-laki  | 2003-07-01    |  Jl Sukarno Hatta       | 1B    | active |
    | Halifah | Indah | perempuan  | 2003-09-01    |  -       | 1B    | active |
    | Hari | Widodo | laki-laki  | 2003-10-01    |  -       | 1B    | active |

    Pastikan kalian menambahkan data **sesuai urutan** dari tabel diatas. Jika terdapat data yang ditulis dengan _slash_ "-" maka data tersebut dianggap kosong / `null`

2. Menghubah data pada tabel `teachers` dimana golongan A menjadi B (naik jabatan)
3. Hapus data pada tabel `teachers` yang sudah tidak aktif (status = **inactive**)

Kerjakan tugas nomor 1 di file `insert_students.sql`. Kerjakan tugas nomor 2 di file `update_data_teachers.sql` dan tugas nomor 3 di file `delete_data_teachers.sql`.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 90) dan **`main_test.go`** (line 23) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

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

Jika berhasil maka akan terdapat tambahan data di tabel `students` sebanyak 15 data. Untuk data di tabel `teachers` akan mengubah data `group` **A** menjadi **B** dan menghapus data yang statusnya **inactive**.
