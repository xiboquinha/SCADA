API:
    CORS: 
        O CORS É UM MIDDLEWARE
        Cross-Origin Resource Sharing(Compartilhamento de recursos multi-plataforma)
        usa de cabeçalhos http para que  uma aplicação em um dominio possa acessar coisinhas de outro dominio

    Multiplexador:
        usado para rotear solicitações HTTP pra diferentes manipuladores com base no URL(pelo que eu entendi é isso que permite que o CORS aconteça)

    middleware:
        é um software que atua como uma ponte entre diferentes aplicações, permitindo que elas se comuniquem de maneira eficiente e inteligente, mas também gerencia o fluxo de tráfego pra evitar sobrecarga 
        pelo que eu entendi: o front nao precisa se comunicar com o back, ele se comunica com o middleware, que se comunica com o back

Go:

    fmt:
        para formatação de texto(tipo o stdio)

    log:
        para registrar erros ou informações

    strconv:
        Para conversões de tipos primitivos

    encoding/json:
        para codificar e decodificar jsons
        
    http:
        ResponseWriter:
            w http.ResponseWriter é um obj que representa uma resposta HTTP que será enviada ao cliente
            é uma interface
            type ResponseWriter interface
            // Header returns the header map that will be sent by
            // [ResponseWriter.WriteHeader]. The [Header] map also is the mechanism with which
            // [Handler] implementations can set HTTP trailers.
            //
            // Changing the header map after a call to [ResponseWriter.WriteHeader] (or
            // [ResponseWriter.Write]) has no effect unless the HTTP status code was of the
            // 1xx class or the modified headers are trailers.
            //
            // There are two ways to set Trailers. The preferred way is to
            // predeclare in the headers which trailers you will later
            // send by setting the "Trailer" header to the names of the
            // trailer keys which will come later. In this case, those
            // keys of the Header map are treated as if they were
            // trailers. See the example. The second way, for trailer
            // keys not known to the [Handler] until after the first [ResponseWriter.Write],
            // is to prefix the [Header] map keys with the [TrailerPrefix]
            // constant value.
            //
            // To suppress automatic response headers (such as "Date"), set
            // their value to nil.
            Header() Header

            // Write writes the data to the connection as part of an HTTP reply.
            //
            // If [ResponseWriter.WriteHeader] has not yet been called, Write calls
            // WriteHeader(http.StatusOK) before writing the data. If the Header
            // does not contain a Content-Type line, Write adds a Content-Type set
            // to the result of passing the initial 512 bytes of written data to
            // [DetectContentType]. Additionally, if the total size of all written
            // data is under a few KB and there are no Flush calls, the
            // Content-Length header is added automatically.
            //
            // Depending on the HTTP protocol version and the client, calling
            // Write or WriteHeader may prevent future reads on the
            // Request.Body. For HTTP/1.x requests, handlers should read any
            // needed request body data before writing the response. Once the
            // headers have been flushed (due to either an explicit Flusher.Flush
            // call or writing enough data to trigger a flush), the request body
            // may be unavailable. For HTTP/2 requests, the Go HTTP server permits
            // handlers to continue to read the request body while concurrently
            // writing the response. However, such behavior may not be supported
            // by all HTTP/2 clients. Handlers should read before writing if
            // possible to maximize compatibility.
            Write([]byte) (int, error)

            // WriteHeader sends an HTTP response header with the provided
            // status code.
            //
            // If WriteHeader is not called explicitly, the first call to Write
            // will trigger an implicit WriteHeader(http.StatusOK).
            // Thus explicit calls to WriteHeader are mainly used to
            // send error codes or 1xx informational responses.
            //
            // The provided code must be a valid HTTP 1xx-5xx status code.
            // Any number of 1xx headers may be written, followed by at most
            // one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
            // headers may be buffered. Use the Flusher interface to send
            // buffered data. The header map is cleared when 2xx-5xx headers are
            // sent, but not with 1xx headers.
            //
            // The server will automatically send a 100 (Continue) header
            // on the first read from the request body if the request has
            // an "Expect: 100-continue" header.
            WriteHeader(statusCode int)

        r * http.Request:
            um ponteiro ao objeto Request HTTP com todas as informações sobre a solicitação recebida pelo servidor

        mux := http.NewServerMux():
            serve pra criar um multiplexador(api -> multiplexador) HTTP.

        mux.HandleFunc("/caminho/url", func): 
            registra um manipulador pra func pra lidar com solicitações pro caminho usado
            pelo que eu entendi ele binda uma função específica(segundo argumento) pra um caminho específico(primeiro argumento)

        http.ListenAndServe(":portaqvcquer", corshandler):
            inicia um sv na porta que vc passou, e usa o corshandler(cors -> corshandler) para configurar o cors para as solicitações feitas nessa porta.

    database/sql:
        serve para interagir com o db SQL com o  SQLite3

        db, err := sql.Open("sqlite3", "nome do db"):
            abre a conexao com o db sem se conectar com ele(só vai conectar qnd fizer um query)

        para executar um query:
            _, err := db.EXEC(`COMANDO SQL`)
            note que o comando está entre acentos graves

    github.com/mattn/go-sqlite3:
        para drivers do sqlite3


    github.com/rs/cors:

        corshandler := cors.Default().Handler(mux):
            aqui estou configurando o CORS para o multiplexador já criado, ou seja, todas as solicitações passadas através desse multiplexador vão ter o CORS configurado automaticamente

    


JS:
    declaração de func:
        function nomefunc(){}

    React:

        pra usar o react vc tem que importar ele, igual uma lib(no site do react ele diz q e uma lib kkkkkkk slk)

        useState:
            é um hook(nao entendi direito oq é um hook mas acho que é tipo uma lib)
            usado para adicionar estado local ao componente, retorna 2 valores: o valor atual do estado, e uma func que permite atualizar o valor dele
                estado local:
                    capacidade de um componente de armazenar infos que podem mudar ao longo do tempo e afetar o que é renderizado no DOM
                        DOM:
                            document Object Model
                            é uma representação estruturada do conteúdo de uma página na web, incluindo textos,imagens, links, scripts, e CSS
        useEffect:
            também é um hook
            é usado para executar efeitos colaterais após a renderização do componente


SQL:

    CREATE TABLE nomedatabela(atributo tipoatributo PRIMARY KEY(opcional) AUTOINCREMENT(se quiser que some 1 unidade a cada um criada NOT NULL(se n puder ser nulo))):

        cria tabela né kkkkkkkkk
        não sei se é exclusivo do sqlite, mas da pra dar um CREATE TABLE IF NOT EXISTS, que ai só cria se nao existir KKKKKKKKKKKKKKKKKK

    "INSERT INTO nomedatabela(atributo, atributo) VALUES (?,?)", variavel, variavel:
        o nome ja é explicativo, mas a interrogação funciona igual o %d no go, o em c, é o valor que vai ser substituido

    

    



    