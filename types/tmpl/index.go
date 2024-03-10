package tmpl

/**
 * @Author: root
 * @Description:
 * @File:  index
 * @Date: 3/10/24 2:52 PM
 */

type Index struct {
	Base
	Swiper   *Swiper   `json:"swiper" comment:"轮播图"`
	Category *Category `json:"category" comment:"导航栏"`
	JSON     any       `json:"json" comment:"json"`
}
