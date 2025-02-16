# Menggunakan image Go resmi sebagai base image
FROM golang:1.23-alpine

# Install dependencies dan air (alat untuk hot-reloading)
RUN apk add --no-cache curl tar && \
    # Download air tar.gz ke home directory
    curl -fLo /root/air.tar.gz https://github.com/air-verse/air/releases/download/v1.61.5/air_1.61.5_linux_amd64.tar.gz && \
    # Ekstrak file tar.gz ke direktori home
    tar -xvzf /root/air.tar.gz -C /root && \
    # Berikan izin eksekusi pada file air
    chmod +x /root/air && \
    # Pindahkan air ke /usr/local/bin
    mv /root/air /usr/local/bin/air && \
    # Hapus file arsip tar.gz untuk menghemat ruang
    rm /root/air.tar.gz

# Set working directory di dalam container
WORKDIR /app

# Menyalin file Go modules (go.mod & go.sum) dan mengunduh dependensi
COPY go.mod go.sum ./
RUN go mod tidy

# Menyalin seluruh file proyek ke dalam container
COPY . .

# Perintah untuk menjalankan air yang akan memantau perubahan file
CMD ["air", "-c", "air-linux.toml"]

# Tentukan port yang akan digunakan
EXPOSE 8080