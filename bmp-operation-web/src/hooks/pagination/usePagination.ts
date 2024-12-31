import {paginationOperate} from 'utils/publicClass.ts';


const usePagination = <T extends Function, K extends {clearSelection: Function}>(list: T, clear: K) => {
    watch(() => [
        paginationOperate.pageSize.value,
        paginationOperate.pageNumber.value
    ], (_: boolean[]): void => {
        clear?.clearSelection();
        if (paginationOperate.routerChange.value) return;
        paginationOperate.total.value = 0;
        list();
    });
}

export default usePagination;