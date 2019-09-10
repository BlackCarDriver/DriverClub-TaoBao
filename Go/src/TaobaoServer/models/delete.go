package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//owner delete a goods 🍏
func DeleteMyGoods(uid, gid string) error {
	if uid == "" || gid == "" {
		err := errors.New("Got a null userid or goodsid")
		logs.Error(err)
		return err
	}
	o := orm.NewOrm()
	if res, err := o.Raw("delete from t_goods where id = '1907150011'", uid, gid).Exec(); err != nil {
		logs.Error(err)
		return err
	} else if af, err := res.RowsAffected(); err == nil && af == 0 {
		err = errors.New("No rows affacted")
		logs.Error("%v : userid:%s, goodsid:%s", err, uid, gid)
		return err
	}
	return nil
}
