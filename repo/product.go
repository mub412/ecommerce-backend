package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID == productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       100,
		ImgUrl:      "https://i.chaldn.com/_mpimage/komola-orange-imported-50-gm-1-kg?src=https%3A%2F%2Feggyolk.chaldal.com%2Fapi%2FPicture%2FRaw%3FpictureId%3D64292&q=best&v=1&m=1600",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       150,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/LargeFeatureImage/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       50,
		ImgUrl:      "https://www.orchardfood.co.za/cdn/shop/files/Bananas-2.png?v=1753697464&width=1946",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Lichu",
		Description: "Lichu is red",
		Price:       10,
		ImgUrl:      "https://media.gettyimages.com/id/566454679/photo/lychee-fruits.jpg?s=1024x1024&w=gi&k=20&c=mb13OLKoGK9QZYHi3j5nMgBfkuhSMue19wMRCrLWy5g=",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)

}
