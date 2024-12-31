package SaleStatus

const (
	/**
	 * 上架中
	 */
	PUTAWAYING = "putawaying"

	/**
	 * 在售
	 */
	SELLING = "selling"

	/**
	 * 已售
	 */
	SOLD = "sold"

	/**
	 * 锁定
	 */
	LOCK = "lock"

	IN           = "in"           //已入库
	PUTAWAY      = "putaway"      //已上架
	CREATED      = "created"      //已创建
	PUTAWAYFAIL  = "putawayfail"  //上架失败
	UNPUTAWAY    = "unputaway"    //下架
	UNPUTAWAYING = "unputawaying" //下架中，将“已上架”的设备执行“下架”操作的中间过程状态
	REMOVED      = "removed"      //已移除
)
