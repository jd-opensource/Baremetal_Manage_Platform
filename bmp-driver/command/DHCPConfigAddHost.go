package command

type DHCPConfigAddHost struct {
	*BaseCommand
}

func init() {

	commands = append(commands, &DHCPConfigAddHost{})
}

func (d *DHCPConfigAddHost) Accept(action string) (Commandor, bool) {
	if action == "DHCPConfigAddHost" {
		return &DHCPConfigAddHost{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *DHCPConfigAddHost) ExecuteBefore() {
	//TODO
}

func (d *DHCPConfigAddHost) Execute() error {
	//TODO
	return nil
}

func (d *DHCPConfigAddHost) ExecuteAfter() {
	//TODO
}
