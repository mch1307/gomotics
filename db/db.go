package db

//import memdb "github.com/hashicorp/go-memdb"
import "fmt"

 var (
	db     *memdb.MemDB
	Schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"nhcLocation": &memdb.TableSchema{
				Name: "nhcLocation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "id"},
					},
				},
			},
/*			"nhcEquipment": &memdb.TableSchema{
				Name: "nhcEquipment",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "name"},
					},
				},
			},
*/
		},
	}
)

// GenericItem holds definition of item equipment
type GenericItem struct {
	id        int
	provider  string
	name      string
	location  string
	value     int
	itemType  string
	switchCmd string
}

// NhcLocation holds NHC location
type NhcLocation map[string]string

// NhcEquipment holds NHC equipment
type NhcEquipment struct {
	id       int
	name     string
	location string
}

func NewStore() error {
	db, _ := buntdb.Open(":memory:")
	db.CreateIndex("age", "*", buntdb.IndexJSON("age"))
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("1", `{"name":{"first":"Tom","last":"Johnson"},"age":38}`, nil)
	}	
}

/* // NewStore instantiate new memdb
func NewStore() error {
	var err error
	db, err = memdb.NewMemDB(Schema)
	if err != nil {
		panic(err)
	}
	return err
} */

/* // SaveToStore save to store
func SaveToStore(rec NhcLocation) error {
	fmt.Println("Starting savetostore")
	txn := db.Txn(true)
	fmt.Println(rec)
	err := txn.Insert("nhcLocation", rec)
	fmt.Println("Starting savetostore")
	txn.Commit()
	if err != nil {
		fmt.Println("Error inserting ", err)
	}
	txn = db.Txn(false)
	record, err1 := txn.First("nhcLocation", "name", "terrasse")
	if err1 != nil {
		panic(err)
	}
	fmt.Println("result: ", record)
	return err
} */

/* func init() {
	db, _ := buntdb.Open(":memory:")
	db.CreateIndex("age", "*", buntdb.IndexJSON("age"))
	db.Update(func(tx *buntdb.Tx) error {
		tx.Set("1", `{"name":{"first":"Tom","last":"Johnson"},"age":38}`, nil)
	}
} */
