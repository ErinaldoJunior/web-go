package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) InsertUsers(nome string, email string, contacto string, senha string) {
	// Abra a conexão com o banco de dados
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/appwebgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicie uma transação
	transacao, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a declaração SQL para a inserção
	stmt, err := transacao.Prepare("INSERT INTO usuario(nome, email, contacto, senha) VALUES(?, ?, ?, ?)")
	if err != nil {
		transacao.Rollback()
		log.Fatal(err)
	}
	defer stmt.Close()

	// encriptando a senha usando o bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar hash da senha:", err)
		return
	}

	// Execute a declaração SQL, deixando que o banco de dados atribua um valor para a coluna 'id'
	result, err := stmt.Exec(nome, email, contacto, hash)
	if err != nil {
		transacao.Rollback()
		log.Fatal(err)
	}

	// Obtenha o ID gerado automaticamente, se necessário
	id, _ := result.LastInsertId()

	log.Print(id)

	// Commit da transação se tudo ocorrer bem
	transacao.Commit()
}

func (app *Application) VerificaUsuario(email, senha string) bool {
	// Abra a conexão com o banco de dados
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/appwebgo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare a declaração SQL para selecionar a senha correspondente ao email
	stmt, err := db.Prepare("SELECT senha FROM usuario WHERE nome = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute a declaração SQL para obter a senha do banco de dados
	var hash string
	err = stmt.QueryRow(email).Scan(&hash)
	if err == sql.ErrNoRows {
		fmt.Println("Usuário não encontrado")
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	// Comparar a senha fornecida com o hash armazenado
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	if err == nil {
		return true
	} else {
		fmt.Println("Senha incorreta!")
		return false
	}
}
