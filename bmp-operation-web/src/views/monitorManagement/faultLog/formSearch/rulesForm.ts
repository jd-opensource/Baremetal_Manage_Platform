/**
 * @file
 * @author
*/

import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';
import {getDateParams} from 'utils/index.ts';

class RuleFormOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        idcId: '',
        sn: '',
        level: '',
        faultReportingTime: '',
        startTime: '',
        endTime: '',
        spec: '',
        dealStatus: ''
    });
    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    btnStatus: Ref<boolean> = ref<boolean>(false);
    fn: Function;

    constructor(fn: Function) {
        const sn: string = sessionStorage?.getItem('sn')?? '';
        this.ruleForm.sn = sn;
        this.fn = fn;
    };

    snChange = () => {
        this.ruleForm.sn = this.ruleForm.sn?.replace(/\s*/g, '');
    };

    disabledDate = (time: Date) => {
        return time.getTime() > Date.now()
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    searchClick = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        const sn = sessionStorage?.getItem('sn')?? '';
        sn?.length ? this.fn('queryParams', this.ruleForm) : this.fn('search');
    };

    clearSn = () => {
        const sn = sessionStorage?.getItem('sn')??'';
        const clearDataOpt = [
            [
                (sn: string) => sn?.length,
                () => {
                    this.fn('queryParams', this.ruleForm)
                    sessionStorage?.removeItem('sn');
                }
            ],
            [
                (sn: string) => !sn?.length,
                () => {
                    this.ruleForm.sn = '';
                    this.searchParams?.sn?.length && this.fn('hasSearch');
                }
            ]
        ];
        for(const key of clearDataOpt) {
            if (key[0](sn)) {
                key[1](sn);
                break;
            }
        }
    };

    singleClear = (type: string) => {
        const that = this;
        (this.ruleForm[`${type}` as keyof typeof that.ruleForm] as string) = '';
        (this.searchParams[`${type}` as keyof typeof that.searchParams] as string)?.length && this.fn('hasSearch')
    };

    clearClick = () => {
        sessionStorage?.removeItem('sn');
        const that = this;
        for (let key in this.ruleForm) {
            (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as string) = '';
        }
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn('clear');
    };

    dateChange = (val: string[] | null) => {
        if (!val) {
            this.ruleForm.startTime = '';
            this.ruleForm.endTime = '';
            this.fn('hasSearch')
            return;
        };
        this.ruleForm.startTime = getDateParams(val[0]);
        this.ruleForm.endTime = getDateParams(val[1]);
    }
};

export default RuleFormOperate;
