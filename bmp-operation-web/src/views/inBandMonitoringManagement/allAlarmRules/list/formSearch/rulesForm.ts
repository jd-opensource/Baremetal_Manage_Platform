/**
 * @file
 * @author
*/

import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType} from '../typeManagement';
import {language, paginationOperate} from 'utils/publicClass.ts';
// import {getDateParams} from 'utils/index.ts';

let that;
class RuleFormOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        userName: '',
        ruleName: '',
        ruleId: ''
    });
    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    btnStatus: Ref<boolean> = ref<boolean>(false);
    fn: Function;
    filter;
    formData = [
        {
            prop: 'userName',
            label: language.t('allAlarmRulesList.search.label.userName'),
            model: 'userName',
            placeholder: language.t('allAlarmRulesList.search.placeholder.userName'),
            change: 'userName',
            clear: 'userName'
        },
        {
            prop: 'ruleName',
            label: language.t('allAlarmRulesList.search.label.ruleName'),
            model: 'ruleName',
            placeholder: language.t('allAlarmRulesList.search.placeholder.ruleName'),
            change: 'ruleName',
            clear: 'ruleName'
        },
        {
            prop: 'ruleId',
            label: language.t('allAlarmRulesList.search.label.ruleId'),
            model: 'ruleId',
            placeholder: language.t('allAlarmRulesList.search.placeholder.ruleId'),
            change: 'ruleId',
            clear: 'ruleId'
        }
    ];
    btnData = [
        {
            type: 'search',
            text: language.t('allAlarmRulesList.search.btn.search')
        },
        {
            type: 'clear',
            text: language.t('allAlarmRulesList.search.btn.clear')
        }
    ];

    constructor(filter: any, fn: Function) {
        this.filter = filter;
        this.fn = fn;
    };

    iptChange = (type: string) => {
        this.#replaceEmpty(type);
    }
    
    #replaceEmpty = (type: string) => {
        that = this;
        this.ruleForm[`${type}` as keyof typeof that.ruleForm] = this.ruleForm[`${type}` as keyof typeof that.ruleForm]?.replace(/\s*/g, '');
    }

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    searchClearClick = (type: string) => {
        Object.is(type, 'search') ? this.searchClick() : this.clearClick();
    }

    searchClick = () => {
        this.ruleFormRef.value?.validate((valid: boolean) => {
            if (valid) {
                if (paginationOperate.pageNumber.value > 1) {
                    paginationOperate.pageNumber.value = 1;
                    return;
                }
                this.fn('search');
            }
        });
    };

    clearIpt = (type: string) => {
        const that = this;
        this.ruleFormRef.value?.clearValidate(type);
        this.ruleForm[`${type}` as keyof typeof that.ruleForm] = '';
        this.fn('hasSearch')
    }

    clearClick = () => {
        const that = this;
        for (let key in this.ruleForm) {
            (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as string) = '';
        }
        this.ruleFormRef.value?.resetFields();
        this.filter?.tableRef?.value?.clearFilter();
        this.filter.reactiveArr.filterParams = {};
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn('clear');
    };
};

export default RuleFormOperate;
