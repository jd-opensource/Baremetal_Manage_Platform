
import 'polyfill-object.fromentries';
import {mount, shallowMount, DOMWrapper} from '@vue/test-utils'; // 对 Vue 组件进行测试
import ElementPlus from 'element-plus';
import i18n from '@/lib/i18n/index.ts'; // 国际化
import UiButton from '@/components/Ui/Button/Button.vue'; // button组件
import UpDownFrame from '@/components/DiaLog/DeviceOperations/UpDownFrame/UpDownFrame.vue'; // 上架、下架、删除设备操作

/**
 * 插槽信息
*/
const Layout = {
    template: `
                <div>
                    <div class="up-down-frame-header">
                        <slot name="title"></slot>
                    </div>
                    <div class="up-down-frame-footer">
                        <slot name="footer"></slot>
                    </div>
                </div>
            `,
        // data() {
        //     return {
        //         item: {
        //             sn: 'test-sn1',
        //             cabinet: 'test-code',
        //             uPosition: 'test-u',
        //             iloIp: 'test-ip',
        //             privateIpv4: 'test-ipv4'
        //         }
        //     }
        // }
};

describe('test-Button', () => {
    test('test-UiButton', () => {
        const btnText = '按钮';
        const wrapper = mount(UiButton, {
            slots: {
                default: `${btnText}`,
            }
        });
        // expect 让你可以使用不同的“匹配器”去验证不同类型的东西。
        expect(wrapper.find('button').text()).toBe('按钮');
    })
}),


describe('test-UpDownFrame', () => {
    it('test-header', async () => {
        const wrapper = mount(Layout, {
            global: {
                plugins: [ElementPlus, i18n]
            },
            slots: {
                // 会自动匹配默认插槽
                default: '<div></div>',
                title: `
                            <p class="title">
                                <img
                                    src="@/assets/img/navigationBarImg/image-manage.png"
                                    alt=""
                                />
                                <!-- 上架下架删除 -->
                                <span class="title-text">${i18n.global.t('upDownFrame.titleTip.up')}</span>
                            </p>
                            <!-- 标题关闭icon -->
                            <div
                                class="close-icon"
                                @click="diaLogClose"
                            >
                                <ui-icon
                                    color="#999"
                                    :size="16"
                                >
                                    <Close />
                                </ui-icon>
                            </div>
                        `
            }
        });
        // 测试是否含有span
        expect(wrapper.find('.title').html()).toContain('span');
        // 测试内容
        expect(wrapper.find('.title-text').text()).toBe('Mount');
        // 获取触发事件的元素,获取到dom
        const closeClick: DOMWrapper<Element> = wrapper.find('.close-icon');
        // 点击事件,然后使用trigger去触发想要测试的事件，在此使用await 是为了确保在断言之前 你的dom操作会执行完成
        await closeClick.trigger('click');
        // 点击事件是否存在
        expect(wrapper.emitted().click).toBeTruthy();
        // expect(wrapper.element).toMatchSnapshot(); // 快照渲染
    }),

    // 测试
    it('test-count', async () => {
        const wrapper = shallowMount(UpDownFrame, {
            props: {
                diaLog: true, // 蒙层
                title: '上架', // 标题
                deleteStatus: true
            },
            global: {
                plugins: [ElementPlus, i18n]
            },
            slots: {
                default: `<div></div>`
            }
        });
        // 判断类是否存在        
        expect(wrapper.classes('up-down-frame')).toBe(true);
        // 检查类名为“count-title-tip”的元素是否为渲染成功
        expect(wrapper.find('.count-title-tip').exists()).toBe(true);
        expect(wrapper.find('.count-title-tip').text()).toContain('After the server');
        expect(wrapper.find('.header-title').html()).toContain('span')
        // 检查类名为“up-down-frame-count-header-title”的元素是否为渲染成功
        expect(wrapper.find('.header-title').exists()).toBe(true);
        // 检查是否包含...内容
        expect(wrapper.find('.header-title').text()).toContain('Please confirm whether');
        const {vm} = wrapper; 
        await vm.$nextTick(); // 同步获取dom元素渲染
        // expect(wrapper.element).toMatchSnapshot(); // 快照渲染
    }),
    
    // 测试
    it('click & footer', async () => {
        const wrapper = mount(Layout, {
            global: {
                plugins: [ElementPlus, i18n]
            },
            slots: {
                // 会自动匹配默认插槽
                default: '<div></div>',
                footer: `
                            <ui-button
                                class="up-down-frame-footer-cancel-btn"
                                @click="cancelClick"
                            >
                                ${i18n.global.t('upDownFrame.btn.cancel')}
                            </ui-button>
                        `
            }
        });
        // 获取触发事件的元素,获取到dom
        const firstNodeContentWrapper: DOMWrapper<Element> = wrapper.find('.up-down-frame-footer-cancel-btn');
        // 测试是否含有button组件
        expect(wrapper.find('.up-down-frame-footer').html()).toContain('button');
        // 点击事件,然后使用trigger去触发想要测试的事件，在此使用await 是为了确保在断言之前 你的dom操作会执行完成
        await firstNodeContentWrapper.trigger('click');
        expect(wrapper.emitted().click).toBeTruthy();
        expect(wrapper.find('.up-down-frame-footer-cancel-btn').text()).toBe('Cancel');
        // expect(wrapper.element).toMatchSnapshot(); // 快照渲染
    })
});
