package afas_test

import (
	"os"
	"testing"

	afas "github.com/tim-online/go-afas-profit-rest"
)

func TestKnSalesRelationOrgInsert(t *testing.T) {
	accountNumber := os.Getenv("AFAS_ACCOUNTNUMBER")
	token := os.Getenv("AFAS_TOKEN")
	api := afas.NewAPI(nil, accountNumber, token)
	api.SetDebug(true)

	req := api.KnSalesRelationOrg().NewInsertRequest()
	rb := req.RequestBody()
	rb.NummerDebiteur = "12345test"
	_, err := req.Do()
	if err != nil {
		t.Error(err)
	}
}
