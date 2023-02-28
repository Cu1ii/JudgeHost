package judge

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestJudgeService(t *testing.T) {

	//go func() {
	//	grpcServer := grpc.NewServer()
	//	RegisterJudgeServiceServer(grpcServer, new(JudgeServiceImpl))
	//	lis, err := net.Listen("tcp", ":8000")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	_ = grpcServer.Serve(lis)
	//}()

	logrus.Info("the judge server begin to run")

	conn, err := grpc.Dial("47.100.227.175:8003", grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()
	client := NewJudgeServiceClient(conn)
	reply, err := client.Judge(context.Background(), &JudgeRequest{ProblemId: 1,
		SubmissionId:    1,
		SubmissionCode:  "#include <iostream>\n int main() { \n std::cout << \"11hello world\" << std::endl; \n return 0; \n }",
		ResolutionPath:  "",
		TimeLimit:       1000,
		MemoryLimit:     64,
		OutputLimit:     0,
		Language:        "C++",
		JudgePreference: 0,
		Spj:             false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(*reply)
}
