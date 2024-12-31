<template>
    <div class="menu">
        <div class="menu-content">
            <div class="button-position" v-show="$route.meta.noShowBack">
                <el-button 
                    type="text" 
                    :icon="ArrowLeft" 
                    @click="requestUser"
                >
                    {{$t('list.back')}}
                </el-button>
            </div>
            <div 
                v-show="!$route.meta.noShowBack" 
                style="width: 155px; height:1px; display: inline; float: left;"
            >
                {{'  '}}
            </div>
            <div class="menu-position">
                <el-menu
                    :default-active="activeIndex"
                    mode="horizontal"
                    @select="handleSelect"
                >
                <el-menu-item index="1">{{changeMenu($route.meta.name)}}</el-menu-item>
                <el-sub-menu index='2' v-if="($route.meta.name === 'instance' || $route.meta.name === 'alarm') && owned === 1">
                    <template #title>{{$t('alarm.alarmManagement')}}</template>
                    <el-menu-item index="2-1">{{$t('alarm.alarmRules')}}</el-menu-item>
                    <el-menu-item index="2-2">{{$t('alarm.alarmHistory')}}</el-menu-item>
                </el-sub-menu>
                </el-menu>
            </div>
        </div>    
    </div>
</template>

<script setup lang="ts">
import { ref, Ref, watch, onMounted, nextTick,onUnmounted, onBeforeUnmount } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { userAPI } from 'api/request.ts'; // 登录接口
import useProjectStore from '@/store/modules/projectId.ts';
import allProjectStore from '@/store/modules/headerDetail.ts';
import {
    useDebouncedAction
} from 'utils/index.ts';
import {
  ArrowLeft,
} from '@element-plus/icons-vue'
/**
 * 路由带值
 */
const route = useRoute();
const router = useRouter();
const projectStore = useProjectStore();
const activeIndex = ref<string>('1');
const owned: Ref<number> = ref<number>(1);
// 设置初始的 activeIndex
onMounted(() => {
    nextTick(() => {
        trigger()
        
    })   
})

/**
 * store库存储的项目列表
*/
const store = allProjectStore();


const trigger = useDebouncedAction(() => {
    activeIndex.value = route.meta.routeDirect;
    owned.value = findOwnedByProjectId(store.projectList, projectStore.projectId) as number;
}, 300);

watch(() => projectStore.projectId, (): Readonly<void> => {
    owned.value = findOwnedByProjectId(store.projectList, projectStore.projectId) as number;
})

const handleSelect = (key: string, keyPath: string[]) => {
    if(key === '1' && route.meta.name === 'alarm') {
        router.push({
            path: `/instance/list`,
            query: {
                projectId: projectStore.projectId,
                projectName: projectStore.projectName
            }
	    });
    }
    if(key === '2-1') {
        router.push({
            path: `/instance/rules/list`,
            query: {
                projectId: projectStore.projectId,
                projectName: projectStore.projectName
            }
        });
    } else if(key === '2-2') {
        router.push({
            path: `/instance/history/list`,
            query: {
                projectId: projectStore.projectId,
                projectName: projectStore.projectName
            }
        });
    }
}

type Project = {
  projectId: string;
  owned: number;
};

const findOwnedByProjectId = (projects: Project[], projectId: string): number  => {
  const project = projects.find(p => p.projectId === projectId);
  return project?.owned as number;
};

watch(() => route.path, (): void => {  
    nextTick(() => {
        trigger()
    })  
});

/**
 * 获取默认项目
 */
 const requestUser = (): void => {
    userAPI({
    }).then(({data}: {data: any}) => {
        if(data?.result){
            goList(data.result.defaultProjectId, data.result.DefaultProjectName)
            projectStore.setProject(data.result.defaultProjectId, data.result.DefaultProjectName);
        }                  
    }).catch (({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
    })
    .finally(() => {

    });
};
/**
 * 返回列表
 * @param id 
 * @param name 
 */
const goList = (id: string, name: string): void => {
    router.push({
		path: `/instance/list`,
        query: {
            projectId: id,
            projectName: name
        }
	});
}

// js中使用国际化
const {t} = useI18n();
/**
 * 改变菜单名称
 */
const changeMenu = (value: string) => {
    switch (value) {
        case 'instance':
            return t('instance.detail.instance');
        case 'alarm':
            return t('instance.detail.instance');
        case 'management':
            return t('list.project.projectMangement');
        case 'user':
            return t('personCentre.personCentre');
        case 'message':
            return t('messageCentre.message');
        default:
            break;
    }
}
</script>

<style lang="scss">
.menu{  
    background-color: #F7F7F7;
    position: fixed;
    width: 100%;
    z-index: 9;
    .menu-content {
        width: 1450px;
        margin: 60px auto 0 auto;
    }
       
    .menu-position{   
        .el-menu {
            background-color: #F7F7F7;
            width: auto;
            height: 50px;
            .el-sub-menu .el-sub-menu__title {
                justify-content: center;
                &:hover {
                    color: #108EF9;
                    background-color: #F7F7F7
                }
            }
            .el-menu-item {
                padding: 0 20px;
                background-color: #F7F7F7 !important;
                &:hover {
                    color: #108EF9;
                }
                &.is-active{
                    color: #108EF9 !important;
                    background-color: #F7F7F7 !important;
                }
            }
        }
        .el-menu--horizontal {
            border-bottom: none;
        }
    }

    .button-position {
        background-color: #F7F7F7;
        display: inline;      
        float: left;
        :hover {
                background-color: #ecf5ff;
            }
        .el-button {
            padding: 24px 55px 27px 50px;
            border: none;
            color: #333333;
            
        }
        .el-button>span {
            margin-top: 1px;
        }

        .visibility-hidden {
            visibility: hidden;
        }
    }
}

.mb30 {
    margin-bottom: 30px;
} 
</style>