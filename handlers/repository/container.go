package billingcyclerepository

import (
	"billing-cycle-ms-app/constant"
	"billing-cycle-ms-app/models"
	"errors"
	"fmt"
	"os"

	"github.com/couchbase/gocb/v2"
	"github.com/labstack/gommon/log"
)

// Container will hold all dependencies for your application.
type billingCycleRepository struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// New returns an eedmpty or an initialized container for your handlers.
func New() models.BillingCycleRepository {
	cluster, err := gocb.Connect(
		"couchbase://couchbase-db-srv",
		gocb.ClusterOptions{
			Username: os.Getenv("COUCHBASE_USERNAME"),
			Password: os.Getenv("COUCHBASE_PASSWORD"),
		})

	if err != nil {
		log.Info("Could not connect to couchbase . Please check if couchbase container is up and running.")
		panic(err)
	}

	bucket := cluster.Bucket("test-bucket")

	c := &billingCycleRepository{
		cluster: cluster,
		bucket:  bucket,
	}

	return c
}

func (c *billingCycleRepository) InsertData(billingData models.BillingCycle) (interface{}, error) {
	collection := c.bucket.DefaultCollection()
	upsertData := billingData
	upsertResult, err := collection.Insert(billingData.Id, upsertData, &gocb.InsertOptions{})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentExists) {
			fmt.Println("Document found")
			return upsertResult, nil
		} else {
			return constant.New("BSS-BILL-CBCI-50000001", err), err
		}
	}
	//fmt.Println(upsertResult.Cas())
	return upsertResult, nil
}

// Example to use Queries
func (c *billingCycleRepository) CheckForBillingId(id string) (interface{}, error) {
	query := "SELECT a.* FROM `test-bucket` a where a.id = $1;"
	rows, err := c.cluster.Query(query, &gocb.QueryOptions{
		Adhoc:                false,
		PositionalParameters: []interface{}{id},
	})
	if err != nil {
		// panic(err)
		log.Errorf("Could not query couchbase - ", err)
		return constant.New("BSS-BILL-CBCI-50000002", err), err

	}
	var billrow interface{}
	for rows.Next() {
		err := rows.Row(&billrow)
		if err != nil {
			// panic(err)
			log.Error("There is file fetching data from couchbase")
			return constant.New("BSS-BILL-CBCI-50000002", err), err
		}
		// fmt.Println(billrow)
	}
	if err := rows.Err(); err != nil {
		// panic(err)
		log.Error("There is some issue while fetching rows object from couchbase .")
		return constant.New("BSS-BILL-CBCI-50000002", err), err
	}
	return billrow, nil
}

func (c *billingCycleRepository) GetBillingCycleById(id string) (interface{}, error) {
	billcycle, err_db := c.bucket.DefaultCollection().Get(id, &gocb.GetOptions{})
	if err_db != nil {
		if errors.Is(err_db, gocb.ErrDocumentNotFound) {
			return constant.New("BSS-BILL-CBCI-50000007", fmt.Errorf("Document with the provided id not found - %s", id)), fmt.Errorf("Document with the provided id not found - %s", id)
		}
		return constant.New("BSS-BILL-CBCI-50000002", err_db), err_db
	}
	var responseData interface{}
	if err := billcycle.Content(&responseData); err != nil {
		panic(err)
	}
	return responseData, nil
}
