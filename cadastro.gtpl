<!DOCTYPE html>
<html>
    <head>
		<title>Cadastro</title>
        <style>
           	html, body {
 			   height: 100%;
			}
			html {
				display: table;
				margin: auto;
				text-align: center;
			}

			body {
				display: table-cell;
				vertical-align: middle;
			}
			
			h1 {
				font-size: 100px;
			}
			
			.content {
				max-width: 1000px;
				margin: auto;
			}
			
			.inputs {
				max-width: 650px;
				margin: auto;
			}
			
			.button {
				background-color: #d9d9d9; 
				border: solid 1px #d9d9d9;
				color: black;
				padding: 15px 32px;
				text-align: center;
				text-decoration: none;
				display: inline-block;
			    width:200px;
				font-size: 17px;
			}
			.button:hover {
				background-color: white;
			}
			
			.input {
			    width: 100%;
			}
			
			.small_input {
			    width: 49%;
			}
		</style>
    </head>
    <body>
		<div class="content">
			{{if .Success}}
				<img src="https://images2.imgbox.com/c6/af/CUsv9zBq_o.gif" width="300px">
				<p style="font-size:25px">Cadastro realizado<br>com sucesso!</p>
				<p><a class="button" href="http://localhost:8080/cadastro">Novo cadastro</a></p>
				<p><a class="button" href="http://localhost:8080/cadastrados">Cadastros realizados</a></p>
				<p><a class="button" href="http://localhost:8080/">Página inicial</a></p>
			{{else}}
				<h1>CADASTRO</h1>
				<div class="inputs">
					<form action="/cadastro" method="post">
						<input class="input" type="text" placeholder="Nome:" name="nome" required><br /><br />
						<input class="input" type="text" placeholder="Nome de usuário:" name="usuario" required><br /><br />
						<input class="input" type="text" placeholder="E-mail:" name="email" required><br /><br />
						<input class="input" type="password" placeholder="Senha:" name="senha" required><br /><br />
						<center><input class="small_input" type="submit" value="Enviar"> 
						<input class="small_input" onClick="window.location.href='http://localhost:8080/'" type="submit" Value="Cancelar"></center>
					</form>
				</div>
			{{end}}
		</div>
    </body>
</html>
