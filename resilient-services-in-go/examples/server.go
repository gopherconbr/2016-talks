type Server struct {
        Addr         string        // TCP address to listen on, ":http" if empty
        Handler      Handler       // handler to invoke, http.DefaultServeMux if nil
        ReadTimeout  time.Duration // maximum duration before timing out read of the request
        WriteTimeout time.Duration // maximum duration before timing out write of the response
}
