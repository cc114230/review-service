package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"
)

type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	SaveReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
	GetReview(context.Context, int64) (*model.ReviewInfo, error)
	ListReviewByUserID(context.Context, int64) ([]*model.ReviewInfo, error)
	AppealReview(context.Context, *AppealParam) (*model.ReviewAppealInfo, error)
}

type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// CreateReview 创建评价,service层调用该方法
func (uc *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReview, req:%v", review)
	// 1.参数业务校验:带业务逻辑的参数校验，比如已经评价过的订单不能再评价
	reviews, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, v1.ErrorDbFailed("查询数据库失败")
	}
	if len(reviews) > 0 {
		// 已经评价过

		return nil, v1.ErrorOrderReviewed("订单:%d已评价", review.OrderID)
	}
	// 2.生成reviewID
	review.ReviewID = snowflake.GenID()
	// 3.查询订单和商品快照信息
	// 4.拼装数据入库
	return uc.repo.SaveReview(ctx, review)
}

func (uc *ReviewUsecase) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] GetReview, reviewID:%v", reviewID)
	return uc.repo.GetReview(ctx, reviewID)

}

func (uc *ReviewUsecase) ListReviewByUserID(ctx context.Context, userID int64) ([]*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] GetReview, UserID:%v", userID)
	return uc.repo.ListReviewByUserID(ctx, userID)
}

// CreateReply 创建评价回复
func (uc *ReviewUsecase) CreateReply(ctx context.Context, param *ReplyParam) (*model.ReviewReplyInfo, error) {
	// 调用data层创建一个评价的回复
	uc.log.WithContext(ctx).Debugf("[biz] CreateReply param:%v", param)
	reply := &model.ReviewReplyInfo{
		ReplyID:   snowflake.GenID(),
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	}
	return uc.repo.SaveReply(ctx, reply)
}

func (uc *ReviewUsecase) AppealReview(ctx context.Context, param *AppealParam) (*model.ReviewAppealInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] AppealReview param:%v", param)
	return uc.repo.AppealReview(ctx, param)
}
