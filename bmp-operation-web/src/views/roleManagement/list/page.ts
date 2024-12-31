const RoleList = () => import(/* webpackChunkName: "RoleList" */ './roleList.vue');

export default {
    path: '/RoleManagement/roleList',
    name: 'RoleList',
    component: RoleList,
    meta: {
        type: 'roleList'
    }
};
