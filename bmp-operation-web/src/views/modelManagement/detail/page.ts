const ModelDetail = () => import(/* webpackChunkName: "ModelDetail" */ './modelDetail.vue');

export default {
    path: '/ModelManagement/modelDetail',
    name: 'ModelDetail',
    component: ModelDetail,
    meta: {
        type: 'modelDetail'
    }
};
