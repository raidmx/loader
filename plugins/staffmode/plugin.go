package staffmode

type Plugin struct {
}

func (Plugin) Name() string {
	return "Staff Mode"
}

func (Plugin) Description() string {
	return "Adds support for Staff Mode"
}

func (Plugin) Author() string {
	return "Crayder"
}

func (Plugin) Version() string {
	return "1.0.0"
}

func (Plugin) OnLoad() {
}

func (Plugin) OnUnload() {

}
