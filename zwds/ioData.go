// ioData.go
package zwds

import (
	"fmt"
)

// Hero - 交互格式, Hero(生日格式)
// type Hero HeroOne

type Hero HeroTwo
type HeroOne struct {
	GongNian int    `json:"gongNian"`
	GongYue  int    `json:"gongYue"`
	GongRi   int    `json:"gongRi"`
	GongShi  int    `json:"gongShi"`
	NongNian int    `json:"nongNian"`
	RunYue   int    `json:"runYue"`
	NongYue  int    `json:"nongYue"`
	NongRi   int    `json:"nongRi"`
	NongShi  int    `json:"nongShi"`
	NianGan  string `json:"nianGan"`
	NianZhi  string `json:"nianZhi"`
	YueGan   string `json:"yueGan"`
	YueZhi   string `json:"yueZhi"`
	RiGan    string `json:"riGan"`
	RiZhi    string `json:"riZhi"`
	ShiGan   string `json:"shiGan"`
	ShiZhi   string `json:"shiZhi"`
	XingMing string `json:"xingMing"`
	XingBie  string `json:"xingBie"`
}
type HeroTwo struct {
	GongNian int    `json:"公年"`
	GongYue  int    `json:"公月"`
	GongRi   int    `json:"公日"`
	GongShi  int    `json:"公时"`
	NongNian int    `json:"农年"`
	RunYue   int    `json:"闰月"`
	NongYue  int    `json:"农月"`
	NongRi   int    `json:"农日"`
	NongShi  int    `json:"农时"`
	NianGan  string `json:"年干"`
	NianZhi  string `json:"年支"`
	YueGan   string `json:"月干"`
	YueZhi   string `json:"月支"`
	RiGan    string `json:"日干"`
	RiZhi    string `json:"日支"`
	ShiGan   string `json:"时干"`
	ShiZhi   string `json:"时支"`
	XingMing string `json:"姓名"`
	XingBie  string `json:"性别"`
}

// Info - 对应Hero, 使用 aDinfoFrom()转换, 构成 HeroInfo, LiuNianInfo;
type Info struct {
	string
	公年, 公月, 公日, 公时, 农年, 闰月, 农月, 农日, 农时     int
	年干, 年支, 月干, 月支, 日干, 日支, 时干, 时支, 姓名, 性别 string
}

// infoFrom - 将Hero转换到Info
func aDinfoFrom(hero Hero) (info Info) {
	info.公年 = hero.GongNian
	info.公月 = hero.GongYue
	info.公日 = hero.GongRi
	info.公时 = hero.GongShi
	info.农年 = hero.NongNian
	info.闰月 = hero.RunYue
	info.农月 = hero.NongYue
	info.农日 = hero.NongRi
	info.农时 = hero.NongShi
	info.年干 = hero.NianGan
	info.年支 = hero.NianZhi
	info.月干 = hero.YueGan
	info.月支 = hero.YueZhi
	info.日干 = hero.RiGan
	info.日支 = hero.RiZhi
	info.时干 = hero.ShiGan
	info.时支 = hero.ShiZhi
	info.姓名 = hero.XingMing
	info.性别 = hero.XingBie
	return info
}

// HeroInfo - zwX.go 数据处理格式 - ioData.go
type HeroInfo struct {
	命主, 身主, 五行 string
	Info
}

func (hi *HeroInfo) bDcopyInfo(info Info) {
	hi.Info = info
	switch info.年干 {
	case "甲", "丙", "戊", "庚", "壬":
		hi.性别 = "阳" + info.性别
	case "乙", "丁", "己", "辛", "癸":
		hi.性别 = "阴" + info.性别
	default:
		fmt.Println("!!! bDcopyInfo() no NianGan.")
	}
	return
}

//////////////////////////////////////////////////////////////////////////
// 以上为 zwX.go 所需数据
// 以下为 zwLx.go 所需数据
//////////////////////////////////////////////////////////////////////////

// LiuHua - 交互格式
type LiuHua struct {
	LyHgSlice []LH   `json:"liuHua"`
	SiHua     string `json:"siHua"`
}

// LH - 取值范围 - LiuYun 流运名 [6]string, HuaGong 化宫简名 [13]string
type LH struct {
	LiuYun  string `json:"liuYun"`
	HuaGong string `json:"huaGong"`
}

// LiuNianInfo - zwLx.go 数据格式 - ioData.go
type LiuNianInfo struct {
	Info
	LiuHua
	四化 string
}

func (lni *LiuNianInfo) bDcopyInfo(info Info) {
	lni.Info = info
	return
}
func (lni *LiuNianInfo) bDcopyLiuHua(liuhua LiuHua) {
	lni.LiuHua = liuhua
	lni.四化 = liuhua.SiHua //
	return
}

//////////////////////////////////////////////////////////////////////////////
// copyLiuHuaNian - 数据转换到LiuNianInfo zwLx.go 处理格式
// !!! 注意前端的LiuHua.LyHgSlice 长度决定后端的长度.
func (lni *LiuNianInfo) ZdCopyLiuHuaNian(liuhua LiuHua, hero Hero) {
	lni.bDcopyInfo(aDinfoFrom(hero))
	lni.bDcopyLiuHua(liuhua)
	return
}
func (lni *LiuNianInfo) ZdChangeLH(lh LH, idLH int) {
	if idLH < len(lni.LyHgSlice) && idLH > 0 {
		lni.LyHgSlice[idLH] = lh
	} else {
		fmt.Println("!!! idLH")
	}
	return
}

// copyHero - 将数值转换到HeroInfo zwX.go格式
func (hi *HeroInfo) ZdCopyHero(hero Hero) {
	hi.bDcopyInfo(aDinfoFrom(hero))
	return
}
