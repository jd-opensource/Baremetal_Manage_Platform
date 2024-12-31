import {defineStore} from 'pinia';
import STORE_NAMES from 'store/storeName.ts';
import {projectListAPI, userAPI} from 'api/request.ts'; // 项目，用户接口

interface StateType {
    projectList: any;
    projectListName: any;
    userForm: any;
    monitor: boolean;
    inMonitor: boolean;
    moveLive: string;
}

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
 const allProjectStore = defineStore(STORE_NAMES.ALLPROJECT, {
    state: (): StateType => {
        return {
            projectList: [],
            projectListName: [],
            userForm: {},
            monitor: false,
            inMonitor: false,
            moveLive: 'test数据'
        }
    },
    getters: {

    },
    actions: {
        test(): void {
            this.moveLive = '1';
        },
        requestObject(): void {
           return  projectListAPI({
                isAll:'1'
            }).then(({data} : {data: any}) => {
                if (data?.result?.projects?.length) {
                    this.projectList = data.result.projects; 
                    this.projectListName = data.result.projects.map((item: any) => item.projectName);                   
                    return
                }
            this.projectList  = [];
            this.projectListName = [];
            }).catch(() => {
                this.projectList  = [];
                this.projectListName = [];
            });
        },
        requestUser(): void {
            userAPI({
            }).then(({data} : {data: any}) => {
                if(data?.result){
                    this.userForm = data.result; 
                    return
                } 
                this.userForm = {}
            }).catch(()=>{
                this.userForm = {}
            })
            .finally(() => {
            });
        }
    },
    // 数据持久化
    persist: true,
});

export default allProjectStore;