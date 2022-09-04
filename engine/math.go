package engine

import "math"

type Vec2 struct {
	X, Y float32
}

type Size struct {
	W, H float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Mat4 [4][4]float32

// 创建正交投影矩阵。
func CreateOrthographicOffCenter(left, right, bottom, top, zNearPlane, zFarPlane float32) Mat4 {
	var mat4 Mat4
	mat4[0][0] = 2.0 / (right - left) //M11
	mat4[0][1] = 0.0                  //M12
	mat4[0][2] = 0.0                  //M13
	mat4[0][3] = 0.0                  //M14

	mat4[1][0] = 0.0                  //M21
	mat4[1][1] = 2.0 / (top - bottom) //M22
	mat4[1][2] = 0.0                  //M23
	mat4[1][3] = 0.0                  //M24

	mat4[2][0] = 0.0                            //M31
	mat4[2][1] = 0.0                            //M32
	mat4[2][2] = 1.0 / (zNearPlane - zFarPlane) //M33
	mat4[2][3] = 0.0                            //M34

	mat4[3][0] = (left + right) / (left - right)       //M41
	mat4[3][1] = (top + bottom) / (bottom - top)       //M42
	mat4[3][2] = zNearPlane / (zNearPlane - zFarPlane) //M43
	mat4[3][3] = 1.0                                   //M44

	return mat4
}

// 创建缩放矩阵。
func CreateScale1f(scale float32) Mat4 {
	var mat4 Mat4
	mat4[0][0] = scale //M11
	mat4[0][1] = 0.0   //M12
	mat4[0][2] = 0.0   //M13
	mat4[0][3] = 0.0   //M14

	mat4[1][0] = 0.0   //M21
	mat4[1][1] = scale //M22
	mat4[1][2] = 0.0   //M23
	mat4[1][3] = 0.0   //M24

	mat4[2][0] = 0.0   //M31
	mat4[2][1] = 0.0   //M32
	mat4[2][2] = scale //M33
	mat4[2][3] = 0.0   //M34

	mat4[3][0] = 0.0 //M41
	mat4[3][1] = 0.0 //M42
	mat4[3][2] = 0.0 //M43
	mat4[3][3] = 1.0 //M44
	return mat4
}

// 创建缩放矩阵。
func CreateScale3f(xScale, yScale, zScale float32) Mat4 {
	var mat4 Mat4
	mat4[0][0] = xScale //M11
	mat4[0][1] = 0.0    //M12
	mat4[0][2] = 0.0    //M13
	mat4[0][3] = 0.0    //M14

	mat4[1][0] = 0.0    //M21
	mat4[1][1] = yScale //M22
	mat4[1][2] = 0.0    //M23
	mat4[1][3] = 0.0    //M24

	mat4[2][0] = 0.0    //M31
	mat4[2][1] = 0.0    //M32
	mat4[2][2] = zScale //M33
	mat4[2][3] = 0.0    //M34

	mat4[3][0] = 0.0 //M41
	mat4[3][1] = 0.0 //M42
	mat4[3][2] = 0.0 //M43
	mat4[3][3] = 1.0 //M44
	return mat4
}

// 两个矩阵相乘。
func Multiply(a, b Mat4) Mat4 {
	var mat4 Mat4
	mat4[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0] + a[0][2]*b[2][0] + a[0][3]*b[3][0] //M11
	mat4[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1] + a[0][2]*b[2][1] + a[0][3]*b[3][1] //M12
	mat4[0][2] = a[0][0]*b[0][2] + a[0][1]*b[1][2] + a[0][2]*b[2][2] + a[0][3]*b[3][2] //M13
	mat4[0][3] = a[0][0]*b[0][3] + a[0][1]*b[1][3] + a[0][2]*b[2][3] + a[0][3]*b[3][3] //M14

	mat4[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0] + a[1][2]*b[2][0] + a[1][3]*b[3][0] //M21
	mat4[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1] + a[1][2]*b[2][1] + a[1][3]*b[3][1] //M22
	mat4[1][2] = a[1][0]*b[0][2] + a[1][1]*b[1][2] + a[1][2]*b[2][2] + a[1][3]*b[3][2] //M23
	mat4[1][3] = a[1][0]*b[0][3] + a[1][1]*b[1][3] + a[1][2]*b[2][3] + a[1][3]*b[3][3] //M24

	mat4[2][0] = a[2][0]*b[0][0] + a[2][1]*b[1][0] + a[2][2]*b[2][0] + a[2][3]*b[3][0] //M31
	mat4[2][1] = a[2][0]*b[0][1] + a[2][1]*b[1][1] + a[2][2]*b[2][1] + a[2][3]*b[3][1] //M32
	mat4[2][2] = a[2][0]*b[0][2] + a[2][1]*b[1][2] + a[2][2]*b[2][2] + a[2][3]*b[3][2] //M33
	mat4[2][3] = a[2][0]*b[0][3] + a[2][1]*b[1][3] + a[2][2]*b[2][3] + a[2][3]*b[3][3] //M34

	mat4[3][0] = a[3][0]*b[0][0] + a[3][1]*b[1][0] + a[3][2]*b[2][0] + a[3][3]*b[3][0] //M41
	mat4[3][1] = a[3][0]*b[0][1] + a[3][1]*b[1][1] + a[3][2]*b[2][1] + a[3][3]*b[3][1] //M42
	mat4[3][2] = a[3][0]*b[0][2] + a[3][1]*b[1][2] + a[3][2]*b[2][2] + a[3][3]*b[3][2] //M43
	mat4[3][3] = a[3][0]*b[0][3] + a[3][1]*b[1][3] + a[3][2]*b[2][3] + a[3][3]*b[3][3] //M44
	return mat4
}

// 创建平移矩阵。
func CreateTranslation(xPosition, yPosition, zPosition float32) Mat4 {
	var mat4 Mat4
	mat4[0][0] = 1.0 //M11
	mat4[0][1] = 0.0 //M12
	mat4[0][2] = 0.0 //M13
	mat4[0][3] = 0.0 //M14

	mat4[1][0] = 0.0 //M21
	mat4[1][1] = 1.0 //M22
	mat4[1][2] = 0.0 //M23
	mat4[1][3] = 0.0 //M24

	mat4[2][0] = 0.0 //M31
	mat4[2][1] = 0.0 //M32
	mat4[2][2] = 1.0 //M33
	mat4[2][3] = 0.0 //M34

	mat4[3][0] = xPosition //M41
	mat4[3][1] = yPosition //M42
	mat4[3][2] = zPosition //M43
	mat4[3][3] = 1.0       //M44
	return mat4
}

// 创建围绕 X 轴旋转的矩阵。
func CreateRotationX(radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))
	var mat4 Mat4
	mat4[0][0] = 1.0
	mat4[0][1] = 0.0
	mat4[0][2] = 0.0
	mat4[0][3] = 0.0

	mat4[1][0] = 0.0
	mat4[1][1] = num1
	mat4[1][2] = num2
	mat4[1][3] = 0.0

	mat4[2][0] = 0.0
	mat4[2][1] = -num2
	mat4[2][2] = num1
	mat4[2][3] = 0.0

	mat4[3][0] = 0.0
	mat4[3][1] = 0.0
	mat4[3][2] = 0.0
	mat4[3][3] = 1.0
	return mat4
}

// 创建围绕 X 轴旋转的矩阵。
func RotationX1f(mat4 Mat4, radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))

	mat4[1][1] *= num1
	mat4[1][2] *= num2

	mat4[2][1] *= -num2
	mat4[2][2] *= num1
	return mat4
}

// 创建围绕 Y 轴旋转的矩阵。
func CreateRotationY(radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))
	var mat4 Mat4
	mat4[0][0] = num1
	mat4[0][1] = 0.0
	mat4[0][2] = -num2
	mat4[0][3] = 0.0

	mat4[1][0] = 0.0
	mat4[1][1] = 1.0
	mat4[1][2] = 0.0
	mat4[1][3] = 0.0

	mat4[2][0] = num2
	mat4[2][1] = 0.0
	mat4[2][2] = num1
	mat4[2][3] = 0.0

	mat4[3][0] = 0.0
	mat4[3][1] = 0.0
	mat4[3][2] = 0.0
	mat4[3][3] = 1.0
	return mat4
}

// 创建围绕 Y 轴旋转的矩阵。
func RotationY1f(mat4 Mat4, radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))
	mat4[0][0] *= num1
	mat4[0][2] *= -num2

	mat4[2][0] *= num2
	mat4[2][2] *= num1
	return mat4
}

// 创建围绕 Z 轴旋转的矩阵。
func CreateRotationZ(radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))
	var mat4 Mat4
	mat4[0][0] = num1 //M11
	mat4[0][1] = num2 //M12
	mat4[0][2] = 0.0  //M13
	mat4[0][3] = 0.0  //M14

	mat4[1][0] = -num2 //M21
	mat4[1][1] = num1  //M22
	mat4[1][2] = 0.0   //M23
	mat4[1][3] = 0.0   //M24

	mat4[2][0] = 0.0 //M31
	mat4[2][1] = 0.0 //M32
	mat4[2][2] = 1.0 //M33
	mat4[2][3] = 0.0 //M34

	mat4[3][0] = 0.0 //M41
	mat4[3][1] = 0.0 //M42
	mat4[3][2] = 0.0 //M43
	mat4[3][3] = 1.0 //M44
	return mat4
}

// 创建围绕 Z 轴旋转的矩阵。
func RotationZ1f(mat4 Mat4, radians float32) Mat4 {
	angle := radians * math.Pi / 180.0
	num1 := float32(math.Cos(float64(angle)))
	num2 := float32(math.Sin(float64(angle)))

	mat4[0][0] *= num1 //M11
	mat4[0][1] *= num2 //M12

	mat4[1][0] *= -num2 //M21
	mat4[1][1] *= num1  //M22
	return mat4
}
