package settings // import "kego.io/process/settings"

type Settings struct {
	Dir       string
	Edit      bool
	Update    bool
	Recursive bool
	Verbose   bool
	Path      string
	Aliases   map[string]string
}

func New() *Settings {
	return &Settings{}
}
