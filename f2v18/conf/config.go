package conf

import (
	"encoding/json"
	"image/color"
	"os"
	"sync"
)

type Config struct {
	GameName    string        // 游戏名
	MarginHor   int           // 窗口水平边距
	MarginVer   int           // 窗口竖直边距
	WidthLine   int           // 棋格线粗
	WidthFrame  int           // 棋盘外围格线粗
	RadiusPiece int           // 棋子半径
	Antialias   float32       // 抗锯齿
	LnHorizon   int           // 棋盘水平线数量
	LnVertical  int           // 棋盘竖直线数量
	ColorBoard  *color.RGBA   // 棋盘背景
	ColorLine   *color.RGBA   // 棋格线
	ColorPieces []*color.RGBA // 棋子
}

func GetInstance() *Config {
	once.Do(func() {
		instance = newConfig()
	})
	return instance
}

var instance *Config
var once sync.Once

func newConfig() *Config {
	c := &Config{
		GameName:    "2v18",
		MarginHor:   100,
		MarginVer:   100,
		ColorBoard:  &color.RGBA{230, 179, 61, 255},
		ColorLine:   &color.RGBA{0, 0, 0, 255},
		ColorPieces: make([]*color.RGBA, 0),
		WidthFrame:  6,
		WidthLine:   4,
		RadiusPiece: 10,
		Antialias:   3,
		LnHorizon:   6,
		LnVertical:  6,
	}
	c.ColorPieces = append(c.ColorPieces, &color.RGBA{255, 255, 255, 255})
	c.ColorPieces = append(c.ColorPieces, &color.RGBA{255, 76, 0, 255})
	return c
}

func (c *Config) Save(filepath string) error {
	var data []byte
	var err error
	if data, err = json.MarshalIndent(c, "", "    "); err == nil {
		return os.WriteFile(filepath, data, 0666)
	}
	return err
}
func (c *Config) Load(filepath string) (*Config, error) {
	var data []byte
	var err error
	if data, err = os.ReadFile(filepath); err == nil {
		err = json.Unmarshal(data, c)
	}
	return c, err
}
