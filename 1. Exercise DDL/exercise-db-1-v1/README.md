# Exercise database

## Connect to database

### Description

Setelah kalian belajar instalasi postgresql, mari kita coba koneksi ke database postgresql dari golang. Disediakan sebuah fungsi bernama `Connection` dengan mengembalikan dua nilai bertipe `*sql.DB` dan `error`. `*sql.DB` merupakan koneksi ke database PostgreSQL.

Lengkapi code tersebut agar berhasil melakukan koneksi ke database postgresql yang sudah kalian pelajari sebelumnya.

Jangan lupa untuk mengisi variable `CAMP_ID` dengan ID CAMP kalian dan variable `credential` dengan credential yang kalian gunakan untuk mengakses database postgresql di local kalian.

Terdapat code untuk melakukan ping ke database menggunakan `Ping()` dan akan menampilkan output `"Successfully connected!"` jika berhasil melakukan koneksi.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan sudah membuat database _credentials_ pada file **`main.go`** variable `credential` dengan akses ke database local kalian. Kita akan mengisi koneksi database ke dalam tipe data _struct_ `DBCredential` yang sudah disediakan.

```go
type DBCredential struct {
    HostName     string
    DatabaseName string
    Username     string
    Password     string
    Port         string
}
```

**Contoh** (ini hanya contoh yaa !!) :

```go
credential = DBCredential{
    HostName:     "localhost",
    DatabaseName: "test_db_camp",
    Username:     "postgres",
    Password:     "postgres",
    Port:         "5432",
}
```

### Test Case

Kalian **harus** menjalakan perintah `go run main.go` untuk mengecek apakah jawaban kalian sudah benar atau belum. kalian akan menanggil fungsi `Connection()` di func `main`

**Input**:

```txt
call Connection()
```

**Expected Output / Behavior**:

```txt
Successfully connected!
```

Jika sudah berhasil silahkan jalankan submit.
