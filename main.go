package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("GET", "/fer", HandleFer)
	server.Listen()

}
