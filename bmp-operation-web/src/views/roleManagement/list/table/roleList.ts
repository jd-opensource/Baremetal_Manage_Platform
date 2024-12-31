import {CurrencyType} from '@utils/publicType';
import {Router, useRouter} from 'vue-router';
import {inspectuserPurview, paginationOperate, CurrentInstance} from 'utils/publicClass.ts';

// 列表数据
class RoleList {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr: {
        tableData: CurrencyType[];
    } = reactive<{
        tableData: CurrencyType[];
    }>({
        tableData: [] // 表格数据-角色列表
    });
    instanceProxy = new CurrentInstance().proxy;

    // 构造器
    constructor () {
        this.init(useRouter());
    };

    /**
     * 初始化-检查userPurview 状态
    */
    init = (router: Router) => {
        const path: string = this.instanceProxy.$defInfo.routerPath('idcList');
        inspectuserPurview.inspectuserPurview()
        .then(() => {
            this.getRoleList();
        })
        .catch(() => {
            router.push(path);
        });
    };

    /**
     * 获取角色列表
    */
    getRoleList = async () => {
        this.isLoading.value = true;
        try {
            const params = {
                pageNumber: paginationOperate.pageNumber.value,
                pageSize: paginationOperate.pageSize.value
            };
            const roleApi = await this.instanceProxy.$roleApi.roleListAPI(params);
            if (roleApi?.data?.result?.roles?.length) {
                const {roles, totalCount}: {roles: CurrencyType[]; totalCount: number;} = roleApi.data.result;
                this.reactiveArr.tableData = roles;
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
        }
        finally {
            this.isLoading.value = false;
        }
    }
};

export default RoleList;
