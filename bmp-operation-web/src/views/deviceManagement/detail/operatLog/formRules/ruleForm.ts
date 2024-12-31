/**
 * @file
 * @author
*/

import {RuleFormRefType} from '@utils/publicType';
// import {getDateParams} from 'utils/index.ts';
import {paginationOperate} from 'utils/publicClass.ts';

class RuleFormOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    ruleForm = reactive({
        // name: '',
        userName: '',
        operateTime: '',
        startTime: '',
        endTime: ''
    })
    searchParams = reactive({
        ...this.ruleForm
    });
    fn;

    constructor(fn: Function) {
        this.fn = fn;
    }

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    disabledDate = (time: Date) => {
        return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 90;
    };

    userNameChange = () => {
        this.ruleForm.userName = this.ruleForm.userName?.replace(/\s*/g, '');
    }

    clearUserName = () => {
        this.ruleForm.userName = '';
        this.searchClick();
    }

    dateChange = (val: string[] | null) => {
        if (!val) {
            this.ruleForm.startTime = '';
            this.ruleForm.endTime = '';
            this.searchClick();
            return;
        };
        this.ruleForm.startTime = String(new Date(val[0]).getTime() / 1000);
        this.ruleForm.endTime = String(new Date(val[1]).getTime() / 1000);
    }

    searchClick = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn('search');
    };

    clearClick = () => {
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

}

export default RuleFormOperate;
