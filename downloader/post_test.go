package downloader

import (
	"encoding/json"
	"testing"
)

func Test_getCsrfToken(t *testing.T) {
	type args struct {
		url   string
		param map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case1",
			args: args{
				url:   GetCsrfTokenUrl,
				param: map[string]int{"portal_entrance": 1},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "failed case1",
			args: args{
				url:   GetCsrfTokenErrUrl,
				param: map[string]int{"portal_entrance": 1},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCsrfToken(tt.args.url, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCsrfToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				return
			}
			if got.Code != tt.want {
				t.Errorf("getCsrfToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_post(t *testing.T) {
	type args struct {
		url   string
		token string
		body  GetPostDataRequest
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success case1",
			args: args{
				url:   GetPostUrl,
				token: "",
				body: GetPostDataRequest{
					Limit:             1,
					Offset:            0,
					PortalType:        6,
					JobFunctionIdList: nil,
					PortalEntrance:    1,
				},
			},
		},
	}
	token, err := getCsrfToken(GetCsrfTokenUrl, map[string]int{})
	if err != nil {
		t.Errorf("getCsrfToken() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := post(tt.args.url, token.Data.Token, tt.args.body)
			if err != nil {
				t.Errorf("post() error = %v", err)
			}
			var postResp GetPostDataResp
			err = json.Unmarshal([]byte(got), &postResp)
			if err != nil {
				t.Errorf("post() error = %v", err)
			}
			if postResp.Code != 0 {
				t.Errorf("post() got = %v", got)
			}
		})
	}
}

func TestMockGetPost(t *testing.T) {
	type args struct {
		fn func(url, token string, body GetPostDataRequest) (string, error)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "mock test1",
			args: args{
				func(url, token string, body GetPostDataRequest) (string, error) {
					return `{
  "code": 0,
  "data": {
    "job_post_list": [
      {
        "id": "7137927772546713869",
        "title": "资深存储研发工程师",
        "sub_title": null,
        "description": "1. 承担块存储网关组骨干员工角色，承担块网关及部分底层功能模块的开发和维护。包括但不限于：IO路径，存储引擎，数据／元数据管理，数据恢复， 网络通信，各种高级功能等。\n2. 负责块存储产品相关功能的设计，编码，以及单元测试等，并协助测试人员完成集成／功能测试，解决测试和生产环境中碰到的各种问题。",
        "requirement": "1. 计算机或相关专业本科及以上学历，具有8年以上存储研发经验，3年以上项目管理经验； \n2. 精通Linux/Unix平台上的C/C++开发，熟悉常用算法和数据结构，熟悉网络编程、多线程编程技术；\n3. 精通存储虚拟化技术，有精简配置、缓存技术、快照技术、复制技术、RAID技术或相关实际开发经验；\n4. 熟悉分布式系统架构，具有大型分布式系统研发，调度，调优经验；\n5. 处事积极主动，具有良好的团队合作精神，自我学习能力强",
        "job_category": {
          "id": "6791702736615426317",
          "name": "研发",
          "en_name": "R&D",
          "i18n_name": "研发",
          "depth": 2,
          "parent": {
            "id": "6791698585114724616",
            "name": "互联网 / 电子 / 网游",
            "en_name": "Internet / Electronics / Games",
            "i18n_name": "互联网 / 电子 / 网游",
            "depth": 1,
            "parent": null,
            "children": null
          },
          "children": null
        },
        "city_info": null,
        "recruit_type": {
          "id": "101",
          "name": "全职",
          "en_name": "Full-time",
          "i18n_name": "全职",
          "depth": 2,
          "parent": {
            "id": "1",
            "name": "社招",
            "en_name": "Experienced",
            "i18n_name": "社招",
            "depth": 1,
            "parent": null,
            "children": null,
            "active_status": 1
          },
          "children": null,
          "active_status": 1
        },
        "publish_time": 1661928496882,
        "job_hot_flag": null,
        "job_subject": null,
        "code": "M4832",
        "department_id": null,
        "job_function": null,
        "job_process_id": "7000320674502068513",
        "recommend_id": null,
        "city_list": [
          {
            "code": "CT_128",
            "name": "深圳",
            "en_name": "Shenzhen",
            "location_type": null,
            "i18n_name": "深圳",
            "py_name": null
          }
        ]
      }
    ],
    "count": 103,
    "extra": "{\"fe_tracking\":{\"log_id\":\"20220903223851010151205084254E8898\",\"query_length\":0,\"total\":103}}"
  },
  "message": "ok",
  "error": null
}`, nil
				},
			},
			want: `{
  "code": 0,
  "data": {
    "job_post_list": [
      {
        "id": "7137927772546713869",
        "title": "资深存储研发工程师",
        "sub_title": null,
        "description": "1. 承担块存储网关组骨干员工角色，承担块网关及部分底层功能模块的开发和维护。包括但不限于：IO路径，存储引擎，数据／元数据管理，数据恢复， 网络通信，各种高级功能等。\n2. 负责块存储产品相关功能的设计，编码，以及单元测试等，并协助测试人员完成集成／功能测试，解决测试和生产环境中碰到的各种问题。",
        "requirement": "1. 计算机或相关专业本科及以上学历，具有8年以上存储研发经验，3年以上项目管理经验； \n2. 精通Linux/Unix平台上的C/C++开发，熟悉常用算法和数据结构，熟悉网络编程、多线程编程技术；\n3. 精通存储虚拟化技术，有精简配置、缓存技术、快照技术、复制技术、RAID技术或相关实际开发经验；\n4. 熟悉分布式系统架构，具有大型分布式系统研发，调度，调优经验；\n5. 处事积极主动，具有良好的团队合作精神，自我学习能力强",
        "job_category": {
          "id": "6791702736615426317",
          "name": "研发",
          "en_name": "R&D",
          "i18n_name": "研发",
          "depth": 2,
          "parent": {
            "id": "6791698585114724616",
            "name": "互联网 / 电子 / 网游",
            "en_name": "Internet / Electronics / Games",
            "i18n_name": "互联网 / 电子 / 网游",
            "depth": 1,
            "parent": null,
            "children": null
          },
          "children": null
        },
        "city_info": null,
        "recruit_type": {
          "id": "101",
          "name": "全职",
          "en_name": "Full-time",
          "i18n_name": "全职",
          "depth": 2,
          "parent": {
            "id": "1",
            "name": "社招",
            "en_name": "Experienced",
            "i18n_name": "社招",
            "depth": 1,
            "parent": null,
            "children": null,
            "active_status": 1
          },
          "children": null,
          "active_status": 1
        },
        "publish_time": 1661928496882,
        "job_hot_flag": null,
        "job_subject": null,
        "code": "M4832",
        "department_id": null,
        "job_function": null,
        "job_process_id": "7000320674502068513",
        "recommend_id": null,
        "city_list": [
          {
            "code": "CT_128",
            "name": "深圳",
            "en_name": "Shenzhen",
            "location_type": null,
            "i18n_name": "深圳",
            "py_name": null
          }
        ]
      }
    ],
    "count": 103,
    "extra": "{\"fe_tracking\":{\"log_id\":\"20220903223851010151205084254E8898\",\"query_length\":0,\"total\":103}}"
  },
  "message": "ok",
  "error": null
}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MockGetPost(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockGetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MockGetPost() got = %v, want %v", got, tt.want)
			}
		})
	}
}
