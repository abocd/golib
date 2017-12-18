package glog
import (
	"fmt"
	"time"
	"bufio"
	"os"
	"sync"
	"runtime"
	"strings"
)

const (
	color_white   = uint8(iota + 90)
	color_red
	color_green
	color_yellow
	color_blue
	color_magenta
	color_cyan
)

const (
	//详情
	verbose = "[VERB]"
	//追踪
	trace   = "[TRAC]"
	//错误
	errors  = "[ERRO]"
	//警告
	warn    = "[WARN]"
	//信息
	info    = "[INFO]"
	//调试
	debug   = "[DBUG]"
	//断言
	assert  = "[ASST]"
)

//级别列表，从 一般到重要
var levels = [7]string{debug,trace,verbose,assert,info,warn,errors}
//级别索引
var levelsIndex = map[string]int{}

const (
	LongFile = 1
	ShortFile
)

type Glog struct{
	//锁
	mu sync.Mutex
	//是否设置了日志保存
	SaveLog bool
	//显示 级别
	ShowLevel string
	showLevelIndex int
	showLevelBool bool
	//保存 级别
	SaveLevel string
	saveLevelIndex int
	saveLevelBool bool
	MaxLogSize int
	// 文件级别
	Flag int
    // 日志保存方式
	out *bufio.Writer
}

var gg *Glog
var levelMu sync.Mutex


func init(){
	gg = NewGlog(&Glog{ShowLevel:debug,SaveLevel:""})
	levelsIndex = make(map[string]int,len(levels))
	for k,v := range levels{
		levelsIndex[v] = k
	}
}

/**
 初始化
 */
func NewGlog(g *Glog)*Glog {
	//if g.ShowLevel != ""{
	//	index := getLevelIndex(g.ShowLevel)
	//	if index >= 0{
	//		g.showLevelIndex = index
	//	} else {
	//		g.showLevelIndex = 999
	//	}
	//} else {
	//	g.showLevelIndex = 999
	//}
	//if g.SaveLevel != ""{
	//	index := getLevelIndex(g.SaveLevel)
	//	if index >= 0{
	//		g.saveLevelIndex = index
	//	} else {
	//		g.saveLevelIndex = 999
	//	}
	//} else {
	//	g.saveLevelIndex = 999
	//}
	return g
}



/**
 使用文件
 */
func NewGLogFile(filename string,g *Glog)*Glog {
	g.mu.Lock()
	defer g.mu.Unlock()
	/**
	 文件必须使用锁
	 */
	file,_ := os.OpenFile(filename,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)
	g.out = bufio.NewWriter(file)
    return NewGlog(g)
}

func (g *Glog)output(calldepth int,s string){
	var shows = []string{s}
	var saves = []string{s}
	g.mu.Lock()
	defer g.mu.Unlock()
	//设置了输出文件
	if g.Flag &(LongFile|ShortFile) != 0{
		showLevelBool := g.showLevelBool
		saveLevelBool := g.saveLevelBool
		g.mu.Unlock()
		if showLevelBool || saveLevelBool {
			_, file, line, ok := runtime.Caller(calldepth)
			if !ok {
				file = "???"
				line = 0
			}
			if g.Flag & ShortFile != 0{
				short := file
				for i := len(file) - 1;i>0;i--{
					if file[i] == '/'{
						short = file[i+1:]
						break
					}
				}
				file  = short
			}
			fs := fmt.Sprintf("file:%s (%d)",file,line)
			if showLevelBool{
				shows = append(shows,fs)
			}
			if saveLevelBool{
				saves = append(saves,fs)
			}
		}
		g.mu.Lock()
	}
	if g.showLevelBool{
		fmt.Println(strings.Join(shows,"\n"))
	}
	if g.saveLevelBool{
		g.out.WriteString(strings.Join(saves,"\n"))
		g.out.WriteByte('\n')
		//需要移除掉 标记颜色的内容
		//增加是否需要立即保存的
		g.out.Flush()
	}
}

func (g *Glog)Debug(format string, a ...interface{}){
	if !g.checkLevelAll(debug){
		return
	}
	level := formatLevel(debug)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (g *Glog)Trace(format string, a ...interface{}){
	if !g.checkLevelAll(trace){
		return
	}
	level := formatLevel(trace)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (g *Glog)Verbose(format string, a ...interface{}){
	if !g.checkLevelAll(verbose){
		return
	}
	level := formatLevel(verbose)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (g *Glog)Asset(format string, a ...interface{}){
	if !g.checkLevelAll(assert){
		return
	}
	level := formatLevel(assert)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}


func (g *Glog)Info(format string, a ...interface{}){
	if !g.checkLevelAll(info){
		return
	}
	level := formatLevel(info)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (g *Glog)Warn(format string, a ...interface{}){
	if !g.checkLevelAll(warn){
		return
	}
	level := formatLevel(warn)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}

func (g *Glog)Error(format string, a ...interface{}){
	if !g.checkLevelAll(errors){
		return
	}
	level := formatLevel(errors)
	g.output(2,fmt.Sprint(formatPrefix(level), fmt.Sprintf(format, a...)))
}







// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13

func Verbose( format string, a ...interface{}) {

	gg.Verbose(format,a...)
}

func Trace(format string, a ...interface{}) {

	gg.Trace(format,a...)
}

func Error(format string, a ...interface{}) {

	gg.Error(format,a...)
}

func Warn( format string, a ...interface{}) {

	gg.Warn(format,a...)
}

func Info( format string, a ...interface{}) {

	gg.Info(format,a...)
}

func Debug( format string, a ...interface{}) {

	gg.Debug(format,a...)
}

func Asset(format string, a ...interface{}) {

	gg.Asset(format,a...)
}

// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func formatLevel(level string) string {

	switch level {
	case verbose:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_white, level)
	case trace:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_cyan, level)
	case errors:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, level)
	case warn:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, level)
	case info:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, level)
	case debug:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, level)
	case assert:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, level)
	default:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_white, level)
	}

}

/**
 格式化输出前缀
 */
func formatPrefix(colorTag string) string {
	return fmt.Sprintf("%s%s:", time.Now().Format("2006-01-02 15:04:05.999"), colorTag)
}


/**
 获取数组索引
 */
func getLevelIndex(s string)int{
	levelMu.Lock()
	defer levelMu.Unlock()
	if value,ok := levelsIndex[s];ok{
		return value
	}
	return -1
}

/**
 判断级别
 @param level1 string 比较的级别
 @param level2 string 欲比较的
 */
func (g *Glog)checkLevel(level1,level2 string)bool{
	if level1 == ""{
		return false
	}
	level1Index := getLevelIndex(level1)
	level2Index := getLevelIndex(level2)
	if level2Index >= level1Index{
		return true
	}  else {
		return false
	}
}

/**
 判断所有的级别
 @return bool   当为false则后面的就不走了
 */
func (g *Glog)checkLevelAll(level string)bool{
	g.mu.Lock()
	defer g.mu.Unlock()
	showBool := g.checkLevel(g.ShowLevel,level)
	saveBool := g.checkLevel(g.SaveLevel,level)
	g.showLevelBool = showBool
	g.saveLevelBool = saveBool
	return showBool || saveBool
}