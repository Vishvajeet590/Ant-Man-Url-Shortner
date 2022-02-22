package api

import (
	"Ant-Man-Url/api/handler/urlHandler"
	"Ant-Man-Url/api/handler/userHandler"
	"Ant-Man-Url/api/middleware"
	Antman "Ant-Man-Url/api/proto/UrlProto"
	Antman2 "Ant-Man-Url/api/proto/UserProto"
	"Ant-Man-Url/infrastructure/repository"
	url "Ant-Man-Url/usecase/url"
	"Ant-Man-Url/usecase/user"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		return
	}

	connPgx, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		println("Errrorr...")
	}
	defer connPgx.Close(context.Background())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverOpt := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.Unary()),
	}
	s := grpc.NewServer(serverOpt...)

	urlRepo := repository.NewUrlMapSql(connPgx)
	urlServ := url.NewService(urlRepo)

	userRepo := repository.NewUserSql(connPgx)
	userServ := user.NewService(userRepo)

	Antman.RegisterAntmanUrlRoutesServer(s, urlHandler.NewUrlServer(urlServ))
	Antman2.RegisterAntmanUserRoutesServer(s, userHandler.NewUserServer(userServ))
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

	err = gwmux.HandlePath("GET", "/", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

		http.ServeFile(w, r, "./never.html")
	})
	if err != nil {
		panic(err)
	}

	add := fmt.Sprintf(":%s", port)
	fmt.Printf("%s", add)
	gwServer := &http.Server{
		Addr:    add,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0: %s", port)
	log.Fatalln(gwServer.ListenAndServe())

}
