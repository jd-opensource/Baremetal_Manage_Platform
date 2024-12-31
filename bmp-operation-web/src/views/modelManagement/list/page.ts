const ModelList = () => import(/* webpackChunkName: "ModelList" */ './modelList.vue');

export default {
    path: '/ModelManagement/modelList',
    name: 'ModelList',
    component: ModelList,
    meta: {
        type: 'modelList'
    }
};
