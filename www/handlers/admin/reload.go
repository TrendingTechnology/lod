package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tile-fund/lod/config"
	"github.com/tile-fund/lod/str"
	"github.com/tile-fund/lod/util"
)

// ReloadCapabilities performs a config reload, picking up any
// changes to the instance capabilities configuration.
func ReloadCapabilities(ctx *fiber.Ctx) error {
	// read config and update config.Cap
	err := config.Read()
	if err != nil {
		util.Error(str.CAdmin, str.EReload, err.Error())
		ctx.Status(500)
		return ctx.JSON(map[string]string{
			"status": "failed",
			"file":   *config.File,
			"error":  err.Error(),
		})
	}

	util.Info(str.CAdmin, str.MReload)
	ctx.Status(200)
	return ctx.JSON(map[string]string{
		"status": "ok",
		"file":   *config.File,
	})
}
