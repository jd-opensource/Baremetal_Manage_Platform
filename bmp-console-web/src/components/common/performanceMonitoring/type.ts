/**
 * @file
 * @author
*/

interface NewDataType {
    map: (args: (ite: {value: string;}) => number) => NewDataType;
}[];

interface NewTimeType {
    map: (args: (ite: {timestamp: number;}) => number) => NewTimeType;
}[];

interface HourRangesType {
    '0': string;
    '6': string;
    '12': string;
    '24': string;
};

type EchartsParamsDataType = {
    title: string;
    showType: string[];
    model: string;
    type: string;
    id: string;
    bigClickVal: {
        type: string;
        title: string;
        model: string;
    }
}[];


type PublicType <T> = {
    typeVal: T;
    id: string;
    data: NewDataType[];
    type: string | string[];
    timeData: NewTimeType[];
    aggregationMethod: string;
    metricName: string | string[];
    unit: string;
    hoverUnit: string;
    status?: string[] | undefined;

    paramsMetricName?: string;
};

interface EchartsDataType extends PublicType<string> {
    filter(arg0: (item: {status: string}) => boolean): [];
};

interface EchartsDataObjType extends PublicType<string> {};

interface ReactiveArrType {
    aggregationMethod: string[];
    isLoadingObj: {[x: string]: boolean} | {};
    echartsParamsData: EchartsParamsDataType;
    echartsData: EchartsDataType;
};

type FormType1 = {
    cpu: string;
    mem: string;
    memUsage: string;
    diskUsage: string;
    diskUsageRate: string;
};

type FormType2 = {
    diskReadWriteTraffic: string;
    diskReadWriteIOPS: string;
    netWorkBps: string;
    netWorkPackagesNum: string;
    averageLoad1Min: string;
};

interface RuleFormType extends FormType1, FormType2 {
    averageLoad5Min: string;
    averageLoad15Min: string;
    totalTCPConnections: string;
    normalTCPConnections: string;
    totalNumberOfProcesses: string;
};

export {
    NewDataType,
    NewTimeType,
    RuleFormType,
    HourRangesType,
    EchartsDataType,
    ReactiveArrType,
    EchartsDataObjType,
    EchartsParamsDataType
}