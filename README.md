# Como usar esse Script
Crie um arquivo chamado `secrets.txt` com o formato de conteúdo parecido com o abaixo:
```
key1=value1
key2=value2
```

Certifique-se de que o Vault está em execução e você tem um token válido. 

Este Token bem deve estar inserido no Código substituindo o campo <TOKEN>:
```
token := "<TOKEN>"
```
Perceba que outras estrtuturas dependem desse customização, atente-se:
```
secretPath := "secret/data/<PATH_TO_KV>"
filePath := "dba.txt"
vaultAddress := "http://<VAULT_ADDRESS>:8200"
```
Note que acima você deve substituir os campos entre <>.

Após isso configurado execute o código:
```
go run main.go
```

## O que o código faz?
Este código lê um arquivo de texto `secrets.txt` onde cada linha contém uma chave e um valor separados por =, cria um mapa com esses dados, e os escreve no Vault no caminho especificado no `secretPath`.