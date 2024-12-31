/**
 * @file
 * @author
*/

import {AxiosRequestConfig} from 'axios';
import {RepeatRequestArrType} from '@utils/publicType';

type LongRequestObjName = {
    length: number;
    forEach(arg0: unknown, index?: number): void;
    splice(arg0: number, arg1: number): Function;
    push(arg0: RepeatRequestArrType): void;
};


// 接收长请求的对象
let longRequestObj: RepeatRequestArrType = {
    methodUrl: '',
    params: {},
    f: () => {}
};

class RemoveLongReq  {

    constructor (name: string, config: AxiosRequestConfig & {islongReq?: boolean}, restData: RepeatRequestArrType) {

        if (config?.islongReq) {
            if (!(longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName)) {
                (longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName | []) = [];
            }
            (longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName).push(restData);
        }
    };    
};

const removeLong = (name: string | number) => {
    if ((longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName)?.length > 0) {
        (longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName).forEach((item: {f(): Function;}, index: number) => {
            item.f();
            (longRequestObj[`${name}` as keyof typeof longRequestObj] as LongRequestObjName).splice(index, 1);
       })
    }
}

export {
    removeLong,
    RemoveLongReq
};
