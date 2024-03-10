package tmpl

import "github.com/xxcheng123/primary-school-system/rpc/pb"

/**
 * @Author: root
 * @Description:
 * @File:  category
 * @Date: 3/10/24 3:09 PM
 */

type CategoryNode struct {
	//Parent *Node
	Child    []*CategoryNode `json:"child"`
	ParentId int64           `json:"parentId"`
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	Status   int64           `json:"status"`
}

type Category struct {
	*pb.ListCategoryResp
	List []*CategoryNode `json:"tree"`
}

func BuildCategoryTree(list []*pb.CategoryInfo) (*CategoryNode, error) {
	m := map[int64]*CategoryNode{}
	var l []*CategoryNode
	for _, item := range list {
		n := &CategoryNode{
			Child:    []*CategoryNode{},
			ParentId: item.ParentId,
			Id:       item.Id,
			Name:     item.Name,
			Status:   item.Status,
		}
		m[item.Id] = n
		l = append(l, n)
	}
	var rootNode = &CategoryNode{}
	for _, item := range l {
		p, ok := m[item.ParentId]
		if item.ParentId == 0 {
			p = rootNode
		} else if !ok {
			continue
		}
		p.Child = append(p.Child, item)
	}
	return rootNode, nil
}
