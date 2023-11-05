// Code generated by goctl. DO NOT EDIT.
package types

type Server struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ServerType    int64  `json:"server_type"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"start_time"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"max_online"`
	MaxQueue      int64  `json:"max_queue"`
	MaxSign       int64  `json:"max_sign"`
	TemplateValue string `json:"template_value"`
}

type GetServerReq struct {
	Id int64 `json:"id"`
}

type GetServerResp struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ServerType    int64  `json:"server_type"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"start_time"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"max_online"`
	MaxQueue      int64  `json:"max_queue"`
	MaxSign       int64  `json:"max_sign"`
	TemplateValue string `json:"template_value"`
}

type SetServerReq struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ServerType    int64  `json:"server_type"`
	Switch        int64  `json:"switch"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"max_online"`
	MaxQueue      int64  `json:"max_queue"`
	MaxSign       int64  `json:"max_sign"`
	TemplateValue string `json:"template_value"`
}

type SetServerResp struct {
	ReturnCode string `json:"return_code"`
}

type CreateServerReq struct {
	Name          string `json:"name"`
	ServerType    int64  `json:"server_type"`
	Switch        int64  `json:"switch"`
	StartTime     int64  `json:"start_time"`
	Area          int64  `json:"area"`
	Tags          string `json:"tags"`
	MaxOnline     int64  `json:"max_online"`
	MaxQueue      int64  `json:"max_queue"`
	MaxSign       int64  `json:"max_sign"`
	TemplateValue string `json:"template_value"`
}

type CreateServerResp struct {
	ReturnCode string `json:"return_code"`
}
