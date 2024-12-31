/**
 * @file
 * @author
*/


import {paginationOperate} from 'utils/publicClass.ts';
import {createRouter, createWebHistory, Router, RouterHistory, RouteLocationNormalized, NavigationGuardNext} from 'vue-router';
import routes from '@/router/routes.ts'; //导入router目录下的routes.ts
import {removeLong} from 'request/RemoveLongReq/remove.ts';
//createWebHistory -- history模式
//createWebHashHistory -- hash模式
const webHistory: RouterHistory = createWebHistory();

const router: Router = createRouter({
    history: webHistory,
    routes
});

// 路由守卫
class RouterEach {

    constructor() {
        this.init();
        this.routerError();
    }

    init = () => {
        router.beforeEach((_: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
            paginationOperate.routerChange.value = true;
            paginationOperate.pageNumber.value = 1;
            paginationOperate.pageSize.value = 20;
            removeLong(from.name);
            next();
        });
    }

    routerError = () => {
        router.onError(() => {
            window.location.reload();
        });
    }
}

new RouterEach();

export default router;
