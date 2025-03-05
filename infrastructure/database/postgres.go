package database

type Postgres struct {
	connString string
}

func NewPostgres(connString string) *Postgres {
	return &Postgres{connString: connString}
}

func (pg *Postgres) Migrate() error {
	return nil
}
