package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

func main() {
	ctx := context.Background()

	g.Log().Infof(ctx, "Auto Add Account...")

	content := gfile.GetContents("acc.txt")
	cline := strings.Split(content, "\n")

	serverContent := gfile.GetContents("addr.txt")
	cServer := strings.Split(serverContent, "\n")

	ctoken := cServer[0]
	cSplit := cServer[1]
	cModel := cServer[2]

	if cSplit == "\\t" {
		cSplit = "\t"
	}

	for _, line := range cline {
		lines := strings.Split(line, cSplit)

		for _, server := range cServer[3:] {
			switch cModel {
			case "rt":
				resp, _ := g.Client().SetHeaderMap(g.MapStrStr{
					"adminapi": ctoken,
				}).Post(ctx, fmt.Sprintf("%s/adminapi/addAccount", server), g.Map{
					"account":       lines[0],
					"password":      lines[1],
					"refresh_token": lines[2],
				})

				g.Log().Infof(ctx, server+"   "+resp.ReadAllString())
			case "rk":
				resp, _ := g.Client().SetHeaderMap(g.MapStrStr{
					"adminapi": ctoken,
				}).Post(ctx, fmt.Sprintf("%s/adminapi/addAccount", server), g.Map{
					"account":         lines[0],
					"password":        lines[1],
					"refresh_cookies": lines[2],
				})

				g.Log().Infof(ctx, server+"   "+resp.ReadAllString())
			case "at":
				resp, _ := g.Client().SetHeaderMap(g.MapStrStr{
					"adminapi": ctoken,
				}).Post(ctx, fmt.Sprintf("%s/adminapi/addAccount", server), g.Map{
					"account":      lines[0],
					"password":     lines[1],
					"access_token": lines[2],
				})
				g.Log().Infof(ctx, server+"   "+resp.ReadAllString())
			}

		}

	}
}
