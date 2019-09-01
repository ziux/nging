// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// VhostGroup 虚拟主机组
type VhostGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*VhostGroup
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name       	string  	`db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *VhostGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *VhostGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *VhostGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *VhostGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *VhostGroup) Objects() []*VhostGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *VhostGroup) NewObjects() *[]*VhostGroup {
	this.objects = []*VhostGroup{}
	return &this.objects
}

func (this *VhostGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *VhostGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *VhostGroup) Short_() string {
	return "vhost_group"
}

func (this *VhostGroup) Struct_() string {
	return "VhostGroup"
}

func (this *VhostGroup) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *VhostGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *VhostGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *VhostGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *VhostGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *VhostGroup) GroupBy(keyField string, inputRows ...[]*VhostGroup) map[string][]*VhostGroup {
	var rows []*VhostGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*VhostGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*VhostGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *VhostGroup) KeyBy(keyField string, inputRows ...[]*VhostGroup) map[string]*VhostGroup {
	var rows []*VhostGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*VhostGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *VhostGroup) AsKV(keyField string, valueField string, inputRows ...[]*VhostGroup) map[string]interface{} {
	var rows []*VhostGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *VhostGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *VhostGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *VhostGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *VhostGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *VhostGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *VhostGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *VhostGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *VhostGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *VhostGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *VhostGroup) Reset() *VhostGroup {
	this.Id = 0
	this.Uid = 0
	this.Name = ``
	this.Description = ``
	this.Created = 0
	return this
}

func (this *VhostGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["Created"] = this.Created
	return r
}

func (this *VhostGroup) Set(key interface{}, value ...interface{}) factory.Model {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Uid": this.Uid = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Description": this.Description = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
			}
	}
	return r
}

func (this *VhostGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["name"] = this.Name
	r["description"] = this.Description
	r["created"] = this.Created
	return r
}

func (this *VhostGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *VhostGroup) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

