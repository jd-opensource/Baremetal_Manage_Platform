package response

//新增加或者删除返给前端的true或者false
type CommonResponse struct {
	Success bool `json:"success"`
}
type FilterList struct {
	Architecture ArchitectureList                     `json:"architecture"`
	CPU          []CPU                                `json:"cpuSpec"`
	Mem          []Mem                                `json:"memSpec"`
	RaidRules    map[int]map[string]map[string]string `json:"raidRules"`
	BootMode     []string                             `json:"bootMode"` //新增bootmode配置@minping，国梁
}
type ArchitectureList struct {
	X86_64    string `json:"x86_64"`
	Arm       string `json:"ARM64(aarch64)"`
	I386      string `json:"i386"`
	LoongArch string `json:"LoongArch™"`
}
type CPU struct {
	Value int    `json:"value"`
	Label string `json:"label"`
	Info  struct {
		CPUManufacturer string `json:"cpuManufacturer"`
		CPUCores        int    `json:"cpuCores"`
		CPUModel        string `json:"cpuModel"`
		CPUAmount       int    `json:"cpuAmount"`
		CPUFrequency    string `json:"cpuFrequency"`
	} `json:"info"`
}
type Mem struct {
	Value int    `json:"value"`
	Label string `json:"label"`
	Info  struct {
		MemType      string `json:"memType"`
		MemSize      int    `json:"memSize"`
		MemAmount    int    `json:"memAmount"`
		MemFrequency int    `json:"memFrequency"`
	} `json:"info"`
}
