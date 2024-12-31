package collector

func init() {
	Register("disk", &DiskCollector{})
}

type DiskCollector struct {
}
