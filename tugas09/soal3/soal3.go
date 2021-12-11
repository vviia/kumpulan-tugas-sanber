package soal3

import "strconv"

func LuasPersegi(sisi uint, showText bool) interface{} {
	switch {
	case sisi > 0 && showText:
		return "luas persegi dengan sisi " + strconv.Itoa(int(sisi)) + " cm adalah " + strconv.Itoa(int(sisi*sisi)) + " cm"
	case sisi > 0 && !showText:
		return sisi * sisi
	case sisi == 0 && showText:
		return "Maaf anda belum menginput sisi dari persegi"
	default:
		return nil
	}
}
