import MockHttpRequest from '../index';
import mockUrl from './config';

/**
 * 设备列表接口
*/
export const mockFaultLogListAPI = () => {
    return MockHttpRequest.service({
        url: mockUrl.faultLogListUrl ,
        method: 'get'
    });
};
