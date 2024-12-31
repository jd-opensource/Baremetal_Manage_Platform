
import { defineStore } from 'pinia';
import STORE_NAMES from 'store/storeName.ts'; // 容器名

const useProjectStore = defineStore(STORE_NAMES.PROJECT_ID, {
    state: () => ({
        projectId: '',
        projectName: ''
    }),
    actions: {
        setProject(id: string, name: string) {
            this.projectId = id;
            this.projectName = name;
            // 同时存储到localStorage以保持页面刷新后数据依然存在
            localStorage.setItem('project', JSON.stringify({ projectId: id, projectName: name }));
        },
        clearProject() {
            this.projectId = '';
            this.projectName = '';
            // 清除localStorage中的数据
            localStorage.removeItem('project');
        },
        initProject() {
            // 从localStorage中读取项目数据
            const project = localStorage.getItem('project');
            if (project) {
                const { projectId, projectName } = JSON.parse(project);
                this.projectId = projectId;
                this.projectName = projectName;
            }
        }
        
    },
    // 数据持久化
    persist: true,
});

export default useProjectStore;