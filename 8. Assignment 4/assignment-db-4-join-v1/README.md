# Exercise DB

## Assignment Join Table

### Description

Toko sembako ingin memberikan promo kepada pembeli dengan menghubungi via email. Promo ini diberikan kepada pembeli yang pernah berbelanja 1 jenis barang dengan  total lebih dari 500 ribu atau pembelian lebih dari 20 unit.

Terdapat 2 tabel yang digunakan untuk menentukan promo tersebut, yaitu tabel `users` dan `orders`. Tabel `users` berisi data-data pembeli, sedangkan tabel `orders` berisi data-data pembelian. Setiap pembeli (_users_) akan mencatat setiap pembelian 1 jenis barang ke dalam 1 data di table `orders`

List kolom tabel `users`:

- `id` (integer) menjadi _primary key_
- `fullname` (varchar(255)) yang berisi nama lengkap pembeli
- `email` (varchar(255)) yang berisi email pembeli
- `address` (varchar(255)) yang berisi alamat pembeli
- `status` (varchar(255)) yang berisi status pembeli (active atau inactive)

List kolom tabel `orders`:

- `id` (integer) menjadi _primary key_
- `product_name` (varchar(255)) yang berisi nama produk yang dibeli
- `unit_price` (integer) yang berisi harga satuan produk
- `quantity` (integer) yang berisi jumlah produk yang dibeli
- `order_date` (date) yang berisi tanggal pembelian
- `user_id` (integer) yang berisi id pembeli (_foreign key_ dari tabel `users`)

Lakukan query untuk menampilkan data-data pembeli yang mendapatkan promo tersebut dengan ketentuan sebagai berikut:

- `status` pembeli harus `active`
- total pembelian (unit price * quantity) harus lebih dari 500 ribu atau total quantity lebih dari 20

Hasil _query_ yang diharapkan:

| order_id | fullname | email | product_name | unit_price | quantity | order_date |
| -------- | -------- | ----- | ------------ | ---------- | -------- | ---------- |
| ... | ... | ... | ... | ... | ... | ... |

Kerjakan _query join_ ini di dalam file `join.sql`.

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

Jika tabel `users` yang dimasukkan adalah berikut:

| id | fullname | email | address | status |
| -- | -------- | ----- | ------- | ------ |
| 1 | John Doe | john@mail.com | Jl. Kebon Jeruk | active |
| 2 | Jane Doe | jane@mail.com  | Jl. Kebon Jeruk | active |
| 3 | Bob Doe | bobdoe@mail.com | Jl. Kebon Jeruk | active |
| 4 | Alice Doe | alice@mail.com | Jl. Kebon Jeruk | active |
| 5 | Bob Marley | marley@mail.com | Jl. Kebon Jeruk | inactive |

Dan table `orders` berisi data:

| id | product_name | unit_price | quantity | order_date | user_id |
| -- | ------------ | ---------- | -------- | ---------- | ------- |
| 1 | Beras 3kg | 30000 | 10 | 2021-01-01 | 1 |
| 2 | Gula 2kg | 20000 | 5 | 2021-01-01 | 2 |
| 3 | Beras 10Kg | 100000 | 6 | 2021-01-01 | 3 |
| 4 | Telur | 5000 | 50 | 2021-01-01 | 4 |
| 5 | Minyak Goreng 1Lt | 30000 | 17 | 2021-01-01 | 5 |

Maka hasil _query_ yang diharapkan adalah:

| order_id | fullname | email | product_name | unit_price | quantity | order_date |
| -------- | -------- | ----- | ------------ | ---------- | -------- | ---------- |
| 3 | Bob Doe | bobdoe@mail.com | Beras 10Kg | 100000 | 6 | 2021-01-01 |
| 4 | Alice Doe | alice@mail.com | Telur  | 5000 | 50 | 2021-01-01 |
