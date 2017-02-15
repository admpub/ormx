package xorm

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestQuery(t *testing.T) {
	db, err := New(`mysql`, `root:root@/blog?charset=utf8;root:root@/blog?charset=utf8`)
	if err != nil {
		t.Fatal(err)
	}
	db.Balancer.TraceOn(`test`, nil)
	views := db.GetOne(`SELECT views FROM webx_post WHERE id=1`)
	fmt.Println(`views:`, views)
	result, err := db.Exec(`UPDATE webx_post SET views=views+1 WHERE id=1`)
	if err != nil {
		t.Fatal(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(`affected:`, affected)
	views = db.Replica().GetOne(`SELECT views FROM webx_post WHERE id=1`)
	fmt.Println(`views:`, views)
}
