package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	//user模块错误
	ErrorExistUser         = 30001
	ErrorFailEncryption    = 30002
	NullPassword           = 30003
	UserNotFind            = 30004
	ErrorPassword          = 30005
	ErrorToken             = 30006
	ErrorCheckTokenTimeOut = 30007
	ErrorUploadFail        = 30008
	ErrorSendEmail         = 30009
	ErrorNotActive         = 30010

	//product模块错误
	ErrorProductImgUpload = 40001
	ProductNotFind        = 40002
	ErrorBossForProduct   = 40003
	ErrorNotSale          = 40004

	//收藏夹模块
	ErrorFavoriteExist = 50001
)
