package util

import (
	"errors"
	"reflect"
)

// 第一页是 page=1
type Pageable interface {
	GetPage() int64
	GetLimit() int64
	GetOffset() int64
}

type GormPage struct {
	page,
	limit int64
}

func NewGormPage(page int64, limit int64) *GormPage {
	return &GormPage{page: page, limit: limit}
}

func (p *GormPage) GetOffset() int64 {
	// gorm offset == -1 的时候是取消 offset 限制
	if p.limit <= 0 || p.page <= 0 {
		return -1
	}
	return (p.page - 1) * p.limit
}

func (p *GormPage) GetPage() int64 {
	return p.page
}

func (p *GormPage) GetLimit() int64 {
	// gorm limit == -1 的时候是取消 limit 限制
	if p.limit <= 0 || p.page <= 0 {
		return -1
	}
	return p.limit
}

func PageSlice(slice interface{}, pageable Pageable) (interface{}, error) {
	pageSize := pageable.GetLimit()
	pageNo := pageable.GetPage()

	// 不分页
	if pageSize <= 0 || pageNo <= 0 {
		return slice, nil
	}

	// 判断是不是 slice 类型
	sliceT := reflect.TypeOf(slice)
	sliceV := reflect.ValueOf(slice)

	// 数据类型判断
	if sliceT.Kind() != reflect.Slice {
		return nil, errors.New("kind of slice is not slice")
	}

	//反射创建
	dSliceV := reflect.MakeSlice(sliceT, 0, 50)
	// 遍历截取数据
	start := int((pageNo - 1) * pageSize)
	end := start + int(pageSize)

	for i := start; i < sliceV.Len(); i++ {
		if i >= end {
			break
		}
		// 添加到新的slice了
		dSliceV = reflect.Append(dSliceV, sliceV.Index(i))
	}

	return dSliceV.Interface(), nil
}
