package main

func ListProducts(options Options) {
	products, err := client.GetProducts()

	if err != nil {
		println("Error:", err.Error())
		return
	}

	for _, p := range products {
		println(p.ID)
	}
}