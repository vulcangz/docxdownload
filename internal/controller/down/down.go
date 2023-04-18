package down

import (
	"context"
	v1 "docxdownload/api/docx/v1"
	"docxdownload/internal/service"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
)

var Down = cDown{}

type cDown struct{}

func (c *cDown) Docx(ctx context.Context, req *v1.DownReq) (res *v1.DownRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err := r.Parse(&req); err != nil {
		g.Log().Fatal(ctx,err)
	}
	filename, err := service.Down().Create(ctx,req.IssueNo)

	w := g.RequestFromCtx(ctx).Response
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Add("Content-Type", "application/octet-stream") // 默认让浏览器下载文件
	w.Header().Add("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Add("Content-Disposition", "attachment; filename="+gurl.RawEncode(filename))
	w.Header().Add("Content-Transfer-Encoding", "binary")

	// g.Dump(gfile.TempDir(), f)
	w.ServeFile(filename)
	return
}
