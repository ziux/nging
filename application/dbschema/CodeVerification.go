// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo/param"
	
	"time"
)

// CodeVerification 验证码
type CodeVerification struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CodeVerification
	namer   func(string) string
	connID  int
	
	Id         	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Code       	string  	`db:"code" bson:"code" comment:"验证码" json:"code" xml:"code"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	OwnerId    	uint64  	`db:"owner_id" bson:"owner_id" comment:"所有者ID" json:"owner_id" xml:"owner_id"`
	OwnerType  	string  	`db:"owner_type" bson:"owner_type" comment:"所有者类型" json:"owner_type" xml:"owner_type"`
	Used       	uint    	`db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Purpose    	string  	`db:"purpose" bson:"purpose" comment:"目的" json:"purpose" xml:"purpose"`
	Start      	uint    	`db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End        	uint    	`db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled   	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	SendMethod 	string  	`db:"send_method" bson:"send_method" comment:"发送方式(mobile-手机;email-邮箱)" json:"send_method" xml:"send_method"`
	SendTo     	string  	`db:"send_to" bson:"send_to" comment:"发送目标" json:"send_to" xml:"send_to"`
}

func (this *CodeVerification) Trans() *factory.Transaction {
	return this.trans
}

func (this *CodeVerification) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CodeVerification) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CodeVerification) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CodeVerification) Objects() []*CodeVerification {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CodeVerification) NewObjects() *[]*CodeVerification {
	this.objects = []*CodeVerification{}
	return &this.objects
}

func (this *CodeVerification) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CodeVerification) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CodeVerification) Short_() string {
	return "code_verification"
}

func (this *CodeVerification) Struct_() string {
	return "CodeVerification"
}

func (this *CodeVerification) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *CodeVerification) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CodeVerification) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CodeVerification) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CodeVerification) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeVerification) GroupBy(keyField string, inputRows ...[]*CodeVerification) map[string][]*CodeVerification {
	var rows []*CodeVerification
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*CodeVerification{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CodeVerification{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *CodeVerification) KeyBy(keyField string, inputRows ...[]*CodeVerification) map[string]*CodeVerification {
	var rows []*CodeVerification
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*CodeVerification{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *CodeVerification) AsKV(keyField string, valueField string, inputRows ...[]*CodeVerification) map[string]interface{} {
	var rows []*CodeVerification
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

func (this *CodeVerification) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeVerification) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *CodeVerification) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CodeVerification) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CodeVerification) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CodeVerification) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["owner_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["owner_type"] = "user" } }
	if val, ok := kvset["send_method"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["send_method"] = "mobile" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CodeVerification) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *CodeVerification) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CodeVerification) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CodeVerification) Reset() *CodeVerification {
	this.Id = 0
	this.Code = ``
	this.Created = 0
	this.OwnerId = 0
	this.OwnerType = ``
	this.Used = 0
	this.Purpose = ``
	this.Start = 0
	this.End = 0
	this.Disabled = ``
	this.SendMethod = ``
	this.SendTo = ``
	return this
}

func (this *CodeVerification) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Code"] = this.Code
	r["Created"] = this.Created
	r["OwnerId"] = this.OwnerId
	r["OwnerType"] = this.OwnerType
	r["Used"] = this.Used
	r["Purpose"] = this.Purpose
	r["Start"] = this.Start
	r["End"] = this.End
	r["Disabled"] = this.Disabled
	r["SendMethod"] = this.SendMethod
	r["SendTo"] = this.SendTo
	return r
}

func (this *CodeVerification) Set(key interface{}, value ...interface{}) factory.Model {
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
				case "Id": this.Id = param.AsUint64(vv)
				case "Code": this.Code = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "OwnerId": this.OwnerId = param.AsUint64(vv)
				case "OwnerType": this.OwnerType = param.AsString(vv)
				case "Used": this.Used = param.AsUint(vv)
				case "Purpose": this.Purpose = param.AsString(vv)
				case "Start": this.Start = param.AsUint(vv)
				case "End": this.End = param.AsUint(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "SendMethod": this.SendMethod = param.AsString(vv)
				case "SendTo": this.SendTo = param.AsString(vv)
			}
	}
	return r
}

func (this *CodeVerification) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["code"] = this.Code
	r["created"] = this.Created
	r["owner_id"] = this.OwnerId
	r["owner_type"] = this.OwnerType
	r["used"] = this.Used
	r["purpose"] = this.Purpose
	r["start"] = this.Start
	r["end"] = this.End
	r["disabled"] = this.Disabled
	r["send_method"] = this.SendMethod
	r["send_to"] = this.SendTo
	return r
}

func (this *CodeVerification) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *CodeVerification) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

