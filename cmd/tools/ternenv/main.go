package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Erro ao carregar o arquivo .env: %v", err))
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando tern: %v\nSaída do comando: %s", err, string(output))
	}

	fmt.Println(string(output))
}
