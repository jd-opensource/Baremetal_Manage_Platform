import {paginationOperate} from 'utils/publicClass.ts';

const customPagination = (list: any) => {

    watch(() => [
        paginationOperate.pageSize.value,
        paginationOperate.pageNumber.value
    ], (_: boolean[]): void => {
        list.filter.tableRef?.value?.clearSelection()
        if (paginationOperate.routerChange.value) return;
        list.init();
        // list.setFilterData(search, filter?.reactiveArr);
    });
};

export {
    paginationOperate,
    customPagination
};
