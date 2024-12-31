import {RuleFormRefType} from '@utils/publicType';
import {RuleFormType} from '../typeManagement';
import {paginationOperate} from 'utils/publicClass.ts';

class RuleFormOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        idcId: '',
        sn: '',
    });
    searchParams: RuleFormType = reactive<RuleFormType>({
        ...this.ruleForm
    });
    fn: Function;
    btnStatus: Ref<boolean> = ref<boolean>(false);

    constructor(fn: Function) {
        this.fn = fn;
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };

    snChange = () => {
        this.ruleForm.sn = this.ruleForm.sn?.replace(/\s*/g, '');
    };

    searchClick = () => {
        this.paginationNumSet('search');
    };

    paginationNumSet = (type: string) => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn(type)
    };

    clearSn = () => {
        this.ruleForm.sn = '';
        this.searchParams?.sn?.length && this.fn('hasSearch', this.searchParams)
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
        // this.fn('clear')
        // this.ruleFormRef.value!.resetFields();
        // console.log(this.ruleForm, 'ruleForm');
    };
};

export default RuleFormOperate;
