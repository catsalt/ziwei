// sjFuWu
package zwds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "time"
	// "os"
	// "path/filepath"
	// "strings"
)

var infoOne = Info{
	公年: 1955, 公月: 3, 公日: 7, 公时: 4, 农年: 1955, 农月: 2, 农日: 14, 农时: 3,
	年干: "乙", 年支: "未", 月干: "己", 月支: "卯", 日干: "丁", 日支: "卯", 时干: "壬", 时支: "寅",
	姓名: "命例", 性别: "女"}
var infoTwo = Info{姓名: "某例", 性别: "女",
	公年: 2019, 公月: 9, 公日: 9, 公时: 19, 农年: 2019, 农月: 8, 农日: 11, 农时: 10,
	年干: "己", 年支: "亥", 月干: "癸", 月支: "酉", 日干: "己", 日支: "酉", 时干: "癸", 时支: "酉"}
var hero = Hero{1955, 3, 7, 4, 1955, 0, 2, 14, 3,
	"乙", "未", "己", "卯", "丁", "卯", "壬", "寅", "命例", "女"}
var liuhuaOne = LiuHua{
	LyHgSlice: []LH{LH{"本命", "原"}, LH{"大运", "原"}, LH{"流年", "原"},
		LH{"流月", "原"}, LH{"流日", "原"}, LH{"流时", "原"}},
	SiHua: "中州",
}

const (
	dir        = "../data"
	pathHeroes = "../data/heroList.txt"
	inFileB    = "../data/inputLiuNian.txt"
)

////////////////////////////////////////////////
// aFreadJson - read Json file to InfoList
type HeroList []Hero

var (
	HL   HeroList
	ID   int
	HI   HeroInfo
	UI   UIZHUX
	XG   XGINDEX //存储主星的所在的UIZHUX的[index]
	LNI  LiuNianInfo
	LXUI UILX
)

func (h *HeroList) aFreadJson(jsonPath string) {
	j, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err, "!!! aFreadJson() 1")
		return
	}
	err = json.Unmarshal(j, h)
	if err != nil {
		fmt.Println(err, "!!! aFreadJson() 2")
		return
	}
	return
}

// aFwriteJson - write InfoList to Json file.
func bFwrite(sb []byte, jsonPath string) {
	err := ioutil.WriteFile(jsonPath, sb, 0777)
	if err != nil {
		fmt.Println(err, "!!! aFwriteJson()")
		return
	}
	return
}

// bFaddHero - add Info to InfoList
func (h *HeroList) cFaddHero(hero Hero) {
	hl := *h
	if len(hl) < 50 {
		hl = append(hl, hero)
		ID = len(hl) - 1
	} else {
		hl[ID] = hero
	}
	j, err := json.MarshalIndent(hl, "", "  ")
	if err != nil {
		fmt.Println(err, "!!! cFaddHero() json")
	}
	bFwrite(j, pathHeroes)
	return
}

// outFile := dir + "/" + hi.姓名 + " 主星盘 " + time.Now().Format("2006-01-02") + ".txt"

// bFreadHero - arrow, "<<" pre, ">>" next, "" this ID,
func (h HeroList) ZfchooseHero(arrow string) Hero {
	if len(h) == 0 {
		// bFaddHero(demo)
	}
	switch arrow {
	case "<<":
		ID--
	case ">>":
		ID++
	default:
	}
	for ID < 0 {
		ID += len(h)
	}
	for ID >= len(h) {
		ID -= len(h)
	}
	return h[ID]
}

func ZfzxPan() {
	var (
		hi HeroInfo
		ui UIZHUX
		xg = make(XGINDEX)
	)
	// HI.bDcopyInfo(infoOne)
	hi.ZdCopyHero(hero)
	ui.ZxPaiPan(&hi, xg)
	for i := 0; i < 13; i++ {
		fmt.Printf("%0.2d: ", i)
		for _, v := range ui[i] {
			fmt.Printf("%s ", v)
		}
		fmt.Println(len(ui[i]))
	}
	HI = hi
	UI = ui
	XG = xg
}
func ZflxPan(liuhua LiuHua, liunian Hero) {
	// LNI.bDcopyInfo(infoTwo)
	// LNI.bDcopyLiuHua(liuhuaOne)
	var (
		lni LiuNianInfo
	)
	lni.ZdCopyLiuHuaNian(liuhua, liunian)
	lxui := lni.ZlxPaiPan(HI, UI, XG)
	LNI = lni
	LXUI = lxui
}
func ZflxChangeLH(lh LH, idLH int) {
	LNI.ZdChangeLH(lh, idLH)
	LXUI.ZlxChangeLH(idLH, LNI, HI, UI, XG)
}

// func (uix UIZHUX) aFwrite(path string) {
// 	var m string
// 	for _, v := range uix {
// 		m += strings.Join(v, ",")
// 		m += "\r\n"
// 	}
// 	ioutil.WriteFile(path, []byte(m), 0777)
// 	return
// }
// cFinit() - if no file, add file, for aFreadJson()
// func cFinit() {
// 	os.MkdirAll(dir, 0777)
// 	fs, _ := ioutil.ReadDir(dir)
// 	noFile := func(f string) bool {
// 		for _, v := range fs {
// 			if filepath.Base(v.Name()) == filepath.Base(f) {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	if noFile(inFileA) || noFile(inFileB) {
// 		bFaddHero(demo)
// 		sb, _ := json.MarshalIndent(InfoList, "", "  ")
// 		aFwriteJson(inFileA, sb)
// 		sb, _ = json.MarshalIndent(demo1, "", "  ")
// 		aFwriteJson(inFileB, sb)
// 	}
// 	return
// }
// func cF() {
// 	cFinit()
// 	aFreadJson(inFileA)
// }

////////////////////////////////////////////////
//都为全局变量: Info->HeroInfo, LiuNian->LiuNianInfo 转换,存储
// var hi HeroInfo
// var InfoList []Info
// var Id int //Id - 全局变量,InfoList 的序号,
// var lni LiuNianInfo
// var ln LiuNian

//ui - 全局变量 - 目前最长35 用来排流化盘;每次使用需初始化;
// var ui UIZHUX
// var uilx UILIUX        //每次使用初始化;90;
// var xg = make(XGINDEX) //存储主星的所在的UIZHUX的[index]
