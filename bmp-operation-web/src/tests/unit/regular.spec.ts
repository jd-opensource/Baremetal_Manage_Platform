import {
    phoneReg, // 中国大陆手机号正则
    emailReg, // 电子邮箱正则
    userNameReg, // 用户名正则
    aOmenPhoneReg, // 中国澳门手机号正则
    taiWanPhoneReg, // 中国台湾手机号正则
    hongKongPhoneReg // 中国香港手机号正则
} from '@/utils/regular.ts'

// 不匹配的手机号
const noPhone: string = '132';

// 不匹配的手机号
const noPhone1: string = '523823722212';

// describe 这个作用域里面的测试都是针对xxxx
describe('test-regular', () => {
    // 中国大陆手机号测试
    it('test-phoneReg', () => {
        // 符合规则的正则
        const correctPhone: string = '13212321231';
        // 匹配正则-toBe-精准匹配
        expect(phoneReg.test(correctPhone)).toBe(true);
        // 不匹配
        expect(phoneReg.test(noPhone)).not.toBe(true);
        // 大于或者等于
        expect(correctPhone.length).toBeGreaterThanOrEqual(11)
        // 小于
        expect(noPhone.length).toBeLessThan(4);
    }),

    // 中国香港手机号测试
    it('test-hongKongPhoneReg', () => {
        const correctPhone: string = '51223222';
        // 匹配正则
        expect(hongKongPhoneReg.test(correctPhone)).toBe(true);
        // 不匹配
        expect(hongKongPhoneReg.test(noPhone)).not.toBe(true);
        // 小于或者等于
        expect(correctPhone.length).toBeLessThanOrEqual(8);
        // 大于
        expect((noPhone1).length).toBeGreaterThan(8);
    }),

    // 中国澳门手机号测试
    it('test-aOmenPhoneReg', () => {
        const correctPhone: string = '62343234';
        // 匹配正则
        expect(aOmenPhoneReg.test(correctPhone)).toBe(true);
        // 不匹配
        expect(aOmenPhoneReg.test(noPhone)).not.toBe(true);
        // 小于或者等于
        expect(correctPhone.length).toBeLessThanOrEqual(8);
        // 大于
        expect((noPhone1).length).toBeGreaterThan(8);
    }),

    // 中国台湾手机号测试
    it('test-taiWanPhoneReg', () => {
        // 符合规则的正则
        const correctPhone: string = '0923213212';
        // 匹配正则
        expect(taiWanPhoneReg.test(correctPhone)).toBe(true);
        // 不匹配
        expect(taiWanPhoneReg.test(noPhone)).not.toBe(true);
        // 大于或者等于
        expect(correctPhone.length).toBeGreaterThanOrEqual(10)
        // 小于
        expect(noPhone.length).toBeLessThan(4);
    }),

    // 邮箱测试
    it('test-emailReg', () => {
        const correctEmail: string = 'testEmal@qq.com';
        const notEmail: string = '11.@qq.com';
        // 匹配正则
        expect(emailReg.test(correctEmail)).toBe(true);
        // 不匹配
        expect(emailReg.test(notEmail)).not.toBe(true);
    })

    // 用户名测试
    it('test-userNameReg', () => {
        const correctUserName: string = '1';
        const notUserName: string = '&';
        // 匹配正则
        expect(userNameReg.test(correctUserName)).toBe(true);
        // 不匹配
        expect(userNameReg.test(notUserName)).not.toBe(true);
    })
});
