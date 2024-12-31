import {createRouter, createWebHistory, Router, RouterHistory} from 'vue-router';
import routes from 'router/routes.ts'; //导入router目录下的routes.ts

// createWebHistory -- history模式
// createWebHashHistory -- hash模式
const webHistory: RouterHistory = createWebHistory();

const router: Router = createRouter({
    history: webHistory,
    routes
});
export default router;
