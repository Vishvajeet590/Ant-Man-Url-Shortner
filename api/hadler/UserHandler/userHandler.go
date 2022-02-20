package UserHandler

import (
	Antman "Ant-Man-Url/api/proto/UserProto"
	"Ant-Man-Url/usecase/user"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"strconv"
)

type userServer struct {
	Antman.UnimplementedAntmanUserRoutesServer
	service *user.Service
}

func NewUserServer(service *user.Service) Antman.AntmanUserRoutesServer {
	return &userServer{
		service: service,
	}
}

func (s *userServer) CreateNewUser(ctx context.Context, req *Antman.SignUpRequest) (*Antman.SignUpResponse, error) {
	log.Printf("Starting to Create new User...\n")
	user, err := s.service.SignUpUser(req.Username, req.Password, req.Email, "user")
	if err != nil {
		log.Printf("Error : %s", err)
		return &Antman.SignUpResponse{
			Success: false,
			Message: fmt.Sprintf("%s", err),
		}, nil
	}

	return &Antman.SignUpResponse{
		Success: true,
		Message: fmt.Sprintf("%s signed up.", user.Username),
	}, nil
}

func (s *userServer) LoginUser(ctx context.Context, req *Antman.LoginRequest) (*Antman.LoginResponse, error) {
	log.Printf("Loging in usere...\n")
	user, err := s.service.LoginUser(req.Email, req.Password)
	if err != nil {
		log.Printf("Error : %s\n", err)
		return &Antman.LoginResponse{
			Success: false,
			Message: fmt.Sprintf("Error while Logining %s ", err),
			Jwt:     "",
		}, err
	}

	return &Antman.LoginResponse{
		Success: true,
		Message: fmt.Sprintf("Login sucessful..."),
		Jwt:     user.Token,
	}, nil
}

func (s *userServer) GetUrlStat(ctx context.Context, req *Antman.GetStatRequest) (*Antman.GetStatResponse, error) {
	log.Printf("Checking JWT...\n")
	md, _ := metadata.FromIncomingContext(ctx)
	id := md["id-key"]
	role := md["role-key"]
	userId, err := strconv.Atoi(id[0])
	if err != nil {
		return nil, err
	}

	if role[0] != "user" {
		fmt.Printf("Error : Role is invalid...\n")
		return &Antman.GetStatResponse{
			Success: false,
			Message: fmt.Sprintf("Error : invalid role. "),
		}, err
	}
	log.Printf("Fetching stat...\n")
	urlStat, err := s.service.GetUrlStat(userId, req.ShortUrl)
	if err != nil {
		log.Printf("Error : %s\n", err)
		return &Antman.GetStatResponse{
			Success: false,
			Message: fmt.Sprintf("Error : %s", err),
		}, err
	}
	return &Antman.GetStatResponse{
		Success:   true,
		Message:   fmt.Sprintf("Stat : "),
		LongUrl:   urlStat.Long_link,
		ShortUrl:  req.ShortUrl,
		Redirects: int32(urlStat.Redirects),
		OwnerId:   int32(urlStat.OwnerId),
		CreatedAt: urlStat.Created_at.String(),
	}, nil
}

func (s *userServer) GetUrlStatList(ctx context.Context, req *Antman.GetStatListRequest) (*Antman.GetStatListResponse, error) {
	//headers, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	//Change Here
	//	return nil, nil
	//}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &Antman.GetStatListResponse{
			Success:   false,
			Message:   fmt.Sprintf("No headers are passed"),
			StatlList: nil,
		}, fmt.Errorf("No header")
	}
	id := md["id-key"]
	role := md["role-key"]
	userId, err := strconv.Atoi(id[0])
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nThis is new Id : %s \n This the new Role : %s \n", id[0], role[0])

	if role[0] != "user" {
		return &Antman.GetStatListResponse{
			Success:   false,
			Message:   fmt.Sprintf("Access not permited..."),
			StatlList: nil,
		}, fmt.Errorf("invalid role")
	}
	statList, err := s.service.GetUrlList(userId)

	if err != nil {
		return &Antman.GetStatListResponse{
			Success:   false,
			Message:   fmt.Sprintf("Something went wrong"),
			StatlList: nil,
		}, fmt.Errorf("eroror")
	}
	list := make([]*Antman.GetStatResponse, len(statList))

	for i, l := range statList {
		curr := new(Antman.GetStatResponse)
		curr.Success = true
		curr.Message = "Stat : "
		curr.Redirects = int32(l.Redirects)
		curr.ShortUrl = l.Keyval
		curr.CreatedAt = l.Created_at.String()
		curr.OwnerId = int32(l.OwnerId)
		list[i] = curr
	}

	return &Antman.GetStatListResponse{
		Success:   true,
		StatlList: list,
		Message:   "List : ",
	}, nil

}
