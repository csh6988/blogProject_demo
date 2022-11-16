package models

import "github.com/csh6988/blogProject_demo/config"

/**
 * @Author: cheney
 * @Date: 2022/11/16 11:11 PM
 * @Desc:
 */

type HomeData struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
