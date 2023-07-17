package main

import (
	"context"
	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"

	"golang.org/x/net/http2/h2c"

	"golang.org/x/net/http2"

	"flag"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func ServerRoutine(myPort int) {
	server := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		fmt.Sprintf("localhost:%v", myPort),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func ClientRoutine(partnerPort int) {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://localhost:%v", partnerPort),
	)
	for {
		var name string
		fmt.Scan(&name)
		res, err := client.Greet(
			context.Background(),
			connect.NewRequest(&greetv1.GreetRequest{
				Name: name,
			}),
		)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(res.Msg.Greeting)
	}
}

func main() {
	var (
		me      = flag.Int("MyPort", 8080, "my port number")
		partner = flag.Int("PartnerPort", 8080, "port number of the communication partner")
	)
	flag.Parse()

	go ServerRoutine(*me)
	ClientRoutine(*partner)
}
