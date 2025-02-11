# Blacklist Project

## Description
Blacklist system to manage who can and cannot buy tickets at specific events or globally, ensuring flexibility, performance and scalability.

## Tecnologias Utilizadas
- Go
- Gin Web framework
- raabbitMQ
- Banco de Dados (PostgreSQL, Redis)
- dig Is a Dependency Injection (DI) library for Go, developed by Uber. Dig helps you manage dependencies declaratively, allowing you to register types and their dependencies, and the framework itself takes care of building the necessary objects.

## Features
- Add blaacklist
- Chek blacklist
- Remove blacklist
- generate reports 
- API para interação com o sistema
- Integração com banco de dados

## Instalação
1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/blacklist.git
   cd blacklist
   ```

2. Crie um ambiente virtual e instale as dependências:
   ```sh
   python -m venv venv
   source venv/bin/activate  # No Windows use: venv\Scripts\activate
   pip install -r requirements.txt
   ```

3. Configure as variáveis de ambiente:
   ```sh
   export DATABASE_URL="sqlite:///blacklist.db"  # Altere conforme o banco utilizado
   ```

4. Execute a aplicação:
   ```sh
   python main.py
   ```

## Uso
- Para adicionar um item à blacklist via API:
  ```sh
  curl -X POST http://localhost:5000/blacklist -H "Content-Type: application/json" -d '{"item": "192.168.1.1"}'
  ```

- Para verificar se um item está na blacklist:
  ```sh
  curl -X GET http://localhost:5000/blacklist/192.168.1.1
  ```

## Contribuição
Sinta-se à vontade para abrir issues e enviar pull requests com melhorias e correções.

## Licença
Este projeto está licenciado sob a [MIT License](LICENSE).

