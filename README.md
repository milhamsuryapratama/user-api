# user-api

## Pengaturan Database

1. buat database mysql (dengan nama terserah)
2. buka project dan cari file `.env`
3. isi `MYSQL_DB` dengan nama database yang telah dibuat sebelumnya

## Endpoint API

1. `/api/login` (LOGIN)
    ## Post Form
    * `username` : admin
    * `password` : admin 
2. `/api/user` ( METHOD GET, untuk menampilkan data user )
    ## Header data
    * `Token`: isi dengan token ketika login

3. `/api/user` ( METHOD POST, untuk menambah data user )
    ## Post Form multipart
    * `nama_lengkap`: isi nama lengkap
    * `username`: isi username
    * `password`: isi password
    * `foto`: isi file image/foto
    ## Header data
    * `Token`: isi dengan token ketika login

4. `/api/user/:id` ( METHOD PUT, untuk mengubah data user )
    ## Post Form multipart
    * `nama_lengkap`: isi nama lengkap
    * `username`: isi username
    * `password`: isi password
    * `foto`: isi file image/foto
    ## Header data
    * `Token`: isi dengan token ketika login

5. `/api/user/:id` ( METHOD DELETE, untuk menghapus data user )
    ## Header data
    * `Token`: isi dengan token ketika login

## Run Project
1. untuk menjalankan project, ketik perintan `go run main.go`