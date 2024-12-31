type FiltersType = {
    withClon(value: string): string;
    withClonZh(value: string): string;
    emptyFilter(value?: string | null | undefined): string;
    paginationEllipsis(value: string): string;
    defaultPassword(val: string): string;
    defaultWidth(val: number): string | number;
};

type DefType = {
    purview(value: string): void;
    routerPath(value: string): void;
    imagePath(value: string): void;
    encryptDecrypt(num: number): string[] | undefined;
};

interface AppType {
    config: {
        globalProperties: {
            $filter: FiltersType;
            $defInfo: DefType;
        };
    };
};

/**
 * ComponentCustomProperties用于扩充组件实例类型以支持自定义全局属性，否则会报警告。
*/
declare module 'vue' {
    export interface ComponentCustomProperties {
        $filter: FiltersType;
        $defInfo: DefType;
    }
}

const filter = (app: AppType) => app.config.globalProperties.$filter = {
    withClon(value: string): string {
        return value + ':';
    },

    withClonZh(value: string): string {
        return value + '：';
    },

    emptyFilter(value?: string | null | undefined): string {
        switch (value) {
            case '':
            case undefined:
            case 'undefined':
            case null:
                return '--';
            default:
                return value;
        }
    },

    paginationEllipsis(value: string): string {
        return value + '...';
    },

    defaultPassword(val: string): string {
        if (val) {
            return val;
        }
        return '********';
    },

    defaultWidth(val: number): number | string {
        if (val) {
            return val;
        }
        return 'auto';
    }
};

export {
    AppType
}

export default filter;