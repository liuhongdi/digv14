package page

import "math"

type PageInfoCont struct {
	PageSum	int  `json:"pageSum"`         // 总页数
	CurrentPage int  `json:"currentPage"` //当前页
	PrevPage  int  `json:"prevPage"`      //前一页
	NextPage  int  `json:"nextPage"`      //后一页
	FirstPage  int  `json:"firstPage"`    //第一页
	LastPage  int  `json:"lastPage"`      //最后一页
}

//根据item总数得到分页数据
func GetPageInfo(currentPage int,pageSize int,itemSum int) (PageInfoCont, error) {
	res := PageInfoCont{}
	res.PageSum = itemSum / pageSize ;
	totalpages := int(math.Ceil(float64(itemSum) / float64(pageSize)))

	if currentPage > totalpages {
		currentPage = totalpages
	}
	if currentPage <= 0 {
		currentPage = 1
	}

	prevPage := currentPage-1
	if (prevPage <= 1){
		prevPage = 0
	}

	nextPage := currentPage+1
	if (nextPage >= totalpages){
		nextPage = 0
	}

	res.PageSum = totalpages
	res.CurrentPage = currentPage
	res.FirstPage = 1
	res.LastPage = totalpages
	res.PrevPage = prevPage
	res.NextPage = nextPage
	return res,nil
}