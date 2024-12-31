import {RouteRecordRaw} from 'vue-router';
import loginRouter from 'views/login/page.ts';
import navBarRouter from 'views/navigationBar/page.ts';
import idcListRouter from 'views/idcManagement/list/page.ts';

/**
 * 路由信息
*/
const routes: Array<RouteRecordRaw> = [
    idcListRouter.def,
    loginRouter,
    {
        ...navBarRouter.nav,
        children: [...navBarRouter.children]
    }
];

export default routes;
