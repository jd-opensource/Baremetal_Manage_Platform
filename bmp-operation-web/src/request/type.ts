/**
 * @file
 * @author
*/

type OtherParamsType = {
    isBusinessProcessing: boolean;
    timeout?: number;
    responseType?: string;
    islongReq?: boolean;
    baseURL?: string;
};

/**
 * 异常信息类
*/
type ErrInfoType = {
    config: {
        isBusinessProcessing: boolean;
    };
    message: string;
};

type DeleteType = Omit<ErrInfoType, 'message'>;

/**
 * 异常返回类型
*/
interface ErrResponseType extends DeleteType {
    status: number;
    data: {
        error: {
            message: string;
        }
    };
};

export {
    ErrInfoType,
    OtherParamsType,
    ErrResponseType
}