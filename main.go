package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

func main() {

	secretPath := "secret/data/<PATH_TO_KV>"
	filePath := "dba.txt"
	vaultAddress := "http://<VAULT_ADDRESS>:8200"
	token := "<TOKEN>"

	config := &vault.Config{
		Address: vaultAddress,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Erro ao criar cliente do Vault: %v", err)
	}

	client.SetToken(token)

	// Abrir e ler o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	secrets := make(map[string]interface{})

	// Ler cada linha do arquivo
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Linha malformada: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		secrets[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	existingSecret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Fatalf("Erro ao verificar a existência do caminho no Vault: %v", err)
	}

	if existingSecret == nil {
		fmt.Printf("O caminho %s não existe, será criado.\n", secretPath)
	} else {
		fmt.Printf("O caminho %s já existe, será atualizado.\n", secretPath)
	}

	// Armazenar os secrets no Vault
	_, err = client.Logical().Write(secretPath, map[string]interface{}{
		"data": secrets,
	})
	if err != nil {
		log.Fatalf("Erro ao escrever no Vault: %v", err)
	}

	fmt.Println("Secrets armazenados com sucesso!")
}
