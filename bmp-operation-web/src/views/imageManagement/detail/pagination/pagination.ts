import {paginationOperate} from 'utils/publicClass.ts';
/**
 * 分页监听
*/
class PaginationOperate {

    // 构造器
    constructor (list: any) {
        this.watchSizeNumber(list);
    };

    /**
     * 监听分页页数、条数变化
    */
    watchSizeNumber = (list: any) => {
        
        watch(() => [
            paginationOperate.pageSize.value,
            paginationOperate.pageNumber.value
        ], (_: boolean[]): void => {
            if (paginationOperate.routerChange.value) return;
            list.getTable();
        });
    };
};

export default PaginationOperate;