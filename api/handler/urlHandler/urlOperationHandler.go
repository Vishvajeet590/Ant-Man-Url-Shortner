package urlHandler

//

import (
	Antman "Ant-Man-Url/api/proto/UrlProto"
	"Ant-Man-Url/entity"
	"Ant-Man-Url/usecase/url"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"strconv"
)

type UrlServer struct {
	Antman.UnsafeAntmanUrlRoutesServer
	service *url.Service
}

func NewUrlServer(service *url.Service) Antman.AntmanUrlRoutesServer {
	return &UrlServer{
		service: service,
	}
}

func (s *UrlServer) GetShortUrl(ctx context.Context, req *Antman.ShortUrlRequest) (*Antman.ShortUrlResponse, error) {
	var shortUrl *entity.Url
	var err error
	isJwt := true
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md["check-key"]) == 0 {
		isJwt = false
	}
	check := md["check-key"]
	if check[0] == "false" {
		isJwt = false
	}
	log.Printf("Starting to short URL")
	longUrl := &entity.Url{
		Long_link: req.OriginalUrl,
		Keyval:    "",
		Redirects: -999,
		OwnerId:   -999,
	}
	if isJwt {
		id := md["id-key"]
		userId, err := strconv.Atoi(id[0])
		if err != nil {
			return nil, err
		}
		shortUrl, err = s.service.MapUrl(longUrl, isJwt, &userId)
	} else {
		shortUrl, err = s.service.MapUrl(longUrl, isJwt, nil)
	}

	if err != nil {
		return &Antman.ShortUrlResponse{
			Success:     false,
			ShortUrlKey: "",
		}, err
	}

	return &Antman.ShortUrlResponse{
		Success:     true,
		ShortUrlKey: "https://urlants.herokuapp.com/" + shortUrl.Keyval,
	}, err

}

func (s *UrlServer) GetLongUrl(ctx context.Context, req *Antman.LongUrlRequest) (*Antman.LongUrlResponse, error) {
	log.Printf("Starting to FETCH long url...\n")
	shortUrl := &entity.Url{
		Long_link: "",
		Keyval:    req.ShortUrl,
		Redirects: -999,
		OwnerId:   -999,
	}

	LongUrl, err := s.service.ResolveUrl(shortUrl)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
		return &Antman.LongUrlResponse{LongUrl: ""}, err
	}
	return &Antman.LongUrlResponse{LongUrl: LongUrl.Long_link}, nil

}

func (s *UrlServer) DeleteUrl(ctx context.Context, req *Antman.DeleteUrlRequest) (*Antman.DeleteUrlResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	id := md["id-key"]
	userId, err := strconv.Atoi(id[0])
	if err != nil {
		return nil, err
	}

	log.Printf("Starting to DELETE url...\n")

	shortUrl := &entity.Url{
		Long_link: "",
		Keyval:    req.ShortUrl,
		Redirects: -999,
		OwnerId:   -999,
	}

	res, err := s.service.DeleteUrl(shortUrl, true, &userId)
	if res == false {
		log.Printf("Error : %s", err)
		return &Antman.DeleteUrlResponse{Success: res}, err
	}
	return &Antman.DeleteUrlResponse{Success: res}, nil
}
