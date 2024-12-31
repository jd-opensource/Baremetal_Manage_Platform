const ModelAdd = () => import(/* webpackChunkName: "ModelAdd" */ './modelAdd.vue');

export default {
    path: '/ModelManagement/modelAdd',
    name: 'ModelAdd',
    component: ModelAdd,
    meta: {
        type: 'modelAdd'
    }
};
