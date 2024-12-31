/**
 * @file
 * @author
*/

import {AxiosError} from 'axios';
import {decrypt, msgTip} from 'utils/index.ts';
import {DataListType, ParamsType, EditUserType} from '../typeManagement';
import {language, inspectuserPurview, paginationOperate, CurrentInstance} from 'utils/publicClass.ts';
import router from 'router/index.ts';
import UserStaticData from 'staticData/user/index.ts';

/**
 * 请求列表，获取数据
*/
class UserList {
    searchTip: Ref<boolean> = ref<boolean>(false);
    sessionId: Ref<string> = ref<string>(''); /// query id
    isLoading: Ref<boolean> = ref<boolean>(true); // loading态
    // 复杂数据类型
    reactiveArr: DataListType = reactive<DataListType>({
        tableData: [] // 表格数据-用户列表
    });
    type: string = '';
    params1: {roleId: string} = {roleId: ''};
    params2: {userName: string} = {userName: ''};
    searchOperate: {hasClear: {value: boolean}} = {hasClear: {value: false}};
    instanceProxy = new CurrentInstance().proxy;
    templateData = [
        {
            prop: 'userName',
            minWidth: '120',
            customInfo: 'userName',
            fixed: 'left',
            label: language.t('userList.label.userName'),
            toolTip: 'userNameTip',
            filterPlacement: '',
        },
        {
            prop: 'userId',
            minWidth: '260',
            customInfo: 'userId',
            fixed: false,
            label: language.t('userList.label.userId'),
            toolTip: 'userIdTip',
            filterPlacement: ''
        },
        {
            prop: 'roleName',
            minWidth: '120',
            customInfo: 'roleName',
            fixed: false,
            label: language.t('userList.label.role'),
            toolTip: '',
            filterPlacement: 'bottom-end',
            key: 'roleId',
            columnKey: 'roleId',
            className: true,
            filterStatus: 'role',
            filterClass: '12',
            filters: 'roleId',
            filterMethod: 'userNameFilter',
            marginRight: '17px',
            setTextClass: false
        },
        {
            prop: 'projectCount',
            minWidth: '80',
            customInfo: 'projectCount',
            fixed: false,
            label: language.t('userList.label.itemNum'),
            toolTip: '',
            setTextClass: false
        },
        {
            prop: 'phoneNumber',
            minWidth: '130',
            customInfo: 'phone',
            fixed: false,
            label: language.t('userList.label.cellPhone'),
            toolTip: 'phoneNumberTip',
        },
        {
            prop: 'email',
            minWidth: '130',
            customInfo: 'email',
            fixed: false,
            label: language.t('userList.label.email'),
            toolTip: 'emailTip',
        },
        {
            prop: 'description',
            minWidth: '160',
            customInfo: 'remark',
            fixed: false,
            label: language.t('userList.label.desc'),
            toolTip: 'descriptionTip',
        },
        {
            prop: 'createdTime',
            minWidth: '160',
            customInfo: 'createTime',
            fixed: false,
            label: language.t('userList.label.createTime'),
            toolTip: '',
            setTextClass: false
        },
        {
            prop: 'operate',
            minWidth: '130',
            customInfo: 'operate',
            fixed: 'right',
            label: language.t('userList.label.operate.name'),
            toolTip: ''
        }
    ];

    // 构造器
    constructor() {
        const id: string | null = sessionStorage?.getItem('roleId');
        const encryptDecrypt: string[] = this.instanceProxy.$defInfo.encryptDecrypt(1)
        this.sessionId.value = id && decrypt(id, encryptDecrypt[0], encryptDecrypt[1]);
        this.init();
    };

    /**
     * 初始化-检查userPurview 状态
    */
    init = () => {
        const path: string = this.instanceProxy.$defInfo.routerPath('idcList');
        inspectuserPurview.inspectuserPurview()
        .then(() => {
            this.getUserList();
        })
        .catch(() => {
            router.push(path)
        });
    };

    /**
      * 请求统一操作
    */
    requestOperate = <T extends {userName: string; roleId: string;}>(params: T, type: string) => {
        this.type = type;
        if (Object.is(type, 'search')) this.params2 = params;
        if (Object.is(type, 'filter')) this.params1 = params;
        this.getUserList();
    };

    resetData = (searchOperate: {hasClear: {value: boolean}}) => {
        this.searchOperate = searchOperate;
        this.searchOperate.hasClear.value = true;
        this.params1 = {roleId: ''};
        this.params2 = {userName: ''};
        this.getUserList();
    }
    /**
     * 请求用户列表接口
    */
    getUserList = async () => {
        this.isLoading.value = true;
        // 搜索id
        const roleId: string = this.params1?.roleId;
        // 初始参数
        let param: ParamsType | {isAll: string;} = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value
        };
        // 如果是从角色进入，要进行数据筛选
        if (this.sessionId?.value || roleId) {
            if (roleId) {
                param = {...param, roleId}
            }
            else {
                param = {...param, roleId: this.sessionId.value};
            }
        }
        if (!this.params1.roleId?.length) {
            Reflect.deleteProperty(this.params1, 'roleId')
        }
        if (!this.params2?.userName?.length) {
            Reflect.deleteProperty(this.params2, 'userName')
        }
        const params = {
            ...param,
            ...this.params1,
            ...this.params2
        };
        try {
            const userApi = await this.instanceProxy.$userApi.userListAPI(params);
            if (userApi?.data?.result?.users?.length) {
                // 解构赋值
                const {users, totalCount} : {users: EditUserType[]; totalCount: number;} = userApi.data.result;
                
                this.setList(users);
                paginationOperate.total.value = totalCount;
                return;
            }
            throw new Error('');
        }
        catch (e) {
            // 异常处理
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            const err = e as AxiosError;
            err.message && msgTip.error(err.message);
        }
        finally {
            // sessionStorage.removeItem('roleId');
            this.searchOperate.hasClear.value = false;
            this.isLoading.value = false;
            this.searchTip.value = Object.keys(params).length > 2;
        }
    };

    setList = (users: EditUserType[], type: string = 'initDefault') => {
        const newVal: EditUserType[] = users?.map((item: EditUserType, index: number) => {
            UserStaticData.userTipData.forEach((t: string) => {
                Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
            });
            return {
                ...item,
                prefix: item.phonePrefix
            }
        });
        this.reactiveArr.tableData = newVal;
        if (type === 'delete') {
            paginationOperate.total.value = paginationOperate.total.value - 1;
        }
    }
};

export default UserList;
