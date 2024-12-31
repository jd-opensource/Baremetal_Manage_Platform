/**
 * @file
 * @author
*/
import {EchartsDataType, EchartsDataObjType} from '../type';

type ObjType = {
    reactiveArr: {
        echartsData: EchartsDataType[];
    };
    diskVal: {
        value: string;
    };
    mountpointVal: {
        value: string;
    };
    nicVal: {
        value: string;
    };
    getEcharts(arg0: EchartsDataObjType, arg1: string, arg2: { device: string; } | { device?: undefined; }): void;
}

class PublicConfigOpt {
    singleColor: string = '#108ee9';
    moreColor: string[] = [
        '#108ee9',
        '#29d1d3'
    ];

    aggregationMethod: string[]  = ['avg', 'max', 'min', 'sum'];

    clearEmptyObj = <T extends object>(deviceVal: T) => {
        if (!Object.keys(deviceVal).length) {
            Reflect.deleteProperty(deviceVal, 'device')
        }
    }

    smallBigSet = (obj: ObjType, val: string, type: string, value: string) => {
        const newData = obj.reactiveArr.echartsData.filter((item: {typeVal: string}) => item.typeVal === val);
        const {diskVal, mountpointVal, nicVal} = obj;
        const deviceVal = diskVal.value ? {device: diskVal.value} : mountpointVal.value ? {device: mountpointVal.value} : nicVal.value ? {device: nicVal.value} : {};
        this.clearEmptyObj(deviceVal);
        for (const key of newData) {
            if (value?.length) {
                key.aggregationMethod = value;
            }
            obj.getEcharts(key, type, deviceVal);
        }
    }
}

const publicConfigOpt = new PublicConfigOpt();

export default publicConfigOpt;