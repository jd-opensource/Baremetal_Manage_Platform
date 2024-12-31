import UserStaticData from 'staticData/user/index.ts';
import {paginationOperate} from 'utils/publicClass.ts';
import {CurrencyType2, CurrencyType3} from '@utils/publicType';
import {TableRef, OSStoreType, UserListType, ReactiveArrType, FilterParamsType} from '../typeManagement';
import store from 'store/index.ts';

class FilterOperate {
    filterEmptyInfo: { // 公共store
        filterParams(
            arg0: CurrencyType2,
            arg1: {roleId: OSStoreType[]},
            arg2: ReactiveArrType,
            arg3: FilterParamsType
    ): Promise<void>;
        deleteEmtpyData(arg0: FilterParamsType): unknown;
    } = store.filterEmpty;
    tableRef: TableRef = {value: {clearFilter(): void{}}};
    // store库存储的oss数据类
    ossStore: {roleId: OSStoreType[]} = {roleId: UserStaticData.userFilterData};
    // 复杂数据类型
    reactiveArr: {
        filterParams: FilterParamsType;
    } = reactive<{
        filterParams: FilterParamsType;
    }>({
        filterParams: {
            roleId: '' // 角色id筛选
        }
    });
    userList: UserListType;

    constructor (list: UserListType) {
        this.reactiveArr.filterParams.roleId = list.sessionId?.value || '';
        this.userList = list;
    };
    

    /**
     * 获取表格ref
     * @param {Object} tableEl 表格信息
    */
    getTableRef = (tableEl: CurrencyType3): void => {
        this.tableRef = tableEl;
    };

    /**
     * 用户名筛选
     * @param {number} value 当前选中的value值
     * @param {Object} row 当前选中的这一项
     * @return {boolean} xxx 返回对应信息
    */
    userNameFilter = (): boolean => {
        // return row.roleName === this.ossStore?.roleId[value - 1]?.text;
        // value: number, row: {roleName: string}
        return true;
    };

    /**
     * 点击【筛选操作】
     * @param {Object} filter 当前点击项
    */
    filterChange = (filter: CurrencyType2) => {
        const filterParams: FilterParamsType = {
            roleId: 'roleId'
        };
        // 筛选完成后重新请求接口，页面赋值为1
        this.filterEmptyInfo.filterParams(filter, this.ossStore, this.reactiveArr, filterParams)
        .then(() => {
            this.filterOperate(this.reactiveArr.filterParams, 'filter');
        })
        .catch(() => {
            this.reactiveArr.filterParams = {roleId: ''};
            this.tableRef?.value?.clearFilter();
            this.filterEmptyInfo.deleteEmtpyData(this.reactiveArr.filterParams);
            this.filterOperate(this.reactiveArr.filterParams, 'filter');
        });
    };

    /**
     * 筛选操作时，页数重置、数据赋值，请求接口
     * @param {Object} filter 当前筛选的数据
     * @param {string} filter.roleId roleId筛选
    */
    filterOperate = (filter: {roleId: string;}, type: string = '') => {
        if (Object.is(type, 'filter') && paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.reactiveArr.filterParams = filter;
        this.userList?.requestOperate(filter, 'filter');
    };
};

export default FilterOperate;
