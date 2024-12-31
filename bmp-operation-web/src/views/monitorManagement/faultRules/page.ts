const FaultRules = () => import(/* webpackChunkName: "FaultRules" */ './faultRules.vue');

export default  {
    path: '/MonitorManagement/faultRules',
    name: 'FaultRules',
    component: FaultRules,
    meta: {
        type: 'faultRules'
    }
};
