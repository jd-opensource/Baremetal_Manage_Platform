const unmountOperate = () => {
    // 销毁的生命周期
    onUnmounted(async () => {
        const roleId = sessionStorage?.getItem('roleId') || '';
        await roleId && sessionStorage.removeItem('roleId');
    });
};

export default unmountOperate;
