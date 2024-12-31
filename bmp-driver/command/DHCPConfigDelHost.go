package command

type DHCPConfigDelHost struct {
	*BaseCommand
}

func init() {

	commands = append(commands, &DHCPConfigDelHost{})
}

func (d *DHCPConfigDelHost) Accept(action string) (Commandor, bool) {
	if action == "DHCPConfigDelHost" {
		return &DHCPConfigDelHost{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *DHCPConfigDelHost) ExecuteBefore() {
	//TODO
}

func (d *DHCPConfigDelHost) Execute() error {
	//TODO
	return nil
}

func (d *DHCPConfigDelHost) ExecuteAfter() {
	//TODO
}
