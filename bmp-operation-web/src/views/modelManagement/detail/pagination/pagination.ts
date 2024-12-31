import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 分页监听
*/
class PaginationOperate {

    // 构造器
    constructor (list: any, filter: any) {
        this.watchSizeNumber(list, filter);
    };

    /**
     * 监听分页页数、条数变化
    */
    watchSizeNumber = (list: any, filter: any) => {
        
        watch(() => [
            paginationOperate.pageSize.value,
            paginationOperate.pageNumber.value
        ], (_: boolean[]): void => {
            if (paginationOperate.routerChange.value) return;
            list.setFilterData(filter.reactiveArr, 'filter');
        });
    };
};

export default PaginationOperate;