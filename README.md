## Pismo-Payments
<p align="center">O objetivo do projeto é simular um mini sistema de transações financeiras</p>
<h4 align="left"> 
	🚧  Interface com o banco de dados para facilitar os testes esta desenvolvimento...  🚧
</h4>

### Features

- [x] Cadastro de contas
- [x] Cadastro de transações
- [x] Mostrar conta



Para rodar o teste é necessário executar os seguintes comandos
```bash
make create_postres (criara a imagem docker com o banco postgresql)
make create_db (criara a tabela)
make migrate_up (ira inserir dados e schemas básicos do projeto)

```
Com isso estamos, em ordem, criara a imagem docker com o banco postgresql,
criara a tabela e por ultimo ira inserir dados e schemas básicos do projeto


Após a realização dos passos anteriores, podemos rodar os testes( coverage abaixo de 10%) com o comando 
make test.
Para rodar o projeto local basta rodar o seguinte:
```bash
make server
```


A seguir temos o exemplo do design do banco de dados utilizado, nesse caso foi o Postgres.

![image](https://user-images.githubusercontent.com/18687651/226506899-a949ba16-38ce-4f1e-b7bc-f81a038a9a12.png)

