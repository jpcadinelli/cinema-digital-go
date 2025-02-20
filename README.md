# Controle de Estoque

Construindo um controle de estoque baseadp no padrão de [API em golang](https://github.com/jpcadinelli/api-pattern-go) que criei.

Para fazer login e ter acesso geral ao sistema utilize o usuário:
 - email: admin_d@_sistema.com
 - senha: Admin!123

Ideia da arquitetura do sistema

api/
│── app/                    # Módulos principais do sistema

│   ├── cinema/             # Gerenciamento de cinemas, salas, filmes e sessões

│   │   ├── model/          # Estruturas (structs) de Cinema, Sala, Filme, Sessão

│   │   ├── repository/     # Operações de banco de dados para cinema

│   │   ├── resource/       # Handlers e controllers dos endpoints de cinema

│   ├── usuario/            # Gerenciamento de usuários

│   │   ├── model/          # Estruturas de usuário

│   │   ├── repository/     # Banco de dados para usuários

│   │   ├── resource/       # Handlers e controllers para usuários

│   ├── permissao/          # Gerenciamento de permissões

│   │   ├── model/          # Estruturas de permissões

│   │   ├── repository/     # Banco de dados para permissões

│   │   ├── resource/       # Handlers e controllers para permissões

│

│── pkg/                    # Pacotes compartilhados

│   ├── database/           # Configuração do banco de dados

│   ├── middleware/         # Middlewares (autenticação, logs, etc.)

│   ├── security/           # Segurança (JWT, criptografia)

│   ├── utils/              # Funções auxiliares

│

│── routes/                 # Definição de rotas da API

│── main.go                 # Ponto de entrada da aplicação

