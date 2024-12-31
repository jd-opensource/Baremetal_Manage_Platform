const UserList = () => import(/* webpackChunkName: "UserList" */ './userList.vue');

export default {
    path: '/UserManagement/userList',
    name: 'UserList',
    component: UserList,
    meta: {
        type: 'userList'
    }
};
