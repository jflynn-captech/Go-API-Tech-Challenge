package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"jf.go.techchallenge/internal/applog"
	"jf.go.techchallenge/internal/config"
	"jf.go.techchallenge/internal/handler"
	"jf.go.techchallenge/internal/services"
	"jf.go.techchallenge/protodata"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			func() *log.Logger { return log.New(os.Stdout, "\r\n", log.LstdFlags) },
			applog.New,

			func(config *config.Configuration) *grpc.ClientConn {
				// todo look into WithInsecure alternative?
				conn, err := grpc.NewClient(config.RpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))

				if err != nil {
					log.Fatalf("Failed to connect to grpc service", err)
				}
				return conn
			},

			func(conn *grpc.ClientConn) protodata.PersonRepositoryClient {
				return protodata.NewPersonRepositoryClient(conn)
			},

			func(conn *grpc.ClientConn) protodata.CourseRepositoryClient {
				return protodata.NewCourseRepositoryClient(conn)
			},
			services.NewPerson,
			services.NewCourse,

			TagRoute(handler.GetOnePerson),
			TagRoute(handler.GetAllPersons),
			TagRoute(handler.CreatePerson),
			TagRoute(handler.UpdateOnePerson),
			TagRoute(handler.DeleteOnePerson),

			TagRoute(handler.GetOneCourse),
			TagRoute(handler.GetAllCourses),
			TagRoute(handler.CreateCourse),
			TagRoute(handler.UpdateOneCourse),
			TagRoute(handler.DeleteOneCourse),

			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Invoke(func(*http.Server) {}),
		fx.Invoke(func(appLog *applog.AppLogger) { appLog.PrintBanner() }),
	).Run()
}

func TagRoute(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(`group:"routes"`),
	)
}

func NewServeMux(routes []handler.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.Pattern(), route.Handler())
	}
	return mux
}

// Start Http Server.
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", 8000), Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
