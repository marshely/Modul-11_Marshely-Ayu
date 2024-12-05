Nama: MARSHELY AYU ISWANTO 


NIM: 2311102073


Kelas: IF-11-03 


MODUL 11


PENCARIAN NILAI ACAK PADA HIMPUNAN DATA


Deskripsi Program


No.1


Program tersebut merupakan program go yaitu untuk menghitung suara dari sebuah pemilihan. Pengguna diminta untuk memasukkan suara dalam bentuk angka yang dipisahkan dengan spasi, dan program akan membaca input tersebut hingga angka 0 dimasukkan sebagai penanda akhir. Suara yang valid hanya angka antara 1 hingga 20, yang merepresentasikan nomor calon. Program ini secara otomatis mengabaikan input yang tidak valid (seperti huruf atau angka di luar jangkauan) dan menghitung jumlah total suara yang terbaca, jumlah suara valid, serta jumlah suara untuk masing-masing calon.

Setelah semua suara diproses, program akan menampilkan outputnya. Misalnya, berapa banyak suara yang masuk, berapa banyak yang valid, dan distribusi suara untuk setiap calon. Jika ada calon yang mendapat suara, program akan mencetak hasilnya dengan nomor calon dan jumlah suara yang diperoleh. 


No.2


Program tersebut merupakan program go yaitu untuk mendukung simulasi pemilihan Ketua dan Wakil Ketua RT dengan cara yang mudah. Pengguna hanya perlu memasukkan suara dalam angka yang dipisahkan oleh spasi, dan setelah semua suara dimasukkan, pengguna menambahkan angka 0 sebagai penanda akhir input. Program ini akan mengelola suara yang diterima, menghitung total suara yang ada, jumlah suara yang valid (angka antara 1 hingga 20), serta mendistribusikan suara kepada setiap calon. Input yang tidak tepat seperti huruf atau angka di luar batasan akan diabaikan secara otomatis.


Selain menghitung suara, program ini juga memilih dua calon dengan suara terbanyak untuk posisi Ketua dan Wakil Ketua RT. Data calon akan diurutkan berdasarkan banyaknya suara, dan bila terdapat suara yang sama, calon dengan nomor yang lebih kecil akan diutamakan. Hasil akhir mencakup total suara yang terhitung, suara yang valid, serta nama Ketua dan Wakil Ketua RT, yang ditampilkan dengan jelas.


No.3  


Pencarian menggunakan binary search untuk menentukan lokasi sebuah angka 𝑘, serta isi dari array tersebut. Data disimpan dalam array global yang bernama data. Program kemudian memanggil fungsi posisi untuk menemukan angka 𝑘 dengan menggunakan algoritma binary search, yang cepat dan efisien untuk data yang sudah terurut. Fungsi posisi bekerja dengan membagi array menjadi dua bagian setiap langkahnya. Ia membandingkan nilai tengah dari array dengan angka 𝑘. Jika nilainya sama, posisi angka ditemukan dan dikembalikan. Jika nilai tengah lebih kecil daripada 𝑘, pencarian dilanjutkan ke separuh kanan; jika lebih besar, pencarian berlanjut ke separuh kiri. Jika angka 𝑘 tidak ditemukan, fungsi akan mengembalikan -1. Hasil akhirnya adalah lokasi angka 𝑘 (dengan indeks dari 0) atau pesan &quot;TIDAK ADA&quot; jika angka tersebut tidak ada. Algoritma ini sangat efisien, khususnya untuk data besar, karena memiliki kompleksitas waktu logaritmik O(logn).

