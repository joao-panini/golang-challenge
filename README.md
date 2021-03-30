# golang-challenge

Install
  Go
  Docker
  docker-compose
  
 Usage
 
    cd $GOPATH/src

    git clone https://github.com/joao-panini/golang-challenge

    cd golang-challenge

    docker-compose up
    
Set up database

    docker exec -it goapp-challenge-database bash -l
    
    mysql -u root -p - password: panini
    
    run commands on sql.sql
    
Endpoints

    localhost:6000/accounts (Post Method) - Cria uma conta no banco
    {
      "name":"...",
      "cpf":"...",
      "secret":"...",
    }
    
    localhost:6000/accounts (Get Method) - Retorna todas as contas 
    
    localhost:6000/login (Post Method) - Retorna o token de autenticação para o usuario
    {
      "cpf": "..."
      "secret": "..."
    }
    
    localhost:6000/transfers (Put method) - Faz uma transferencia da conta logada para uma conta destino.
    
    *Para realizar transferencias utilize as contas criadas no arquivo sql.sql pois as contas novas são salvas com o saldo 0 no banco.
    *Senha para as contas criadas: 1234
    {     
      "ToAccountId": 1,
      "amount": 100.0
    }
    
    localhost:6000/transfers (Get Method) - Retorna todas as transações da conta logada.
    
    
    
    
