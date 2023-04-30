package productscontroller

import (
	"jwtMux/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id": 1,
			"nama_product": "kemeja",
			"stok": 854,
		},
		{
			"id": 2,
			"nama_product": "celana",
			"stok": 353,
		},
		{
			"id": 3,
			"nama_product": "topi",
			"stok": 542,
		},
	}

	helper.ResponJSON(w, http.StatusOK, data)
}