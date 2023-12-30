package biz

// ReplyParam 商家回复评价的参数
type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// AppealParam 商家申诉评价的参数
type AppealParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
	OpUser    string
}
