const unmountOperate = () => {
    // 销毁的生命周期
    onUnmounted(() => {
        const sn = sessionStorage?.getItem('sn')??'';
        sn && sessionStorage.removeItem('sn');
    });
};

export default unmountOperate;
