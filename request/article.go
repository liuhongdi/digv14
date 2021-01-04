package request

type ArticleRequest struct {
	ID    uint64 `form:"id" binding:"required,gte=1"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Page int `form:"page" binding:"gte=1"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
