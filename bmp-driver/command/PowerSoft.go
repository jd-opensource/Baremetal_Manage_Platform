package command

type PowerSoft struct {
	*BaseCommand
}

func init() {

	commands = append(commands, &PowerSoft{})
}

func (d *PowerSoft) Accept(action string) (Commandor, bool) {
	if action == "PowerSoft" {
		return &PowerSoft{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PowerSoft) ExecuteBefore() {
	//TODO

}

func (d *PowerSoft) Execute() error {
	//TODO
	return nil
}

func (d *PowerSoft) ExecuteAfter() {
	//TODO

}
