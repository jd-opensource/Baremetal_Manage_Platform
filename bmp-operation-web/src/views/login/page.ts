const LoginPage = () => import(/* webpackChunkName: "Login" */ './login.vue');

export default {
    name: 'Login',
    path: '/Login/login',
    component: LoginPage
};
