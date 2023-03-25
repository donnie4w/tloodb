package tloodb

import (
	"fmt"
	"testing"
	"time"

	"github.com/donnie4w/simplelog/logging"
)

type TestObj struct {
	Id   *int64
	Name *string `idx`
	Age_ *int
}

func init() {
	UseSimpleDB("test.db")
}

func _Test_DB(t *testing.T) {
	var err error
	// name := "dongdong"
	age := 333
	err = Insert(&TestObj{Name: nil, Age_: &age})
	var s string
	// err, s = BuildIndex[TestObj]()
	fmt.Println("————————————————————————————————————————————", err)
	fmt.Println("————————————————————————————————————————————", s)
	// id := int64(3)
	// err = Update(&TestObj{&id, &name, &age})
	// Delete(TestObj{Id: 3})
	// Delete(&TestObj{Id: &id})
	// err = DeleteWithKey("0_testobj_id_2")
	time.Sleep(3 * time.Second)
	ts := Selects[TestObj](0, 10)
	for i, v := range ts {
		logging.Debug(i+1, "----", v)
	}
	logging.Debug("max idx==>", GetIdSeqValue[TestObj]())

	fmt.Println("------------------------------------------------")
	ts = SelectByIdxName[TestObj]("name", "wuxiaodong")
	for i, v := range ts {
		logging.Debug(i+1, "=====", v)
	}
	fmt.Println("------------------------------------------------")
	ts = SelectByIdxNameLimit[TestObj]("age_", []string{"111", "333"}, 0, 2)
	for i, v := range ts {
		logging.Debug(i+1, "=========>", v)
	}
	o := SelectOneByIdxName[TestObj]("Age", "11")
	logging.Debug("o==>", o)
	fmt.Println("")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	IterDB()
}

func _Test_Backup(t *testing.T) {
	SimpleDB().BackupToDisk("bak", nil)
}

func _Test_Load(t *testing.T) {
	SimpleDB().LoadDataFile("bak")
	IterDB()
}

func Benchmark_Alloc(b *testing.B) {
	var i int
	for i = 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
		// Insert(&TestObj{Name_: "wuxiaodong", Age_: i})
		ts := Selects[TestObj](0, 10)
		for i, v := range ts {
			logging.Debug(i+1, "----", v)
		}
		ts = SelectByIdxName[TestObj]("Age_", "3370")
		for i, v := range ts {
			logging.Debug(i+1, "=====", v)
		}
		ts = SelectByIdxNameLimit[TestObj]("age", []string{"215", "216", "333"}, 0, 2)
		for i, v := range ts {
			logging.Debug(i+1, "=========>", v)
		}
	}
	logging.Debug("i===>", i)
}

func IterDB() {
	keys, _ := SimpleDB().GetKeys()
	for i, v := range keys {
		logging.Debug("key", i+1, "==", v)
		value, _ := SimpleDB().GetString([]byte(v))
		logging.Debug(v, "==>", value)
	}
}

func _Test_snap(t *testing.T) {
	SimpleDB().Put([]byte("d"), []byte("3"))
	logging.Debug(SimpleDB().GetKeys())
	er := SimpleDB().BackupToDisk("snap.lb", []byte("d"))
	logging.Debug(er)
	logging.Debug(RecoverBackup("snap.lb"))
	for _, v := range RecoverBackup("snap.lb") {
		logging.Debug(string(v.Key), " == ", string(v.Value))
	}
}

func _Test_tran(t *testing.T) {
	db, _ := NewDB("test.db")
	tx, _ := db.db.OpenTransaction()

	key := []byte("key5")
	value := []byte("value5")
	// key1 := []byte("key4")
	// value1 := []byte("value6")
	tx.Put(key, value, nil)
	logging.Debug("11111111111")
	db.Del([]byte("1"))
	// logging.Debug(db.Put(key1, value1))
	logging.Debug("2222222222")
	bs, err := db.Get(key)
	logging.Error(err)
	logging.Debug(err, string(bs))
	if err == nil {
		logging.Debug(err, string(bs))
	}
	tx.Commit()
	bs, err = db.Get(key)
	logging.Error(err)
	if err == nil {
		logging.Debug(err, string(bs))
	}

}

func TestView(t *testing.T) {
	OpenView(4343)
	select {}
}
