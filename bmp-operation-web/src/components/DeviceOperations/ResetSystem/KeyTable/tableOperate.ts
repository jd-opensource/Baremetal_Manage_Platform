import {ChooseItemType} from '../type';
import {CurrencyType, CurrencyType3} from '@utils/publicType';
import {RuleFormType, RulesCheckType} from '../type';

const tableOperate = (ruleForm: RuleFormType, rulesCheck: RulesCheckType) => {
    let tableRef: {
        [x: string]: unknown;
        value?: {
            toggleRowSelection(arg1: CurrencyType, arg2: undefined): void;
            clearFilter(): unknown;
            clearSelection(): void;
        }
    } = {
        value: {
            clearFilter(): void{},
            toggleRowSelection(): void{},
            clearSelection(): void{}
        }
    };

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    const getTableRef = (tableEl: CurrencyType3): void => {
        tableRef = tableEl;
    };

    const cellClick = (row: CurrencyType) => {
        toggleSelection([row]);
    };

    const toggleSelection = (rows?: CurrencyType[]) => {
        if (rows) {
            rows.forEach((row: CurrencyType) => {
                tableRef.value?.toggleRowSelection(row, undefined);
            });
            return;
        }
        tableRef.value?.clearSelection()
    };

    const handleSelectionChange = (val: ChooseItemType[]) => {
        const id: string[] = val.map((item: ChooseItemType) => item.keypairId);
        ruleForm.sshKeyId = id.join(',');
        rulesCheck.hasKeyFlag.value = !(val.length)
    };

    return {
        tableRef,
        getTableRef,
        cellClick,
        handleSelectionChange
    };
};

export default tableOperate;
