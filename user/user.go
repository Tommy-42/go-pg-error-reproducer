package user

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" sql:"-"`
	XXX_unrecognized     []byte   `json:"-" sql:"-"`
	XXX_sizecache        int32    `json:"-" sql:"-"`
}
