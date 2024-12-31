/**
 * @file
 * @author
*/

import {TableRef} from '../typeManagement';

interface ClassParamsType {
    filterOperates: {
        // filterEmptyInfo: {
        //     deleteEmtpyData: void;
        // }
        reactiveArr: {
            filterParams: {roleId: string;};
        };
        tableRef: TableRef;
    };
    searchOperate: {
        hasClear: {
            value: boolean;
        };
        reactiveArr: {
            searchParams: {};
        };
    };
};

interface ListType {
    reactiveArr: { tableData: string | unknown[]; };
    getUserList: () => void; sessionId: { value: string; };
    resetData: (arg0: {
        hasClear: {
            value: boolean;
        };
        reactiveArr: {
            searchParams: {};
        };
    }) => void;
};

const refreshResetOperate = <T extends ListType, K extends ClassParamsType['filterOperates'], Z extends ClassParamsType['searchOperate']>(list: T, filter: K, search: Z) => {
    const refresh = () => {
        if (!list.reactiveArr.tableData.length) {
            reset();
            return;
        }
        list.getUserList();
    };

    const reset = () => {
        list.sessionId.value = ''
        filter.reactiveArr.filterParams = {roleId: ''};
        search.hasClear.value = true;
        search.reactiveArr.searchParams = {userName: ''};
        filter.tableRef?.value?.clearFilter();
        list.resetData(search);
    };

    return {
        refresh,
        reset
    };
};

export default refreshResetOperate;
