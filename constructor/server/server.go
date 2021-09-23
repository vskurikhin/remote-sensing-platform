package server

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/config"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
	"os"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	Dao    *domain.DAO
	PoolRo *pgxpool.Pool
	PoolRw *pgxpool.Pool
	Server *sa.Atreugo
}

// New инициализирует сервер для ответа на сетевые запросы HTTP.
func New(cfg *config.Config) *Server {
	c := sa.Config{
		Addr:             cfg.Server.Host + `:` + cfg.Server.Port,
		Compress:         true,
		Name:             "httpd",
		GracefulShutdown: true,
	}
	poolRo := openDBRo(cfg)
	poolRw := openDBRw(cfg)
	go gracefulClose(poolRo, poolRw)
	versionDB(poolRw)

	return &Server{
		Dao:    domain.New(poolRo, poolRw),
		Server: sa.New(c),
	}
}

func (s *Server) UseBefore(fns sa.Middleware) *sa.Router {
	return s.Server.UseBefore(fns)
}

func (s *Server) StaticCustom() *sa.Path {
	return s.Server.StaticCustom("/", &sa.StaticFS{
		Root:               "web/public",
		GenerateIndexPages: true,
		AcceptByteRange:    true,
		PathRewrite: func(ctx *sa.RequestCtx) []byte {
			return ctx.Path()
		},
		PathNotFound: func(ctx *sa.RequestCtx) error {
			return ctx.TextResponse("File not found", 404)
		},
	})
}

// GET устанавливает обработчик для GET запросов
func (s *Server) GET(url string, viewFn sa.View) *sa.Path {
	return s.Server.GET(url, viewFn)
}

// POST устанавливает обработчик для POST запросов
func (s *Server) POST(url string, viewFn sa.View) *sa.Path {
	return s.Server.POST(url, viewFn)
}

// PUT устанавливает обработчик для PUT запросов
func (s *Server) PUT(url string, viewFn sa.View) *sa.Path {
	return s.Server.PUT(url, viewFn)
}

func (s *Server) DELETE(url string, viewFn sa.View) *sa.Path {
	return s.Server.DELETE(url, viewFn)
}

// ListenAndServe запускает сервер для ответа на сетевые запросы HTTP.
func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}

func gracefulClose(poolRo *pgxpool.Pool, poolRw *pgxpool.Pool) {
	// Настраиваем канал для отправки сигнальных уведомлений.
	// Нужно использовать буферизованный канал или есть риск пропустить сигнал
	// если не готовы принять сигнал при отправке.
	c := make(chan os.Signal, 1)

	// Блокировать до получения сигнала.
	s := <-c
	fmt.Println("Got signal:", s)
	poolRw.Close()
	poolRo.Close()
}

func openDBRw(cfg *config.Config) *pgxpool.Pool {
	dbCFG := cfg.DataBase
	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRw, dbCFG.PortRw, dbCFG.DBName)
	con, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if con.MinConns < 1 {
		con.MinConns = 1
	}
	con.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		localAddr := conn.PgConn().Conn().LocalAddr()
		remoteAddr := conn.PgConn().Conn().RemoteAddr()
		logger.Debugf("AfterConnect: ConnInfo() = %v -> %v", localAddr, remoteAddr)
		return nil
	}
	con.AfterRelease = func(conn *pgx.Conn) bool {
		// do something with every new connection
		localAddr := conn.PgConn().Conn().LocalAddr()
		remoteAddr := conn.PgConn().Conn().RemoteAddr()
		logger.Debugf("AfterRelease: ConnInfo() = %v -> %v", localAddr, remoteAddr)
		return true
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), con)
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}

func openDBRo(cfg *config.Config) *pgxpool.Pool {
	dbCFG := cfg.DataBase
	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRo, dbCFG.PortRo, dbCFG.DBName)
	con, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if con.MinConns < 1 {
		con.MinConns = 1
	}
	con.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		localAddr := conn.PgConn().Conn().LocalAddr()
		remoteAddr := conn.PgConn().Conn().RemoteAddr()
		logger.Debugf("AfterConnect: ConnInfo() = %v -> %v", localAddr, remoteAddr)
		return nil
	}
	con.AfterRelease = func(conn *pgx.Conn) bool {
		// do something with every new connection
		localAddr := conn.PgConn().Conn().LocalAddr()
		remoteAddr := conn.PgConn().Conn().RemoteAddr()
		logger.Debugf("AfterRelease: ConnInfo() = %v -> %v", localAddr, remoteAddr)
		return true
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), con)
	if err != nil {
		logger.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}

func versionDB(db *pgxpool.Pool) {
	if logger.DebugEnabled() {
		var version string
		err := db.QueryRow(context.Background(), "SELECT VERSION()").Scan(&version)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		logger.Debug(version)
	}
}
