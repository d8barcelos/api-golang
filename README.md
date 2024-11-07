Crud em Go
Este é um guia abrangente para o projeto, uma aplicação de exemplo que implementa operações básicas de CRUD (Create, Read, Update, Delete) para usuários. O projeto inclui um Dockerfile para facilitar a execução em contêineres.

Informações
Título: Crud em Go
Versão: 1.0
Host: localhost:8080

Pré-requisitos
Antes de começar, certifique-se de que possui os seguintes pré-requisitos instalados em seu sistema:

Go: A linguagem de programação Go.
Docker: Necessário caso deseje executar a aplicação em um contêiner.
Instalação
Siga os passos abaixo para instalar o projeto em seu ambiente de desenvolvimento:

Clone o repositório:

bash
Copiar código
git clone https://github.com/d8barcelos/api-golang
Navegue até o diretório do projeto:

bash
Copiar código
cd api-golang
Construa a aplicação usando o Docker Compose:

bash
Copiar código
docker compose up
Executando a Aplicação
Após a instalação, você pode executar a aplicação com o seguinte comando (caso deseje executá-la diretamente com Golang):

bash
Copiar código
docker container run --name meuprimeirocrudgo -p 27017:27017 -d mongo
go run main.go
A aplicação estará acessível em http://localhost:8080.

Testando a Aplicação
Se preferir, após executar o projeto, visite: http://localhost:8080/swagger/index.html# para ver e testar todos os contratos de rota.

A aplicação oferece endpoints REST para criar, listar, atualizar e excluir usuários. Você pode utilizar ferramentas como curl ou Postman para testar os endpoints. Aqui estão alguns exemplos de comandos curl para testar os endpoints:

Criar um usuário:

curl -X POST -H "Content-Type: application/json" -d '{"name": "João", "email": "joao@example.com", "age": 30, "password": "password$#@$#323"}' http://localhost:8080/createUser
Atualizar um usuário:

curl -X PUT -H "Content-Type: application/json" -d '{"name": "João Silva"}' http://localhost:8080/updateUser/{userId}
Excluir um usuário:

curl -X DELETE http://localhost:8080/deleteUser/{userID}
Lembre-se de ajustar os comandos conforme suas necessidades e requisitos.

Modelos de Dados
request.UserLogin
Estrutura contendo os campos necessários para login do usuário.

email (string, obrigatório): O email do usuário (deve ser um email válido).
password (string, obrigatório): A senha do usuário (deve ter no mínimo 6 caracteres e conter pelo menos um dos caracteres: !@#$%*).
request.UserRequest
Estrutura contendo os campos obrigatórios para criar um novo usuário.

age (inteiro, obrigatório): A idade do usuário (deve ser entre 1 e 140).
email (string, obrigatório): O email do usuário (deve ser um email válido).
name (string, obrigatório): O nome do usuário (deve ter no mínimo 4 caracteres e no máximo 100 caracteres).
password (string, obrigatório): A senha do usuário (deve ter no mínimo 6 caracteres e conter pelo menos um dos caracteres: !@#$%*).
request.UserUpdateRequest
Estrutura contendo os campos para atualizar as informações do usuário.

age (inteiro, obrigatório): A idade do usuário (deve ser entre 1 e 140).
name (string, obrigatório): O nome do usuário (deve ter no mínimo 4 caracteres e no máximo 100 caracteres).
response.UserResponse
Estrutura de resposta contendo as informações do usuário.

age (inteiro): A idade do usuário.
email (string): O email do usuário.
id (string): O ID único do usuário.
name (string): O nome do usuário.
rest_err.Causes
Estrutura que representa as causas de um erro.

field (string): O campo associado à causa do erro.
message (string): Mensagem de erro descrevendo a causa.
rest_err.RestErr
Estrutura que descreve por que ocorreu um erro.

causes (array de rest_err.Causes): Causas do erro.
code (inteiro): Código do erro.
error (string): Descrição do erro.
message (string): Mensagem de erro.
Endpoints
Nota: Para autenticação, você deve incluir o token de acesso no cabeçalho Authorization como "Bearer " para endpoints protegidos.

A API oferece os seguintes endpoints:

POST /createUser

Descrição: Cria um novo usuário com as informações fornecidas.
Parâmetros:
userRequest (body, obrigatório): Informações do usuário para o registro.
Respostas:
200: OK (Usuário criado com sucesso)
400: Bad Request (Erro de solicitação)
500: Internal Server Error (Erro interno do servidor)
DELETE /deleteUser/{userId}

Descrição: Exclui um usuário com base no ID fornecido.
Parâmetros:
userId (path, obrigatório): ID do usuário a ser excluído.
Respostas:
200: OK (Usuário excluído com sucesso)
400: Bad Request (Erro de solicitação)
500: Internal Server Error (Erro interno do servidor)
GET /getUserByEmail/{userEmail}

Descrição: Recupera os detalhes do usuário com base no email fornecido.
Parâmetros:
userEmail (path, obrigatório): Email do usuário a ser recuperado.
Respostas:
200: Informações do usuário recuperadas com sucesso
400: Erro: ID de usuário inválido
404: Usuário não encontrado
GET /getUserById/{userId}

Descrição: Recupera os detalhes do usuário com base no ID fornecido.
Parâmetros:
userId (path, obrigatório): ID do usuário a ser recuperado.
Respostas:
200: Informações do usuário recuperadas com sucesso
400: Erro: ID de usuário inválido
404: Usuário não encontrado
POST /login

Descrição: Permite que um usuário faça login e receba um token de autenticação.
Parâmetros:
userLogin (body, obrigatório): Credenciais de login do usuário.
Respostas:
200: Login bem-sucedido, token de autenticação fornecido
403: Erro: Credenciais de login inválidas
PUT /updateUser/{userId}

Descrição: Atualiza os detalhes do usuário com base no ID fornecido.
Parâmetros:
userId (path, obrigatório): ID do usuário a ser atualizado.
userRequest (body, obrigatório): Informações do usuário para atualização.
Respostas:
200: OK (Usuário atualizado com sucesso)
400: Bad Request (Erro de solicitação)
500: Internal Server Error (Erro interno do servidor)
Contribuindo
Se deseja contribuir para o projeto MeuPrimeiroCRUD em Go, sinta-se à vontade para enviar pull requests ou relatar problemas no repositório oficial.

Licença
Este projeto é distribuído sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.

Esperamos que esta documentação Swagger tenha sido útil para entender e interagir com a API do projeto MeuPrimeiroCRUD em Go. Se tiver dúvidas ou precisar de suporte adicional, não hesite em nos contatar. Aproveite a API!