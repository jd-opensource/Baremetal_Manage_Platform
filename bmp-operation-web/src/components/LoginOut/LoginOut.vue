<template>
    <!-- 退出登录删除操作 -->
    <custom-dia-log
        :custom-class="'login-out'"
        :dia-log="diaLog"
        :src="($defInfo.imagePath('loginOut') as unknown as string)"
        :is-loading="loginOut.isLoading"
        @dia-log-close="loginOut.cancelClick"
        @determine-click="loginOut.determineClick"
    >
        <div class="currency-count">
            <!-- headerTitle提示 -->
            <p class="header-title">
                <span>{{$t('loginOut.tip')}}?</span>
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {AxiosError} from 'axios';
import {customTip} from 'utils/index.ts';
import {RouterOperate, locationItem, CurrentInstance} from 'utils/publicClass.ts';

// props类型
type PropsType = {diaLog: boolean;};

withDefaults(defineProps<PropsType>(), {});

// emit 事件
const emitValue = defineEmits(['dia-log-close']);

/**
 * 退出登录
*/
class LoginOut {
    router = new RouterOperate().router
    isLoading: Ref<boolean> = ref<boolean>(false);
    instanceProxy = new CurrentInstance().proxy;

    /**
     * 确定按钮弹窗
    */
    determineClick = () => {
        this.isLoading.value = true;
        this.requestLoginOut();
    };

    /**
     * 请求退出登录接口，成功后把事件回传，关闭弹窗
    */
    requestLoginOut = async () => {
        try {
            const loginOutApi = await this.instanceProxy.$loginUserApi.loginOutAPI();
            const status: boolean = (loginOutApi?.data?.result?.success)?? false;
            this.#dealWithRespones(status);
        }
        catch (e) {
            const err = e as AxiosError;
            customTip('error', err.message, 800, () =>  this.isLoading.value = false);
        }
    };

    #dealWithRespones = (status: boolean) => {
        const optData = [
            [
                () => status,
                () => {
                    const pathUrl: string = window.btoa(this.router?.currentRoute?.value?.fullPath);
                    const path: string = this.instanceProxy!.$defInfo.routerPath('login') as unknown as string;
                    locationItem.cookie.set('path_url', pathUrl);
                    localStorage.clear();
                    sessionStorage?.removeItem('roleId');
                    window.open(path, '_self');
                    this.isLoading.value = false;
                    this.cancelClick();
                }
            ]
        ];
        for (const key of optData) {
            if (key[0]) {
                key[1]();
                break;
            }
        }
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 关闭蒙层
        emitValue('dia-log-close', false);
    };
};
// new 实例化
const loginOut: LoginOut = new LoginOut();

</script>
<style lang="scss">
@import './loginOut';
</style>
