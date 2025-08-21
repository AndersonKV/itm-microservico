FROM golang:1.24-bullseye

# Instala dependências do Oracle
RUN apt-get update && \
    apt-get install -y unzip libaio1 curl && \
    rm -rf /var/lib/apt/lists/*

# Baixa Oracle Instant Client
WORKDIR /opt/oracle
RUN curl -L -o instantclient-basiclite.zip https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linux.x64-21.11.0.0.0.zip && \
    curl -L -o instantclient-sdk.zip https://download.oracle.com/otn_software/linux/instantclient/instantclient-sdk-linux.x64-21.11.0.0.0.zip && \
    unzip instantclient-basiclite.zip && \
    unzip instantclient-sdk.zip && \
    rm *.zip && \
    ln -s /opt/oracle/instantclient_21_11/libclntsh.so.21.1 /opt/oracle/instantclient_21_11/libclntsh.so

ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_21_11:$LD_LIBRARY_PATH
ENV ORACLE_HOME=/opt/oracle/instantclient_21_11

# Copia código
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build
RUN go build -o auth-service ./cmd/main.go

# Rodar
CMD ["./auth-service"]
