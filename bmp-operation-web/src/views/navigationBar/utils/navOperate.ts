/**
 * @file
 * @author
*/

import store from 'store/index.ts';
import {CurrencyType3} from '@utils/publicType';
import {RouterOperate, CurrentInstance, TimerOperate, paginationOperate} from 'utils/publicClass.ts';

class NavOperate extends TimerOperate {
    routerOperate: {router: {currentRoute: {value: {path: string}}; push: any; replace: (arg0: { path: string; }) => void}} = new RouterOperate();
    navTimer: null | number = null;
    init: Ref<boolean> = ref<boolean>(true);
    initShort: Ref<boolean> = ref<boolean>(true);
    userShort: Ref<boolean> = ref<boolean>(true);
    initUser: Ref<boolean> = ref<boolean>(true);
    menuRef: {
        [x: string]: unknown;
        value?: {
            close(arg0: string): unknown;
        }
    } = {value: {close(): void{}}};
    status: {
        InBandMonitoringManagement: number;
        MonitorManagement: number;
        MessageNotification: number;
        PersonalCenter: number;
    } = reactive({
        'InBandMonitoringManagement': 1,
        'MonitorManagement': 1,
        'MessageNotification': 1,
        'PersonalCenter': 1
    })
    instanceProxy = new CurrentInstance().proxy;

    // 构造器
    constructor (purviewInfoOpt: any) {
        super();
        purviewInfoOpt.getPurviewInfo();
        this.watchPath(purviewInfoOpt);        
    };

    select = (path: string) => {
        this.routerOperate.router.push(path);
        // this.routerOperate.router.replace({
        //     path: index
        // });
    };

    /**
     * 监听路径path
    */
    watchPath = (purviewInfoOpt: any) => {
        /**
         * 监听路由path
         * @param {string} newValue 每次更新的路由的值
        */
        watch(() => this.routerOperate.router.currentRoute.value.path, (newValue: string) => {   
            clearTimeout((this.navTimer as number));
            clearTimeout((this.timer as number));
            if (newValue) {
                const {hasNavigationBar} = store.navigationBarStatus;
                this.init.value = hasNavigationBar;
                this.initUser.value = hasNavigationBar;
                this.initShort.value = !hasNavigationBar;
                this.userShort.value = !hasNavigationBar;
                paginationOperate.total.value = 0;
                this.routerOperate.router.currentRoute.value.path = newValue;
                purviewInfoOpt.getPurviewInfo();
                // const path = this.instanceProxy.$defInfo.routerPath('myProfile');
                // const publicPath = this.instanceProxy.$defInfo.routerPath;
                // const path = [publicPath('myProfile'), publicPath('securitySettings'), publicPath('apiKey'), publicPath('license')]
                // if (!path.includes(window.location.pathname)) {
                //     store.userPurviewInfo().userInfo();
                // }
                const typeCount: string[] = ['InBandMonitoringManagement', 'MonitorManagement', 'PersonalCenter', 'MessageNotification'];
                for (const key of typeCount) {
                    // @ts-ignore
                    this.#setClose(newValue, key);
                }
            }
        }, {deep: true});
    };

    #setClose = <T extends string, K extends keyof T>(newValue: T, type: K) => {
        const that = this;
        // @ts-ignore
        if (!newValue.includes(type) && !this.status[type]) {
            this.menuRef.value?.close(String(type));
            (that.status[`${String(type)}` as keyof typeof that.status] as number) ++;
        }
    }

    closeClick = (index: string) => {
        if (['InBandMonitoringManagement', 'MonitorManagement', 'PersonalCenter', 'MessageNotification'].indexOf(index) > -1) {
            const that = this;
            (this.status[`${index}` as keyof typeof that.status] as number) = 1;
        }
    };

    openClick = (index: string) => {
        const that = this;
        
        // if (!['MonitorManagement', 'PersonalCenter', 'MessageNotification'].includes(index)) {
        //     (this.status[`${index}` as keyof typeof that.status] as number) = 1;
        //     return;
        // }
        (this.status[`${index}` as keyof typeof that.status] as number) = 0;
    };

    /**
     * 获取Menu ref
     * @param {Object} menuEl menu信息
    */
    getMenuRef = (menuEl: CurrencyType3): void => {
        this.menuRef = menuEl;
    };

    unfoldCollapseClick = () => {
        for (const index in this.status) {
            // @ts-ignore
            this.status[index] = 0;
        }
        this.init.value = false;
        this.initUser.value = false;
        this.initShort.value = false;
        this.userShort.value = false;
        // this.menuRef.value?.close('MonitorManagement');
        // this.menuRef.value?.close('PersonalCenter');
        // this.navTimer = setTimeout(() => {
            store.navigationBarStatus.hasNavigationBarClick();
        // }, 200);
    };
};

export default NavOperate;
