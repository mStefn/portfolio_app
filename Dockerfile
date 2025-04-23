# Etap budowania aplikacji Go
FROM golang:1.21 AS builder

WORKDIR /app

# Skopiuj tylko pliki potrzebne do pobrania zależności najpierw
COPY app/go.mod app/go.sum ./
RUN go mod tidy

# Skopiuj pozostałe pliki aplikacji
COPY app/ ./

# Buduj aplikację — plik wynikowy w katalogu głównym buildera
RUN go build -o /my_app ./main.go

# Etap finalny — minimalny obraz z tylko potrzebnymi plikami
FROM debian:bookworm-slim

WORKDIR /app

# Skopiuj skompilowaną aplikację z poprzedniego etapu
COPY --from=builder /my_app ./

# Skopiuj zasoby aplikacji
COPY app/templates ./templates
COPY app/static ./static

# Wymuś użycie shella i podaj pełną ścieżkę do binarki
CMD ["/app/my_app"]

