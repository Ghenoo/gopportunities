# GoJob Opportunities API

<p align="center">
  <img src="./assets/gopportunities.png" alt="GoJob Header">
</p>

Este projeto é uma API moderna de oportunidades de emprego construída usando Golang. A API é desenvolvida por Go-Gin como roteador, GoORM para comunicação de banco de dados, SQLite como banco de dados e Swagger para documentação e teste de API. O projeto segue uma estrutura de pacotes moderna para manter a base de código organizada e de fácil manutenção.

---

## Características

- Introdução ao Golang e construção de APIs modernas
- Configuração do ambiente de desenvolvimento para criação da API
- Usando Go-Gin como roteador para gerenciamento de rotas
- Implementando SQLite como banco de dados para a API
- Utilização do GoORM para comunicação com o banco de dados
- Integração do Swagger para documentação e testes de API
- Criação de uma estrutura de pacotes moderna para organização do projeto
- Implementação de uma API completa de oportunidades de emprego com endpoints para busca, criação, edição e exclusão de oportunidades
- Implementação de testes automatizados para garantir a qualidade da API

## Live version

Você pode checar a documentação e realizar testes na API visitando a versão ativa hosteada em []

## Installation

To use this project, you need to follow these steps:

1. Clone o repositório: `git clone https://github.com/Ghenoo/gopportunities.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`


## Makefile Commands

O projeto inclui um Makefile para ajudar você manejar tarefas comuns mais facil. Veja a lista de comandos e uma breve descrição do que elas fazem:

- `make run`: Executa a aplicação generando a documentação da API.
- `make run-with-docs`: Gera a documentação API usando Swag, após executar a aplicação.
- `make build`: Constrói a aplicação e cria um arquivo executável com o nome de `gopportunities`.
- `make test`: Executa testes em todos os packages do projeto.
- `make docs`: Gera a documentação da API usando Swag.
- `make clean`: Remove o `gopportunities` executável e deleta o `./docs` do diretório.

```sh
make run
```

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

Para mais informações sobre Docker e Docker Compose, entre no link da documentação oficial:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Ferramentas Usadas

Esse projeto uso e seguiu as seguintes ferramentas:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

Esse projeto foi inspirado no [arthur404dev](https://github.com/arthur404dev).