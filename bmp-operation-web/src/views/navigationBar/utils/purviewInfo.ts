/**
 * @file
 * @author
*/

import {decrypt} from 'utils/index.ts';
import {CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import eyeryData from 'staticData/navBar/index.ts';

class PurviewInfoOpt {
    purview: Ref<string> = ref<string>('');
    userPurview: {
        userInfo(): void;
        getUserPurview: any;
        userPurview: string;
    } = {
        userInfo: () => '',
        userPurview: '',
        getUserPurview: store.userPurviewInfo()
    };
    instanceProxy = new CurrentInstance().proxy;
    navData;

    constructor() {
        this.navData = eyeryData(this.instanceProxy)
    }

    navLeft1 = () => {
        return eyeryData(this.instanceProxy).slice(0, 4);
    };

    navLeft2 = () => {
        const purview: string = this.instanceProxy!.$defInfo.purview('admin');
        const data = eyeryData(this.instanceProxy).filter((item: {isF: boolean}) => item.isF);
        const status = this.purview.value === purview;
        const info = data.filter((item: {hasAdmin: boolean}) => status ? item.hasAdmin : (!item.hasAdmin))
        if (info[0].path === '/UserManagement/userList') {
            return info.reverse();
        }
        return info;
    };

    outOfMonitoring = () => {
        return this.navData.filter((item: {childrenCentF: boolean}) => item.childrenCentF);
    }

    messagePersonal = () => {
        const purview: string = this.instanceProxy!.$defInfo.purview('admin');
        const data = this.navData.filter((item: {isChildrenF: boolean}) => item.isChildrenF);
        const status = this.purview.value === purview;
        return status ? data : data.filter((item: {childrenHasAdmin: boolean}) => item.childrenHasAdmin);
    }

    /**
     * 菜单栏默认态
     * @param {Object} item 当前点击的这一项路由
     * @param {string} item.path 当前点击的路由地址
     * @return {string} /xxx/xxx 路由地址
    */
    defaultActive = (item: {path: string;}): string =>  {
        for (const key of eyeryData(this.instanceProxy)) {
            if (key?.children?.length) {
                for (const listKey of key.children) {
                    if (listKey.otherPath.includes(item.path)) {
                        return listKey.path;
                    }
                }
            }
            if (key?.otherPath.includes(item.path)) {
                return key.path;
            }
        }
        return item.path;
    };

    /**
     * 获取权限信息
    */
    getPurviewInfo = () => {
        try {
            const status: string = localStorage?.getItem('userPurview')?? '';
            this.purview.value = (decrypt(status));
            const operator: string = this.instanceProxy!.$defInfo.purview('operator')?? '';
            const admin: string = this.instanceProxy!.$defInfo.purview('admin')?? '';
            const userName: string = localStorage?.getItem('userName')?? '';
            const email: string = localStorage?.getItem('email')?? '';
            const requestData = [
                [
                    (operator: string, admin: string) => ![operator, admin].includes(this.purview.value as never),
                    () => {
                        this.userPurview?.getUserPurview.getUserPurview(this.instanceProxy.$loginUserApi)
                        .then((res: string) => {
                            this.userPurview.userPurview = res;
                            store.userPurviewInfo().userInfo()
                        });
                    }
                ],
                [
                    () => ([userName, email].includes('')),
                    async () => store.userPurviewInfo().userInfo()
                ]
            ];
            for (const key of requestData) {
                if (key[0](operator, admin)) {
                    key[1](operator, admin);
                    break;
                }
            }
        }
        catch {
            throw 'error';
        }
    };
}

export default PurviewInfoOpt;
