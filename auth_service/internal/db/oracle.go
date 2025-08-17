package db

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func ConnectOracle(user, password, host, port, serviceName string) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s/%s@%s:%s/%s", user, password, host, port, serviceName)
    db, err := sql.Open("godror", dsn)
    if err != nil {
        return nil, err
    }
    // Testa a conexão
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
	