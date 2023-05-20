package jobs

import (
	"airhass/global"
	"strings"
)

func Process(csvStr string) {
	// Separate by values
	values := strings.Split(csvStr, ",")

	global.Logger.Debugf("> 0.3 微米颗粒数:\t%s\t个/0.1L", values[0])
	global.Logger.Debugf("PM 2.5 浓度值:\t\t%s\tμg/m³", values[1])
	global.Logger.Debugf("甲醛浓度值:\t\t%s\tμg/m³", values[2])
	global.Logger.Debugf("CO₂ 浓度值:\t\t%s\tppm", values[3])
	global.Logger.Debugf("VOC:\t\t\t%s\tppb", values[6])

}
