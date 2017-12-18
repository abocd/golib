package glog
import (
	"fmt"
	"time"
	"bufio"
	"os"
	"sync"
	"runtime"
	"strings"
	"path/filepath"
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
	//是否需要手动flush写入文件
	NeedFlush bool
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
	filename,err := filepath.Abs(filename)
	if err != nil{
		Error(err.Error())
	}
	dir := filepath.Dir(filename)

	_,err = os.Stat(dir)
	if err != nil{
		Error(err.Error())
	}
	if os.IsNotExist(err){
		os.MkdirAll(dir,0777)
	}
	file,_ := os.OpenFile(filename,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)
	g.out = bufio.NewWriter(file)
    return NewGlog(g)
}

func (g *Glog)output(level,s string,calldepth int){
	var shows = []string{fmt.Sprintf("%s%s",formatPrefix(formatLevel(level)),s)}
	var saves = []string{fmt.Sprintf("%s%s",formatPrefix(level),s)}
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
		g.out.WriteString("\n")
		//需要移除掉 标记颜色的内容
		//增加是否需要立即保存的
		if !g.NeedFlush {
			g.out.Flush()
		}
	}
}

func (g *Glog)Flush(){
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.NeedFlush {
		g.out.Flush()
	}
}

func (g *Glog)Debug( a ...interface{}){
	if !g.checkLevelAll(debug){
		return
	}
	g.output(debug,fmt.Sprint(a...),2)
}

func (g *Glog)Trace(a ...interface{}){
	if !g.checkLevelAll(trace){
		return
	}
	g.output(trace,fmt.Sprint( a...),2)
}

func (g *Glog)Verbose(a ...interface{}){
	if !g.checkLevelAll(verbose){
		return
	}
	g.output(verbose,fmt.Sprint(a...),2)
}

func (g *Glog)Asset(a ...interface{}){
	if !g.checkLevelAll(assert){
		return
	}
	g.output(assert,fmt.Sprint( a...),2)
}


func (g *Glog)Info(a ...interface{}){
	if !g.checkLevelAll(info){
		return
	}
	g.output(info,fmt.Sprint(a...),2)
}

func (g *Glog)Warn(a ...interface{}){
	if !g.checkLevelAll(warn){
		return
	}
	g.output(warn,fmt.Sprint(a...),2)
}

func (g *Glog)Error(a ...interface{}){
	if !g.checkLevelAll(errors){
		return
	}
	g.output(errors,fmt.Sprint(a...),2)
}


// https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13

func Verbose( a ...interface{}) {

	gg.Verbose(a...)
}

func Trace(a ...interface{}) {

	gg.Trace(a...)
}

func Error(a ...interface{}) {

	gg.Error(a...)
}

func Warn( a ...interface{}) {

	gg.Warn(a...)
}

func Info( a ...interface{}) {

	gg.Info(a...)
}

func Debug( a ...interface{}) {

	gg.Debug(a...)
}

func Asset(a ...interface{}) {

	gg.Asset(a...)
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