# Assignment DB 3

## Assignment Data Query Language

### Description

Sekolah **SMP CAMP** sedang mengadakan ujian, mata pelajaran yang diuji adalah Bahasa Indonesia, Bahasa Inggris, Matematika, dan IPA. Sekolah tersebut sudah menggunakan aplikasi untuk mencatat hasil nilai ujian siswa. Aplikasi tersebut menyimpan data siswa dalam suatu tabel _SQL_ bernama `final_scores`. Tabel tersebut berbentuk sebagai berikut:

| id | exam_id | first_name | last_name | bahasa_indonesia | bahasa_inggris | matematika | ipa | exam_status | fee_status |
|----|------|------------|-----------|------------------|----------------|------------|-----|-------------|------------|
| 1  | 1A-002 | John       | Doe       | 90               | 80             | 70         | 90  | pass  | full |
| 2  | 1A-001 | Jane       | Doe       | 80               | 70             | 60         | 80  | fail  | full |
| 3  | 1B-003 | John       | Smith     | 70               | 60             | 50         | 70  | pass  | installment |
| 4  | 1C-002 | Jane       | Smith     | 60               | 50             | 40         | 60  | pass  | full |
| ... | ...     | ...        | ...       | ...              | ...            | ...        | ... | ...   | ...  | ... |

Terdapat kode pengenal ujian siswa di kolom `exam_id` yang memiliki pola `<kelas>-<nomor urut>`, dimana `kelas` adalah kelas dari siswa tersebut dan `nomor urut` adalah nomor urut dari siswa tersebut. Jadi jika ada contoh code `1A-001` maka siswa tersebut berada di kelas `1A` dan nomor urutnya adalah `001`.

Terdapat juga kolom `exam_status` dan `fee_status` yang berisi status ujian dan status pembayaran siswa. Exam status akan berisi 2 nilai yaitu `pass` dan `fail` dengan keterangan sebagai berikut:

- "**pass**" jika selama ujian tidak ada indikasi mencontek atau curang
- "**fail**" jika selama ujian ketahuan mencontek atau curang

Fee status adalah status pembayaran uang SPP minimal 3 bulan terakhir. Terdapat 3 nilai yaitu `full`, `installment`, dan `not paid` dengan keterangan sebagai berikut:

- "**full**" jika sudah membayar uang SPP secara penuh
- "**installment**" jika sudah membayar uang SPP secara mencicil / angsuran
- "**not paid**" jika belum membayar uang SPP

Setelah melakukan ujian, sekolah ingin mendapatkan top 5 nilai rata-rata tertinggi (_average score_). Nilai `average_score` dihitung dengan cara mendapatkan rata-rata dari nilai-nilai di kolom `bahasa_indonesia`, `bahasa_inggris`, `matematika`, dan `ipa`.

Sekolah akan memberikan hadiah kepada 5 nilai tinggi tersebut. Namun, sekolah hanya mengambil top 5 nilai tertinggi dengan ketentuan sebagai berikut:

- Jika siswa tersebut tidak ada indikasi mencontek atau curang
- Jika siswa tersebut sudah membayar uang SPP secara penuh atau secara mencicil / angsuran

Tampilkan hasil query dari tabel `final_scores` dengan format sebagai berikut:

| id |  fullname| class  | average_score |
|----|----------|----------|---------------|
| 1  | John Doe | 1A    | 85            |
| 2  | Jane Smith | 1B    | 75            |
| ...| ...      | ...      | ...           |

Kolom `fullname` merupakan gabungan dari `first_name` dan `last_name` yang dipisahkan oleh spasi `" "`.

> **hint**: gunakan fungsi `SUBSTRING()` atau `SPLIT_PART()` pada SQL postgres untuk mendapatkan potongan string

Buatlah query `select` untuk mengambil data sesuai ketentuan di atas pada file `select.sql`.

### **Perhatian**

Sebelum kalian menjalankan `grader-cli test`, pastikan kalian sudah mengubah database credentials pada file **`main.go`** (line 109) dan **`main_test.go`** (line 12) sesuai dengan database kalian. Kalian cukup mengubah nilai dari  `"username"`, `"password"` dan `"database_name"`saja.

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

Jika di dalam tabel `final_scores` terdapat data sebagai berikut:

| No | Exam ID | First Name | Last Name | Bahasa Indonesia | Bahasa Inggris | Matematika | IPA | Exam Status | Fee Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 1A-001|John|Doe|80|90|70|80|pass|full |
| 2 | 1A-002|Jane|Doe|90|80|90|80|pass|notpaid |
| 3 | 1A-003|John|Smith|70|80|70|80|pass|installment |
| 4 | 1A-004|Jane|White|80|70|80|80|pass|full |
| 5 | 1A-005|Abrams|White|80|70|80|80|pass|full |
| 6 | 1A-006|Herdi|White|80|70|80|80|fail|notpaid |
| 7 | 1A-007|Wendy|White|100|95|80|80|fail|installment |
| 8 | 1A-008|Ardi|White|100|95|80|80|pass|notpaid |
| 9 | 1A-009|Abrams|Smith|95|93|80|80|fail|notpaid |
| 10 | 1A-010|Welly|White|95|93|80|80|fail|notpaid |
| 11 | 1B-001|Indah|Sudarni|95|93|80|80|fail|full |
| 12 | 1B-002|Aren|White|80|70|80|80|pass|full |
| 13 | 1B-003|John|Bernard|80|90|70|80|faid|installment |
| 14 | 1B-004|Jane|Abrams|90|80|90|80|pass|full |
| 15 | 1B-005|John|Albert|70|80|70|80|pass|installment |

Maka kita akan mendapatkan output sebagai berikut:

| id |  fullname| class  | average_score |
|----|----------|----------|---------------|
| 14 | Jane Abrams | 1B    | 85            |
| 1  | John Doe | 1A    | 80            |
| 12  | Aren White | 1B    | 77            |
| 5  | Abrams White | 1A    | 77            |
| 4 | Jane White | 1A    | 77         |

Hasil di atas didapat setelah mengeliminasi siswa yang tidak memenuhi kriteria, yaitu siswa yang tidak membayar uang SPP dan yang terindicasi mencontek atau curang. Kemudian diurutkan berdasarkan nilai rata-rata tertinggi ke terendah dan diambil 5 siswa dengan nilai rata-rata tertinggi.
