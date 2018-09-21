package dot

import (
	"reflect"
)

//TypeId dot type guid
type TypeId string

//LiveId dot live guid
type LiveId string

//String convert typeid to string
func (c *TypeId) String() string {
	return string(*c)
}

//String convert liveid to string
func (c *LiveId) String() string {
	return string(*c)
}

//Metadata dot metadata
type Metadata struct {
	TypeId      TypeId
	Version     string
	Name        string
	ShowName    string
	Single      bool
	RelyTypeIds []TypeId
	NewDoter    Newer        `json:"-"`
	RefType     reflect.Type `json:"-"`
}

//Live live/instance
type Live struct {
	TypeId    TypeId
	LiveId    LiveId
	RelyLives []LiveId
	Dot       Dot
}

//NewMetadata new Metadata
func NewMetadata() *Metadata {
	return &Metadata{}
}

//Clone clone Metadata
func (m *Metadata) Clone() *Metadata {
	c := *m
	c.RelyTypeIds = make([]TypeId, len(m.RelyTypeIds))
	copy(c.RelyTypeIds, m.RelyTypeIds)
	return &c
}

//NewDot new a dot
func (m *Metadata) NewDot(args interface{}) (dot Dot, err error) {
	dot = nil
	err = nil
	if m.NewDoter != nil {
		dot, err = m.NewDoter(args)
	} else if m.RefType != nil {
		v := reflect.New(m.RefType)
		dot = v.Interface()
	}
	return
}

//Newer instace for new dot
type Newer = func(args interface{}) (dot Dot, err error)

//Dot componet
type Dot interface{}

//Lifer life cycle
// Create, Start,Stop,Destroy
// Create and Start are separate, in order to resolve the dependencies between different dot instances,
// if there is no problem with the dependencies, then you can directly null in Start
type Lifer interface {
	//Create 在这个方法在进行初始，也运行或监听相同内容，最好放在Start方法中实现
	Create(conf SConfig) error
	//Start
	//ignore 在调用其它Lifer时，true 出错出后继续，false 出现一个错误直接返回
	Start(ignore bool) error
	//Stop
	//ignore 在调用其它Lifer时，true 出错出后继续，false 出现一个错误直接返回
	Stop(ignore bool) error
	//Destroy 销毁 Dot
	//ignore 在调用其它Lifer时，true 出错出后继续，false 出现一个错误直接返回
	Destroy(ignore bool) error
}

//Tager dot自己的标签数据，dot自己使用
type Tager interface {
	//SetTag set tag
	SetTag(tag interface{})
	//GetTag get tag
	GetTag() (tag interface{})
}

//StatusType status type
type StatusType int

//Statuser Status
type Statuser interface {
	Status() StatusType
}

//HotConfiger hot change config
type HotConfiger interface {
	//Update 更新配置信息， 返回true表示成功
	HotConfig(newConf SConfig) bool
}

//Checker 检测dot，运行一些验证或测试数据，返回对应的结果
type Checker interface {
	Check(args interface{}) interface{}
}

const (
	//TagDot tag dot
	TagDot = "dot"
)
