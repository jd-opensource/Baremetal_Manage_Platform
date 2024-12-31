/**
 * @file
 * @author
*/

// keyInfoAPI,systemRaidsAPI
import {queryImagesAPI} from 'api/public/request.ts';
import {instanceDetailAPI} from 'api/device/request.ts';
import {language, CurrentInstance} from 'utils/publicClass.ts';
import {CurrencyType, RuleFormRefType, StatusParamsType, SystemPartitionType} from '@utils/publicType';
import store from 'store/index.ts';
import * as allType from '../type';

class RuleOperate {
    rules: allType.RuleType = {
        hostName: [
            {
                required: false,
                trigger: 'blur',
                validator: ''
            }
        ],
        password: [ // 邮箱
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        confirmPassword: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
    };
    rulesCheck: allType.RulesCheckType;

    constructor (rulesCheck: allType.RulesCheckType) {
        this.rulesCheck = rulesCheck;
        this.rules.hostName[0].validator = rulesCheck.hostNameCheck;
        this.rules.password[0].validator = rulesCheck.passwordCheck;
        this.rules.confirmPassword[0].validator = rulesCheck.confirmPasswordCheck;
    }
};

class DataOperate extends RuleOperate {
    isLoading: Ref<boolean> = ref<boolean>(true);
    selectStatus: Ref<boolean> = ref<boolean>(false);
    showRemaining: Ref<boolean> = ref<boolean>(false);
    currIndex: Ref<number> = ref<number>(0);
    selectCurrIndex: Ref<number> = ref<number>(0);
    systemCurrIndex: Ref<number> = ref<number>(0);
    bootModeCurrIndex: Ref<number> = ref<number>(0);
    reactiveArr: allType.ReactiveArrType = reactive<allType.ReactiveArrType>({
        tableData: [],
        centOSData: [],
        imageTypekeyData: [
            {
                name: '--',
                initName: '--'
            }
        ],
        systemDiskData: [
            {
                raidId: '--',
                raidName: '--',
                volumeDetail: '--',
                diskType: '--'
            }
        ],
        keyPariData: [],
        systemPartitionData: [
            {
                format: '--',
                size: 0,
                point: '--'
            }
        ],
        bootModeData: ['--']
    });
}

class FormRulesOperate extends DataOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    ruleForm: allType.RuleFormType = reactive<allType.RuleFormType>({
        instanceName: '',
        instanceId: '',
        imageType: '',
        bootMode: '',
        systemVolumeRaidId: '',
        hostName: '',
        setType: 'customPassword',
        password: '',
        confirmPassword: '',
        sshKeyId: '',
        imageId: '',
        systemPartition: [],
        userName: '',
        installBmpAgent: true
    });
    proxy = new CurrentInstance().proxy;

    setData = (item: allType.ItemType, query: Pick<allType.QueryType, 'deviceTypeId' | 'idcId'>) => {
        const {instanceName, instanceId} = item;
        this.ruleForm.instanceName = instanceName;
        this.ruleForm.instanceId = instanceId;
        this.reactiveArr.tableData.push({
            ...item,
            showTooltip: false
        })
        this.#getSystemRaids();
        this.#getImages(query)
    };

    
    #getImages = (query: Pick<allType.QueryType, 'deviceTypeId' | 'idcId'>) => {
        queryImagesAPI(
            query
        )
        .then(({data}: {data: {result: {[x: string]: {systemPartition: allType.ReactiveArrType['systemPartitionData']} & allType.DropDownItemType[]}}}) => {
            if (data?.result && Object.values(data.result).length) {
                const val: allType.ReactiveArrType['imageTypekeyData'] = Object.keys(data.result).map((item: string, index: number) => {
                    return {
                        name: item,
                        initName: index === this.currIndex.value ? data.result[item][index].imageName : language.t('resetSystem.placeholder.select'),
                        childrenData: data.result[item],
                        img: this.setImg(item)
                    }
                });
                this.reactiveArr.imageTypekeyData = val;
                this.ruleForm.imageType = data.result[val[0].name][0].imageName;
                this.ruleForm.imageId = data.result[val[0].name][0].imageId;
                this.showRemaining.value = !(this.reactiveArr.imageTypekeyData.length > 2);
                this.reactiveArr.bootModeData = data.result[val[0].name][0].bootMode.split(',');
                this.ruleForm.bootMode = (this.reactiveArr.bootModeData as never)[0];
                const systemPartitionData = JSON.parse((data.result[val[0].name])[0].systemPartition as string)?? this.reactiveArr.systemPartitionData;
                this.reactiveArr.systemPartitionData = this.#setSystemPartitionData(systemPartitionData);
                this.ruleForm.systemPartition = systemPartitionData;
                this.ruleForm.userName = this.#setUserName(data.result[val[0].name] as unknown as string);
                // this.#getSystemRaids(query);
            }
        })
    };

    #setSystemPartitionData = (val: SystemPartitionType[]) => {
        const newData = val?.map((item: SystemPartitionType) => {
            return {
                ...item,
                size: store.sysPartitionInfo.setSizeNoUnit(item.size)
            }
        });
        return newData;
    };

    setImg = (type: string) => {
        const {$defInfo} = this.proxy;
        const imgUrl = new Map([
            ['CentOS', $defInfo.imagePath('centos')],
            ['Ubuntu', $defInfo.imagePath('ubantu')],
            ['Debian', $defInfo.imagePath('debian')],
            ['OpenEuler', $defInfo.imagePath('openEuler')],
            ['Windows', $defInfo.imagePath('windows')]
        ]);
        return imgUrl.get(type)?? $defInfo.imagePath('genericVersion');
    };

    #getSystemRaids = () => {
        instanceDetailAPI(
            {
                instanceId: this.ruleForm.instanceId
            }
        )
        .then(({data}: {data: {result: {instance: CurrencyType}}}) => {
            if (data?.result?.instance?.idcName?.length) {
                const {systemVolumeRaidId, systemVolumeRaidName, systemInfo, systemVolumeType} = data.result.instance;
                this.ruleForm.systemVolumeRaidId = systemVolumeRaidId?? '';
                this.reactiveArr.systemDiskData[0].raidName = systemVolumeRaidName?? '';
                const newVal = systemInfo.split(')')[0];
                this.reactiveArr.systemDiskData[0].volumeDetail = newVal + ' | ' + this.reactiveArr.systemDiskData[0].raidName + ')';
                this.reactiveArr.systemDiskData[0].diskType = systemVolumeType;
                this.reactiveArr.systemDiskData[0].raidId = systemVolumeRaidId;
                // 暂时先注释
                // this.#getKeyInfo();
                return;
            }
        })
    }

    // #getSystemRaids = (query: Omit<allType.QueryType, 'idcId' | 'imageId'>) => {
    //     systemRaidsAPI(
    //         {
    //             deviceTypeId: query.deviceTypeId,
    //             volumeType: 'system'
    //         }
    //     )
        // .then(({data}: {data: {result: {raidId: string}[]}}) => {
        //     if (data?.result?.length) {
        //         this.ruleForm.systemVolumeRaidId = data.result[0].raidId;
        //         this.reactiveArr.systemDiskData = data.result;
        //         // 暂时先注释
        //         // this.#getKeyInfo();
        //         return;
        //     }
        // })
    // };

    // #getKeyInfo = () => {
    //     keyInfoAPI(
    //         {
    //             isAll: '1'
    //         }
    //     )
    //     .then(({data}: {data: {result: {keypairs: CurrencyType[]}}}) => {
    //         if (data?.result?.keypairs?.length) {
    //             this.reactiveArr.keyPariData = data.result.keypairs;
    //             return;
    //         }
    //     })
    //     .finally(() => this.isLoading.value = false)
    // };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    visibleChange: StatusParamsType <(val: boolean) => boolean> = (val: boolean): boolean => {
        return this.selectStatus.value = val;
    };

    visibleClick = (item: allType.DropDwomType, index: number) => {
        this.currIndex.value = index;
        this.ruleForm.userName = this.#setUserName(item.name);
        if (item.initName === language.t('resetSystem.placeholder.select')) {
            this.#setRuleForm(item.childrenData[0]);
            this.reactiveArr.imageTypekeyData.forEach((t) => {
                if (t.name === item.name) {
                    t.initName = t.childrenData![0].imageName;
                }
            });
            return;
        };
        const val = item.childrenData.find((ite: {imageName: string;}) => ite.imageName === item.initName);
        this.#setRuleForm(val!);
    };

    #setUserName = (name: string) => {
        const newUserName = new Map([
            ['Windows', 'administrator']
        ]);
        return newUserName.get(name)?? 'root';
    };
 
    activeClick = () => {
       this.showRemaining.value = !this.showRemaining.value;
       const height: allType.HeightType = document.querySelector('.custom-image') as HTMLDivElement;
        if (this.showRemaining.value) {
            height.style.maxHeight = height.scrollHeight + 'px';
            return;
        }
        height.style.maxHeight = '155px';
    };

    systemDiskClick = (item: {raidId: string}, index: number) => {
        this.ruleForm.systemVolumeRaidId = item.raidId;
        this.systemCurrIndex.value = index;
    };

    radioChange = (item: string, index: number) => {
        this.bootModeCurrIndex.value = index;
        this.ruleForm.bootMode = item;
    };

    dropdownItemChange = (item: allType.ChildrenDataType, index: number) => {
        this.reactiveArr.imageTypekeyData[index].initName = item.imageName;
        this.#setRuleForm(item!);
    };

    #setRuleForm = (itemData: allType.ChildrenDataType) => {
        this.ruleForm.imageId = itemData!.imageId;
        this.ruleForm.imageType = itemData!.imageName;
        this.reactiveArr.bootModeData = itemData!.bootMode.split(',');
        this.ruleForm.bootMode = (this.reactiveArr.bootModeData as never)[0];
        const systemPartitionData = (itemData?.systemPartition && JSON.parse(itemData.systemPartition))?? [{
            format: '--',
            size: 0,
            point: '--'
        }];
        this.ruleForm.systemPartition = systemPartitionData;
        this.reactiveArr.systemPartitionData = this.#setSystemPartitionData(systemPartitionData);
    };

    setPasswordChange = () => {
        this.ruleForm.password = '';
        this.ruleForm.confirmPassword = '';
        this.ruleForm.sshKeyId = '';
        this.rulesCheck.confirmPasswordFlag.value = false;
        this.rulesCheck.colorStatus.value = false;
        this.rulesCheck.hasKeyFlag.value = false;
        this.rulesCheck.errorTip.value = language.t('resetSystem.regCheckTip.password');
    };
};

export default FormRulesOperate;
