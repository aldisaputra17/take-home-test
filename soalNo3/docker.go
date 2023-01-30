package soalNo3

/* No 3.  Apakah ada kesalahan dari script di bawah ini?
Jika ada tolong jelaskan dimana letak kesalahannya dan bagaimana anda memperbaikinya.
Jika tidak ada, tolong jelaskan untuk apa script di bawah ini.

FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80
*/

/* Jawaban:
Soal di atas tidak ada kesalahan, script tersebut merupakan DockerFile. Dockerfile di gunakan
untuk membuat image di dalam docker. Image docker ini di gunakan untuk menjalankan aplikasi
"Indihome" yg di tulis dengan bhs pemograman golang. Dockerfile ini akan membuat sebuah image
yg menjalankan aplikasi indihome yg berjalan di port 80.
penjelasan function :
1. from golang, adalah sebuah image docker bahasa golang berdasarkan versi latest.
2. WORKDIR /go/src/github.com/telkomdev/indihome, di gunakan untuk menyalin file dan direktori
dari sumber tujuan yg di tentukan pada image docker.
3. WORKDIR /go/src/github.com/telkomdev/indihome,  menetapkan direktori kerja untuk instruksi
Dockerfile lainnya, seperti RUN , CMD , dan juga direktori kerja untuk
menjalankan instans gambar kontainer.
4. RUN go get github.com/tools/godep, Instruksi RUN akan menjalankan perintah pada layer baru di
 atas image saat ini.
5. sama seperti no 4
6. sama seperti no 4
7.ENTRYPOINT /go/bin/indihome, sama saja dengan CMD, instruksi ini mendefinisikan perintah apa
yang akan dieksekusi ketika menjalankan sebuah container.
8. LISTEN 80, adalah mendengar aplikasi tersebut berjalan di port 80
*/
