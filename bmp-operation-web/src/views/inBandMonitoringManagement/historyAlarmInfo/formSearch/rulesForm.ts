/**
 * @file
 * @author
*/

import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType} from '../typeManagement';
import {language, paginationOperate} from 'utils/publicClass.ts';

class RuleFormOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        userName: '',
        alarmTime: '',
        startTime: null,
        endTime: null
    });
    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    btnStatus: Ref<boolean> = ref<boolean>(false);
    btnData = [
        {
            type: 'search',
            text: language.t('historyAlarmInfo.search.btn.search')
        },
        {
            type: 'clear',
            text: language.t('historyAlarmInfo.search.btn.clear')
        }
    ];
    fn: Function;

    constructor(fn: Function) {
        this.fn = fn;
        // watch(() => this.ruleForm.alarmTime, (newValue) => {
        //     if (!newValue) {
        // this.ruleFormRef.value?.clearValidate('alarmTime')

        //     }
        // }, {deep: true})
    };

    userNameChange = () => {
        this.ruleForm.userName = this.ruleForm.userName?.replace(/\s*/g, '');
    };

    disabledDate = (time: Date) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 30;
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    searchClear = (type: string) => {
        Object.is(type, 'search') ? this.searchClick() : this.clearClick();
    }

    searchClick = () => {
        this.ruleFormRef.value!.validate((valid: boolean) => {
            if (valid) {
                if (paginationOperate.pageNumber.value > 1) {
                    paginationOperate.pageNumber.value = 1;
                    return;
                }
                this.fn('search');
            }
        })
    };

    clearUserName = () => {
        this.ruleForm.userName = '';
        this.ruleFormRef.value?.clearValidate('userName')
        this.fn('hasSearch');
    };

    clearClick = () => {
        const that = this;
        this.ruleFormRef.value?.resetFields();
        for (let key in this.ruleForm) {
            (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as string) = '';
        }
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn('clear');
    };

    handleBlur = () => {
        this.ruleFormRef.value?.clearValidate('alarmTime')
    }

    dateChange = (val: string[] | null) => {
        if (!val) {
            this.ruleForm.startTime = null;
            this.ruleForm.endTime = null;
            this.fn('hasSearch')
            return;
        };
        this.ruleForm.startTime = new Date(val[0]).getTime() / 1000;
        this.ruleForm.endTime = new Date(val[1]).getTime() / 1000;
    }
};

export default RuleFormOperate;
