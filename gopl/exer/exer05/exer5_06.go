package main

import (
	"fmt"
	"math"
)

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Surface는 3-D 표면 함수의 SVG 렌더링을 계산합니다.

const (
	width, height = 600, 320            // 픽셀 단위의 캔버스 크기
	cells         = 100                 // 그리드 셀의 수
	xyrange       = 30.0                // 축 범위 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x나 y의 단위 픽셀
	zscale        = height * 0.4        // z 단위 픽셀
	angle         = math.Pi / 6         // x, y축의 각도 (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func ComputeSurface() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64) {
	// 셀 (i,j)의 모퉁이점 (x,y)를 찾는다.
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 표면 높이 z를 계산한다.
	z := f(x, y)

	// (x,y,z)를 2-D SVG 캔버스 (sx, sy)에 등각 투영한다.
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)에서의 거리
	return math.Sin(r) / r
}
