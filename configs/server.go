package configs

type Config struct {
	Server   Server
	Database Database
	Auth     Auth
	AI       AI
}

type Server struct {
	Host, Port string
}

type Database struct {
	Host, Port, Name, User, Pass, Tz string
}
