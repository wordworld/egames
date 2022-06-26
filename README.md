# 简单的 2D 游戏

Simple 2D games powered by [Ebitengine](https://ebiten.org).

## 战斗队形 2v18

战斗队形 2v18 是一款经典的模拟围攻棋盘游戏。红蓝双方在 6x6 的网格棋盘上，分别拥有 2 枚和 18 枚棋子。

### 开局

* 18 枚蓝方棋子（用 0-9 a-h 表示）占满蓝方底线开始的 3 行格点
* 2 枚红方棋子（用 I、J 表示）可沿红方底线格点自由摆放 
* 形如：

```
0 ── 1 ── 2 ── 3 ── 4 ── 5
│    │    │    │    │    │
6 ── 7 ── 8 ── 9 ── a ── b
│    │    │    │    │    │
c ── d ── e ── f ── g ── h
│    │    │    │    │    │
│ ── ┼ ── ┼ ── ┼ ── ┼ ── │
│    │    │    │    │    │
│ ── ┼ ── ┼ ── ┼ ── ┼ ── │
│    │    │    │    │    │
└────┴─── I ── J ───┴────┘
```

### 玩法

* 每个回合双方在棋盘内各走一步
* 每步只能走一格，且落子处不能有棋子
* 当红方落子后，若有一枚蓝方棋子与两枚红方棋子共线相邻、且不在红方棋子中间，则红方形成战斗队形，该蓝子被从棋盘上淘汰

### 胜负标准

* 若蓝方棋子全被淘汰，则红方胜
* 若红方棋子无处可走，则蓝方胜

## 参考

[Ebiten Tour](https://ebiten.org/tour/)
| [Documents](https://ebiten.org/documents/)
| [API](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2)
| [Examples](https://ebiten.org/examples/)
| [go-inovation](https://github.com/hajimehoshi/go-inovation)

[Android Studio 安装、配置 NDK](https://developer.android.google.cn/studio/projects/install-ndk?hl=zh-cn)
| [AndroidDevTools](https://www.androiddevtools.cn/)

[Cocos Creator Shader Effect 系列 - 6 - 内发光特效](https://www.jianshu.com/p/326b73f86ecc)

[RGB 颜色参考](https://tool.oschina.net/commons?type=3)

## Errors

[gomobile: no usable NDK in $ANDROID_HOME](https://githubwyb.github.io/blogs/2022-05-24-gomobile/)

[Gradle sync 报错 Plugin was not found](https://blog.csdn.net/qq_41624557/article/details/123848212)

[Build was configured to prefer settings repositories over project repositories](https://blog.csdn.net/hpp_1225/article/details/119888981)