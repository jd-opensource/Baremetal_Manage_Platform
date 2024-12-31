package command

type PingHost struct {
	*BaseCommand
}

func init() {

	commands = append(commands, &PingHost{})
}

func (d *PingHost) Accept(action string) (Commandor, bool) {
	if action == "PingHost" {
		return &PingHost{BaseCommand: &BaseCommand{}}, true
	}
	return nil, false
}

func (d *PingHost) ExecuteBefore() {
	//TODO

}

func (d *PingHost) Execute() error {
	//TODO
	return nil
}

func (d *PingHost) ExecuteAfter() {
	//TODO

}
