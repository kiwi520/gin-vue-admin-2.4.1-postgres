package postgres

type InitDBFunc interface {
	Init() (err error)
}
