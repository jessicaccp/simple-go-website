<!DOCTYPE html>
<html>
    <head>
    	<title>Página Inicial</title>
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
		</style>
    </head>
    <body>
    	<div class="content">
			<p><img src="https://images2.imgbox.com/fb/a9/QwnBLedZ_o.gif"></p>
			<p><a class="button" href="http://localhost:8080/cadastro">Cadastrar-se</a></p>
			<p><a class="button" href="http://localhost:8080/cadastrados">Cadastros realizados</a></p>
		</div>
    </body>
</html>
