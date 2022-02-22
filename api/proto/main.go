package main

import (
	"Ant-Man-Url/api/handler/UrlHandler"
	"Ant-Man-Url/api/handler/UserHandler"
	"Ant-Man-Url/api/middleware"
	Antman "Ant-Man-Url/api/proto/UrlProto"
	Antman2 "Ant-Man-Url/api/proto/UserProto"
	"Ant-Man-Url/infrastructure/repository"
	"Ant-Man-Url/usecase/Url"
	"Ant-Man-Url/usecase/user"
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {

	StartServer()
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func StartServer() {
	dbUrl := "postgres://vishwajeet:vishvapriya123@localhost:5432/keystore"
	connPgx, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		println("Errrorr...")
	}
	defer connPgx.Close(context.Background())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverOpt := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.Unary()),
	}
	s := grpc.NewServer(serverOpt...)

	urlRepo := repository.NewUrlMapSql(connPgx)
	urlServ := Url.NewService(urlRepo)

	userRepo := repository.NewUserSql(connPgx)
	userServ := user.NewService(userRepo)

	Antman.RegisterAntmanUrlRoutesServer(s, UrlHandler.NewUrlServer(urlServ))
	Antman2.RegisterAntmanUserRoutesServer(s, UserHandler.NewUserServer(userServ))
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = Antman.RegisterAntmanUrlRoutesHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = Antman2.RegisterAntmanUserRoutesHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = gwmux.HandlePath("GET", "/{keyval}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		var keyval = pathParams["keyval"]
		var longUrl = ""

		tx, errs := connPgx.Begin(context.Background())
		if errs != nil {
			return
		}
		row := tx.QueryRow(context.Background(), "SELECT url from keys_out where keyval = $1 limit 1", keyval)
		errs = row.Scan(&longUrl)
		if errs != nil {
			return
		}
		http.Redirect(w, r, longUrl, http.StatusSeeOther)
	})
	if err != nil {
		panic(err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
