package logic

import (
	"fmt"
	"gin_blog/dao"
	"gin_blog/logger"
	"gin_blog/models"
	sql "github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// 文章相关逻辑

const (
	KeyArticleCountPerDay = "article:count:%s"
)

func ArticleReadCountIncr(articleId int)(err error){
	// 为指定文章增加阅读数
	date := time.Now().Format("20060102")
	redisKey := fmt.Sprintf(KeyArticleCountPerDay, date)
	logger.Debug("ArticleReadCountIncr", zap.String("redisKey", redisKey), zap.Int("articleId", articleId))
	err = dao.Client.ZIncrBy(redisKey, 1, fmt.Sprintf("%d",articleId)).Err()
	return
}


func ArticleTopN(n int64)([]int64, error){
	// 获取阅读数排名前n位的
	date := time.Now().Format("20060102")
	redisKey := fmt.Sprintf(KeyArticleCountPerDay, date)
	idStrs, err := dao.Client.ZRevRange(redisKey, 0, n-1).Result()
	if err != nil {
		return nil, err
	}
	var ids = make([]int64, len(idStrs))
	for _, idStr := range idStrs{
		id, err := strconv.ParseInt(idStr, 0, 16)
		if err != nil {
			logger.Warn("ArticleTopN:strconv.ParseInt failed", zap.Any("error", err))
			continue
		}
		ids = append(ids, id)
	}
	return ids, nil
}


func QueryArticlesByIds(ids []int64)([]*models.Article, error){

	query, args, err := sql.In("select id, title from article where id in (?)", ids)
	if err != nil {
		logger.Error("QueryArticlesByIds", zap.Any("error", err))
		return nil, err
	}
	//sqlStr := "select id, title from article where id in (?)"
	var dest []*models.Article
	err = dao.QueryRows(&dest, query, args...)
	return dest, err
}