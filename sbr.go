package sbr

import "reflect"

type Sbr struct {
}

// 遍历struct并且自动进行赋值
func (s *Sbr) StructByReflect(data map[string]any, inStructPtr any) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if v, ok := data[t.Name]; ok {
			f.Set(reflect.ValueOf(v))
		} else {
			panic(t.Name + " not found")
		}
	}
}
