package main

type GetTokenRespData struct {
	Code int `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type GetPostDataRequest struct {
	Keyword           string        `json:"keyword"`
	Limit             int           `json:"limit"`
	Offset            int           `json:"offset"`
	JobCategoryIdList []interface{} `json:"job_category_id_list"`
	LocationCodeList  []interface{} `json:"location_code_list"`
	SubjectIdList     []interface{} `json:"subject_id_list"`
	RecruitmentIdList []interface{} `json:"recruitment_id_list"`
	PortalType        int           `json:"portal_type"`
	JobFunctionIdList []interface{} `json:"job_function_id_list"`
	PortalEntrance    int           `json:"portal_entrance"`
}
type GetPostDataResp struct {
	Code int `json:"code"`
	Data struct {
		JobPostList []struct {
			Id          string      `json:"id"`
			Title       string      `json:"title"`
			SubTitle    interface{} `json:"sub_title"`
			Description string      `json:"description"`
			Requirement string      `json:"requirement"`
			JobCategory struct {
				Id       string `json:"id"`
				Name     string `json:"name"`
				EnName   string `json:"en_name"`
				I18NName string `json:"i18n_name"`
				Depth    int    `json:"depth"`
				Parent   struct {
					Id       string      `json:"id"`
					Name     string      `json:"name"`
					EnName   string      `json:"en_name"`
					I18NName string      `json:"i18n_name"`
					Depth    int         `json:"depth"`
					Parent   interface{} `json:"parent"`
					Children interface{} `json:"children"`
				} `json:"parent"`
				Children interface{} `json:"children"`
			} `json:"job_category"`
			CityInfo    interface{} `json:"city_info"`
			RecruitType struct {
				Id       string `json:"id"`
				Name     string `json:"name"`
				EnName   string `json:"en_name"`
				I18NName string `json:"i18n_name"`
				Depth    int    `json:"depth"`
				Parent   struct {
					Id           string      `json:"id"`
					Name         string      `json:"name"`
					EnName       string      `json:"en_name"`
					I18NName     string      `json:"i18n_name"`
					Depth        int         `json:"depth"`
					Parent       interface{} `json:"parent"`
					Children     interface{} `json:"children"`
					ActiveStatus int         `json:"active_status"`
				} `json:"parent"`
				Children     interface{} `json:"children"`
				ActiveStatus int         `json:"active_status"`
			} `json:"recruit_type"`
			PublishTime  int64       `json:"publish_time"`
			JobHotFlag   interface{} `json:"job_hot_flag"`
			JobSubject   interface{} `json:"job_subject"`
			Code         string      `json:"code"`
			DepartmentId interface{} `json:"department_id"`
			JobFunction  interface{} `json:"job_function"`
			JobProcessId string      `json:"job_process_id"`
			RecommendId  interface{} `json:"recommend_id"`
			CityList     []struct {
				Code         string      `json:"code"`
				Name         string      `json:"name"`
				EnName       string      `json:"en_name"`
				LocationType interface{} `json:"location_type"`
				I18NName     string      `json:"i18n_name"`
				PyName       interface{} `json:"py_name"`
			} `json:"city_list"`
		} `json:"job_post_list"`
		Count int    `json:"count"`
		Extra string `json:"extra"`
	} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
