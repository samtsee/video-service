package repository

type PostRepository interface {
	Connect(connectionString string) error
	Disconnect() error
}
