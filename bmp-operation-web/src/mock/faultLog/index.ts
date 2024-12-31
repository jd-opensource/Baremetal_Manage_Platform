import Mock from 'mockjs';
import mockUrl from './config';

//使用setup配置请求的响应时间，单位是毫秒
Mock.setup({
    timeout: '2000-10000' // 意味这接口响应时间介于2000毫秒-10s之间
})

// 获取 mock.Random 对象
const Random = Mock.Random;

const mockFaultLogList = Mock.mock({
    'data|20': [
        { // 生成30条数据
            'id|+1': 1, // id会自增
            logid: Random.guid(), // 获取全局唯一标识符
            sn: '210200A00JN174002217',
            manageStatusName: '已创建',
            deviceId: 'd-qwird47pir8ejqqvojz6lbtrvp75',
            status: true,
            createTime: '@datetime(yyyy-MM-dd HH:mm:ss)', // 随机生成创建时间
            role: "@pick(['ttt', 'sss', 'rrr', 'fff'])", // 随机生成角色
        }
    ]
});

Mock.mock(mockUrl.faultLogListUrl, 'get', () => {
    return {
        code: 200,
        ...mockFaultLogList,
        totalCount: 20,
    }
});
