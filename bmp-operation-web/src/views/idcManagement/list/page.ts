const IdcList = () => import(/* webpackChunkName: "IdcList" */ './idcList.vue');

export default {
    def: {
        path: '/:pathMatch(.*)',
        redirect: {
			name: 'IdcList'
		}
    },
    noDef: {
        path: '/IdcManagement/idcList',
        name: 'IdcList',
        component: IdcList,
        meta: {
            type: 'idcList'
        }
    }
    // path: '/IdcManagement/idcList',
    //     name: 'IdcList',
    //     component: IdcList,
    //     meta: {
    //         type: 'idcList'
    //     }
};
