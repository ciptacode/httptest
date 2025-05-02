# httptest

Aplikasi HTTP testing sederhana yang dibuat dengan Golang, berfungsi untuk mengembalikan response sesuai dengan status code HTTP yang diminta.

Aplikasi ini berguna untuk kebutuhan pengujian HTTP client, seperti simulasi status code error (404, 500, 422, dll) dan delay response.

---

## Format URL:

```
GET /<status_code>?sleep=<milliseconds>
```
- <status_code>: status HTTP yang ingin diuji (contoh: 200, 404, 422).
- sleep: (opsional) delay dalam milidetik sebelum mengembalikan response.

## Contoh:

```
curl https:://httptest.fahrigunadi.dev/200
# Response: 200 OK

curl https:://httptest.fahrigunadi.dev/422
# Response: 422 Unprocessable Entity

curl https:://httptest.fahrigunadi.dev/500?sleep=3000
# Response setelah 3 detik: 500 Internal Server Error
```

Contoh penggunaannya cocok untuk:

- Menguji penanganan timeout di aplikasi.

- Mengetes behavior client terhadap berbagai HTTP status code.