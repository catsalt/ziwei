package zwds

//安主星模块
import (
	"fmt"
)

// HeroInfo - zwX.go 数据处理格式 - ioData.go
// type HeroInfo struct {
// 	命主, 身主, 五行 string
// 	Info
// }
// type Info struct {
// 	string
// 	公年, 公月, 公日, 公时, 农年, 闰月, 农月, 农日, 农时     int
// 	年干, 年支, 月干, 月支, 日干, 日支, 时干, 时支, 姓名, 性别 string
// }
/////////////////////////////////////////////////////////////////////////////
// UI系列数组,第0组存放姓名, 时间等信息, (1-12)组为十二宫位每个宫位的输出内容, 供GUI使用;
// XGINDEX字典,输出内容(主星)对应的宫位,v值存放UISHIZU十二宫位的序列号, key值为输出内容;
type (
	XGINDEX map[string]int //定盘主星
	UIZHUX  [13][]string   //定盘主星
)

func (ui UIZHUX) checkShengRi(hi HeroInfo) (srGeShi bool) {
	check := 0
	gz := hi.年干 + hi.年支
	for _, v := range 干支序 {
		if v == gz {
			check++
		}
	}
	// fmt.Println(check)
	if hi.农月 > 0 && hi.农月 < 13 {
		check++
	}
	// fmt.Println(check)
	if hi.农日 > 0 && hi.农日 < 31 {
		check++
	}
	// fmt.Println(check)
	if hi.农时 > 0 && hi.农时 < 13 {
		check++
	}
	// fmt.Println(check)
	for _, v := range 性别序 {
		if v == hi.性别 {
			check++
		}
	}
	// fmt.Println(check)
	if check == 5 {
		srGeShi = true
	} else {
		srGeShi = false
	}
	return srGeShi
}

// AsetId - correct Index - 返回1-12;
func AsetId(id int) int {
	for id > 12 {
		id -= 12
	}
	for id < 1 {
		id += 12
	}
	return id
}
// AstrId - 返回id, 没有则返回-1;
func AstrId(str string, s []string) int {
	for k, v := range s {
		if str == v {
			return k
		}
	}
	fmt.Printf("!!! no %v in %v.", str, s)
	return -1
}
func indexGan(gan string) (index int) {
	return AstrId(gan, 干序[:])
}
func indexZhi(zhi string) (index int) {
	return AstrId(zhi, 支序[:])
}
func (ui *UIZHUX) aGongGan(hi HeroInfo) {
	var indexGongGan int
	switch hi.年干 {
	case "甲", "己":
		indexGongGan = 0
	case "乙", "庚":
		indexGongGan = 1
	case "丙", "辛":
		indexGongGan = 2
	case "丁", "壬":
		indexGongGan = 3
	case "戊", "癸":
		indexGongGan = 4
		// default:
		// 	fmt.Println("Wrong! in func dnGongGan().")
	}
	j := 0
	for i := 1; i < 13; i++ {
		ui[i] = append(ui[i], 宫干[indexGongGan][j], 宫支[j])
		j++
	}
	return
}
func (xg XGINDEX) bGongMing(hi HeroInfo) {
	//十二宫
	index := 1 + hi.农月 - hi.农时
	var i int
	for i = 0; i < 12; i++ {
		xg[宫名[i]] = AsetId(index)
		index++
	}
	//身宫
	index = hi.农月 + hi.农时 - 1
	xg[宫名[i]] = AsetId(index)
	return
}
func (xg XGINDEX) cWuXing(ui UIZHUX) {
	i := xg["命宫"]
	gongGan := ui[i][0]
	gongZhi := ui[i][1]
	var idGan int
	switch gongGan {
	case "甲", "乙":
		idGan = 0
	case "丙", "丁":
		idGan = 1
	case "戊", "己":
		idGan = 2
	case "庚", "辛":
		idGan = 3
	case "壬", "癸":
		idGan = 4
	}
	var idZhi int
	switch gongZhi {
	case "子", "丑", "午", "未":
		idZhi = 0
	case "寅", "卯", "申", "酉":
		idZhi = 1
	case "辰", "巳", "戌", "亥":
		idZhi = 2
	}
	xg["五行局"] = 五行局表[idGan][idZhi]
	return
}
func (xg XGINDEX) dZiFuX(hi HeroInfo) {
	xg["紫微"] = 紫微表[xg["五行局"]][hi.农日-1]
	for i := 0; i < 13; i++ {
		xg[紫府系[i]] = 紫府系表[xg["紫微"]-1][i]
	}
}
func (xg XGINDEX) eYueX(hi HeroInfo) {
	for i := 0; i < 8; i++ {
		xg[月系[i]] = 月系表[hi.农月-1][i]
	}
}
func (xg XGINDEX) fShiX(hi HeroInfo) {
	for i := 0; i < 6; i++ {
		xg[时系[i]] = 时系表[hi.农时-1][i]
	}
}
func (xg XGINDEX) gRiX(hi HeroInfo) {
	xg["三台"] = AsetId(xg["左辅"] + hi.农日 - 1)
	xg["八座"] = AsetId(xg["右弼"] - hi.农日 + 1)
	xg["恩光"] = AsetId(xg["文昌"] + hi.农日 - 2)
	xg["天贵"] = AsetId(xg["文曲"] + hi.农日 - 2)
}
func (xg XGINDEX) hNianZhiX(hi HeroInfo) {
	indexNianZhi := indexZhi(hi.年支)
	for i := 0; i < 19; i++ {
		xg[年支系[i]] = 年支系表[indexNianZhi][i]
	}
	xg["天才"] = AsetId(xg["命宫"] + indexNianZhi)
	xg["天寿"] = AsetId(xg["身宫"] + indexNianZhi)
	xg["天殇"] = xg["仆役"]
	xg["天使"] = xg["疾厄"]
	xg["龙德"] = AsetId(未 + indexNianZhi)
}
func (xg XGINDEX) iNianGanX(hi HeroInfo) {
	indexNianGan := indexGan(hi.年干)
	for i := 0; i < 10; i++ {
		xg[年干系[i]] = 年干系表[indexNianGan][i]
	}
}
func (xg XGINDEX) jHuoLingX(hi HeroInfo) {
	var idNianZhi int
	switch hi.年支 {
	case "寅", "午", "戌":
		idNianZhi = 0
	case "申", "子", "辰":
		idNianZhi = 1
	case "巳", "酉", "丑":
		idNianZhi = 2
	case "亥", "卯", "未":
		idNianZhi = 3
	}
	xg["火星"] = 火星表[idNianZhi][hi.农时-1]
	xg["铃星"] = 铃星表[idNianZhi][hi.农时-1]

}
func (xg XGINDEX) kChangShengX(hi HeroInfo) {
	var cs int
	switch xg["五行局"] {
	case 水二局, 土五局:
		cs = 申
	case 木三局:
		cs = 亥
	case 金四局:
		cs = 巳
	case 火六局:
		cs = 寅
	}
	for i := 0; i < 12; i++ {
		switch hi.性别 {
		case "阳男", "阴女":
			xg[长生系[i]] = AsetId(cs + i)
		case "阳女", "阴男":
			xg[长生系[i]] = AsetId(cs - i)
		}
	}
}
func (xg XGINDEX) lXunKongX(hi HeroInfo) {
	idNianGan := indexGan(hi.年干)
	idNianZhi := indexZhi(hi.年支)
	xg["旬空"] = 旬空表[idNianGan][idNianZhi]
	switch hi.年干 {
	case "甲", "丙", "戊", "庚", "壬":
		xg["副旬"] = xg["旬空"] + 1
	case "乙", "丁", "己", "辛", "癸":
		xg["副旬"] = xg["旬空"] - 1
	}
}
func (ui *UIZHUX) mDaXian(hi HeroInfo, xg XGINDEX) {
	var uid int
	for k, v := range 大限表[xg["五行局"]] {
		switch hi.性别 {
		case "阳男", "阴女":
			uid = AsetId(xg["命宫"] + k)
		case "阳女", "阴男":
			uid = AsetId(xg["命宫"] - k)
		default:
			fmt.Println(hi.性别, "!!! mDaXian().")
		}
		ui[uid] = append(ui[uid], v)
	}
}

// nZhux - for a~m functions above all.
func (ui *UIZHUX) nZhuX(hi HeroInfo, xg XGINDEX) {
	if !ui.checkShengRi(hi) {
		fmt.Println("检查生日格式失败.")
		return
	}
	ui.aGongGan(hi)
	xg.bGongMing(hi)    //13.
	xg.cWuXing(*ui)     //1.此步骤必须在dnGongGan(),anGongMing()之后; 在安星之前;
	xg.dZiFuX(hi)       //14
	xg.eYueX(hi)        //8
	xg.fShiX(hi)        //6
	xg.gRiX(hi)         //4 此步骤必须在anYueX(),anShiX()之后;
	xg.hNianZhiX(hi)    //24
	xg.iNianGanX(hi)    //10
	xg.jHuoLingX(hi)    //2
	xg.kChangShengX(hi) //12
	xg.lXunKongX(hi)    //2
	ui.mDaXian(hi, xg)
	//安星按照公式总共有96个, 今后还可以增加, 形成了一个字典什么星在什么宫位;
}

//////////////////////////////////////////////////////////////////////////////
//根据HeroInfo的农历信息, 安星入字典XGINDEX;
//根据constOut的表格顺序, 输出字典到UIZHUX(十二宫位)
func (ui *UIZHUX) oUiZhuX(hi HeroInfo, xg XGINDEX) {
	ui.nZhuX(hi, xg)
	for xid, x := range 输出本表 {
		gid, ok := xg[x]
		if !ok {
			fmt.Printf("XGINDEX没有 %s 星->中断.\n", x)
			return
		}
		ui[gid] = append(ui[gid], x)
		if xid > 23 && xid < 54 {
			ui[gid] = append(ui[gid], 庙旺本表[xid-24][gid-1])
		} //!!!修改一定要注意
		if xid >= 54 && xid < 88 {
			ui[gid] = append(ui[gid], "") //为没有输出庙旺的留空;
		}
	}
}

// 完善HeroInfo, 安命主,身主星,!!!!!!必须在 oUiZhuX之后
// func (hi *HeroInfo) pXingHero(ui UIZHUX, xg XGINDEX) {
func (ui UIZHUX) pXingHero(hi *HeroInfo, xg XGINDEX) {
	gongZhi := ui[xg["命宫"]][1]
	hi.命主 = 命主星[indexZhi(gongZhi)]
	hi.身主 = 身主星[indexZhi(hi.年支)]
	hi.五行 = 五行局[xg["五行局"]]
}

//!!!!!!pXingHero()之后
// 将HeroInfo 个人信息放入数组UIZHUX[0];
func (ui *UIZHUX) qUiHeroInfo(hi HeroInfo) {
	ui[0] = append(ui[0], fmt.Sprintf("姓名: %v  性别: %v", hi.姓名, hi.性别))
	ui[0] = append(ui[0], fmt.Sprintf("命主: %v  身主: %v", hi.命主, hi.身主))
	ui[0] = append(ui[0], fmt.Sprintf("五行局: %v", hi.五行))
	ui[0] = append(ui[0], fmt.Sprintf("八字: %v%v %v%v %v%v %v%v",
		hi.年干, hi.年支, hi.月干, hi.月支, hi.日干, hi.日支, hi.时干, hi.时支))
	ui[0] = append(ui[0], fmt.Sprintf("公历: %v %v %v %v", hi.公年, hi.公月, hi.公日, hi.公时))
	ui[0] = append(ui[0], fmt.Sprintf("农历: %v %v%v %v %v(1-12)",
		hi.农年, hi.闰月, hi.农月, hi.农日, hi.农时))
	return
}

//////////////////////////////////////////////////////////////////////////////
func (uix *UIZHUX) ZxPaiPan(hi *HeroInfo, xg XGINDEX) {
	uix.oUiZhuX(*hi, xg) //顺序执行!
	uix.pXingHero(hi, xg)
	uix.qUiHeroInfo(*hi)
	return
}
