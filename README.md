📘 Crud em Go
Este repositório contém uma aplicação de exemplo em Go que implementa operações básicas de CRUD (Create, Read, Update, Delete) para gerenciamento de usuários. A aplicação é acompanhada de um Dockerfile para facilitar a execução em contêineres.

📌 Informações do Projeto
Título: Crud em Go
Versão: 1.0
Host: localhost:8080
🚀 Pré-requisitos
Certifique-se de que seu ambiente possui os seguintes itens instalados:

Go: A linguagem de programação Go.
Docker: Necessário para executar a aplicação em um contêiner.
🛠️ Instalação
1. Clone o Repositório
bash
Copiar código
git clone https://github.com/d8barcelos/api-golang
2. Navegue até o Diretório do Projeto
bash
Copiar código
cd api-golang
3. Construa a Aplicação com Docker Compose
bash
Copiar código
docker compose up
💻 Executando a Aplicação
Após instalar as dependências, você pode executar a aplicação diretamente com Go:

bash
Copiar código
docker container run --name meuprimeirocrudgo -p 27017:27017 -d mongo
go run main.go
A aplicação estará acessível em http://localhost:8080.

🧪 Testando a Aplicação
Acesse a documentação completa dos endpoints em: Swagger UI. Utilize ferramentas como curl ou Postman para testar os endpoints.

Exemplos de Comandos curl
Criar um usuário:

bash
Copiar código
curl -X POST -H "Content-Type: application/json" -d '{"name": "João", "email": "joao@example.com", "age": 30, "password": "password$#@$#323"}' http://localhost:8080/createUser
Atualizar um usuário:

bash
Copiar código
curl -X PUT -H "Content-Type: application/json" -d '{"name": "João Silva"}' http://localhost:8080/updateUser/{userId}
Excluir um usuário:

bash
Copiar código
curl -X DELETE http://localhost:8080/deleteUser/{userId}
📑 Modelos de Dados
request.UserLogin
Campos para login do usuário:

email (string, obrigatório): Email válido.
password (string, obrigatório): Mínimo 6 caracteres, incluindo um dos caracteres especiais !@#$%*.
request.UserRequest
Campos obrigatórios para criar um novo usuário:

age (inteiro, obrigatório): Idade entre 1 e 140.
email (string, obrigatório): Email válido.
name (string, obrigatório): Nome entre 4 e 100 caracteres.
password (string, obrigatório): Mínimo 6 caracteres, incluindo !@#$%*.
response.UserResponse
Campos de resposta com informações do usuário:

age (inteiro): Idade do usuário.
email (string): Email do usuário.
id (string): ID único do usuário.
name (string): Nome do usuário.
🔗 Endpoints
POST /createUser
Descrição: Cria um novo usuário.
Parâmetros: userRequest (body, obrigatório)
Respostas: 200 OK, 400 Bad Request, 500 Internal Server Error
DELETE /deleteUser/{userId}
Descrição: Exclui um usuário pelo ID.
Parâmetros: userId (path, obrigatório)
Respostas: 200 OK, 400 Bad Request, 500 Internal Server Error
GET /getUserByEmail/{userEmail}
Descrição: Recupera um usuário pelo email.
Parâmetros: userEmail (path, obrigatório)
Respostas: 200 OK, 400 Bad Request, 404 Not Found
GET /getUserById/{userId}
Descrição: Recupera um usuário pelo ID.
Parâmetros: userId (path, obrigatório)
Respostas: 200 OK, 400 Bad Request, 404 Not Found
POST /login
Descrição: Realiza login e retorna um token.
Parâmetros: userLogin (body, obrigatório)
Respostas: 200 OK, 403 Forbidden
PUT /updateUser/{userId}
Descrição: Atualiza as informações de um usuário.
Parâmetros: userId (path, obrigatório), userRequest (body, obrigatório)
Respostas: 200 OK, 400 Bad Request, 500 Internal Server Error
🤝 Contribuindo
Contribuições são bem-vindas! Sinta-se à vontade para abrir um pull request ou relatar problemas.

📜 Licença
Distribuído sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

Esperamos que esta documentação ajude você a explorar e interagir com a API do projeto. Em caso de dúvidas, entre em contato. Aproveite! 🎉
