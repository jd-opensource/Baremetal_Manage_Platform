const ApiKeyApi = () => import(/* webpackChunkName: "ApiKeyApi" */ './apiKeyApi.vue');

export default {
    path: '/PersonalCenter/apiKeyApi',
    name: 'ApiKeyApi',
    component: ApiKeyApi,
    meta: {
        type: 'apiKeyApi'
    }
};
