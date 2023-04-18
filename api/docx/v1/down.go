package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DownReq struct {
	g.Meta  `path:"/down" tags:"Down" method:"post" summary:"Create and download docx file"`
	IssueNo uint `json:"issueNo" in:"query" d:"100" v:"required#期号不能为空" dc:"期号,默认100"`
}

type DownRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
