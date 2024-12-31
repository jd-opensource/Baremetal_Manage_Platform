<template>
    <div class="bmp-content">
        <el-header class="header-content">
            <div 
                class="bmp-logo"
                @click="goback"
            >
                <div class="logo-content"><strong>BMP</strong> console                
                </div>
            </div>
            <!-- 项目下拉列表 -->
            <div class="menu-size" >
                <el-dropdown size="default">
                    <span class="project-dropdown">
                        <img :src="folder" class="icon-size mr10">
                        <el-tooltip  
                            :disabled="isShowNameTooltip"  
                            :content= "projectItemName"
                            placement="right-start"
                        >
                            <span 
                                class="long-project-title"  
                                ref="refName" 
                                @mouseover="onMouseOverName(refName)"
                            >
                                {{projectItemName ? projectItemName : $t('list.project.projectMangement')}}
                            </span> 
                        </el-tooltip>                  
                        <img
                            alt=""
                            class="arrow-img"
                            src="@/assets/img/arrow.png"
                        />
                    </span>                   
                    <template #dropdown>
                        <div>
                            <el-dropdown-menu>
                                <el-dropdown-item  
                                    v-for="(item, index) in store.projectList"
                                    v-model="selectedProjectId"
                                    :key="index"
                                    :command="item" 
                                    @click.native="openProject(item)"
                                >
                                    <p class="long-project-name ml10">
                                        {{item.projectName}}
                                    </p>
                                    <el-tooltip
                                        v-if="item.projectId === store.userForm.defaultProjectId"
                                        :content= "$t('personCentre.form.defaultProject')"
                                        placement="right"
                                    >
                                        <img   
                                            alt=""
                                            class="default-img"
                                            src="@/assets/img/default.png" 
                                        />
                                    </el-tooltip>
                                </el-dropdown-item>                                
                            </el-dropdown-menu>
                            <div 
                                class="quit" 
                                @click="jumpToProject"
                            >
                                <div class="quit el-dropdown-menu__item">{{$t('list.project.projectMangement')}}</div>
                            </div>
                        </div>
                    </template>        
                </el-dropdown> 
            </div>
            <!-- 中英文切换 -->
            <div class="button-change">
                <el-button 
                    class="mr15"
                    :class="changeLanguage === 'zh_CN' ? 'header-optional' : 'header-disabled'"
                    key="CN"
                    type="text"
                    @click="changeLang('zh_CN')"
                >
                    {{ "中文" }}
                </el-button>
                    <span class="font-style">{{'/'}}</span>
                <el-button
                    key="EN"
                    type="text"
                    :class="changeLanguage === 'en_US' ? 'header-optional' : 'header-disabled'"
                    @click="changeLang('en_US')"
                >
                    {{ "EN" }}
                </el-button>
                <span class="font-style-line">{{'|'}}</span>
            </div>

            <div class="bmp-logo">
                <img
                    v-if="!messageUnread"
                    class="message-img"
                    src= "@/assets/img/no-message-tip.png"
                    alt="暂无"
                    @click="goMessageList"
                />
                <img
                    v-else
                    class="message-img"
                    src= "@/assets/img/have-message-tip.png"
                    alt="暂无"
                    @click="goMessageList"
                />
            </div>
            <!-- 用户下拉列表 -->
            <img
                class="operate-header-login"
                src="@/assets/img/user-icon.png"
                alt=""
            />
            <div class="personal-menu-content">
                <el-dropdown size="large" class="dropdown-content">
                    <span class="el-dropdown-link project-dropdown">
                        {{'Hello, '}}
                        <span style="color:#108EF9; margin-left: 6px;">{{store.userForm.userName}}</span>
                        <img
                            alt=""
                            class="arrow-img"
                            src="@/assets/img/arrow.png"
                        />
                    </span>
                    <template #dropdown>
                    <el-dropdown-menu >
                        <div class="user-mid">
                            <img
                                class="operate-header-down"
                                src="@/assets/img/user-icon.png"
                                alt=""
                            />
                            <p>{{store.userForm.userName}}</p>
                            <p>{{store.userForm.email}}</p>
                            <el-divider class="divider-type"></el-divider>
                        </div>                        
                        <el-dropdown-item @click="jumpToDetail('account')">{{$t('personCentre.account')}}</el-dropdown-item>
                        <el-dropdown-item @click="jumpToDetail('securitySettings')">{{$t('personCentre.securitySettings')}}</el-dropdown-item>
                        <el-dropdown-item @click="jumpToDetail('sshKey')">{{$t('personCentre.sshKey')}}</el-dropdown-item>
                        <el-dropdown-item @click="jumpToDetail('apiKey')">{{$t('personCentre.apiKey')}}</el-dropdown-item>
                        <div class="quit">
                            <el-dropdown-item @click="quitUser">{{$t('personCentre.quit')}}</el-dropdown-item>
                        </div>                        
                    </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>            
        </el-header>
        <div v-if="close">
            <log-out
                :openVisible="close"
                @close="logoutCancel"
            >
            </log-out>
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    ref,
    Ref,
    watch,
    WritableComputedRef,
    getCurrentInstance,
    onMounted
} from 'vue';
import { useRouter, useRoute} from 'vue-router';
import {useI18n} from 'vue-i18n';
import folder from '@/assets/img/folder.png';
import logOut from 'components/logout/logout.vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import useProjectStore from '@/store/modules/projectId.ts';
import VueCookies from 'vue-cookies'; // cookie
import type {
    FormInstance, // element-plus ui类
} from 'element-plus';
import {userAPI, messageUnreadAPI} from 'api/request.ts'; // 登录接口
import i18n from 'lib/i18n/index.ts'; // 国际化
/**
 * 国际化
*/
const {global} = i18n;
/**
 * 路由
 */
const router = useRouter();
/**
 * 路由带值
 */
 const route = useRoute();
/**
 * store库取值
 */
const store = allProjectStore();

const projectStore = useProjectStore();

// js中使用国际化
const {locale} : {locale: WritableComputedRef<string>} = useI18n();
/**
 * 设置类-
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: string | null = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 用户ID
*/
let userId: string = (VueCookies as any).get('user_id');

/**
 * 用户ID解码
 */
userId = window.atob(userId);

/**
 * 改变的语言
*/
const changeLanguage: Ref<string> = (ref(getLocationItem) as any);

/**
 * 改变语言类型
 * @param {string} lang 语言类型
*/
const changeLang = (lang: string): void => {
    locale.value = lang;
    (VueCookies as any).set('X-Jdcloud-Language', locale.value);
    changeLanguage.value = (VueCookies as any).get('X-Jdcloud-Language');
    location.reload()
};

/**
 * 项目item类型
 */
type ProjectType = {
    createdTime: string;
    projectName: string;
    instanceCount: number;
    projectId: string;
    updatedTime: string;
};
/**
 * 项目名称展示
 */
const projectItemName: Ref<any> = ref<string>('');

/**
 * 选中项目ID
 */
const selectedProjectId: Ref<any> = ref<string>('');
/**
 * 在编辑时，名称赋值
 */
onMounted(() => {
    projectItemName.value = route.query.projectName ? route.query.projectName : '';
})

/**
 * 消息未读标志
 */
 const messageUnread: Ref<boolean> = ref<boolean>(false)

/**
 * 请求是否有信息未读
*/
const requestRead = (): void => {
    messageUnreadAPI({
    }).then(({data}: any) => {
        if(data?.result) {
            messageUnread.value = data.result.hasUnread
            return
        }
        messageUnread.value = false
    }).catch(() => {
        messageUnread.value = false
    }).finally(() => {
    });
}; 

/**
 * 路由
 */
 watch(() => route.path, (): void => {
    if(route.path !== '/Login/login') {
        requestRead()
    }   
});

/**
 * 获取默认项目
 */
 const goback = (): void => {
    userAPI({
    }).then(({data}: {data: any}) => {
        if(data?.result){
            goList(data.result.defaultProjectId, data.result.DefaultProjectName)
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
    projectStore.setProject(id, name);
    router.push({
		path: `/instance/list`,
        query: {
            projectId: id,
            projectName: name
        }
	});
    
}

/**
 * 进入消息中心
 */
 const goMessageList = (): void => {
    router.push({
		path: `/message/list`,
        
	});
    (instanceMitt as any).proxy.$Bus.emit('project-name', '');
};
/**
 * 登出操作打开标志
 */
const close: Ref<boolean> = ref<boolean>(false)
/**
 * 获取项目名称
 * @param {string} item 项目名称
*/
const getProjectName: unknown = (item: string) => {
    projectItemName.value = item
};

/**
 * 监听project-name，获取项目名称
*/
(instanceMitt as any)?.proxy?.$Bus?.on('project-name', getProjectName);

watch(() => store.projectList, (): Readonly<void> => { 
    
    selectedProjectId.value = route.query.projectId
    let projectIds: any = [];
    store.projectList.map((item:any) => {
        if(item.projectId === selectedProjectId.value && route.query.type  !== 'management') {
            projectItemName.value = item.projectName
        }
        projectIds.push(item.projectId)
    });
    if(!projectIds.includes(selectedProjectId.value) && selectedProjectId.value) {
        selectedProjectId.value =  store.projectList[0].projectId;
        projectItemName.value = store.projectList[0].projectName;
        projectStore.setProject(store.projectList[0].projectId, store.projectList[0].projectName);
        router.push({
            path: `/instance/list`,
            query: {
                projectId: store.projectList[0].projectId,
                projectName: store.projectList[0].projectName,
            }
        });
    }

},{deep:true});

// 跳转到项目管理
const jumpToProject = () =>{
    projectItemName.value = '';
    router.push({
        path: `/management`,
    });
}
// 跳转到个人中心
const jumpToDetail = (type: string) =>{
    router.push({
        path: `/user`,
        query: {
            type:type,   
        }
    });
}
/**
 * 名称ref
*/
let refName: Ref<any> = ref<FormInstance>();
/**
 * 判断开启tooltip功能
 */
const isShowNameTooltip: Ref<boolean> = ref<boolean>(true);
const onMouseOverName = (str:any) :void =>  {
    let contentWidth = str.offsetWidth;
    // 判断是否开启tooltip功能
    if (contentWidth > 149) {
        isShowNameTooltip.value = false;
    } else {
        isShowNameTooltip.value = true;
    }

}
/**
 * 展示项目名称
 * @param item 
 */
const openProject = (item: ProjectType) => {
    selectedProjectId.value =  item.projectId
    projectItemName.value = item.projectName;
    projectStore.setProject(item.projectId, item.projectName);
    router.push({
        path: `/instance/list`,
        query: {
            projectId: item.projectId,
            projectName: item.projectName
        }
    });
}
/**
 * logout弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} close.value 弹窗关闭
 */
 const logoutCancel = (type: boolean): boolean => {
    return close.value = type;
};
/**
 * 用户登出
 */
const quitUser = () => {
    close.value = !close.value;
}
</script>

<style lang="scss">
@import './bmpHeader.scss'
</style>