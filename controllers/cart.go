package controllers

import (
	"database/sql"
	"errors"
	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Shopping interface {
	Add(name string, qty int) string
	Remove(name string)
}

type Cart struct {
	Items map[string]Item
}
type Item struct {
	Id    string
	Name  string
	Price string
	Qty   int
}

func getPrice(name string) int {
	return len(name) * 10
}

func getItem(id string) (EcsGoods, error) {
	var i EcsGoods
	sqlUser := beego.AppConfig.String("mysqluser")
	sqlPass := beego.AppConfig.String("mysqlpass")
	sqlUrl := beego.AppConfig.String("mysqlurls")
	sqlDb := beego.AppConfig.String("mysqldb")
	dataSourceName := sqlUser + ":" + sqlPass + "@tcp(" + sqlUrl + ")/" + sqlDb
	db, err := sql.Open("mysql", dataSourceName+"?charset=utf8")
	if err != nil {
		return i, err
	}
	beedb.OnDebug = true
	orm := beedb.New(db)

	orm.Where("goods_id=?", id).Find(&i)
	if i.GoodsName == "" {
		return i, errors.New("Cannot find item")
	} else {
		return i, nil
	}
}

func (c *Cart) Add(id string, qty int) string {
	beego.Trace("Add:", id, qty)
	item, err := getItem(id)
	if err != nil {
		beego.Error(err.Error())
		return err.Error()
	}
	beego.Trace(item.GoodsName, item.ShopPrice)
	b, ok := c.Items[item.GoodsId]
	if !ok {
		b = Item{
			Id:    item.GoodsId,
			Name:  item.GoodsName,
			Price: item.ShopPrice,
			Qty:   qty,
		}
	} else {
		b.Qty += qty
	}
	c.Items[id] = b
	return "Success add to cart."
}

func (c *Cart) Remove(id string) {
	_, ok := c.Items[id]
	beego.Trace("Remove:", id, " ok:", ok)
	if ok {
		delete(c.Items, id)
	}
}
