import {RuleFormRefType} from '@utils/publicType';
// 
import {RuleFormType, RulesType, ProcessingOperationsType} from '../typeManagement'
import reactiveData from './reactiveData';
import {regExpCheck} from './regExpCheck';
import {locationItem, language} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import RegularContent from 'utils/regular.ts';
import ModelStaticData from 'staticData/model/index.ts'
import {addModelAPI} from 'api/model/request.ts'; // 添加机型接口
// import StorageOperate from './storage';
import {Router, useRouter} from 'vue-router';
import {msgTip, methodsTotal} from 'utils/index.ts';

interface DataCountType {
    hasAddFlag: Ref<boolean>;
    hasAddFlag2: Ref<boolean>;
    raidCanFlag: Ref<boolean>;
    isLoading: Ref<boolean>;
    volumeNameTip: Ref<string>;
    volumeTypeTip: Ref<string>;
    raidCanTip: Ref<string>;
    sysRaidTip: Ref<string>;
    diskTypeTip: Ref<string>;
    interfaceTypeTip: Ref<string>;
    minSizeNumTip: Ref<string>;
    amountTip: Ref<string>;
    tableHasError: Ref<boolean>;
    tableHasError2: Ref<boolean>;
    tableErrorTip: Ref<string>;
    tableErrorTip2: Ref<string>;
    ruleFormRef: Ref<RuleFormRefType | undefined>;
    ruleForm: RuleFormType;
    rules: RulesType;
};
const reactiveDataOperate = reactiveData();
const formRules = () => {
    const router: Router = useRouter();
    const idcData = store.idcInfo.idcDataNoOpt;
    const languages = locationItem.getLocationItem === 'zh_CN';
    const modelFormStore: ProcessingOperationsType | any = store.modeFormInfo();

    const val = localStorage.getItem('model-item')?? '';
    const newVal = JSON.parse(val);
    // new StorageOperate(newVal);
    const dataCount: DataCountType = {
        volumeNameTip: ref(language.t('modelForm.label.volumeManage.empty.name')),
        volumeTypeTip: ref(language.t('modelForm.label.volumeManage.empty.type')),
        raidCanTip: ref(language.t('modelForm.label.volumeManage.empty.raidCan')),
        sysRaidTip: ref(language.t('modelForm.label.volumeManage.empty.raid')),
        diskTypeTip: ref(language.t('modelForm.label.volumeManage.empty.disk')),
        interfaceTypeTip: ref(language.t('modelForm.label.volumeManage.empty.interfaceType')),
        minSizeNumTip: ref(language.t('modelForm.label.volumeManage.empty.minNum')),
        amountTip: ref(''),
        hasAddFlag: ref<boolean>(true),
        isLoading: ref<boolean>(false),
        hasAddFlag2: ref<boolean>(true),
        raidCanFlag: ref<boolean>(false),
        // 表单ref
        ruleFormRef: ref<RuleFormRefType>(),
        tableHasError: ref<boolean>(false),
        tableHasError2: ref<boolean>(false),
        tableErrorTip: ref(language.t('modelForm.label.volumeManage.save.empty')),
        tableErrorTip2: ref(language.t('modelForm.label.volumeManage.save.empty')),
        // 表单数据
        ruleForm: reactive<RuleFormType>({
            // nicAmount: newVal.nicAmount,
            // gpuModel: newVal.gpuModel,
            // gpuAmount: newVal.gpuAmount,
            // interfaceMode: newVal.interfaceMode,
            // gpuManufacturer: newVal.gpuManufacturer,
            // nicRate: newVal.nicRate,
            // height: newVal.height,
            arrayCard: newVal.isNeedRaid === 'no_need_raid' ? '2' : '1',
            // machineRoomName: languages ? idcData[0]?.name : idcData[0]?.nameEn,
            machineRoomName: (languages ? idcData[0]?.name : idcData[0]?.nameEn),
            // idcId: formData?.value?.idcId || idcData[0]?.idcId, // idcId
            idcId: newVal?.idcId || idcData[0]?.idcId, // idcId
            modelType: reactiveDataOperate.modelFormStore.initModelType(newVal.deviceSeries),
            cpuData: [], // cpu数据
            memData: [], // 内存数据
            bootMode: [],
            volumeManagerTableData: [],
            noConfigurationData: [],
            ...reactiveDataOperate.modelFormStore.ruleForm,
        }),
        // 表单规则
        rules: reactive<RulesType>({
            // height: [
            //     {
            //         required: true,
            //         trigger: 'change',
            //         message: '请选择高度'
            //     }
            // ],
            cpuModel: [
                {
                    required: true,
                    trigger: 'blur',
                    validator: regExpCheck.cpuModelCheck
                }
            ],
            cpuFrequency: [
                {
                    required: true,
                    trigger: 'blur',
                    validator: regExpCheck.cpuFrequencyCheck
                }
            ], // 主频
            name: [ // 机型名称
                {
                    required: true,
                    trigger: 'blur',
                    validator: regExpCheck.nameCheck
                }
            ],
            deviceType: [ // 机型规格
                {
                    required: true,
                    trigger: 'blur',
                    validator: regExpCheck.deviceTypeCheck
                }
            ],
            volumeManager: [
                {
                    required: false,
                    trigger: 'blur',
                    message: ''
                }
            ],
            ...reactiveDataOperate.modelFormStore.rules
        })
    };

    const initVolumeType = (type: string) => {
        const obj = new Map([
            ['system', language.t('modelForm.label.system')],
            ['data', language.t('modelForm.label.data')]
        ]);
        return obj.get(type);
    }

    if (newVal.isNeedRaid === 'no_need_raid') {
        const noConfigData = dataCount.ruleForm.noConfigurationData = newVal.volumes.map((item: any, index: number) => {
            return {
                ...item,
                id: new Date().getTime() + index,
                editFlag: true,
                minSizeNum: item.volumeSize,
                minSize: item.volumeUnit,
                amount: item.volumeAmount,
                // diskType: item.diskType,
                // interfaceType: item.interfaceType,
                // diskType: item.diskType === 'notLimited' ? language.t('modelForm.unrestricted') : item.diskType,
                // interfaceType: item.interfaceType === 'notLimited' ? language.t('modelForm.unrestricted') : item.interfaceType,
                volumeType: initVolumeType(item.volumeType),
                newVolumeType: initVolumeType(item.volumeType)
            }
        });
        // , 'raidCanFlag', 'sysRaidFlag'
        const type = ['volumeNameFlag', 'volumeTypeFlag', 'diskTypeFlag', 'interfaceTypeFlag', 'minSizeNumFlag', 'minSizeNumFlag', 'amountFlag']
        noConfigData.forEach((t: any) => {
            type.forEach((a: any) => {
                t[`${a}${t.id}`] = 'success'
            })
        });
        dataCount.ruleForm.noConfigurationData = noConfigData;
    }
    else {
        const volumeTableData = dataCount.ruleForm.volumeManagerTableData = newVal.volumes.map((item: any, index: number) => {
            return {
                ...item,
                id: new Date().getTime() + index,
                editFlag: true,
                minSizeNum: item.volumeSize,
                sysRaid: item.raid,
                // raidCan: item.raidCan,
                // diskType: item.diskType,
                // interfaceType: item.interfaceType,
                // diskType: item.diskType === 'notLimited' ? language.t('modelForm.unrestricted') : item.diskType,
                // interfaceType: item.interfaceType === 'notLimited' ? language.t('modelForm.unrestricted') : item.interfaceType,
                raidCan: ModelStaticData.singleRaid0.includes(item.raidCan) ? language.t('modelForm.raid') : item.raidCan,
                newSysRaid: item.raidId,
                minSize: item.volumeUnit,
                amount: item.volumeAmount,
                volumeType: initVolumeType(item.volumeType),
                newVolumeType: initVolumeType(item.volumeType)
            }
        });

        const type = ['volumeNameFlag', 'volumeTypeFlag', 'diskTypeFlag', 'interfaceTypeFlag', 'minSizeNumFlag', 'minSizeNumFlag', 'amountFlag', 'raidCanFlag', 'sysRaidFlag']
        volumeTableData.forEach((t: any) => {
            type.forEach((a: any) => {
                t[`${a}${t.id}`] = 'success'
            })
        });
        dataCount.ruleForm.volumeManagerTableData = volumeTableData;
    }
    const setDiskInterfaceType = (item: string) => {
        return item === 'notLimited' ? language.t('modelForm.unrestricted') : item;
    }

    const addVolumeClick = () => {
        if (!dataCount.hasAddFlag.value) return;
        dataCount.hasAddFlag.value = false;
        dataCount.tableHasError.value = false;
        const obj = {
            id: new Date().getTime(),
            editFlag: false,
            volumeName: '',
            volumeType: '',
            newVolumeType: '',
            diskType: '',
            minSize: 'GB',
            raidCan: '',
            amount: 0,
            minSizeNum: '',
            sysRaid: '',
            interfaceType: '',
        }
        dataCount.ruleForm.volumeManagerTableData.push(obj);
    };

    const addVolumeClick2 = () => {
        if (!dataCount.hasAddFlag2.value) return;
        dataCount.hasAddFlag2.value = false;
        dataCount.tableHasError2.value = false;
       const obj = {
            id: new Date().getTime(),
            editFlag: false,
            volumeName: '',
            volumeType: '',
            newVolumeType: '',
            diskType: '',
            minSize: 'GB',
            amount: 1,
            minSizeNum: '',
            interfaceType: '',
        }
        dataCount.ruleForm.noConfigurationData.push(obj);
    }

    const volumeNameBlur = (val: any, index: number) => {
        val[`volumeNameFlag${index}`] = '';
        if (!val.volumeName?.length) {
            val[`volumeNameFlag${index}`] = 'error';
            dataCount.volumeNameTip.value = language.t('modelForm.label.volumeManage.empty.name');
        }
        else if (!RegularContent.name1Reg.test(val.volumeName)){
            val[`volumeNameFlag${index}`] = 'error';
            dataCount.volumeNameTip.value = language.t('modelForm.errorTip.currency');
        }
        else {
            val[`volumeNameFlag${index}`] = 'success';
            dataCount.volumeNameTip.value = '';
        }
    }

    const volumeTypeChange = (val: any, index: number) => {
        val[`volumeTypeFlag${index}`] = '';
        if (!val.volumeType?.length) {
            val[`volumeTypeFlag${index}`] = 'error';
        }
        else {
            if (dataCount.ruleForm.arrayCard === '1') {
                const dataTable = dataCount.ruleForm.volumeManagerTableData;

                if (val.volumeType === language.t('modelForm.label.data') && dataTable.length === 1) {
                    val.volumeType = '';
                    val.newVolumeType = '';
                    dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.sys.tip1')
                    val[`volumeTypeFlag${index}`] = 'error';
                }
                else if (val.volumeType ===  language.t('modelForm.label.system') && dataTable.length > 1 && dataTable.map((item: {volumeType: string}) => item.volumeType).filter((ite: string) => ite === language.t('modelForm.label.system')).length > 1) {
                    val.volumeType = '';
                    val.newVolumeType = '';
                    dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.sys.tip2')
                    val[`volumeTypeFlag${index}`] = 'error';
                }
                else {
                    val[`volumeTypeFlag${index}`] = 'success';
                    dataCount.volumeTypeTip.value  = ''
                    val.newVolumeType = val.volumeType;
                }
            }
            else {
                const noRaidData = dataCount.ruleForm.noConfigurationData;
                if (val.volumeType === language.t('modelForm.label.data') && noRaidData.length === 1) {
                    val.volumeType = '';
                    dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.sys.tip1')
                    val[`volumeTypeFlag${index}`] = 'error';
                    val.newVolumeType = '';
                }
                else if (val.volumeType === language.t('modelForm.label.system') && noRaidData.length > 1 && noRaidData.map((item: {volumeType: string}) => item.volumeType).filter((ite: string) => ite === language.t('modelForm.label.system')).length > 1) {
                    val.volumeType = '';
                    dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.sys.tip2')
                    val[`volumeTypeFlag${index}`] = 'error';
                    val.newVolumeType = '';
                }
                else {
                    val[`volumeTypeFlag${index}`] = 'success';
                    dataCount.volumeTypeTip.value  = ''
                    val.newVolumeType = val.volumeType;
                }
            }
        }
    };

    const raidCanChange = (val: any, index: number) => {
        val[`raidCanFlag${index}`] = '';
        if (!val.raidCan?.length) {
            val[`raidCanFlag${index}`] = 'error';
        }
        else {
            val[`raidCanFlag${index}`] = 'success';
            dataCount.raidCanTip.value  = ''
        }
        if (val.sysRaid?.length) {
            val.sysRaid = ''
            val[`sysRaidFlag${index}`] = void 0;
        }
        if (val.interfaceType) {
            val.interfaceType = '';
            val[`interfaceTypeFlag${index}`] = void 0;
        }
    };

    const sysRaidChange = (val: any, index: number, data: any) => {
        val[`sysRaidFlag${index}`] = '';
        if (typeof val.sysRaid === 'string') {
            // if (data[0] === language.t('modelForm.none')) {
            //     val['newSysRaid'] = ['']
            // }
            // else {
            //     let newId = data[0].value;
            //     val['newSysRaid'] = [newId];
            // }
            let newId = data[0].value;
            val['newSysRaid'] = [newId];
        }
        else {
            let d: any = []
            data.forEach((item: any) => {
                val.sysRaid.forEach((t: any) => {
                    if (item.label === t) {
                        d.push(item.value)
                    }

                })
            })
            val['newSysRaid'] = d;
        }
        if (Array.isArray(val.sysRaid)) {
            if (val.sysRaid.length > 0) {
                val[`sysRaidFlag${index}`] = 'success';
                dataCount.sysRaidTip.value  = '';
            }
            else {
                val[`sysRaidFlag${index}`] = 'error';
                dataCount.sysRaidTip.value  = language.t('modelForm.label.volumeManage.empty.raid')
                // dataCount.sysRaidTip.value  = '请选择RAID模式'
            }
        }
        else {
            if (!val.sysRaid?.length) {
                val[`sysRaidFlag${index}`] = 'error';
            }
            else {
                val[`sysRaidFlag${index}`] = 'success';
                dataCount.sysRaidTip.value  = ''
            }
        }
    };

    const diskTypeChange = (val: any, index: number) => {
        val[`diskTypeFlag${index}`] = '';
        if (!val.diskType?.length) {
            val[`diskTypeFlag${index}`] = 'error';
            // dataCount.raidCanTip.value  = '请选择RAID配置'
        }
        else {
            val[`diskTypeFlag${index}`] = 'success';
            dataCount.diskTypeTip.value  = ''
        }
    }

    const interfaceTypeChange = (val: any, index: number) => {
        val[`interfaceTypeFlag${index}`] = '';
        if (!val.interfaceType?.length) {
            val[`interfaceTypeFlag${index}`] = 'error';
            // dataCount.raidCanTip.value  = '请选择RAID配置'
        }
        else {
            val[`interfaceTypeFlag${index}`] = 'success';
            dataCount.interfaceTypeTip.value  = ''
        }
    }
    
    const minSizeNumBlur = (val: any, index: number) => {
        val[`minSizeNumFlag${index}`] = '';
        if (!String(val.minSizeNum)?.length) {
            val[`minSizeNumFlag${index}`] = 'error';
            dataCount.minSizeNumTip.value = language.t('modelForm.label.volumeManage.empty.minNum')
            // dataCount.raidCanTip.value  = '请选择RAID配置'
        }
        else if (val.minSize === 'TB') {
            val[`minSizeNumFlag${index}`] = 'error';
            if (!RegularContent.capacityReg.test(val.minSizeNum)) {
                dataCount.minSizeNumTip.value  = language.t('modelForm.errorTip.number2')
            }
            else if ((Number(val.minSizeNum) > 100000)) { 
                dataCount.minSizeNumTip.value = language.t('modelForm.errorTip.number3');
            }
            else {
                dataCount.minSizeNumTip.value = ''
                val[`minSizeNumFlag${index}`] = 'success';
            }
        }
        else {
            val[`minSizeNumFlag${index}`] = 'error';
            if (!RegularContent.numberReg.test(val.minSizeNum)) {
                dataCount.minSizeNumTip.value  = language.t('modelForm.errorTip.number1');
            }
            else if ((Number(val.minSizeNum) > 100000)) {
                dataCount.minSizeNumTip.value  = language.t('modelForm.errorTip.number3');
            }
            else {
                dataCount.minSizeNumTip.value = ''
                val[`minSizeNumFlag${index}`] = 'success';
            }
        }
    }

    const minSizeChange = (val: {minSizeNum: string;}) => {
        val.minSizeNum = '';
    }

    watch(() => dataCount.ruleForm.modelCPU, (newValue) => {
        const cpuStatus = ModelStaticData.modelSpecData.includes(newValue);
        if (cpuStatus) {
            regExpCheck.cpuManufacturerFlag.value = false;
            regExpCheck.cpuModelFlag.value = false;
            regExpCheck.cpuInfoFlag.value = false;
            regExpCheck.cpuGHzFlag.value = false;
        }
    })

    watch(() => dataCount.ruleForm.modelStorage, (newValue) => {
        const memeStatus = ModelStaticData.modelSpecData.includes(newValue);
        if (memeStatus) {
            regExpCheck.memTypeFlag.value = false;
            regExpCheck.memFrequencyFlag.value = false;
            regExpCheck.memSizeFlag.value = false;
            regExpCheck.memInfoFlag.value = false;
        }
    })

    const setAmountMaxSize = (item: any, index: number) => {
        if (dataCount.ruleForm.arrayCard === '2' || [language.t('modelForm.raid'), 'NO RAID'].includes(item.raidCan)) {
            item.amount = 1;
            dataCount.amountTip.value = ''
            item[`amountFlag${index}`] = 'success';
            return 1;
        }
        if (!Array.isArray(item.sysRaid)) return;
        // r-l6pounvfife0njlinhk6ztf2xyio
        if (item.sysRaid.includes('RAID10')) {
            if (item.amount % 4 === 0) {
                // item.amountFlag
                item[`amountFlag${index}`] = 'success';
                dataCount.amountTip.value = ''
            }
            return;
        }
        if (item.sysRaid.includes('RAID1')) {
            if (item.amount % 2 === 0) {
                item[`amountFlag${index}`] = 'success';
                dataCount.amountTip.value = ''
            }
            return;
        }

        if (item.sysRaid.includes('RAID0')) {
            if (item.amount >= 2) {
                item[`amountFlag${index}`] = 'success';
                dataCount.amountTip.value = ''
            }
            return;
        }
        // // else 
        // return 1000;
    }

    const setAmountMinSize = (item: any, index: number) => {
        if (dataCount.ruleForm.arrayCard === '2' || [language.t('modelForm.raid'), 'NO RAID'].includes(item.raidCan)) {
            item[`amountFlag${index}`] = 'success';
            return 1;
        }
        if (!Array.isArray(item.sysRaid) || !item.sysRaid?.length) {
            return 1;
        };
        // if (item.sysRaid.length === 1) {
            // r-l6pounvfife0njlinhk6ztf2xyio
            if (item.sysRaid.includes('RAID10')) {
                if (item.amount < 4) {
                    item.amount = 4;
                }
                // else if (item.amount % 4 !== 0) {
                //     // item.amountFlag
                //     item[`amountFlag${index}`] = 'error';
                //     dataCount.amountTip.value = '最低4或者大于4的偶数'
                // }
                return 4;
            }
            if (item.sysRaid.length === 1 && item.sysRaid.includes('RAID5')) {
                item[`amountFlag${index}`] = 'success';
                dataCount.amountTip.value = '';
                if (item.amount < 3) {
                    item.amount = 3;
                }
                return 3;
                
            }
            // if (item.sysRaid.includes('RAID1') && item.sysRaid.includes('RAID5')) {
            //     if (item.amount < 4) {
            //         item.amount = 4;
            //     }
            //     else {
            //         if (item.amount % 2 === 0) {
            //             item[`amountFlag${index}`] = 'success';
            //             dataCount.amountTip.value = ''
            //         }
            //         else {
            //             item[`amountFlag${index}`] = 'error';
            //         dataCount.amountTip.value = '最低2块，或者大于2的偶数'
            //         }
            //     }
            //     return 4;
            // }
            if (item.sysRaid.includes('RAID0') && item.sysRaid.includes('RAID5')) {
                if (item.amount < 3) {
                    item.amount = 3;
                }
                return 3;
            }
            if (item.sysRaid.includes('RAID1')) {
                if (item.amount < 2) {
                    item.amount = 2;
                }
                // else if (item.amount % 2 !== 0) {
                //     // item.amountFlag
                //     item[`amountFlag${index}`] = 'error';
                //     dataCount.amountTip.value = '最低2或者大于2的偶数'
                // }
                return 2;
            }
            if (item.sysRaid.includes('RAID0')) {
                if (item.amount < 2) {
                    item.amount = 2;
                }
                return 2;
            }
    }

    const sure = async () => {
        await dataCount.ruleFormRef.value!.validate((valid: Required<boolean>) => {
            if (!valid) {
                regExpCheck.heightFlag.value = true;
                regExpCheck.interfaceModeFlag.value = true;
                regExpCheck.nicRateFlag.value = true;
            }
            const {volumeManagerTableData, noConfigurationData} = dataCount.ruleForm;
            const cpuStatus = ModelStaticData.modelSpecData.includes(dataCount.ruleForm.modelCPU);
            const memStatus = ModelStaticData.modelSpecData.includes(dataCount.ruleForm.modelStorage);
            setFlag();
            setOtherEmpty(cpuStatus, memStatus)
            if (dataCount.ruleForm.arrayCard === '1') {
                dataCount.tableHasError.value = !(volumeManagerTableData?.length > 0);
                if (!dataCount.hasAddFlag.value) {
                    dataCount.tableHasError.value = true;
                    dataCount.tableErrorTip.value = language.t('modelForm.label.volumeManage.save.volumeManagement');
                    return
                };
                if (dataCount.tableHasError.value) {
                    dataCount.tableHasError.value = true;
                    dataCount.tableErrorTip.value = language.t('modelForm.label.volumeManage.save.empty');
                    return;
                }
            }
            else {
                dataCount.tableHasError2.value = !(noConfigurationData?.length > 0);
                if (!dataCount.hasAddFlag2.value) {
                    dataCount.tableHasError2.value = true;
                    dataCount.tableErrorTip2.value = language.t('modelForm.label.volumeManage.save.volumeManagement');
                    return
                };
                if (dataCount.tableHasError2.value) {
                    dataCount.tableHasError2.value = true;
                    dataCount.tableErrorTip2.value = language.t('modelForm.label.volumeManage.save.empty');
                    return;
                }
            }
            if (valid) {
                regExpCheck.heightFlag.value = false;
                // let params: any = {}
                let cpuBol = false;
                let memBol = false;
                if (cpuStatus && memStatus) {
                    cpuBol = false;
                    memBol = false;
                }
                else if (cpuStatus && !memStatus) {
                    cpuBol = false;
                    memBol = true;
                }
                else if (!cpuStatus && memStatus) {
                    cpuBol = true;
                    memBol = false;
                }
                else {
                    cpuBol = true;
                    memBol = true;
                }
                
                modelFormStore.setCpuMemData(cpuBol, memBol, dataCount.ruleForm)
                .then(() => {
                    requestEditModel(cpuBol, memBol)
                })
            }
            else {
                regExpCheck.heightFlag.value = true;
            }
        })
    }

    const setOtherEmpty = (cpuStatus: boolean, memStatus: boolean) => {
        if (cpuStatus) {
            regExpCheck.cpuManufacturerFlag.value = false;
        }
        else {
            regExpCheck.cpuInfoFlag.value = false;
        }
        if (memStatus) {
            regExpCheck.memTypeFlag.value = false;
            regExpCheck.memFrequencyFlag.value = false;
            regExpCheck.memSizeFlag.value = false;
        }
        else {
            regExpCheck.memInfoFlag.value = false;
        }
    }


    const setVolumeType = (type: string) => {
        const data = [
            [
                (value: string) => value === language.t('modelForm.label.system'),
                () => 'system'
            ],
            [
                (value: string) => value === language.t('modelForm.label.data'),
                () => 'data'
            ]
        ];
        for (const key of data) {
            if (key[0](type)) {
                return key[1](type)
            }
        }
    }


    /**
     * 请求添加机型接口，成功后把事件回传，关闭弹窗
    */
    const requestEditModel = (cpuBol: boolean, memBol: boolean) => {
        const {arrayCard, noConfigurationData, gpuManufacturer, gpuModel, gpuAmount, volumeManagerTableData} = dataCount.ruleForm;
        const obj = gpuAmount > 0 ? {
            gpuAmount,
            gpuModel,
            gpuManufacturer
        } : {}
        const volumesData = arrayCard === '1' ? volumeManagerTableData.map((item: any) => {
            
            return {
                volumeName: item.volumeName,
                volumeType: setVolumeType(item.newVolumeType),
                diskType: item.diskType,
                interfaceType: item.interfaceType,
                volumeSize: String(item.minSizeNum),
                volumeUnit: item.minSize,
                volumeAmount: item.amount,
                raidCan: ModelStaticData.singleRaid0.includes(item.raidCan) ? 'RAID0-stripping' : item.raidCan,
                raid: typeof item.newSysRaid === 'string' ? item.newSysRaid : item.newSysRaid.join(','),
            }
        }) : noConfigurationData.map((item: any) => {
            return {
                volumeName: item.volumeName,
                volumeType: setVolumeType(item.newVolumeType),
                diskType: item.diskType,
                interfaceType: item.interfaceType,
                volumeSize: String(item.minSizeNum),
                volumeUnit: item.minSize,
                volumeAmount: item.amount,
            }
        })
        if (gpuAmount <= 0) {
            const emptyData = ['gpuAmount', 'gpuModel', 'gpuManufacturer'];
            for (const key of emptyData) {
                Reflect.deleteProperty(dataCount.ruleForm, key)
            }
        }
        const newParams = {
            ...obj,
            ...dataCount.ruleForm,
            volumes: volumesData,
            isNeedRaid: setIsNeedRaid(arrayCard),
            cpuSpec: cpuBol ? 'user_defined' : 'common',
            memSpec: memBol ? 'user_defined' : 'common',
            deviceSeries: modelFormStore.setModelType(dataCount.ruleForm.modelType)
        }
        if (arrayCard === '2') {
            const status = noConfigurationData.some((item: {volumeType: string}) => item.volumeType ===  language.t('modelForm.label.system'));
            if (!status) {
                dataCount.tableHasError2.value = true;
                dataCount.tableErrorTip2.value = language.t('modelForm.label.volumeManage.sys.tip1');
                return;
            }
        }
        else {
            const status = volumeManagerTableData.some((item: {volumeType: string}) => item.volumeType ===  language.t('modelForm.label.system'));
            if (!status) {
                dataCount.tableHasError.value = true;
                dataCount.tableErrorTip.value = language.t('modelForm.label.volumeManage.sys.tip1');
                return;
            }
        }
        dataCount.isLoading.value = true;
        addModelAPI(
            {
                ...newParams,
                deviceTypeId: newVal.deviceTypeId
            }
        ).then(({data} : {data: {result: {deviceTypeId: string;}}}) => {
            if(data?.result?.deviceTypeId) {
                msgTip.success(language.t('operate.success'));
                methodsTotal.sendMsg('model-add-template-success', 'success');
                cancel();
            }
        }).finally(() => {
            dataCount.isLoading.value = false;
        });
    };

    // const setVolumeType = (type: string) => {
    //     const obj = new Map([
    //         [language.t('modelForm.label.system'), 'system'],
    //         [language.t('modelForm.label.data'), 'data']
    //     ])
    //     // @ts-ignore
    //     return obj[type];
    // }

    const setIsNeedRaid = (type: string) => {
        const obj = {
            '1': 'need_raid',
            '2': 'no_need_raid'
        };
        // @ts-ignore
        return obj[type];
    }

    const nicAmountIpt = () => {
        if (dataCount.ruleForm.interfaceMode) {
            dataCount.ruleForm.interfaceMode = '';
            regExpCheck.interfaceModeFlag.value = true;
        }
    };


    const setFlag = () => {
        const data = ['cpuInfo', 'memInfo', 'cpuManufacturer', 'memType', 'memFrequency', 'memSize'];
        for (const key of data) {
            // @ts-ignore
            regExpCheck[`${key}Flag`].value = !(dataCount.ruleForm[key])
        }
    }

    const setAddBtn = () => {
        const status = dataCount.ruleForm.volumeManagerTableData.every((item: {editFlag: boolean}) => item.editFlag);
        dataCount.hasAddFlag.value = status;
        return status ? 'active' : 'default-s';
    }

    const setAddBtn2 = () => {
        const status = dataCount.ruleForm.noConfigurationData.every((item: {editFlag: boolean}) => item.editFlag);
        dataCount.hasAddFlag2.value = status;
        return status ? 'active' : 'default-s';
    }

    watch(() => dataCount.ruleForm.arrayCard, (_) => {
        // if (newValue === '2') {
            // dataCount.ruleForm.volumeManagerTableData = [];
            // dataCount.hasAddFlag.value = true;
        // }
    })

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType}) => {
        dataCount.ruleFormRef.value = formEl.value;
        // setFormRef('ruleFormRef', formEl)
        // emitValue('rule-form-ref', dataCount.ruleFormRef);
    };

    const saveClick = (val: any, index: number) => {
        const status = val[`volumeNameFlag${index}`] === 'success' && val[`volumeTypeFlag${index}`] === 'success' && val[`diskTypeFlag${index}`] === 'success' && val[`interfaceTypeFlag${index}`] === 'success' && val[`minSizeNumFlag${index}`] === 'success' && val[`amountFlag${index}`] === 'success';
        if (status && val[`raidCanFlag${index}`] === 'success' && val[`sysRaidFlag${index}`] === 'success') {
            val.editFlag = true;
            dataCount.hasAddFlag.value = true;
            const hasEverySave = dataCount.ruleForm.volumeManagerTableData.every((item: {editFlag: boolean}) => item.editFlag);
            if (hasEverySave) {
                dataCount.tableErrorTip.value = '';
                dataCount.tableHasError.value = false;
            }
        }
        if (dataCount.ruleForm.arrayCard === '1') {
            if (val[`raidCanFlag${index}`] === void 0) {
                dataCount.raidCanTip.value = language.t('modelForm.label.volumeManage.empty.raidCan')
                val[`raidCanFlag${index}`] = '';
            }
            
            if (val[`sysRaidFlag${index}`] === void 0) {
                dataCount.sysRaidTip.value = language.t('modelForm.label.volumeManage.empty.raid')
                val[`sysRaidFlag${index}`] = '';
            }
    
        }
        if (val[`volumeNameFlag${index}`] === void 0) {
            dataCount.volumeNameTip.value = language.t('modelForm.label.volumeManage.empty.name')
            val[`volumeNameFlag${index}`] = '';
            
        }
        if (val[`volumeTypeFlag${index}`] === void 0) {
            dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.empty.type')
            val[`volumeTypeFlag${index}`] = '';
        }
   
        if (val[`diskTypeFlag${index}`] === void 0) {
            dataCount.diskTypeTip.value = language.t('modelForm.label.volumeManage.empty.disk');
            val[`diskTypeFlag${index}`] = '';
        }

        if (val[`interfaceTypeFlag${index}`] === void 0) {
            dataCount.interfaceTypeTip.value = language.t('modelForm.label.volumeManage.empty.interfaceType');
            val[`interfaceTypeFlag${index}`] = '';
        }

        if (val[`minSizeNumFlag${index}`] === void 0) {
            dataCount.minSizeNumTip.value = language.t('modelForm.label.volumeManage.empty.minNum')
            val[`minSizeNumFlag${index}`] = '';
        }
    }

    const noConfigurationSaveClick = (val: any, index: number) => {
        const status = val[`volumeNameFlag${index}`] === 'success' && val[`volumeTypeFlag${index}`] === 'success' && val[`diskTypeFlag${index}`] === 'success' && val[`interfaceTypeFlag${index}`] === 'success' && val[`minSizeNumFlag${index}`] === 'success';
        if (status) {
            val.editFlag = true;

            dataCount.hasAddFlag2.value = true;
            const hasEverySave = dataCount.ruleForm.noConfigurationData.every((item: {editFlag: boolean}) => item.editFlag);
            if (hasEverySave) {
                dataCount.tableErrorTip2.value = '';
                dataCount.tableHasError2.value = false;
            }
            return;
        }

        if (val[`volumeNameFlag${index}`] === void 0) {
            dataCount.volumeNameTip.value = language.t('modelForm.label.volumeManage.empty.name')
            val[`volumeNameFlag${index}`] = '';
            
        }
        if (val[`volumeTypeFlag${index}`] === void 0) {
            dataCount.volumeTypeTip.value = language.t('modelForm.label.volumeManage.empty.type')
            val[`volumeTypeFlag${index}`] = '';
        }
   
        if (val[`diskTypeFlag${index}`] === void 0) {
            dataCount.diskTypeTip.value = language.t('modelForm.label.volumeManage.empty.disk');
            val[`diskTypeFlag${index}`] = '';
        }

        if (val[`interfaceTypeFlag${index}`] === void 0) {
            dataCount.interfaceTypeTip.value = '请选择接口类型';
            val[`interfaceTypeFlag${index}`] = '';
        }

        if (val[`minSizeNumFlag${index}`] === void 0) {
            dataCount.interfaceTypeTip.value = language.t('modelForm.label.volumeManage.empty.interfaceType');
            val[`minSizeNumFlag${index}`] = '';
        }

    }

    watch(() => [dataCount.ruleForm.gpuManufacturer, dataCount.ruleForm.gpuModel], (newValue) => {
        const status = newValue.some(item => item === '');
        if (status) {
            dataCount.ruleForm.gpuAmount = 0;   
        }
    })

    const editClick = (val: {
        split(arg0: string): string[];editFlag: boolean; raidCan: string; raid: string; sysRaid: string[];
}) => {
        
        if (val.raidCan === 'RAID') {
            val.sysRaid = val?.raid ? val.raid.split(',') : val.sysRaid;
        }
        val.editFlag = false;
    }

    const noConfigurationEditClick = (val: {editFlag: boolean;}) => {
        val.editFlag = false;
    }

    const deleteClick = (index: number) => {
        dataCount.ruleForm.volumeManagerTableData.splice(index, 1);
        dataCount.hasAddFlag.value = true;
    }

    const noConfigurationDeleteClick = (index: number) => {
        dataCount.ruleForm.noConfigurationData.splice(index, 1);
        dataCount.hasAddFlag2.value = true;
    }

    const cancel = () => {
        router.push('/ModelManagement/modelList')
    }

    const amountMinValue = () => {
        // const {gpuModel, gpuManufacturer} = dataCount.ruleForm;
        // if (gpuModel && gpuManufacturer) {
        //     return 1;
        // }
        // return 0;
        return 0;
    }


    return {
        sure,
        cancel,
        ...dataCount,
        getFormRef,
        addVolumeClick,
        setDiskInterfaceType,
        amountMinValue,
        volumeNameBlur,
        volumeTypeChange,
        // setRaidInfo,
        raidCanChange,
        sysRaidChange,
        diskTypeChange,
        noConfigurationSaveClick,
        interfaceTypeChange,
        minSizeNumBlur,
        minSizeChange,
        addVolumeClick2,
        setAmountMinSize,
        setAmountMaxSize,
        setAddBtn,
        setAddBtn2,
        saveClick,
        nicAmountIpt,
        editClick,
        noConfigurationEditClick,
        deleteClick,
        noConfigurationDeleteClick
    }
}

export default formRules;