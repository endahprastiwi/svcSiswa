package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	pb "github.com/endahprastiwi/svcSiswa/siswa"
	"google.golang.org/grpc"
)

const (
	//port = ":50051"
	defaultName = "Berlalu"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer jadi siswa.ServerSide
func (s *server) ViewSiswa(ctx context.Context, in *pb.SiswaRequest) (*pb.SiswaReply, error) {
	return &pb.SiswaReply{Message: "Skripsi " + in.Name}, nil
}

func ViewSiswa(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // pemanggilan diri sendiri
	fmt.Println(r.Form) // mencetak informasi form di dalam serverside
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// pemanggilan grpc client yg akan ditampilkan di client
	//fmt.Fprintf(w, "SKRIPSI BERLALU!!!!") // mengirim data ke clientside
}

func main() {

	//lis, err := net.Listen("tcp", port)
	http.HandleFunc("/", ViewSiswa)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSiswaServer(s, &server{})

	// Register reflection service on gRPC server.
	/*reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}*/
}
