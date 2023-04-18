package down

import (
	"context"
	"docxdownload/internal/service"
	"fmt"

	"github.com/gingfrederik/docx"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDown struct{}

func NewDown() *sDown {
	return &sDown{}
}

func init() {
	service.RegisterDown(NewDown())
}

// Create 趋势周刊生成
func (c *sDown) Create(ctx context.Context, issueNo uint) (filename string, err error) {
	f := docx.NewFile()
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("ff0000")

	nextPara := f.AddParagraph()
	nextPara.AddLink("worldlink", `https://www.worldlink.com.cn`)

	nextPara = f.AddParagraph()
	nextPara.AddText("Documentation").Size(22).Color("red")

	nextPara = f.AddParagraph()
	nextPara.AddText("Vue’s ")
	nextPara.AddText("")
	nextPara.AddLink("official documentation", `https://vuejs.org/`)
	nextPara.AddText("")
	nextPara.AddText(" provides you with all information you need to get started.")

	t := gtime.Now().Format("Ymd")
	filename = fmt.Sprintf("周刊第%d期（%s）.docx", issueNo, t)
	err = f.Save(filename)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	return
}
