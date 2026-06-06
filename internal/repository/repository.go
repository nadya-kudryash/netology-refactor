package repository

type RepositoryWriter interface {
	SaveOrder(customer string, products string, total float64, status string) error
}
