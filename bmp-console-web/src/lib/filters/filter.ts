// ComponentCustomProperties用于扩充组件实例类型以支持自定义全局属性，否则会报警告。

declare module 'vue' {
    export interface ComponentCustomProperties {
        $filter: any;
    }
}
const filter = (app: any) => app.config.globalProperties.$filter = {
/* eslint-disable indent */
    withClonZh(value: string): string {
        return value + '：';
    },

    withClon(value: string): string {
        return value + ':';
    },

    emptyFilter(value: string | null | undefined): string {
        switch (value) {
            case '':
            case undefined:
            case null:
                return '--';
            default:
                return value;
        }
    },

    numberEmptyFilter(value: string | null | undefined): string {
        switch (value) {
            case '':
            case undefined:
            case null:
                return '0';
            default:
                return value;
        }
    },
    
    paginationEllipsis(value: string): string {
        return value + '...';
    },
};

export default filter;
/* eslint-enable indent */