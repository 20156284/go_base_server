package interfaces

type InitDateFunc interface {
	Init() error
}

//@description: 批量初始化表数据
func InitDb(inits ...InitDateFunc) error {
	for _, init := range inits {
		if err := init.Init(); err != nil {
			return err
		}
	}
	return nil
}
