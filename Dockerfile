# Etap 1: Budowanie aplikacji Go
FROM golang:1.21 AS builder

# Ustawienie katalogu roboczego
WORKDIR /app

# Kopiowanie plików go.mod i go.sum z katalogu app
COPY app/go.mod app/go.sum ./

# Pobieranie zależności Go
RUN go mod tidy

# Kopiowanie reszty kodu aplikacji z katalogu app
COPY app/ .

# Budowanie aplikacji Go
RUN go build -o my_app ./main.go

# Etap 2: Tworzenie obrazu produkcyjnego
FROM debian:bookworm-slim

# Ustawienie katalogu roboczego
WORKDIR /app

# Skopiowanie skompilowanego pliku aplikacji z etapu budowania
COPY --from=builder /app/my_app .

# Skopiowanie katalogu szablonów i statycznych plików do kontenera
COPY app/templates ./templates
COPY app/static ./static

# Określenie portu, na którym aplikacja nasłuchuje
EXPOSE 8080  

CMD ["./my_app"]

