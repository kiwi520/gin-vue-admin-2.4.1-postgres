package response

import (
	"gin-vue-admin/model/postgres"
)

type ExaCustomerResponse struct {
	Customer postgres.ExaCustomer `json:"customer"`
}
