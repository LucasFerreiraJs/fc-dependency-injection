package product

type ProductUsecase struct {
	repository ProductRepositoryInterface
}

func NewProductUsecase(r ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repository: r,
	}

}

func (u *ProductUsecase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
