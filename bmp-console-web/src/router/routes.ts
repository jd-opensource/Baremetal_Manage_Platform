import {RouteRecordRaw} from 'vue-router';

const LoginPage: () => void = () => import (/* webpackChunkName: "Login" */ 'views/Login/login.vue');
const PersonalCentrePage: () => void = () => import (/* webpackChunkName: "PersonalCentre" */ 'views/PersonalCentre/personalCentre.vue');
const ProjectManagement: () => void = () => import (/* webpackChunkName: "ProjectManagement" */ 'views/ProjectManagement/projectManagement.vue');
const ProjectManagementDetail: () => void = () => import (/* webpackChunkName: "ProjectManagementDetail" */ 'views/ProjectManagementDetail/projectManagementDetail.vue');
const InstanceList: () => void = () => import (/* webpackChunkName: "InstanceList" */ 'views/InstanceList/instanceList.vue');
const InstanceDetail: () => void = () => import (/* webpackChunkName: "InstanceDetail" */ 'views/InstanceDetail/instanceDetail.vue');
const InstanceCreate: () => void = () => import (/* webpackChunkName: "InstanceCreate" */ 'views/InstanceCreate/instanceCreate.vue');
const MessageList: () => void = () => import (/* webpackChunkName: "MessageList" */ 'views/MessageList/messageList.vue');
const MessageDetail: () => void = () => import (/* webpackChunkName: "MessageDetail" */ 'views/MessageDetail/messageDetail.vue');
const AlarmHistory: () => void = () => import (/* webpackChunkName: "AlarmHistory" */ 'views/AlarmHistory/alarmHistory.vue');
const AlarmRuleList: () => void = () => import (/* webpackChunkName: "AlarmRuleList" */ 'views/AlarmRuleList/alarmRuleList.vue');
const AlarmRuleDetail: () => void = () => import (/* webpackChunkName: "AlarmRuleDetail" */ 'views/AlarmRuleDetail/alarmRuleDetail.vue');
const AlarmRuleCreate: () => void = () => import (/* webpackChunkName: "AlarmRuleCreate" */ 'views/AlarmRuleCreate/alarmRuleCreate.vue');


const routes: Array<RouteRecordRaw> = [
    {
        path: '/:pathMatch(.*)',
        redirect: '/management'
    },
    {
        name: 'Login',
        path: '/Login/login',
        meta: {
            showvalues: true,
        },
        component: LoginPage
    },
    {
        name: 'PersonalCentrePage',
        path: '/user',
        meta: {
            name: 'user',
            noShowBack: true,
            routeDirect: '1'
        },
        component: PersonalCentrePage
    },
    {
        name: 'ProjectManagement',
        path: '/management',
        meta: {
            name: 'management',
            noShowBack: true,
            routeDirect: '1'
        },
        component: ProjectManagement
    },
    {
        name: 'ProjectManagementDetail',
        path: '/management/detail',
        meta: {
            name: 'management',
            noShowBack: false,
            routeDirect: '1'
        },
        component: ProjectManagementDetail
    },
    {
        name: 'InstanceList',
        path: '/instance/list',
        meta: {
            name: 'instance',
            noShowBack: false, 
            routeDirect: '1'  
        },
        component: InstanceList
    },
    {
        name: 'InstanceDetail',
        path: '/instance/detail',
        meta: {
            name: 'instance',
            noShowBack: false,   
            routeDirect: '1'
        },
        component: InstanceDetail
    },
    {
        name: 'InstanceCreate',
        path: '/instance/create',
        meta: {
            name: 'instance',
            noShowBack: false,  
            routeDirect: '1'         
        },
        component: InstanceCreate
    },
    {
        name: 'MessageList',
        path: '/message/list',
        meta: {
            name: 'message',
            noShowBack: true,
            routeDirect: '1'
        },
        component: MessageList
    },
    {
        name: 'MessageDetail',
        path: '/message/detail',
        meta: {
            name: 'message',
            noShowBack: false,
            routeDirect: '1'
        },
        component: MessageDetail
    },
    {
        name: 'AlarmRule',
        path: '/instance/rules/list',
        meta: {
            name: 'alarm',
            noShowBack: false,
            routeDirect: '2-1'
        },
        component: AlarmRuleList
    },
    {
        name: 'AlarmRuleDetail',
        path: '/instance/rules/detail',
        meta: {
            name: 'alarm',
            noShowBack: false,
            routeDirect: '2-1'
        },
        component: AlarmRuleDetail
    },
    {
        name: 'AlarmRuleCreate',
        path: '/instance/rules/create',
        meta: {
            name: 'alarm',
            noShowBack: false,
            routeDirect: '2-1'
        },
        component: AlarmRuleCreate
    },
    {
        name: 'AlarmHistory',
        path: '/instance/history/list',
        meta: {
            name: 'alarm',
            noShowBack: false,
            routeDirect: '2-2'
        },
        component: AlarmHistory
    },
    
];

export default routes;