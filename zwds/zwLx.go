// zwLX 流星模块
package zwds

import (
	"fmt"
)

// LiuNianInfo - zwLx.go 数据格式 - ioData.go
// type LiuNianInfo struct {
// 	Info
// 	LiuHua
// 	四化 string
// }
/////////////////////////////////////////////////////////////////////////////
// UI系列数组,第0组存放姓名, 时间等信息, (1-12)组为十二宫位每个宫位的输出内容, 供GUI使用;
// XGINDEX字典,输出内容(主星)对应的宫位,v值存放UISHIZU十二宫位的序列号, key值为输出内容;
// UILIUX格式跟UIZHUX一样, 某流运名(如大运...等)的流化星曜的输出宫位表示;
// UILX, 存放,不同流运名(如大运,流年,流月等)流化星曜的输出;
// 四化星UILIUX[宫位][四化]必须对应到主星UIZHUX的具体位置[宫位][主星], 不采用字典;
type (
	// XGINDEX map[string]int //定盘主星,个别流化类星,
	// UIZHUX  [13][]string   //定盘主星
	UILIUX UIZHUX //流化类星
	UILX   []UILIUX
)

// 先算大运数, 再算流运命宫所在宫位; 大运数, 出生年虚岁1, 虚岁五行局数起大运, 所以要减一;
// 流运命宫 = [6]string{"命宫", "运命宫", "年命宫", "月命宫", "日命宫", "时命宫"}
func (lni LiuNianInfo) aLiuMingGong(hi HeroInfo, xg XGINDEX) {
	xg["大运数"] = (lni.农年-(xg["五行局"]+2+hi.农年-1))/10 + 1
	switch hi.性别 {
	case "阳男", "阴女":
		xg[流运命宫[1]] = AsetId(xg[流运命宫[0]] + xg["大运数"] - 1) //daXian(1-12)
	case "阳女", "阴男":
		xg[流运命宫[1]] = AsetId(xg[流运命宫[0]] - xg["大运数"] + 1)
	}
	xg[流运命宫[2]] = 岁前系表[indexZhi(lni.年支)][0] //跟岁建同
	xg[流运命宫[3]] = AsetId(                   //算法见子年斗君表注释;月命宫,日,时; 先算正月然后顺布;
		indexZhi(hi.年支) + 子年斗君表[hi.农时-1][hi.农月-1] + lni.农年 - hi.农年 + lni.农月 - 1)
	xg[流运命宫[4]] = AsetId(xg[流运命宫[3]] + lni.农日 - 1)
	xg[流运命宫[5]] = AsetId(xg[流运命宫[4]] + lni.农时 - 1)
	return
}

// bLiuGanZhi - 流年干支, 以及四化宫干;
func (lni LiuNianInfo) bLiuGanZhi(liuyun, huagong string, hi HeroInfo, ui UIZHUX, xg XGINDEX) (gan, zhi, huaGan string) {
	idLiuYun, idHuaGong := AstrId(liuyun, 流运名[:]), AstrId(huagong, 化宫简名[:])
	switch liuyun {
	case 流运名[0]:
		gan, zhi = hi.年干, hi.年支
	case 流运名[1]: //大运取,宫干支
		gan, zhi = ui[xg[流运命宫[1]]][0], ui[xg[流运命宫[1]]][1]
	case 流运名[2]:
		gan, zhi = lni.年干, lni.年支
	case 流运名[3]:
		gan, zhi = lni.月干, lni.月支
	case 流运名[4]:
		gan, zhi = lni.日干, lni.日支
	case 流运名[5]:
		gan, zhi = lni.时干, lni.时支
	}
	if idHuaGong == 12 {
		huaGan = gan
	} else {
		huaGan = ui[AsetId(xg[流运命宫[idLiuYun]]+idHuaGong)][0]
	}
	return gan, zhi, huaGan
}

//dLiuGongMing - 流年宫名
func (uitmp *UILIUX) dLiuGongMing(huagong, liuYunJianMing string, uid int) {
	for id := 0; id < 12; id++ {
		idx := AsetId(uid + id)
		uitmp[idx] = append(uitmp[idx], huagong, liuYunJianMing, 宫名[id])
	}
}

// cHuaGanX - 安化干系
func (uitmp *UILIUX) cHuaGanX(huaGan string, lni LiuNianInfo, ui UIZHUX, xg XGINDEX) {
	for k, _ := range uitmp {
		uitmp[k] = make([]string, 6) //6 根据GUI制定的.
	}
	var zhuX string
	for k, v := range 化干系 {
		switch lni.四化 {
		case "中州":
			zhuX = 中州化干表[indexGan(huaGan)][k]
		case "全书":
			zhuX = 全书化干表[indexGan(huaGan)][k]
		}
		gid, _ := xg[zhuX]
		xid := AstrId(zhuX, ui[gid][5:]) // 用于安四化.主星zhuX在gid宫位的对应序列号.
		uitmp[gid][xid/2] = v            // 除2 是因为zhuX后面有一庙旺
	}
}

//iLiuGanX - 安流干系
func (uitmp *UILIUX) iLiuGanX(gan string) {
	idGan := indexGan(gan)
	for k, v := range 流干系 {
		uitmp[流干系表[idGan][k]] = append(uitmp[流干系表[idGan][k]], v)
	}
}

//iLiuZhiX - 安流支系
func (uitmp *UILIUX) iLiuZhiX(zhi string) {
	idZhi := indexZhi(zhi)
	for k, v := range 流支系 {
		uitmp[流支系表[idZhi][k]] = append(uitmp[流支系表[idZhi][k]], v)
	}
}

//eSuiQianX - 安岁前
func (uitmp *UILIUX) eSuiQianX(nianZhi string) {
	for k, v := range 岁前系 {
		id := 岁前系表[indexZhi(nianZhi)][k]
		uitmp[id] = append(uitmp[id], v)
	}
}

//eBoShiX - 安博士
func (uitmp *UILIUX) eBoShiX(nianGan string, hi HeroInfo) {
	id := 流干系表[indexGan(nianGan)][0]
	var uid int
	for k, v := range 博士系 {
		switch hi.性别 {
		case "阳男", "阴女":
			uid = AsetId(id + k)
		case "阳女", "阴男":
			uid = AsetId(id - k)
		}
		uitmp[uid] = append(uitmp[uid], v)
	}
}

//eJiangQianX - 安将前
func (uitmp *UILIUX) eJiangQianX(nianZhi string) {
	var id int
	switch nianZhi {
	case "寅", "午", "戌":
		id = 0
	case "申", "子", "辰":
		id = 1
	case "巳", "酉", "丑":
		id = 2
	case "亥", "卯", "未":
		id = 3
	}
	for k, v := range 将前系 {
		uid := 将前系表[id][k]
		uitmp[uid] = append(uitmp[uid], v)
	}
}

//eXiaoXian - 生日的年支定1岁所在宫位, 定小限, 虚岁(长生年为1岁)
func (uitmp *UILIUX) eXiaoXian(hi HeroInfo, lni LiuNianInfo) {
	var idZero, idXiaoXian int
	switch hi.年支 {
	case "寅", "午", "戌":
		idZero = 辰
	case "申", "子", "辰":
		idZero = 戌
	case "巳", "酉", "丑":
		idZero = 未
	case "亥", "卯", "未":
		idZero = 丑
	}
	switch hi.性别 {
	case "阳男", "阴男":
		idXiaoXian = AsetId(idZero + lni.农年 - hi.农年)
	case "阳女", "阴女":
		idXiaoXian = AsetId(idZero - lni.农年 + hi.农年)
	}
	for id := 1; id < 13; id++ {
		if id == idXiaoXian {
			uitmp[id] = append(uitmp[id], "小限")
		} else {
			uitmp[id] = append(uitmp[id], "")
		}
	}
}

//////////
//mUiLiuX - for a~l functions above all
func (uitmp *UILIUX) mUiLiuX(liuyun, huagong string, lni LiuNianInfo, hi HeroInfo, ui UIZHUX, xg XGINDEX) {
	id, idh := AstrId(liuyun, 流运名[:]), AstrId(huagong, 化宫简名[:])
	if id < 0 || idh < 0 {
		return //没有
	}
	lni.aLiuMingGong(hi, xg)
	gan, zhi, huaGan := lni.bLiuGanZhi(liuyun, huagong, hi, ui, xg)
	uitmp.cHuaGanX(huaGan, lni, ui, xg)
	uitmp.dLiuGongMing(huagong, 流运名[id], xg[流运命宫[id]])
	switch liuyun { //或者id ==0, id ==2
	case 流运名[0]: //本命
		return
	case 流运名[2]: //流年
		uitmp.eSuiQianX(zhi)
		uitmp.eBoShiX(gan, hi)
		uitmp.eJiangQianX(zhi)
		uitmp.eXiaoXian(hi, lni)
	}
	uitmp.iLiuGanX(gan)
	uitmp.iLiuZhiX(zhi)
	return
}
func (uitmp *UILIUX) mUiLiuNianInfo(liuyun string, lni LiuNianInfo, xg XGINDEX) {
	uitmp[0][0] = fmt.Sprintf("%v %v四化", lni.姓名, lni.四化)
	if liuyun == 流运名[0] || liuyun == 流运名[1] {
		return
	}
	uitmp[0][1] = fmt.Sprintf("公历: %v", lni.公年)
	uitmp[0][2] = fmt.Sprintf("农历: %v", lni.农年)
	uitmp[0][3] = fmt.Sprintf("八字: %v%v", lni.年干, lni.年支)
	if liuyun == 流运名[2] {
		return
	}
	uitmp[0][1] += fmt.Sprintf(" %v", lni.公月)
	uitmp[0][2] += fmt.Sprintf(" %v %v", lni.闰月, lni.农月)
	uitmp[0][3] += fmt.Sprintf(" %v%v", lni.月干, lni.月支)
	if liuyun == 流运名[3] {
		return
	}
	uitmp[0][1] += fmt.Sprintf(" %v", lni.公日)
	uitmp[0][2] += fmt.Sprintf(" %v", lni.农日)
	uitmp[0][3] += fmt.Sprintf(" %v%v", lni.日干, lni.日支)
	if liuyun == 流运名[4] {
		return
	}
	uitmp[0][1] += fmt.Sprintf(" %v", lni.公时)
	uitmp[0][2] += fmt.Sprintf(" %v", lni.农时)
	uitmp[0][3] += fmt.Sprintf(" %v%v", lni.时干, lni.时支)
	return
}

//////////////////////////////////////////////////////////////////////////////
// func (lni LiuNianInfo) nUiLiX(hi HeroInfo, ui UIZHUX, xg XGINDEX) (uilx UILX) {
// 	for _, lh := range lni.LiuHua.LyHgSlice {
// 		liuyun := lh.LiuYun
// 		huagong := lh.HuaGong
// 		var uit UILIUX
// 		uit.mUiLiuX(liuyun, huagong, lni, hi, ui, xg)
// 		uit.mUiLiuNianInfo(liuyun, lni, xg)
// 		uilx = append(uilx, uit)
// 	}
// 	return
// }
func (lni LiuNianInfo) nUitmp(idLH int, hi HeroInfo, ui UIZHUX, xg XGINDEX) (uitmp UILIUX) {
	lh := lni.LyHgSlice[idLH]
	liuyun := lh.LiuYun
	huagong := lh.HuaGong
	uitmp.mUiLiuX(liuyun, huagong, lni, hi, ui, xg)
	uitmp.mUiLiuNianInfo(liuyun, lni, xg)
	return uitmp
}
func (lni LiuNianInfo) nUiLiX(hi HeroInfo, ui UIZHUX, xg XGINDEX) (uilx UILX) {
	for i := 0; i < len(lni.LyHgSlice); i++ {
		uitmp := lni.nUitmp(i, hi, ui, xg)
		uilx = append(uilx, uitmp)
	}
	return
}

//////////////////////////////////////////////////////////////////////////////
// 根据前端的lni, 排出所有流年选项的流年盘
func (lni LiuNianInfo) ZlxPaiPan(hi HeroInfo, ui UIZHUX, xg XGINDEX) (uilx UILX) {
	uilx = lni.nUiLiX(hi, ui, xg)
	return uilx
}

// 如果某个流年盘的liuhua LH变化, 只改变某个流年盘
func (uilx UILX) ZlxChangeLH(idLH int, lni LiuNianInfo, hi HeroInfo, ui UIZHUX, xg XGINDEX) {
	if idLH < len(uilx) {
		uilx[idLH] = lni.nUitmp(idLH, hi, ui, xg)
	}

	return
}
