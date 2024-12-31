export interface AggregateType {
    DeleteType: {
        Params: {
            name: string;
            age: number;
            sex?: string | number;
        }
    },
    DeepCopyType: {
        Data: {
            name: number;
        }
    },
    OperateType: {
        PropsInfo: {
            title: string;
            diaLog: boolean;
            instanceInfo: string;
        }
    }
};
