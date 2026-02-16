package products

import (
	"context"
)

type Service interface {
	ListProducts(ctx context.Context) ([]Product,error)
	GetProduct(ctx context.Context,id string)(*Product,error)
	CreateProduct(ctx context.Context,product *Product)error
	UpdateProduct(ctx context.Context,id string,product *Product)error
	DeleteProduct(ctx context.Context,id string)error
}

type svc struct {
	//repository
	repo Repository
}

//depandancy injection - we inject the service layer into the handler layer. The handler layer depends on the service layer to get the list of products from the database. By injecting the service layer into the handler layer, we can easily change the implementation of the service layer without affecting the handler layer. It also makes it easier to test the handler layer by mocking the service layer.
//we return interface here because we want to hide the implementation details of the service layer from the handler layer. The handler layer only needs to know about the interface and not the implementation of the service layer. This allows us to change the implementation of the service layer without affecting the handler layer. It also makes it easier to test the handler layer by mocking the service layer.
func NewService(repo Repository) Service {
	return &svc{
		repo:repo,
	}
}

func(s *svc) ListProducts(ctx context.Context) ([]Product,error) {
	//call the repo layer to get the list of the products from the database and return it to the handler layer
	products,err := s.repo.ListProducts(ctx)
	if err != nil {
		return nil,err
	}
	return products,nil
}

func(s * svc) GetProduct(ctx context.Context,id string)(*Product,error){
	product,err := s.repo.GetProduct(ctx,id)
	if err != nil {
		return nil,err
	}
	return product,nil
}

func (s *svc) CreateProduct(ctx context.Context,product *Product) error{
	err :=s.repo.CreateProduct(ctx,product)
	if err != nil {
		return err
	}
	return nil
}

func (s *svc) UpdateProduct(ctx context.Context,id string,product *Product) error {
	err := s.repo.UpdateProduct(ctx,id,product)
	if err != nil {
		return err
	}
	return nil
}

func(s *svc)DeleteProduct(ctx context.Context,id string)error {
	err := s.repo.DeleteProduct(ctx,id)
	if err != nil {
		return err
	}
	return nil
}


//in this Handler -> depands on -> Service interface
//service interface ->implemented by -> svc struct
//svc -> call repo layer to get the lsit of products