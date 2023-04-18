package main

import (
	_ "docxdownload/internal/packed"

	_ "docxdownload/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"docxdownload/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
