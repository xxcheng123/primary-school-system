package fcp

import (
	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/rpc/pb"
)

/**
 * @Author: root
 * @Description:
 * @File:  swiper
 * @Date: 3/8/24 7:35 PM
 */

//用于快速复制swiper

func CPSwiper(src *ent.Swiper, dst *pb.InfoSwiper) error {
	dst.Title = src.Title
	dst.Url = src.URL
	dst.Id = src.ID
	dst.Order = src.Order
	dst.Img = src.Img
	dst.Status = src.Status
	dst.UpdateTime = src.UpdateTime.UnixMilli()
	return nil
}
func ConvSwiper(src *ent.Swiper) (*pb.InfoSwiper, error) {
	p := &pb.InfoSwiper{}
	err := CPSwiper(src, p)
	return p, err
}
