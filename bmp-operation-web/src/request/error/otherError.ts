/**
 * @file
 * @author
*/
import {msgTip} from 'utils/index.ts';

class OtherError {
    setError = (message: string, config: {isBusinessProcessing: boolean}, response: {status: number}) => {
        const dataError = [
            [
                (status: boolean) => status,
                () => {
                    return msgTip.error(message);
                }
            ],
            [
                (status: boolean) => !status,
                () => {
                    return Promise.reject({message, code: response?.status});
                }
            ]
        ];
        for (const key of dataError) {
            if (key[0](config?.isBusinessProcessing?? false)) {
                return key[1](config?.isBusinessProcessing?? false);
            }
        }
    };

    cancelError = () => {
        return Promise.reject('');
    }
}

export default OtherError;
