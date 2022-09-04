package engine

import (
	_ "embed"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

//go:embed vertex.glsl
var vertexStr string

//go:embed fragment.glsl
var fragmentStr string

// 引擎
type Object struct {
	window              *glfw.Window
	camera              Camera
	shader              Shader
	vao                 VAO
	vbo                 VBO
	ebo                 EBO
	isRunning           bool
	deltaTime           float64
	totalElapsedSeconds float64
	closeCb             CloseCallback
	sizeCb              SizeCallback
	keyCb               KeyCallback
	mouseButtonCb       MouseButtonCallback
	cursorPosCb         CursorPosCallback
	scrollCb            ScrollCallback
}

func (g *Object) GetCamera() *Camera {
	return &g.camera
}

// 初始化
func (g *Object) Init(w, h int) bool {
	//初始化日志
	if err := os.MkdirAll("./log", 0777); err != nil {
		panic(err)
	}
	cfg := &logCfg{
		Level:      "debug",
		FileName:   "./log/log.txt",
		MaxSize:    1024,
		MaxAge:     90,
		MaxBackups: 10,
	}
	if err := logInit(cfg); err != nil {
		panic(err)
	}
	//初始化glfw
	if err := glfw.Init(); err != nil {
		Log("glfw.Init", err)
		return false
	}
	glfw.WindowHint(glfw.RedBits, 8)
	glfw.WindowHint(glfw.GreenBits, 8)
	glfw.WindowHint(glfw.BlueBits, 8)
	glfw.WindowHint(glfw.AlphaBits, 8)
	glfw.WindowHint(glfw.DepthBits, 24)
	glfw.WindowHint(glfw.DoubleBuffer, 1)
	//设置版本
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//创建窗口
	var err error
	if g.window, err = glfw.CreateWindow(1, 1, "", nil, nil); err != nil {
		Log("glfw.CreateWindow", err)
		return false
	}
	//居中显示
	_, _, px, py := glfw.GetPrimaryMonitor().GetWorkarea()
	g.window.SetPos((px-w)/2, (py-h)/2)
	g.window.SetSize(w, h) //解决出口一闪的问题
	//g.window.SetSizeLimits()//大小限制
	g.window.SetAspectRatio(16, 9) //窗口比例
	g.window.SetAttrib(glfw.Resizable,glfw.False)//禁止调整大小
	//回调函数
	g.window.SetCloseCallback(g.closeCallback)             //关闭回调
	g.window.SetSizeCallback(g.sizeCallback)               //调整回调
	g.window.SetKeyCallback(g.keyCallback)                 //按键回调
	g.window.SetMouseButtonCallback(g.mouseButtonCallback) //鼠标按键
	g.window.SetCursorPosCallback(g.cursorPosCallback)     //鼠标位置
	g.window.SetScrollCallback(g.scrollCallback)           //鼠标滚动
	//上下文
	g.window.MakeContextCurrent()
	//垂直同步
	glfw.SwapInterval(1)
	//初始化opengl
	if err := gl.Init(); err != nil {
		Log("gl.Init", err)
		return false
	}
	//gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Viewport(0, 0, int32(w), int32(h))
	//摄像机
	g.camera = NewCamera(g)
	//着色器
	var ok bool
	if g.shader, ok = NewShader(vertexStr, fragmentStr); !ok {
		return false
	}
	g.shader.Use()
	//VAO、VBO、EBO
	g.vao = NewVao(g)
	g.vbo = NewVbo(g)
	g.vbo.Upload([]float32{
		//position       //uv
		0.5, -0.5, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.0, 1.0, 1.0,
		-0.5, -0.5, 0.0, 0.0, 0.0,
	})
	g.ebo = NewEbo(g)
	g.ebo.Upload([]int32{
		2, 1, 0,
		0, 1, 3,
	})
	positionSize := int32(3)
	uvSize := int32(2)
	vertexSizeBytes := positionSize + uvSize
	g.vbo.Bind(0, positionSize, vertexSizeBytes, 0)
	g.vbo.Bind(1, uvSize, vertexSizeBytes, positionSize)
	//日志
	Log("glfw:" + glfw.GetVersionString())
	Log("opengl:" + gl.GoStr(gl.GetString(gl.VERSION)))
	g.SetRunning(true)
	return true
}

// 获取运行状态
func (g *Object) GetRunning() bool {
	return g.isRunning
}

// 设置运行状态
func (g *Object) SetRunning(run bool) {
	g.isRunning = run
}
