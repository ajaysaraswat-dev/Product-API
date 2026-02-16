package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) (error)
}

type svc struct {
	//repository
}

//depandancy injection - we inject the service layer into the handler layer. The handler layer depends on the service layer to get the list of products from the database. By injecting the service layer into the handler layer, we can easily change the implementation of the service layer without affecting the handler layer. It also makes it easier to test the handler layer by mocking the service layer.
//we return interface here because we want to hide the implementation details of the service layer from the handler layer. The handler layer only needs to know about the interface and not the implementation of the service layer. This allows us to change the implementation of the service layer without affecting the handler layer. It also makes it easier to test the handler layer by mocking the service layer.
func NewService() Service {
	return &svc{
		//initialize repo here
	}
}

func(s *svc) ListProducts(ctx context.Context) (error) {
	//call the repo layer to get the list of the products from the database and return it to the handler layer
	return nil
}


//in this Handler -> depands on -> Service interface
//service interface ->implemented by -> svc struct
//svc -> call repo layer to get the lsit of products