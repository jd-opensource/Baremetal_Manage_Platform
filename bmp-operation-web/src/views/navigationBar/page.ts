/**
 * @file
 * @author
*/

import {CurrencyType} from "@utils/publicType";

type NavBarType = () => Promise<typeof import("*.vue")>

const NavigationBar: NavBarType = () => import(/* webpackChunkName: "NavigationBar" */ './navigationBar.vue');

// 创建一个 require 上下文
const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

// 将上下文合并为一个数组
const allRouterData: {default: CurrencyType; noDef: CurrencyType}[] = [
    ...requireAll(require.context('../', true, /\/page.ts$/)),
]

let arr: {default: CurrencyType; noDef: CurrencyType;}[][] = [];

allRouterData.forEach((t, i) => {
    if (!t?.default || t.default?.def || Object.is(t.default.name, 'Login')) {
        arr.push(allRouterData.splice(i, 1));
    }
})

export default {
    nav: {
        name: 'NavigationBar',
        path: '/IdcManagement/idcList',
        component: NavigationBar
    },
    children: [
        arr[0][0].default.noDef,
        ...allRouterData.map((item) => item.default)
    ]
};
