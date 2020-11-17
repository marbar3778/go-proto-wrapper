// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/foo.proto

package wrapper_examples

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/marbar3778/go-proto-wrapper"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FooMsg) Unwrap() proto.Message {
	if x := this.GetFoo(); x != nil {
		return x
	}
	if x := this.GetBar(); x != nil {
		return x
	}
	if x := this.GetBaz(); x != nil {
		return x
	}
	return nil
}

func (this *FooMsg) Wrap(value proto.Message) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *Foo:
		this.Sum = &FooMsg_Foo{vt}
		return nil
	case *Bar:
		this.Sum = &FooMsg_Bar{vt}
		return nil
	case *Baz:
		this.Sum = &FooMsg_Baz{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message FooMsg", value)
}
