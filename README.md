# Design Architecture
- design architecture yang di gunakan adalah konsep clean architecture, setiap komponen yang ada bersifat independen dan tidak bergantung pada library external yang spesifik. Seperti tidak tergantung pada spesifik framework atau tidak bergantung pada spesifik database yang dipakai. Seperti yang di paparkan Uncle Bob dan konsep clean achitecture yang di jelaskan oleh beliau, yakni tentang arsitektur kue lapis.
- Oleh Uncle Bob beliau menyebutkan 4 layer pada arsitekturnya:
1. Entities
2. Usecase
3. Controller
4. Framework dan Driver
- Disini project ini memiliki 4 layer:
1. Entities
2. Repository
3. Usecase
4. Handler

- Entities
Layer ini merupakan layer yang menyimpan model & schema yang dipakai pada domain lainnya. Layer ini dapat diakses oleh semua layer dan oleh semua domain.
- Repository
Layer ini merupakan layer yang menyimpan database handler. Querying, Inserting, Deleting akan dilakukan pada layer ini. Tidak ada business logic disini. Yang ada hanya fungsi standard untuk input-output dari datastore.
- Usecase
Layer ini merupakan layer yang akan bertugas sebagai pengontrol, yakni menangangi business logic pada setiap domain. Layer ini juga bertugas memilih repository apa yang akan digunakan, dan domain ini bisa memiliki lebih dari satu repository layer.
- Handler
Tugas dari layer ini menjadi dinding penghubung antara user dan sistem. Menerima segala input dan validasi input sesuai standar yang digunakan.

# Catatan
- Sorting
    1. Product terbaru 
        - sortBy = created_at atau kosong ("")
        - sortType = desc atau kosong ("")
    2. Product harga termurah
        - sortBy = price
        - sortType = asc
    3. Product harga terendah
        - sortBy = quantity
        - sortType = asc
    4. Sort by product name (A-Z)
        - sortBy = name
        - sortType = asc
      Sort by product name (Z-A)
        - sortBy = name
        - sortType = desc

- Run aplikasi
    1. buat database erajaya
    2. jalankan redis
    3. go run main.go atau air . (jika menggunakan air) atau gunakan docker 

