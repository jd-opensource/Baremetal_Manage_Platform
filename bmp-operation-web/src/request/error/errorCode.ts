/**
 * @file
 * @author
*/
import {msgTip} from 'utils/index.ts';

class ErrorCode {

    setErrorCode = (code: number) => {
        const obj = {
            400: '失败请求',
            404: '参数错误',
            500: '接口异常',
            502: '服务异常'
        };
        const str = (obj[`${code}` as unknown as keyof typeof obj])
        return msgTip.error(str);
    }
}

export default ErrorCode;
