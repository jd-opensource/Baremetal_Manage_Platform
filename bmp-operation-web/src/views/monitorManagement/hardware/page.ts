const HardwareStatus = () => import(/* webpackChunkName: "HardwareStatus" */ './hardwareStatus.vue');

export default {
    // def: {
    //     path: '/MonitorManagement',
    //     redirect: {
	// 		name: 'HardwareStatus'
	// 	}
    // },
    // noDef: {
    //     path: '/MonitorManagement/hardwareStatus',
    //     name: 'HardwareStatus',
    //     component: HardwareStatus,
    //     meta: {
    //         type: 'hardwareStatus'
    //     }
    // }
    path: '/MonitorManagement/hardwareStatus',
        name: 'HardwareStatus',
        component: HardwareStatus,
        meta: {
            type: 'hardwareStatus'
        }
};
