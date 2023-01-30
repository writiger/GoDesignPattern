package facade

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantRes string
		wantErr bool
	}{{
		name:    "姓名查询测试张三",
		args:    "zhangsan@qq.com",
		wantRes: "Zhang San",
		wantErr: false,
	}, {
		name:    "姓名查询测试赵四",
		args:    "zhaosi@qq.com",
		wantErr: true,
	}, {
		name:    "姓名查询测试Bob",
		args:    "bob@qq.com",
		wantErr: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &database{}
			res, err := d.GetProperty(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProperty() error = %v, wanterr = %v", err, tt.wantErr)
				return
			}
			if res != tt.wantRes {
				t.Errorf("GetProperty() res = %v, wantRes = %v", err, tt.wantRes)
				return
			}
		})
	}
}
