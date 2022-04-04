package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	pb "server/protos/user"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 50051, "The server port")
	db_url = "postgresql://postgres:password@localhost:5432/grpc_assignment_db"
)

type userAuthServer struct {
	pb.UnimplementedUserAuthServer
}

func (s *userAuthServer) Register(ctx context.Context, in *pb.RegReq) (*pb.ApiRes, error) {
	// establish connection to postgresql database
	conn := initDbConn()
	defer conn.Close()

	// verify registration request
	if err := verifyRegReq(in, conn); err != nil {
		return &pb.ApiRes{ResCode: 400, Message: err.Error()}, nil
	}

	// register user
	sqlStatement := `INSERT INTO "public".user (email, username, password) VALUES ($1, $2, $3)`

	hashed, _ := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	_, err := conn.Exec(sqlStatement, in.GetEmail(), in.GetUsername(), hashed)
	if err != nil {
		return &pb.ApiRes{ResCode: 400, Message: err.Error()}, nil
	}
	log.Printf("Registered: %v", in.GetUsername())
	return &pb.ApiRes{ResCode: 200, Message: fmt.Sprintf("Registration successful. Registered as: %v", in.GetUsername())}, nil
}

func (s *userAuthServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.ApiRes, error) {
	// establish connection to postgresql database
	conn := initDbConn()
	defer conn.Close()

	// verify login request
	if err := verifyLoginReq(in, conn); err != nil {
		return &pb.ApiRes{ResCode: 400, Message: err.Error()}, nil
	}

	// login user
	log.Printf("Logged in: %v", in.GetUsername())
	return &pb.ApiRes{ResCode: 200, Message: fmt.Sprintf("Login successful. Logged in as: %v", in.GetUsername())}, nil
}

func (s *userAuthServer) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.ApiRes, error) {
	return &pb.ApiRes{ResCode: 200, Message: "Successfully logged out"}, nil
}

func verifyRegReq(in *pb.RegReq, conn *sql.DB) error {
	if in.GetEmail() == "" && in.GetUsername() == "" {
		return fmt.Errorf("Email and username cannot be empty")
	}
	if in.GetEmail() == "" {
		return fmt.Errorf("Email cannot be empty")
	}
	if in.GetUsername() == "" {
		return fmt.Errorf("Username cannot be empty")
	}
	if len(in.GetPassword()) < 6 {
		return fmt.Errorf("Password cannot be less than 6 characters")
	}

	// is email already in database? if yes, return error
	sqlStatement := `SELECT email FROM "public".user WHERE email = $1`
	var email string
	err := conn.QueryRow(sqlStatement, in.GetEmail()).Scan(&email)
	switch err {
	case nil:
		return fmt.Errorf("Email already exists")
	case sql.ErrNoRows:
		// email is not in database
	default:
		return fmt.Errorf("Could not query database: %v", err)
	}

	// is username already in database? if yes, return error
	sqlStatement = `SELECT username FROM "public".user WHERE username = $1`
	var username string
	err = conn.QueryRow(sqlStatement, in.GetUsername()).Scan(&username)
	switch err {
	case nil:
		return fmt.Errorf("Username already exists")
	case sql.ErrNoRows:
		// username is not in database
	default:
		return fmt.Errorf("Could not query database: %v", err)
	}

	return nil
}

func verifyLoginReq(in *pb.LoginReq, conn *sql.DB) error {
	if in.GetUsername() == "" {
		return fmt.Errorf("Username cannot be empty")
	}

	sqlStatement := `SELECT password FROM "public".user WHERE username = $1`
	var hashedPass string
	err := conn.QueryRow(sqlStatement, in.GetUsername()).Scan(&hashedPass)
	switch err {
	case nil:
		if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(in.GetPassword())); err != nil {
			return fmt.Errorf("Password is incorrect")
		}
	case sql.ErrNoRows:
		return fmt.Errorf("No registered user found by this username")
	default:
		log.Fatalf("could not scan row: %v", err)
	}

	return nil
}

func initDbConn() *sql.DB {
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("could not open postgresql connection: %v", err)
	}
	log.Printf("Connected to database at port 5432")
	return conn
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserAuthServer(s, &userAuthServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
