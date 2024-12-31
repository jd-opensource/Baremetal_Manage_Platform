/**
 * @file
 * @author
*/

import {language} from '@utils/publicClass';
import {RuleFormType} from '../typeManagement';
import {RuleFormRefType} from '@utils/publicType';
import store from 'store/index.ts';
import ResourceMonitorStaticData from 'staticData/inBandMonitoring/resourceMonitor/index.ts'

class RuleFormOperate {
    monitoringIndicators: {name: string; value: string}[] = ResourceMonitorStaticData.monitoringIndicators;
    checkAll: Ref<boolean> = ref<boolean>(false);
    indeterminate: Ref<boolean> = ref<boolean>(false);
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        keyData: [],
        userName: '',
        resourceType: language.t('resourceMonitor.search.count.resourceType'),
        idcId: '',
        instanceId: '',
        dimension: language.t('resourceMonitor.search.count.dimensions'),
        monitoringIndicators: []
    });

    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    btnStatus: Ref<boolean> = ref<boolean>(false);
    fn: Function;

    constructor(fn: Function) {
        this.fn = fn;
        watch(() => store.idcInfo.idcData, (newValue: {idcId: string}[]) => {
            this.ruleForm.idcId = newValue[0]?.idcId;
        })
        this.watchMonitoringIndicators();
    };

    watchMonitoringIndicators = () => {
        watch(() => this.ruleForm.monitoringIndicators, (newValue) => {
            this.ruleFormRef.value!.clearValidate('monitoringIndicators')
            if (!newValue?.length) {
                this.checkAll.value = false;
                this.indeterminate.value = false;
            }
            else if (newValue.length === this.monitoringIndicators.length) {
                this.checkAll.value = true;
                this.indeterminate.value = false
            }
            else {
                this.indeterminate.value = true
            }
        })
    }

    handleCheckAll = (val: boolean) => {
        this.ruleForm.monitoringIndicators = val ? this.monitoringIndicators.map((item: {value: string}) => item.value) : [];
    }

    userNameChange = () => {
        this.ruleForm.userName = this.ruleForm.userName?.replace(/\s*/g, '');
    };

    instanceIdChange = () => {
        this.ruleForm.instanceId = this.ruleForm.instanceId?.replace(/\s*/g, '');
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    disabledDate = (time: any) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 30;
    }

    searchClick = () => {
        this.ruleFormRef.value?.validate((valid: boolean) => {
            if (valid) this.fn('search', this.ruleForm);
        });
    };

    clearInstanceId = () => {
        this.ruleForm.instanceId = '';
    }

    clearUserName = () => {
        this.ruleForm.userName = '';
    };

    clearClick = () => {
        const that = this;
        for (let key in this.ruleForm) {
            if (Array.isArray(this.ruleForm[key as keyof RuleFormType])) {
                (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as []) = [];
            }
            if (!['resourceType', 'idcId', 'dimension'].includes(key)) {
                (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as string) = '';
                this.ruleFormRef.value!.clearValidate('userName');
                this.ruleFormRef.value!.clearValidate('instanceId');
            }
        }
        this.fn('', this.ruleForm);
    };
};

export default RuleFormOperate;
