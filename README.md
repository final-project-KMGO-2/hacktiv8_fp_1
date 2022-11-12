# hacktiv8_fp_1

### Final Project Hacktiv8 Kelompok 2
- Tasya Gracinia
- Muhammad Khoirul Anam
- Alexander

### Deskripsi
Source code pada direktori ini merupakan hasil dari pengerjaan final project 1 studi independen Hacktiv8 - Golang for Backend Engineer kelompok 2. Tugas yang diberikan adalah membuat API sebuah aplikasi todos yang memiliki endpoint untuk Create, Read, Update, dan Delete todos. Sebagai tambahan untuk pengerjaan proyek ini dan juga sebagai pilot project ke depannya, API memiliki jwt auth yang mengharuskan pengguna login sebelum menggunakan API.
Untuk token, digunakan Bearer token sehingga diperlukan penambahan string "Bearer " (dengan spasi) yang dimasukkan ke dalam header Authorization agar dapat masuk.

Proyek ini dideploy menggunakan heroku pada link berikut :

https://todo-list-fph8.herokuapp.com/

### Dokumentasi
Untuk melihat dokumentasi swagger proyek ini, link berikut dapat diakses :

https://todo-list-fph8.herokuapp.com/swagger/index.html

### Pembagian tugas
Pada pengerjaan final project ini, dilakukan pembagian tugas sebagai berikut :

- Tasya Gracinia
    - Desain model untuk database
    - Membuat endpoint delete todos by id
    - Membuat endpoint update todos by id
- Muhammad Khoirul Anam
    - Membuat GitHub Organization
    - Membuat endpoint post todos
    - Membuat endpoint get todos by id
    - Membuat middleware auth
    - Melakukan deployment ke heroku
- Alexander
    - Membuat boilerplate project
    - Melakukan setup Swaggo
    - Membuat endpoint dan service auth
    - Membuat endpoint get todos
