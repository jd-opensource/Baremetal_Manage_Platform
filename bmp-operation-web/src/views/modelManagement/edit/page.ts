const ModelEdit = () => import(/* webpackChunkName: "ModelEdit" */ './modelEdit.vue');

export default {
    path: '/ModelManagement/modelEdit',
    name: 'ModelEdit',
    component: ModelEdit,
    meta: {
        type: 'modelEdit'
    }
};
