package command

type DHCPRestart struct {
	*BaseCommand
}

func init() {

	commands = append(commands, &DHCPRestart{})
}

func (d *DHCPRestart) Accept(action string) (Commandor, bool) {
	if action == "DHCPRestart" {
		return &DHCPRestart{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *DHCPRestart) ExecuteBefore() {
	//TODO
}

func (d *DHCPRestart) Execute() error {
	//TODO
	return nil
}

func (d *DHCPRestart) ExecuteAfter() {
	//TODO
}
