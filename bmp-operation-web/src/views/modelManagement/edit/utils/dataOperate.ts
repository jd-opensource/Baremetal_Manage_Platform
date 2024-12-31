// deepCopy
import {filterData} from 'utils/index.ts';
import {language, locationItem, CurrentInstance} from 'utils/publicClass.ts';
import {CurrencyType} from '@utils/publicType';
// RadiosPropsType
import {CpuDataType, MemDataType, RuleFormType, PropsReactiveType} from '../typeManagement';
import store from 'store/index.ts';
// import ModelStaticData from 'staticData/model/index.ts';

// let raidList: any = [];
// let sysRaidData: any = [];
class DataOperate {
    idcData = store.idcInfo.idcDataNoOpt;
    raidsStore = store.modeFormInfo();
    // 复杂数据类型
    reactiveData: PropsReactiveType = reactive<PropsReactiveType>({
        architecture: [],
        raidConfig: [],
        raidData: {},
        volumeTypeData: [
            language.t('modelForm.label.system'),
            language.t('modelForm.label.data')
        ],
        diskTypeData: [
            {
                label: 'SSD',
                value: 'SSD'
            },
            {
                label: 'HDD',
                value: 'HDD'
            },
            {
                label: language.t('modelForm.unrestricted'),
                value: 'notLimited'
            }
        ],
        noRaidInterfaceData: [
            {
                label: 'NVME',
                value: 'NVME'
            }
        ],
        interfaceData: [
            {
                label: 'SATA',
                value: 'SATA'
            },
            {
                label: 'SAS',
                value: 'SAS'
            },
            {
                label: language.t('modelForm.unrestricted'),
                value: 'notLimited'
            }
        ],
        sysRaid: [],
        noSysRaid: [],
        sysDiskRaid0: [],
        bootModeData: [
            {
                select: true,
                label: '--'
            },
            {
                select: true,
                label: '--'
            }
        ],
        computerRoomRadioBtn: [ // 机房名称
            {
                title: '',
                idcId: '',
                showTooltip: false
            }
        ]
    });
    ruleForm: any;
    proxy = new CurrentInstance().proxy;
    // props: {cpuSpec: string; memSpec: string; hasTemplate: boolean;} & RadiosPropsType
    constructor (ruleForm: RuleFormType) {
        this.ruleForm = ruleForm;
        const val = localStorage.getItem('model-item')?? '';
        const newVal = JSON.parse(val);
        ruleForm.modelCPU = this.#setDefText(newVal.cpuSpec);
        ruleForm.modelStorage = this.#setDefText(newVal.memSpec);
        // props
        this.architectureData(ruleForm);
        if (!this.idcData?.length) {
            this.idcList(ruleForm);
            return;
        };
        this.getIdcData(this.idcData, ruleForm);
    };

    #setDefText = (value: string) => {
        const text = new Map([
            ['common', language.t('modelForm.specificationsType.presetSpecifications')]
        ]);
        return text.get(value)?? language.t('modelForm.specificationsType.otherSpecifications');
    };
    
    // , props: any
    architectureData = (ruleForm: RuleFormType) => {
        const val = localStorage.getItem('model-item')?? '';
        const newVal = JSON.parse(val);
        this.raidsStore.getCommonData()
        .then((res: {architecture: CurrencyType; cpuSpec: CpuDataType[]; memSpec: MemDataType[]; raidRules: CurrencyType; bootMode: string[]}) => {
            const {architecture, cpuSpec, memSpec, raidRules, bootMode} = res;
            this.reactiveData.raidData = raidRules;
            // 无
            const noneInfo = Object.values(raidRules['1']);
            // RAID数据
            const raidInfo = Object.values(raidRules['2']);
            // 单盘RAID0
            const singleInfo = Object.values(raidRules['3']);
            const newNone: {label: string; value: string;}[] = [];
            const newRaid: {label: string; value: string;}[] = [];
            const singleRaid: {label: string; value: string;}[] = [];
            this.reactiveData.noSysRaid = this.#setRaidData(noneInfo, newNone);
            this.reactiveData.sysRaid = this.#setRaidData(raidInfo, newRaid);
            this.reactiveData.sysDiskRaid0 = this.#setRaidData(singleInfo, singleRaid);
            const raidRulesData = Object.values(raidRules);
            this.reactiveData.raidConfig = this.#setRaidCan(raidRulesData);
            this.reactiveData.architecture = Object.keys(architecture);
            this.reactiveData.bootModeData = bootMode.map((item: string) => {
                return {
                    label: item,
                    select: true
                }
            });
            ruleForm.cpuData = cpuSpec;
            ruleForm.memData = memSpec;
            ruleForm.bootMode = bootMode.join(',');
            // ruleForm, props
            // this.setCpuMemStatus(ruleForm, newVal.cpuSpec, newVal.memSpec)
            this.raidsData(ruleForm, newVal);
        });
    };

    idcList = async (ruleForm: {machineRoomName: string; idcId: string}) => {
        try {
            const res = await this.proxy.$idcApi.idcListAPI({isAll: '1'});
            if (res?.data?.result?.idcs?.length) {
                const {idcs} = res.data.result;
                this.getIdcData(idcs, ruleForm);
            }
        }
        catch {
            throw new Error();
        }
    };

    getIdcData = (data: CurrencyType[], ruleForm: {machineRoomName: string; idcId: string}) => {
        this.reactiveData.computerRoomRadioBtn = filterData(data, 'name')
        .map((item: {name: string; nameEn: string; idcId: string}) => {
            return {
                ...item,
                title: this.#setName(locationItem.getLocationItem, item),
                idcId: item.idcId
            }
        });
        if (!ruleForm.machineRoomName?.length) {
            ruleForm.machineRoomName = this.reactiveData.computerRoomRadioBtn[0].title;
            ruleForm.idcId = this.reactiveData.computerRoomRadioBtn[0].idcId;
        }
    };

    #setName = (name: string, item: CurrencyType) => {
        const keyName: Map<string, string> = new Map([
            ['zh_CN', item.name],
            ['en_US', item.nameEn]
        ]);
        return keyName.get(name);
    };

    #setRaidCan = (raidRulesData: string[]): string[] => {
        const raidCan: string[] = [];
        raidRulesData.forEach((t: string) => {
            Object.keys(t).forEach((g: string) => {
                raidCan.push(g);
            })
        })
        return raidCan;
    }

    #setRaidData = (dataInfo: string[], raid: {label: string; value: string;}[]) => {
        dataInfo.forEach((item: string) => {
            Object.keys(item).forEach((t: any) => {
                raid.push({
                    label: t,
                    value: item[t]
                })
            })
        })
        return raid;
    }

    /**
     * 获取raids数据
    */
    raidsData = (ruleForm: any, newVal: any) => {
        for (const key in newVal) {
            for (const index in ruleForm) {
                if (key === index) {
                    this.setCpuMemStatus(ruleForm, newVal.cpuSpec, newVal.memSpec);
                    (ruleForm[`${index}` as keyof typeof ruleForm] as string) = newVal[index];
                }
            }
        }
    };

    setCpuMemStatus = (ruleForm: RuleFormType, status: string, status2: string) => {
        const sCpuMemData = [
            [
                (status: string, _: string) => status === 'common',
                () => {
                    ruleForm.cpuAmount = 1;
                    ruleForm.cpuCores =  1;
                    ruleForm.cpuFrequency =  '';
                    ruleForm.cpuManufacturer =  '';
                    ruleForm.cpuModel = '';
                }
            ],
            [
                (status: string, _: string) => status !== 'common',
                () => ruleForm.cpuInfo = ''
            ],
            [
                (_: string, status2: string) => status2 === 'common',
                () => {
                    ruleForm.memSize = '';
                    ruleForm.memAmount =  1;
                    ruleForm.memFrequency =  '';
                    ruleForm.memType = '';
                }   
            ],
            [
                (_: string, status2: string) => status2 !== 'common',
                () => ruleForm.memInfo = ''
            ]
        ];

        for (const key of sCpuMemData) {
            if (key[0](status, status2)) {
                key[1](status, status2);
            }
        }
    };
};

export default DataOperate;
