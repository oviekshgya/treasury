# Latihan Go

Lima soal di sini masing-masing punya fungsi bernama `GO`. Karena Go tidak mendukung beberapa fungsi dengan nama yang sama dalam satu package, setiap soal ditempatkan di folder sendiri, dari `soal1` sampai `soal5`. File `main.go` memanggil semuanya dengan input dari contoh soal.

## Menjalankan program

Gunakan Go 1.22 atau yang lebih baru. Tidak ada library tambahan yang perlu dipasang.

Buka terminal di folder project, lalu jalankan:

```sh
go run .
```

Hasilnya:

```text
Soal 1:
0
-2
3

Soal 2 (hasil, error):
0 not found
62952 <nil>

Soal 3:
[1]
[2 1]

Soal 4:
1 -> 1
2 -> 2
3 -> 3
4 -> 1

Soal 5:
14
39
```

`<nil>` pada soal 2 berarti tidak ada error. Kalau mau mencoba input lain, ubah pemanggilan fungsi di `main.go`, lalu jalankan perintah yang sama.

## Soal 1: mencari bilangan yang belum ada

Implementasi ada di [soal1/solution.go](soal1/solution.go).

Teks soal dan contohnya berbeda. Jika yang dicari benar-benar nilai terkecil atau terbesar dari isi array, `[-1, -7, -5]` akan menghasilkan `-1`, dan `[1, 2, 1, 6]` akan menghasilkan `1`. Namun contoh meminta `-2` dan `3`. Kode ini mengikuti contoh: mencari bilangan yang belum ada, dimulai dari yang paling dekat dengan nol.

Untuk input positif, pencarian berjalan dari `1, 2, 3, ...`. Untuk input negatif, arahnya `-1, -2, -3, ...`. Angka yang sudah ada dicatat dalam map. Begitu menemukan angka yang belum ada, angka itu dikembalikan.

Pada `[-1, -7, -5]`, angka `-1` sudah ada dan `-2` belum ada, jadi hasilnya `-2`. Pada `[1, 2, 1, 6]`, angka `1` dan `2` sudah ada, jadi hasilnya `3`.

Jika positif dan negatif muncul bersamaan, hasilnya `0`. Nol diabaikan saat menentukan tanda. Input kosong atau berisi nol saja juga menghasilkan `0`.

## Soal 2: menggabungkan angka, membalik, lalu membuang digit

Implementasi ada di [soal2/solution.go](soal2/solution.go).

Contoh `1259, 2601` dengan `exclude` berisi `0` dan `1` bisa dibaca seperti ini:

```text
Gabung       : 12592601
Balik        : 10629521
Buang 0 dan 1: 62952
```

Kode membaca urutan angka dan digit dari belakang, sambil melewati digit yang dikecualikan. Hasil akhirnya baru dikonversi ke `int`. Jadi, angka gabungan tidak perlu ditampung dalam `int` sementara yang bisa lebih dulu mengalami overflow.

Karena gabungan diperlakukan sebagai bilangan, nol di depan tidak dipertahankan. Input `0, 12` menjadi `12`, lalu dibalik menjadi `21`. Nol di tengah tetap dipakai. Jika semua input nol, masih ada satu digit `0` yang bisa diproses.

Jika tidak ada digit tersisa, fungsi mengembalikan `0` dan error dengan pesan `"not found"`. Error ini juga tersedia sebagai `soal2.ErrNotFound`. Kalau yang tersisa hanya digit nol, hasilnya `0, nil`. Jika hasil akhir melebihi kapasitas `int`, fungsi mengembalikan `0` dan error konversi.

Soal tidak menentukan perlakuan untuk bilangan negatif atau isi `exclude` di luar digit desimal. Di sini tanda minus diabaikan, begitu juga nilai `exclude` di luar `0` sampai `9`. Nilai yang berulang dalam `exclude` tidak mengubah hasil.

## Soal 3: mencari angka yang ada di kedua slice

Implementasi ada di [soal3/solution.go](soal3/solution.go).

Isi `numsB` dicatat dalam map, lalu setiap angka di `numsA` diperiksa terhadap map tersebut. Angka yang cocok dimasukkan ke hasil, lalu dihapus dari map supaya tidak terambil dua kali. Setelah itu hasil diurutkan dari besar ke kecil.

Itulah sebabnya `[1, 2, 2]` dan `[1, 2, 4]` menghasilkan `[2, 1]`. Angka `2` hanya muncul sekali dalam hasil. Jika tidak ada angka yang cocok, hasilnya slice kosong. Kedua slice masukan tidak diubah.

## Soal 4: menghitung kemunculan digit

Implementasi ada di [soal4/solution.go](soal4/solution.go).

Bilangan diubah menjadi teks, lalu digitnya dibaca satu per satu. Setiap digit menjadi key pada map, dan jumlah kemunculannya menjadi value. Untuk `1223334`, hasilnya `1:1`, `2:2`, `3:3`, dan `4:1`.

Input `0` dihitung sebagai satu digit nol. Jika input negatif, tanda minus dilewati. Cara ini juga bisa menangani nilai `int` negatif terkecil tanpa harus mengubahnya menjadi positif terlebih dahulu.

Map tidak menjamin urutan saat dibaca. Karena itu, `main.go` menampilkan hasil dengan memeriksa digit `0` sampai `9` agar urutannya tetap sama.

## Soal 5: perkalian dengan goroutine

Implementasi ada di [soal5/solution.go](soal5/solution.go).

Angka dipasangkan berdasarkan indeks: `0` dengan `1`, `2` dengan `3`, dan seterusnya. Setiap pasangan dikerjakan oleh satu goroutine. Jika angka terakhir tidak punya pasangan, angka itu dikalikan dengan dirinya sendiri.

Untuk `[1, 2, 3, 4, 5]`, ada tiga goroutine yang menghitung `1*2`, `3*4`, dan `5*5`. Masing-masing mengirim hasil ke channel. Goroutine pemanggil menerima ketiganya dan menjumlahkan `2 + 12 + 25`, sehingga hasil akhirnya `39`. Saat dipanggil dari `main.go`, penjumlahan ini berlangsung di goroutine utama.

Urutan hasil masuk ke channel bisa berbeda, tetapi jumlah akhirnya tetap sama. Fungsi menunggu satu hasil untuk setiap pasangan sebelum mengembalikan total. Input kosong menghasilkan `0`. Perkalian dan penjumlahan memakai tipe `int` biasa, tanpa pemeriksaan overflow.

## Menjalankan pengujian

Setiap folder soal punya file `solution_test.go`. Jalankan semuanya dengan:

```sh
go test -v ./...
```

Pengujian mencakup contoh soal, input kosong, nol, duplikat, bilangan negatif, overflow pada soal 2, dan nilai `int` negatif terkecil pada soal 4. Untuk fungsi yang menerima slice, ada pemeriksaan bahwa isi input tidak berubah. Soal 5 juga diuji dengan beberapa pemanggilan bersamaan dan 1.001 angka untuk memeriksa banyak pasangan.

Untuk pemeriksaan statis dan data race:

```sh
go vet ./...
go test -race -cover ./...
```

Opsi `-race` memerlukan C compiler yang didukung Go. Di Windows, pastikan compiler seperti GCC tersedia di `PATH`. Coverage dari perintah terakhir menunjukkan bagian kode yang dijalankan oleh tes; `main.go` hanya berisi contoh pemanggilan dan belum punya tes tersendiri.

Konfigurasi `.idea`, file pengaturan dan instruksi agent lokal, serta hasil build sudah dimasukkan ke `.gitignore`.
