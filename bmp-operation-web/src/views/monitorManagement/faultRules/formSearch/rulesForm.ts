import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';
import {methodsTotal} from 'utils/index.ts';
class RuleFormOperate {
    btnDisabled: Ref<boolean> = ref<boolean>(true);
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        alertName: '',
        alertSpec: '',
        alertLevel: '',
    });
    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    btnStatus: Ref<boolean> = ref<boolean>(false);
    fn: Function

    constructor(fn: Function) {
        this.fn = fn;
    };

    faultNameChange = () => {
        this.ruleForm.alertName = this.ruleForm.alertName?.replace(/\s*/g, '');
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    searchClick = () => {
        // if (paginationOperate.pageNumber.value > 1) {
        //     paginationOperate.pageNumber.value = 1;
        //     return;
        // }
        // this.fn('search')
        this.paginationNumSet('search');
        // this.list.search();
    };

    singleClear = (type: string) => {
        const that = this;
        (this.ruleForm[`${type}` as keyof typeof that.ruleForm] as string) = '';
       (this.searchParams[`${methodsTotal.toLine(type)}` as keyof typeof that.searchParams] as string)?.length && this.fn('hasSearch');
    };

    clearClick = () => {
        const that = this;
        for (let key in this.ruleForm) {
            (this.ruleForm[`${key}` as keyof typeof that.ruleForm] as string) = '';
        }
        this.paginationNumSet('clear');
        // if (paginationOperate.pageNumber.value > 1) {
        //     paginationOperate.pageNumber.value = 1;
        //     return;
        // }
        // this.fn('clear');
    };

    paginationNumSet = (type: string) => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn(type)
    };
};

export default RuleFormOperate;
