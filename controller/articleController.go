package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv14/pkg/page"
	"github.com/liuhongdi/digv14/pkg/result"
	"github.com/liuhongdi/digv14/pkg/validCheck"
	"github.com/liuhongdi/digv14/request"
	"github.com/liuhongdi/digv14/service"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}
//得到一篇文章的详情
func (a *ArticleController) GetOne(c *gin.Context) {
	result := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		result.Error(400,errs.Error())
		return
	}

	if (param.ID == 100) {
		var z int = 0
		var i int = 100 / z
		fmt.Println("i:%i",i)
	}

	articleOne,err := service.GetOneArticle(param.ID);
	if err != nil {
		result.Error(404,"数据查询错误")
	} else {
		result.Success(&articleOne);
	}
	return
}

//得到多篇文章，按分页返回
func (a *ArticleController) GetList(c *gin.Context) {
	fmt.Println("--------------------begin ArticleController GetList")
	result := result.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
    //if curPage not exist
    if (len(curPage) == 0) {
		pageInt = 1
	} else {
		param := request.ArticleListRequest{Page: validCheck.StrTo(c.Param("page")).MustInt()}
		valid, errs := validCheck.BindAndValid(c, &param)
		if !valid {
			result.Error(400,errs.Error())
			return
		}
		pageInt = param.Page
	}

	pageSize := 2;
	pageOffset := (pageInt-1) * pageSize

	articles,err := service.GetArticleList(pageOffset,pageSize)
	if err != nil {
		result.Error(404,"数据查询错误");
		fmt.Println(err.Error())
	} else {
		//sum,_ := dao.SelectcountAll()
		sum,_ := service.GetArticleSum()
		pageInfo,_ := page.GetPageInfo(pageInt,pageSize,sum)
		result.Success(gin.H{"list":&articles,"pageinfo":pageInfo});
	}
	return
}

