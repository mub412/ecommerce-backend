package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}
	err = h.productRepo.Delete(pId)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal server Error")
		return
	}
	util.SendData(w, http.StatusOK, "Successfully deleted product")
}
