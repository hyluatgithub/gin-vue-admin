package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	CasbinService
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
	ChatGptService
}
