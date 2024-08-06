# GO Application WEB

## Projeto de aprendizagem de uma web app utilizando GO.

Esse projeto foi desenvolvido para se aprender como criar uma aplicação web utilizando GO + HTML + JS,
utilizamos o padrao de MVC (Model, View, Controller) para criar essa aplicação.

Seguindo a seguinte base 

Routes > Controller > Models > Base > Models > Controller > View

Cada um tem sua responsabilidade no projeto ou seja.

### Routes
É responsável por inicializar  as rotas  que são chamadas e os métodos que vao ser utilizados nos controllers(endpoints).

### Controller
É o responsável por ver a requisição do usuario, o tipo dela, e o que deve ser entregue para o cliente,
chamando a classe de Serviço que executara toda regra de negócio para nós, que nesse projeto a regra está dentro de models.

### Models
É o modulo que executa todas ações no banco que tenha relação com o tipo Produto, deletar, alterar, inserir ou Atualizar.

### Base
É onde o modelo bate para conseguir as informações sobre o Produto e devolve o mesmo para ele, 
em nosso package db, só temos as configuraçoes de conexão com o banco.