/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2013-12-16 21:38
 * description :
 * history :
 */

package app

import (
	"github.com/jrsix/gof"
	"go2o/src/core/service/dps"
)

func Init(app gof.App) {
	dps.Init(app)
	gof.CurrentApp = app
}
