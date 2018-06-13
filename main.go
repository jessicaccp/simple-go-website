package main

// bibliotecas
import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
	"os"
	"log"
)

// pagina de clientes cadastrados
func listaClientes(w http.ResponseWriter, r *http.Request) {
	// se o caminho nao for / (raiz), exibe pagina de erro
	if r.URL.Path != "/cadastrados" {
		paginaErro(w, r, http.StatusNotFound)
		return
	}

	// faz a leitura dos dados do arquivo onde foram salvos os cadastros
	dados, err := ioutil.ReadFile("clientes.txt")
	
	// transforma de []byte para string
	sDados := string(dados)
	buttons := `<p><a class="button" href="http://localhost:8080/">Página inicial</a></p><p><a class="button" href="http://localhost:8080/cadastro">Novo cadastro</a></p>`
		
	// trata um possivel erro na leitura do arquivo
    if err != nil {
    	log.Fatal(err)
    }
    
    // define o template da pagina e executa
	t, _ := template.ParseFiles("cadastrados.gtpl")
	t.Execute(w, nil)
	
	// exibe os dados cadastrados na pagina
	fmt.Fprintf(w, sDados)
	fmt.Fprintf(w,	buttons)
}

// pagina de erro
func paginaErro(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, _ := template.ParseFiles("404.gtpl")
		t.Execute(w, nil)
	}
}

// pagina principal
func paginaPrincipal(w http.ResponseWriter, r *http.Request) {
	// se o caminho nao for / (raiz), exibe pagina de erro
	if r.URL.Path != "/" {
		paginaErro(w, r, http.StatusNotFound)
		return
	}
	
	// define o template da pagina e executa
	t, _ := template.ParseFiles("paginaInicial.gtpl")
	t.Execute(w, nil)
}

// estrutura usada para organizar os dados recebidos durante o cadastro
type Cadastro struct {
	Nome		string
	Usuario		string
	Email		string
	Senha		string
}

// pagina de cadastro
func cadastro(w http.ResponseWriter, r *http.Request) {
	// se o caminho nao for /cadastro, exibe pagina de erro
	if r.URL.Path != "/cadastro" {
		paginaErro(w, r, http.StatusNotFound)
		return
	}

	// define o template da pagina
    t, _ := template.ParseFiles("cadastro.html")
    
    // exibe no terminal o metodo de requisicao http
	fmt.Println("\nMétodo HTTP:", r.Method)
	
	// se metodo http GET, exibe o formulario de cadastro
    if r.Method == "GET" {
        t.Execute(w, nil)
    
    // se metodo http POST, lida com os dados cadastrais recebidos
    } else {
    	// atualiza os valores de r.Form
        r.ParseForm()
        
        // exibe no terminal os dados recebidos
        fmt.Println("Nome:",			r.Form["nome"])
        fmt.Println("Nome de usuário:",	r.Form["usuario"])
        fmt.Println("E-mail:",			r.Form["email"])
        fmt.Println("Senha:",			r.Form["senha"])
        
        // organiza os dados do cliente
       	cliente := Cadastro {
			Nome:		r.FormValue("nome"),
			Usuario:	r.FormValue("usuario"),
			Email:		r.FormValue("email"),
			Senha:		r.FormValue("senha"),
		}
		
		// caso exista o arquivo, salva os dados no final
		// caso nao exista, cria um novo arquivo e adiciona os dados
		f, err := os.OpenFile("clientes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		
		// tratamento de possivel erro na abertura do arquivo
		if err != nil {
			log.Fatal(err)
		}
		
		// escreve os dados no arquivo
		// e trata possivel erro de escrita no arquivo
		if _, err := f.Write([]byte("Nome: "			+ cliente.Nome + "<br />\n" +
									"Nome de usuário: "	+ cliente.Usuario + "<br />\n" +
        							"E-mail: "			+ cliente.Email + "<br />\n" +
        							"Senha: "			+ cliente.Senha + "<br /><br />\n\n")); err != nil {
			log.Fatal(err)
		}
		
		// fecha o arquivo
		// e trata possivel erro de arquivo
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	    
	    // mostra na pagina a mensagem de que o cadastro foi bem sucedido
        t.Execute(w, struct{ Success bool }{true})
    }
}


func main() {
	// define as paginas do site
	http.HandleFunc("/", paginaPrincipal)
	http.HandleFunc("/cadastro", cadastro)
	http.HandleFunc("/cadastrados", listaClientes)

	// inicia o servidor http na porta 8080
	// caso nao funcione, log.Fatal() printa o erro no terminal e termina o programa
	log.Fatal(http.ListenAndServe(":8080", nil))
}

